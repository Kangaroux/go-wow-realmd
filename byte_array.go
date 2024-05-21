package main

import "bytes"

type ByteArray struct {
	bigEndian bool
	data      []byte
}

func NewByteArray(data []byte, bigEndian bool) *ByteArray {
	return &ByteArray{bigEndian: bigEndian, data: data}
}

func (ba *ByteArray) Bytes() []byte {
	return ba.data
}

func (ba *ByteArray) Clone() *ByteArray {
	return &ByteArray{
		bigEndian: ba.bigEndian,
		data:      bytes.Clone(ba.data),
	}
}

func (ba *ByteArray) AsBigEndian() *ByteArray {
	if !ba.bigEndian {
		ba.swapEndian()
	}
	return ba
}

func (ba *ByteArray) AsLittleEndian() *ByteArray {
	if ba.bigEndian {
		ba.swapEndian()
	}
	return ba
}

func (ba *ByteArray) swapEndian() {
	ba.bigEndian = !ba.bigEndian
	reverseBytesNoCopy(ba.data)
}

// reverses the byte array in place
func reverseBytesNoCopy(data []byte) {
	n := len(data)
	for i := 0; i < n/2; i++ {
		data[i], data[n-i-1] = data[n-i-1], data[i]
	}
}
