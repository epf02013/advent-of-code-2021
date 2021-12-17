package day_16_parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHexToBinary(t *testing.T) {
	hexString := "D2FE28"
	bitString := HexToBinary(hexString)
	assert.Equal(t, strings.Split("110100101111111000101000", ""),bitString)
}
