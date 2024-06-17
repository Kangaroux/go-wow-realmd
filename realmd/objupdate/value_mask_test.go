package objupdate_test

import (
	"testing"

	o "github.com/kangaroux/gomaggus/realmd/objupdate"
	"github.com/stretchr/testify/assert"
)

func TestValueMask(t *testing.T) {
	cases := []struct {
		masks    []o.FieldMask
		expected []uint32
	}{
		{
			masks:    []o.FieldMask{},
			expected: []uint32{0},
		},
		{
			masks:    []o.FieldMask{{Size: 1, Offset: 0}},
			expected: []uint32{1, 0x1},
		},
		{
			masks:    []o.FieldMask{{Size: 32, Offset: 0}},
			expected: []uint32{1, 0xFFFFFFFF},
		},
		{
			masks:    []o.FieldMask{{Size: 33, Offset: 0}},
			expected: []uint32{2, 0xFFFFFFFF, 0x1},
		},
		{
			masks:    []o.FieldMask{{Size: 1, Offset: 32}},
			expected: []uint32{2, 0x0, 0x1},
		},
		{
			masks:    []o.FieldMask{{Size: 2, Offset: 31}},
			expected: []uint32{2, 0x80000000, 0x1},
		},
		{
			masks:    []o.FieldMask{{Size: 1, Offset: 0}, {Size: 2, Offset: 2}},
			expected: []uint32{1, 0xD},
		},
		{
			// Largest field offset
			masks: []o.FieldMask{o.FieldMaskPlayerPetSpellPower},
			expected: []uint32{
				42,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x2000,
			},
		},
	}

	for _, c := range cases {
		vm := o.ValueMask{}

		for _, fm := range c.masks {
			vm.SetFieldMask(fm)
		}

		assert.Equal(t, c.expected, vm.Mask())
	}
}