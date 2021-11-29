package loadgen

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/pyroscope-io/pyroscope/benchmark/internal/config"
	"github.com/pyroscope-io/pyroscope/pkg/agent/upstream"
	"github.com/pyroscope-io/pyroscope/pkg/agent/upstream/remote"
	"github.com/pyroscope-io/pyroscope/pkg/structs/transporttrie"
	"github.com/sirupsen/logrus"
)

// how many retries to check the pyroscope server is up
const MaxReadinessRetries = 10

var CapacityExceeded = errors.New("capacity exceeded")

type Fixtures [][]*transporttrie.Trie

type LoadGen struct {
	Config    *config.LoadGen
	Rand      *rand.Rand
	SymbolBuf []byte

	runProgressMetric prometheus.Gauge
	uploadErrors      prometheus.Counter
	successfulUploads prometheus.Counter
	pusher            GatewayPusher
}

type GatewayPusher interface {
	Push() error
}
type NoopGatewayPusher struct{}

func (NoopGatewayPusher) Push() error {
	return nil
}

func Cli(cfg *config.LoadGen) error {
	r := rand.New(rand.NewSource(int64(cfg.RandSeed)))

	var pusher GatewayPusher
	if cfg.PushgatewayAddress == "" {
		logrus.Debug("no pushgateway configured")
		pusher = NoopGatewayPusher{}
	} else {
		logrus.Debug("will push metrics to ", cfg.PushgatewayAddress)
		pusher = push.New(cfg.PushgatewayAddress, cfg.ServerAddress).Gatherer(prometheus.DefaultGatherer)
	}

	l := &LoadGen{
		Config:    cfg,
		Rand:      r,
		SymbolBuf: make([]byte, cfg.ProfileSymbolLength),

		runProgressMetric: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "pyroscope",
			Subsystem: "benchmark",
			Name:      "progress",
			Help:      "",
		}),
		uploadErrors: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "pyroscope",
			Subsystem: "benchmark",
			Name:      "upload_errors",
			Help:      "",
		}),
		successfulUploads: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "pyroscope",
			Subsystem: "benchmark",
			Name:      "successful_uploads",
			Help:      "",
		}),
		pusher: pusher,
	}

	promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: "pyroscope",
			Subsystem: "benchmark",
			Name:      "requests_total",
			Help:      "",
		},
		func() float64 { return float64(cfg.Apps * cfg.Requests * cfg.Clients) },
	)

	if cfg.CapacityBenchmark {
		for {
			if err := l.Run(cfg); err == CapacityExceeded {
				break
			} else if err != nil {
				return err
			}
			cfg.Clients *= 2
		}

		//////////////////////////////////////////////////////////////////////////
		/*
			memProfile, err := os.Create("mem.prof")
			if err != nil {
				log.Fatal(err)
			}
			if err := pprof.WriteHeapProfile(memProfile); err != nil {
				log.Fatal(err)
			}
			memProfile.Close()
		*/
		//////////////////////////////////////////////////////////////////////////

		minClients := cfg.Clients / 2
		//logrus.Infof("minClients: %d", minClients)
		maxClients := cfg.Clients
		//logrus.Infof("maxClients: %d", maxClients)
		logrus.Infof(
			"client capacity: %d",
			minClients-
				10+ // substract 10 to get the highest good result, not the lowest bad result
				10*sort.Search((maxClients-minClients)/10, func(n int) bool {
					cfg.Clients = minClients + 10*n
					if err := l.Run(cfg); err == CapacityExceeded {
						return true
					} else if err != nil {
						panic(err) // wish we could do a Ruby-style return here
					}
					return false
				}),
		)
		return nil
	}

	return l.Run(cfg)
}

func (l *LoadGen) Run(cfg *config.LoadGen) error {
	runtime.GC() // so fresh and so clean
	logrus.Info("checking server is available...")
	err := waitUntilEndpointReady(cfg.ServerAddress)
	if err != nil {
		return err
	}

	logrus.Info("generating fixtures")
	fixtures := l.generateFixtures()
	logrus.Debug("done generating fixtures.")

	logrus.Info("starting sending requests")
	logrus.Infof("cfg %+v\n", cfg)
	wg := sync.WaitGroup{}
	wg.Add(l.Config.Apps * l.Config.Clients)
	appNameBuf := make([]byte, 25)

	// Memory-inefficient but obviously correct histogram.
	var latencyHistogram []uint32
	var errorCounter *uint64
	if cfg.CapacityBenchmark {
		latencyHistogram = make(
			[]uint32,
			7500+ // number of milliseconds we care about per use-case
				1+ // an extra to capture samples that exceed the threshold
				1, // and one more for slice length being one more than the highest index
		)
		errorCounter = new(uint64)
	}

	for i := 0; i < l.Config.Apps; i++ {
		// generate a random app name
		l.Rand.Read(appNameBuf)
		appName := hex.EncodeToString(appNameBuf)
		for j := 0; j < l.Config.Clients; j++ {
			go l.startClientThread(appName, &wg, fixtures[i], latencyHistogram, errorCounter)
		}
	}
	wg.Wait()

	logrus.Debug("done sending requests")

	// Latency quantiles and error rate. If latency's too high or error rate
	// is non-zero, we've exceeded this instance's capacity.
	if cfg.CapacityBenchmark {
		var p50, p90, p99 time.Duration
		requests, total := uint32(0), uint32(cfg.Apps*cfg.Clients*cfg.Requests)

		for milliseconds, samples := range latencyHistogram {
			requests += samples
			if requests >= total*99/100 {
				p99 = time.Duration(milliseconds) * time.Millisecond
				break
			} else if requests >= total*90/100 {
				p90 = time.Duration(milliseconds) * time.Millisecond
			} else if requests >= total*50/100 {
				p50 = time.Duration(milliseconds) * time.Millisecond
			}
		}
		if p90 == 0 {
			p90 = p99
		}
		if p50 == 0 {
			p50 = p90
		}
		errorRate := float64(atomic.LoadUint64(errorCounter)) / float64(total)

		logrus.Infof("50th percentile latency: %v", p50)
		logrus.Infof("90th percentile latency: %v", p90)
		logrus.Infof("99th percentile latency: %v", p99)
		logrus.Infof("errors: %d, error rate: %.2f%%, availability: %.2f%%", atomic.LoadUint64(errorCounter), 100.0*errorRate, 100.0*(1.0-errorRate))

		if p99 > 7500*time.Millisecond { // 7.5 seconds but keep the math in integers
			return CapacityExceeded
		}
		if errorRate > 0.001 { // 0.1% error rate is 99.9% available
			//if errorRate > 0.01 { // 1% error rate is 99% available
			return CapacityExceeded
		}
	}

	return nil
}

