// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: metastore/v1/types.proto

package metastorev1

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *BlockList) CloneVT() *BlockList {
	if m == nil {
		return (*BlockList)(nil)
	}
	r := new(BlockList)
	r.Tenant = m.Tenant
	r.Shard = m.Shard
	if rhs := m.Blocks; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.Blocks = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *BlockList) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *BlockMeta) CloneVT() *BlockMeta {
	if m == nil {
		return (*BlockMeta)(nil)
	}
	r := new(BlockMeta)
	r.FormatVersion = m.FormatVersion
	r.Id = m.Id
	r.Tenant = m.Tenant
	r.Shard = m.Shard
	r.CompactionLevel = m.CompactionLevel
	r.MinTime = m.MinTime
	r.MaxTime = m.MaxTime
	r.CreatedBy = m.CreatedBy
	r.Size = m.Size
	if rhs := m.Datasets; rhs != nil {
		tmpContainer := make([]*Dataset, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Datasets = tmpContainer
	}
	if rhs := m.StringTable; rhs != nil {
		tmpContainer := make([]string, len(rhs))
		copy(tmpContainer, rhs)
		r.StringTable = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *BlockMeta) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Dataset) CloneVT() *Dataset {
	if m == nil {
		return (*Dataset)(nil)
	}
	r := new(Dataset)
	r.Tenant = m.Tenant
	r.Name = m.Name
	r.MinTime = m.MinTime
	r.MaxTime = m.MaxTime
	r.Size = m.Size
	if rhs := m.TableOfContents; rhs != nil {
		tmpContainer := make([]uint64, len(rhs))
		copy(tmpContainer, rhs)
		r.TableOfContents = tmpContainer
	}
	if rhs := m.ProfileTypes; rhs != nil {
		tmpContainer := make([]int32, len(rhs))
		copy(tmpContainer, rhs)
		r.ProfileTypes = tmpContainer
	}
	if rhs := m.Labels; rhs != nil {
		tmpContainer := make([]int32, len(rhs))
		copy(tmpContainer, rhs)
		r.Labels = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Dataset) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *BlockList) EqualVT(that *BlockList) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Tenant != that.Tenant {
		return false
	}
	if this.Shard != that.Shard {
		return false
	}
	if len(this.Blocks) != len(that.Blocks) {
		return false
	}
	for i, vx := range this.Blocks {
		vy := that.Blocks[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *BlockList) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*BlockList)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *BlockMeta) EqualVT(that *BlockMeta) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.FormatVersion != that.FormatVersion {
		return false
	}
	if this.Id != that.Id {
		return false
	}
	if this.Tenant != that.Tenant {
		return false
	}
	if this.Shard != that.Shard {
		return false
	}
	if this.CompactionLevel != that.CompactionLevel {
		return false
	}
	if this.MinTime != that.MinTime {
		return false
	}
	if this.MaxTime != that.MaxTime {
		return false
	}
	if this.CreatedBy != that.CreatedBy {
		return false
	}
	if this.Size != that.Size {
		return false
	}
	if len(this.Datasets) != len(that.Datasets) {
		return false
	}
	for i, vx := range this.Datasets {
		vy := that.Datasets[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &Dataset{}
			}
			if q == nil {
				q = &Dataset{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if len(this.StringTable) != len(that.StringTable) {
		return false
	}
	for i, vx := range this.StringTable {
		vy := that.StringTable[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *BlockMeta) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*BlockMeta)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Dataset) EqualVT(that *Dataset) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Tenant != that.Tenant {
		return false
	}
	if this.Name != that.Name {
		return false
	}
	if this.MinTime != that.MinTime {
		return false
	}
	if this.MaxTime != that.MaxTime {
		return false
	}
	if len(this.TableOfContents) != len(that.TableOfContents) {
		return false
	}
	for i, vx := range this.TableOfContents {
		vy := that.TableOfContents[i]
		if vx != vy {
			return false
		}
	}
	if this.Size != that.Size {
		return false
	}
	if len(this.ProfileTypes) != len(that.ProfileTypes) {
		return false
	}
	for i, vx := range this.ProfileTypes {
		vy := that.ProfileTypes[i]
		if vx != vy {
			return false
		}
	}
	if len(this.Labels) != len(that.Labels) {
		return false
	}
	for i, vx := range this.Labels {
		vy := that.Labels[i]
		if vx != vy {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Dataset) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Dataset)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *BlockList) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockList) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BlockList) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Blocks[iNdEx])
			copy(dAtA[i:], m.Blocks[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Blocks[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Shard != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Shard))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockMeta) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMeta) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BlockMeta) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.StringTable) > 0 {
		for iNdEx := len(m.StringTable) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StringTable[iNdEx])
			copy(dAtA[i:], m.StringTable[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.StringTable[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Datasets) > 0 {
		for iNdEx := len(m.Datasets) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Datasets[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Size != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Size))
		i--
		dAtA[i] = 0x48
	}
	if m.CreatedBy != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.CreatedBy))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxTime != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MaxTime))
		i--
		dAtA[i] = 0x38
	}
	if m.MinTime != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MinTime))
		i--
		dAtA[i] = 0x30
	}
	if m.CompactionLevel != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.CompactionLevel))
		i--
		dAtA[i] = 0x28
	}
	if m.Shard != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Shard))
		i--
		dAtA[i] = 0x20
	}
	if m.Tenant != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Tenant))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if m.FormatVersion != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.FormatVersion))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Dataset) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dataset) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Dataset) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Labels) > 0 {
		var pksize2 int
		for _, num := range m.Labels {
			pksize2 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize2
		j1 := i
		for _, num1 := range m.Labels {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA[j1] = uint8(num)
			j1++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize2))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ProfileTypes) > 0 {
		var pksize4 int
		for _, num := range m.ProfileTypes {
			pksize4 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize4
		j3 := i
		for _, num1 := range m.ProfileTypes {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA[j3] = uint8(num)
			j3++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize4))
		i--
		dAtA[i] = 0x3a
	}
	if m.Size != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Size))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TableOfContents) > 0 {
		var pksize6 int
		for _, num := range m.TableOfContents {
			pksize6 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize6
		j5 := i
		for _, num := range m.TableOfContents {
			for num >= 1<<7 {
				dAtA[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA[j5] = uint8(num)
			j5++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize6))
		i--
		dAtA[i] = 0x2a
	}
	if m.MaxTime != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MaxTime))
		i--
		dAtA[i] = 0x20
	}
	if m.MinTime != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MinTime))
		i--
		dAtA[i] = 0x18
	}
	if m.Name != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Name))
		i--
		dAtA[i] = 0x10
	}
	if m.Tenant != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Tenant))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockList) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Shard != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Shard))
	}
	if len(m.Blocks) > 0 {
		for _, s := range m.Blocks {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *BlockMeta) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FormatVersion != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.FormatVersion))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.Tenant != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Tenant))
	}
	if m.Shard != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Shard))
	}
	if m.CompactionLevel != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.CompactionLevel))
	}
	if m.MinTime != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MinTime))
	}
	if m.MaxTime != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MaxTime))
	}
	if m.CreatedBy != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.CreatedBy))
	}
	if m.Size != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Size))
	}
	if len(m.Datasets) > 0 {
		for _, e := range m.Datasets {
			l = e.SizeVT()
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if len(m.StringTable) > 0 {
		for _, s := range m.StringTable {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Dataset) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tenant != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Tenant))
	}
	if m.Name != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Name))
	}
	if m.MinTime != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MinTime))
	}
	if m.MaxTime != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MaxTime))
	}
	if len(m.TableOfContents) > 0 {
		l = 0
		for _, e := range m.TableOfContents {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	if m.Size != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Size))
	}
	if len(m.ProfileTypes) > 0 {
		l = 0
		for _, e := range m.ProfileTypes {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	if len(m.Labels) > 0 {
		l = 0
		for _, e := range m.Labels {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	n += len(m.unknownFields)
	return n
}

func (m *BlockList) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shard", wireType)
			}
			m.Shard = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shard |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockMeta) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FormatVersion", wireType)
			}
			m.FormatVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FormatVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			m.Tenant = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tenant |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shard", wireType)
			}
			m.Shard = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shard |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactionLevel", wireType)
			}
			m.CompactionLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompactionLevel |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			m.MinTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			m.MaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedBy", wireType)
			}
			m.CreatedBy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedBy |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size", wireType)
			}
			m.Size = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datasets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Datasets = append(m.Datasets, &Dataset{})
			if err := m.Datasets[len(m.Datasets)-1].UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringTable", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringTable = append(m.StringTable, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Dataset) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Dataset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dataset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			m.Tenant = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tenant |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			m.Name = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Name |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			m.MinTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			m.MaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.TableOfContents = append(m.TableOfContents, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return protohelpers.ErrInvalidLength
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return protohelpers.ErrInvalidLength
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.TableOfContents) == 0 {
					m.TableOfContents = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.TableOfContents = append(m.TableOfContents, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field TableOfContents", wireType)
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size", wireType)
			}
			m.Size = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ProfileTypes = append(m.ProfileTypes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return protohelpers.ErrInvalidLength
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return protohelpers.ErrInvalidLength
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ProfileTypes) == 0 {
					m.ProfileTypes = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ProfileTypes = append(m.ProfileTypes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileTypes", wireType)
			}
		case 8:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Labels = append(m.Labels, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return protohelpers.ErrInvalidLength
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return protohelpers.ErrInvalidLength
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Labels) == 0 {
					m.Labels = make([]int32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Labels = append(m.Labels, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}