// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polypb/metadata.proto

/*
	Package polypb is a generated protocol buffer package.

	It is generated from these files:
		polypb/metadata.proto

	It has these top-level messages:
		DiskMeta
		NodeMeta
		BackupMeta
*/
package polypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StorageType
type StorageType int32

const (
	StorageType_LOCAL_DISK StorageType = 0
	StorageType_LOCAL_MEM  StorageType = 1
)

var StorageType_name = map[int32]string{
	0: "LOCAL_DISK",
	1: "LOCAL_MEM",
}
var StorageType_value = map[string]int32{
	"LOCAL_DISK": 0,
	"LOCAL_MEM":  1,
}

func (x StorageType) String() string {
	return proto.EnumName(StorageType_name, int32(x))
}
func (StorageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{0} }

// BackupType
type BackupType int32

const (
	BackupType_FULL BackupType = 0
	BackupType_INC  BackupType = 1
)

var BackupType_name = map[int32]string{
	0: "FULL",
	1: "INC",
}
var BackupType_value = map[string]int32{
	"FULL": 0,
	"INC":  1,
}

func (x BackupType) String() string {
	return proto.EnumName(BackupType_name, int32(x))
}
func (BackupType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{1} }

// DiskMeta is a metadata about a disk.
type DiskMeta struct {
	Total uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Avail uint64 `protobuf:"varint,2,opt,name=avail,proto3" json:"avail,omitempty"`
}

func (m *DiskMeta) Reset()                    { *m = DiskMeta{} }
func (m *DiskMeta) String() string            { return proto.CompactTextString(m) }
func (*DiskMeta) ProtoMessage()               {}
func (*DiskMeta) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{0} }

func (m *DiskMeta) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *DiskMeta) GetAvail() uint64 {
	if m != nil {
		return m.Avail
	}
	return 0
}

// NodeMeta is a metadata about a node.
type NodeMeta struct {
	Disk     *DiskMeta `protobuf:"bytes,1,opt,name=disk" json:"disk,omitempty"`
	Addr     string    `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	StoreDir string    `protobuf:"bytes,3,opt,name=store_dir,json=storeDir,proto3" json:"store_dir,omitempty"`
	NodeId   NodeID    `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3,casttype=NodeID" json:"node_id,omitempty"`
}

func (m *NodeMeta) Reset()                    { *m = NodeMeta{} }
func (m *NodeMeta) String() string            { return proto.CompactTextString(m) }
func (*NodeMeta) ProtoMessage()               {}
func (*NodeMeta) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{1} }

func (m *NodeMeta) GetDisk() *DiskMeta {
	if m != nil {
		return m.Disk
	}
	return nil
}

func (m *NodeMeta) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *NodeMeta) GetStoreDir() string {
	if m != nil {
		return m.StoreDir
	}
	return ""
}

func (m *NodeMeta) GetNodeId() NodeID {
	if m != nil {
		return m.NodeId
	}
	return nil
}

// BackupMeta is a metadata about a backup file.
type BackupMeta struct {
	StoredTime    *time.Time  `protobuf:"bytes,1,opt,name=stored_time,json=storedTime,stdtime" json:"stored_time,omitempty"`
	NodeId        NodeID      `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3,casttype=NodeID" json:"node_id,omitempty"`
	Host          string      `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	StorageType   StorageType `protobuf:"varint,4,opt,name=storage_type,json=storageType,proto3,enum=polypb.StorageType" json:"storage_type,omitempty"`
	EndTime       *time.Time  `protobuf:"bytes,5,opt,name=end_time,json=endTime,stdtime" json:"end_time,omitempty"`
	FileSize      int64       `protobuf:"varint,6,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	BackupType    BackupType  `protobuf:"varint,7,opt,name=backup_type,json=backupType,proto3,enum=polypb.BackupType" json:"backup_type,omitempty"`
	Db            DatabaseID  `protobuf:"bytes,8,opt,name=db,proto3,casttype=DatabaseID" json:"db,omitempty"`
	ToLsn         string      `protobuf:"bytes,9,opt,name=to_lsn,json=toLsn,proto3" json:"to_lsn,omitempty"`
	Key           Key         `protobuf:"bytes,10,opt,name=key,proto3,casttype=Key" json:"key,omitempty"`
	BaseTimePoint TimePoint   `protobuf:"bytes,11,opt,name=base_time_point,json=baseTimePoint,proto3,casttype=TimePoint" json:"base_time_point,omitempty"`
}

