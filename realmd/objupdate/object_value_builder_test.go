package objupdate

import (
	"math"
	"testing"

	"github.com/kangaroux/gomaggus/internal"
	"github.com/mixcode/binarystruct"
	"github.com/stretchr/testify/assert"
)

type valueBlock struct {
	MaskSize byte
	Mask     []uint32
	Values   []uint32
}

func TestObjectGuid(t *testing.T) {
	b := ObjectValueBuilder{buf: &ValueBuffer{}}
	b.Guid(0xDEADBEEF11C0FFEE)
	expected := internal.MustMarshal(
		&valueBlock{
			MaskSize: 1,
			Mask:     []uint32{0x3},
			Values:   []uint32{0x11C0FFEE, 0xDEADBEEF},
		},
		binarystruct.LittleEndian,
	)

	assert.Equal(t, expected, b.buf.Bytes())
}

func TestObjectType(t *testing.T) {
	b := ObjectValueBuilder{buf: &ValueBuffer{}}
	b.Type(ObjectTypePlayer)
	expected := internal.MustMarshal(
		&valueBlock{
			MaskSize: 1,
			Mask:     []uint32{0x4},
			Values:   []uint32{1 << uint32(ObjectTypePlayer)},
		},
		binarystruct.LittleEndian,
	)

	assert.Equal(t, expected, b.buf.Bytes())
}

func TestObjectScaleX(t *testing.T) {
	b := ObjectValueBuilder{buf: &ValueBuffer{}}
	b.ScaleX(123.45)
	expected := internal.MustMarshal(
		&valueBlock{
			MaskSize: 1,
			Mask:     []uint32{0x10},
			Values:   []uint32{math.Float32bits(123.45)},
		},
		binarystruct.LittleEndian,
	)

	assert.Equal(t, expected, b.buf.Bytes())
}
