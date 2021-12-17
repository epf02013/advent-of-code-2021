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

func TestParsePacketVersionAndId(t *testing.T) {
	bitString := strings.Split("110100101111111000101000", "")
	var version, typeId int = ParsePacketVersionAndId(bitString)

	assert.Equal(t, 6, version)
	assert.Equal(t, 4, typeId)
}

func TestParseValueAndEndFromLiteralPacket(t *testing.T) {
	bitString := strings.Split("110100101111111000101000", "")
	value,endPosition := ParseValueAndEndFromLiteralPacket(bitString)

	assert.Equal(t, 2021,int(value))
	assert.Equal(t, len(bitString) -1,endPosition)
}