func (m *BackupMeta) Reset()                    { *m = BackupMeta{} }
func (m *BackupMeta) String() string            { return proto.CompactTextString(m) }
func (*BackupMeta) ProtoMessage()               {}
func (*BackupMeta) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{2} }

func (m *BackupMeta) GetStoredTime() *time.Time {
	if m != nil {
		return m.StoredTime
	}
	return nil
}

func (m *BackupMeta) GetNodeId() NodeID {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *BackupMeta) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *BackupMeta) GetStorageType() StorageType {
	if m != nil {
		return m.StorageType
	}
	return StorageType_LOCAL_DISK
}

func (m *BackupMeta) GetEndTime() *time.Time {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *BackupMeta) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *BackupMeta) GetBackupType() BackupType {
	if m != nil {
		return m.BackupType
	}
	return BackupType_FULL
}

func (m *BackupMeta) GetDb() DatabaseID {
	if m != nil {
		return m.Db
	}
	return nil
}

func (m *BackupMeta) GetToLsn() string {
	if m != nil {
		return m.ToLsn
	}
	return ""
}

func (m *BackupMeta) GetKey() Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *BackupMeta) GetBaseTimePoint() TimePoint {
	if m != nil {
		return m.BaseTimePoint
	}
	return nil
}

func init() {
	proto.RegisterType((*DiskMeta)(nil), "polypb.DiskMeta")
	proto.RegisterType((*NodeMeta)(nil), "polypb.NodeMeta")
	proto.RegisterType((*BackupMeta)(nil), "polypb.BackupMeta")
	proto.RegisterEnum("polypb.StorageType", StorageType_name, StorageType_value)
	proto.RegisterEnum("polypb.BackupType", BackupType_name, BackupType_value)
}
func (m *DiskMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiskMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.Total))
	}
	if m.Avail != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.Avail))
	}
	return i, nil
}

func (m *NodeMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Disk != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.Disk.Size()))
		n1, err := m.Disk.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Addr) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if len(m.StoreDir) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.StoreDir)))
		i += copy(dAtA[i:], m.StoreDir)
	}
	if len(m.NodeId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.NodeId)))
		i += copy(dAtA[i:], m.NodeId)
	}
	return i, nil
}

func (m *BackupMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StoredTime != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.StoredTime)))
		n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StoredTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.NodeId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.NodeId)))
		i += copy(dAtA[i:], m.NodeId)
	}
	if len(m.Host) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if m.StorageType != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.StorageType))
	}
	if m.EndTime != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime)))
		n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EndTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.FileSize != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.FileSize))
	}
	if m.BackupType != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.BackupType))
	}
	if len(m.Db) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Db)))
		i += copy(dAtA[i:], m.Db)
	}
	if len(m.ToLsn) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.ToLsn)))
		i += copy(dAtA[i:], m.ToLsn)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.BaseTimePoint) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.BaseTimePoint)))
		i += copy(dAtA[i:], m.BaseTimePoint)
	}
	return i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DiskMeta) Size() (n int) {
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovMetadata(uint64(m.Total))
	}
	if m.Avail != 0 {
		n += 1 + sovMetadata(uint64(m.Avail))
	}
	return n
}

