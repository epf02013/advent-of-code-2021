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
	packetInfo := ParsePacketInfoFromLiteralPacket(bitString)

	assert.Equal(t, 2021,packetInfo.Value)
	assert.Equal(t, len(bitString) -1,packetInfo.Length)
}

func TestParseOperatorPacketInfo(t *testing.T) {
	bitString := strings.Split("00111000000000000110111101000101001010010001001000000000", "")
	var packetInfo PacketInfo = ParseOperatorPacketInfo(bitString)

	assert.Equal(t, 9, packetInfo.VersionSums)
	assert.Equal(t, len(bitString) - 7, packetInfo.Length)
}
