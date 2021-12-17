package parser

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

func TestParseOperatorPacketInfoWhenLengthIdIs0(t *testing.T) {
	bitString := strings.Split("00111000000000000110111101000101001010010001001000000000", "")
	var packetInfo PacketInfo = ParseOperatorPacketInfo(bitString)

	assert.Equal(t, 9, packetInfo.VersionSums)
	assert.Equal(t, len(bitString) - 7, packetInfo.Length)
}

func TestParseOperatorPacketInfoWhenLengthIdIs1(t *testing.T) {
	bitString := strings.Split("11101110000000001101010000001100100000100011000001100000", "")
	var packetInfo PacketInfo = ParseOperatorPacketInfo(bitString)

	assert.Equal(t, 14, packetInfo.VersionSums)
	assert.Equal(t, len(bitString) - 5, packetInfo.Length)
}


func TestGetPacketInfo(t *testing.T) {
	assert.Equal(t, 16, GetPacketInfo("8A004A801A8002F478").VersionSums)
	assert.Equal(t, 12, GetPacketInfo("620080001611562C8802118E34").VersionSums)
	assert.Equal(t, 23, GetPacketInfo("C0015000016115A2E0802F182340").VersionSums)
	assert.Equal(t, 31, GetPacketInfo("A0016C880162017C3686B18A3D4780").VersionSums)
}