func (m *NodeMeta) Size() (n int) {
	var l int
	_ = l
	if m.Disk != nil {
		l = m.Disk.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.StoreDir)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func (m *BackupMeta) Size() (n int) {
	var l int
	_ = l
	if m.StoredTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StoredTime)
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.StorageType != 0 {
		n += 1 + sovMetadata(uint64(m.StorageType))
	}
	if m.EndTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime)
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.FileSize != 0 {
		n += 1 + sovMetadata(uint64(m.FileSize))
	}
	if m.BackupType != 0 {
		n += 1 + sovMetadata(uint64(m.BackupType))
	}
	l = len(m.Db)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.ToLsn)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.BaseTimePoint)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DiskMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DiskMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiskMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avail", wireType)
			}
			m.Avail = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Avail |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Disk == nil {
				m.Disk = &DiskMeta{}
			}
			if err := m.Disk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoreDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = append(m.NodeId[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeId == nil {
				m.NodeId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StoredTime == nil {
				m.StoredTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StoredTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = append(m.NodeId[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeId == nil {
				m.NodeId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageType", wireType)
			}
			m.StorageType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StorageType |= (StorageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSize", wireType)
			}
			m.FileSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BackupType", wireType)
			}
			m.BackupType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BackupType |= (BackupType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Db", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Db = append(m.Db[:0], dAtA[iNdEx:postIndex]...)
			if m.Db == nil {
				m.Db = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToLsn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToLsn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseTimePoint", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseTimePoint = append(m.BaseTimePoint[:0], dAtA[iNdEx:postIndex]...)
			if m.BaseTimePoint == nil {
				m.BaseTimePoint = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetadata
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("polypb/metadata.proto", fileDescriptorMetadata) }

var fileDescriptorMetadata = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0xc5, 0x89, 0x71, 0xec, 0x31, 0x50, 0x6b, 0x5b, 0x24, 0x17, 0xa4, 0x18, 0xd1, 0x1e, 0x10,
	0x6a, 0x1d, 0x09, 0x54, 0x2e, 0x3d, 0x11, 0xd2, 0x4a, 0x11, 0x81, 0x56, 0x86, 0x9e, 0xad, 0x35,
	0xbb, 0x98, 0x55, 0x1c, 0xaf, 0x15, 0x2f, 0x95, 0xc2, 0x07, 0xf4, 0xdc, 0x63, 0x3f, 0xa9, 0xc7,
	0x1e, 0x7b, 0x4a, 0x2b, 0xfa, 0x17, 0x39, 0x55, 0x3b, 0x8b, 0x81, 0x4b, 0xa5, 0xde, 0xe6, 0xbd,
	0x99, 0xd9, 0x79, 0xef, 0x69, 0x61, 0xbd, 0x92, 0xc5, 0xac, 0xca, 0x7a, 0x13, 0xae, 0x28, 0xa3,
	0x8a, 0xc6, 0xd5, 0x54, 0x2a, 0x49, 0x1c, 0x43, 0x6f, 0x44, 0xb9, 0x94, 0x79, 0xc1, 0x7b, 0xc8,
	0x66, 0xd7, 0x97, 0x3d, 0x25, 0x26, 0xbc, 0x56, 0x74, 0x52, 0x99, 0xc1, 0x8d, 0xd7, 0xb9, 0x50,
	0x57, 0xd7, 0x59, 0x7c, 0x21, 0x27, 0xbd, 0x5c, 0xe6, 0xf2, 0x61, 0x52, 0x23, 0x04, 0x58, 0x99,
	0xf1, 0xed, 0x03, 0x70, 0x07, 0xa2, 0x1e, 0x9f, 0x70, 0x45, 0xc9, 0x33, 0x58, 0x56, 0x52, 0xd1,
	0x22, 0xb4, 0xb6, 0xac, 0x1d, 0x3b, 0x31, 0x40, 0xb3, 0xf4, 0x33, 0x15, 0x45, 0xd8, 0x32, 0x2c,
	0x82, 0xed, 0x2f, 0x16, 0xb8, 0xa7, 0x92, 0x71, 0x5c, 0x7c, 0x09, 0x36, 0x13, 0xf5, 0x18, 0xf7,
	0xfc, 0xbd, 0x20, 0x36, 0x5a, 0xe3, 0xe6, 0xe1, 0x04, 0xbb, 0x84, 0x80, 0x4d, 0x19, 0x9b, 0xe2,
	0x3b, 0x5e, 0x82, 0x35, 0xd9, 0x04, 0xaf, 0x56, 0x72, 0xca, 0x53, 0x26, 0xa6, 0x61, 0x1b, 0x1b,
	0x2e, 0x12, 0x03, 0x31, 0x25, 0x2f, 0xa0, 0x53, 0x4a, 0xc6, 0x53, 0xc1, 0x42, 0x7b, 0xcb, 0xda,
	0x59, 0xe9, 0xc3, 0x62, 0x1e, 0x39, 0xfa, 0xea, 0x70, 0x90, 0x38, 0xba, 0x35, 0x64, 0xdb, 0x3f,
	0xdb, 0x00, 0x7d, 0x7a, 0x31, 0xbe, 0xae, 0x50, 0xca, 0x21, 0xf8, 0xb8, 0xcf, 0x52, 0x1d, 0xcc,
	0x9d, 0xa2, 0x8d, 0xd8, 0xa4, 0x16, 0x37, 0x59, 0xc4, 0xe7, 0x4d, 0x6a, 0x7d, 0xfb, 0xeb, 0xaf,
	0xc8, 0x4a, 0xc0, 0x2c, 0x69, 0xfa, 0xf1, 0xd9, 0xd6, 0xbf, 0xce, 0x6a, 0x33, 0x57, 0xb2, 0x56,
	0x77, 0x9a, 0xb1, 0x26, 0x07, 0xb0, 0xa2, 0x9f, 0xa1, 0x39, 0x4f, 0xd5, 0xac, 0xe2, 0x28, 0x7a,
	0x6d, 0xef, 0x69, 0x13, 0xc7, 0x99, 0xe9, 0x9d, 0xcf, 0x2a, 0x9e, 0xf8, 0xf5, 0x03, 0x20, 0x6f,
	0xc1, 0xe5, 0xe5, 0x9d, 0xe0, 0xe5, 0xff, 0x14, 0xdc, 0xe1, 0xa5, 0x51, 0xbb, 0x09, 0xde, 0xa5,
	0x28, 0x78, 0x5a, 0x8b, 0x1b, 0x1e, 0x3a, 0x5b, 0xd6, 0x4e, 0x3b, 0x71, 0x35, 0x71, 0x26, 0x6e,
	0x38, 0xd9, 0x07, 0x3f, 0xc3, 0x6c, 0x8c, 0xa0, 0x0e, 0x0a, 0x22, 0x8d, 0x20, 0x13, 0x1b, 0xea,
	0x81, 0xec, 0xbe, 0x26, 0x5d, 0x68, 0xb1, 0x2c, 0x74, 0xd1, 0xfa, 0xda, 0x62, 0x1e, 0xc1, 0x80,
	0x2a, 0x9a, 0xd1, 0x5a, 0xdb, 0x6f, 0xb1, 0x8c, 0xac, 0x83, 0xa3, 0x64, 0x5a, 0xd4, 0x65, 0xe8,
	0xa1, 0xf9, 0x65, 0x25, 0x47, 0x75, 0x49, 0x9e, 0x43, 0x7b, 0xcc, 0x67, 0x21, 0xe0, 0x5e, 0x67,
	0x31, 0x8f, 0xda, 0xc7, 0x7c, 0x96, 0x68, 0x8e, 0xbc, 0x81, 0x27, 0x7a, 0x1f, 0x1d, 0xa6, 0x95,
	0x14, 0xa5, 0x0a, 0x7d, 0x1c, 0x5b, 0x5d, 0xcc, 0x23, 0x4f, 0xdb, 0xf8, 0xa8, 0xc9, 0x64, 0x55,
	0x4f, 0xdd, 0xc3, 0xdd, 0x57, 0xe0, 0x3f, 0xca, 0x8c, 0xac, 0x01, 0x8c, 0x3e, 0x1c, 0x1d, 0x8e,
	0xd2, 0xc1, 0xf0, 0xec, 0x38, 0x58, 0x22, 0xab, 0xe0, 0x19, 0x7c, 0xf2, 0xee, 0x24, 0xb0, 0x76,
	0xa3, 0xe6, 0x1f, 0xe0, 0xb0, 0x0b, 0xf6, 0xfb, 0x4f, 0xa3, 0x51, 0xb0, 0x44, 0x3a, 0xd0, 0x1e,
	0x9e, 0x1e, 0x05, 0x56, 0x3f, 0xf8, 0x7e, 0xdb, 0xb5, 0x7e, 0xdc, 0x76, 0xad, 0xdf, 0xb7, 0x5d,
	0xeb, 0xdb, 0x9f, 0xee, 0x52, 0xe6, 0x60, 0xbc, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7,
	0xbd, 0x63, 0x4f, 0x74, 0x03, 0x00, 0x00,
}
