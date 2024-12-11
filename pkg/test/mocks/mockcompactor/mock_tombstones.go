// Code generated by mockery. DO NOT EDIT.

package mockcompactor

import (
	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"
	iter "github.com/grafana/pyroscope/pkg/iter"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockTombstones is an autogenerated mock type for the Tombstones type
type MockTombstones struct {
	mock.Mock
}

type MockTombstones_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTombstones) EXPECT() *MockTombstones_Expecter {
	return &MockTombstones_Expecter{mock: &_m.Mock}
}

// ListTombstones provides a mock function with given fields: before
func (_m *MockTombstones) ListTombstones(before time.Time) iter.Iterator[*metastorev1.Tombstones] {
	ret := _m.Called(before)

	if len(ret) == 0 {
		panic("no return value specified for ListTombstones")
	}

	var r0 iter.Iterator[*metastorev1.Tombstones]
	if rf, ok := ret.Get(0).(func(time.Time) iter.Iterator[*metastorev1.Tombstones]); ok {
		r0 = rf(before)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iter.Iterator[*metastorev1.Tombstones])
		}
	}

	return r0
}

// MockTombstones_ListTombstones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTombstones'
type MockTombstones_ListTombstones_Call struct {
	*mock.Call
}

// ListTombstones is a helper method to define mock.On call
//   - before time.Time
func (_e *MockTombstones_Expecter) ListTombstones(before interface{}) *MockTombstones_ListTombstones_Call {
	return &MockTombstones_ListTombstones_Call{Call: _e.mock.On("ListTombstones", before)}
}

func (_c *MockTombstones_ListTombstones_Call) Run(run func(before time.Time)) *MockTombstones_ListTombstones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time))
	})
	return _c
}

func (_c *MockTombstones_ListTombstones_Call) Return(_a0 iter.Iterator[*metastorev1.Tombstones]) *MockTombstones_ListTombstones_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTombstones_ListTombstones_Call) RunAndReturn(run func(time.Time) iter.Iterator[*metastorev1.Tombstones]) *MockTombstones_ListTombstones_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTombstones creates a new instance of MockTombstones. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTombstones(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTombstones {
	mock := &MockTombstones{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}