func (l *LoadGen) generateFixtures() Fixtures {
	var f Fixtures

	for i := 0; i < l.Config.Apps; i++ {
		f = append(f, []*transporttrie.Trie{})

		randomGen := rand.New(rand.NewSource(int64(l.Config.RandSeed + i)))
		p := l.generateProfile(randomGen)
		for j := 0; j < l.Config.Fixtures; j++ {
			f[i] = append(f[i], p)
		}
	}

	return f
}

func (l *LoadGen) startClientThread(appName string, wg *sync.WaitGroup, appFixtures []*transporttrie.Trie, latencyHistogram []uint32, errorCounter *uint64) {
	rc := remote.RemoteConfig{
		UpstreamThreads:        1,
		UpstreamAddress:        l.Config.ServerAddress,
		UpstreamRequestTimeout: 10 * time.Second,
	}

	r, err := remote.New(rc, logrus.New())
	if err != nil {
		panic(err)
	}

	requestsCount := l.Config.Requests

	threadStartTime := time.Now().Truncate(10 * time.Second)
	threadStartTime = threadStartTime.Add(time.Duration(-1*requestsCount) * (10 * time.Second))

	st := threadStartTime

	for i := 0; i < requestsCount; i++ {
		t := appFixtures[i%len(appFixtures)]

		st = st.Add(10 * time.Second)
		et := st.Add(10 * time.Second)
		t0 := time.Now()
		err := r.UploadSync(&upstream.UploadJob{
			Name:            appName + "{}",
			StartTime:       st,
			EndTime:         et,
			SpyName:         "gospy",
			SampleRate:      100,
			Units:           "samples",
			AggregationType: "sum",
			Trie:            t,
		})
		if latencyHistogram != nil {
			milliseconds := uint32(time.Since(t0).Milliseconds())
			if milliseconds >= uint32(len(latencyHistogram)) {
				milliseconds = uint32(len(latencyHistogram) - 1)
			}
			atomic.AddUint32(&latencyHistogram[milliseconds], 1)
		}
		if err != nil {
			l.uploadErrors.Add(1)
			if errorCounter != nil {
				atomic.AddUint64(errorCounter, 1)
			} else {
				time.Sleep(time.Second) // lessens the load and corrupts the result
			}
		} else {
			l.successfulUploads.Add(1)
		}

		err = l.pusher.Push()
		if err != nil {
			logrus.Error(err)
		}
	}

	wg.Done()
}

func (l *LoadGen) generateProfile(randomGen *rand.Rand) *transporttrie.Trie {
	t := transporttrie.New()

	for w := 0; w < l.Config.ProfileWidth; w++ {
		symbol := []byte("root")
		for d := 0; d < 2+l.Rand.Intn(l.Config.ProfileDepth); d++ {
			randomGen.Read(l.SymbolBuf)
			symbol = append(symbol, byte(';'))
			symbol = append(symbol, []byte(hex.EncodeToString(l.SymbolBuf))...)
			if l.Rand.Intn(100) <= 20 {
				t.Insert(symbol, uint64(l.Rand.Intn(100)), true)
			}
		}

		t.Insert(symbol, uint64(l.Rand.Intn(100)), true)
	}
	return t
}

// TODO(eh-am) exponential backoff and whatnot
func waitUntilEndpointReady(url string) error {
	client := http.Client{Timeout: 10 * time.Second}
	retries := 0

	for {
		_, err := client.Get(url)

		// all good?
		if err == nil {
			return nil
		}
		if retries >= MaxReadinessRetries {
			break
		}

		time.Sleep(time.Second)
		retries++
	}

	return fmt.Errorf("maximum retries exceeded ('%d') waiting for server ('%s') to respond", retries, url)
}
