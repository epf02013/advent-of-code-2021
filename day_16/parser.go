package day_16_parser

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToBinary(hexString string) []string {
	parts := []string{}
	for _, character := range hexString {
		i, _ := strconv.ParseInt(string(character), 16, 64)
		bitString := fmt.Sprintf("%04b", i)
		parts = append(parts, strings.Split(bitString,"")...)
	}
	return parts
}

func ParsePacketVersionAndId(bitString []string) (int, int) {
	version, _ := strconv.ParseInt(strings.Join(bitString[0:3], ""), 2, 64)
	typeId, _ := strconv.ParseInt(strings.Join(bitString[3:6], ""), 2, 4)

	return int(version),int(typeId)
}

type PacketInfo struct {
	VersionSums int
	Length      int
	Value int
}

func ParsePacketInfoFromLiteralPacket(bitString []string) PacketInfo {
	literalValueLength := len(bitString) - ((len(bitString) - 6)%5)
	valueBits := []string{}

	for i := 6; i < literalValueLength-4; i+=5 {
		elems := bitString[i+1 : i+5]
		valueBits = append(valueBits, elems...)
		if bitString[i] == "0" {
			break
		}
	}
	numberOfParsedBits := (5 * len(valueBits)) / 4
	numberOfTrailingZeros := numberOfParsedBits % 4
	endPosition := numberOfParsedBits + numberOfTrailingZeros - 1 + 6
	value, _ := strconv.ParseInt(strings.Join(valueBits, ""), 2, 64)
	version, _ := ParsePacketVersionAndId(bitString)
	return PacketInfo{ Value: int(value), Length: endPosition, VersionSums: version }
}

func ParseOperatorSubPacketInfosByLength(bitString []string) ([]PacketInfo, int) {
	lengthOfSubPackets, _ := strconv.ParseInt(strings.Join(bitString[7:22], ""), 2, 64)
	var packetInfos []PacketInfo
	for i := 22; i < 22 + int(lengthOfSubPackets); {
		_, typeId := ParsePacketVersionAndId(bitString[i:i+7])
		var packetInfo PacketInfo
		if typeId == 4 {
			packetInfo = ParsePacketInfoFromLiteralPacket(bitString[i:])
			packetInfos = append(packetInfos, packetInfo)
		}
		i += packetInfo.Length
	}
	return packetInfos, int(lengthOfSubPackets) + 22
}


func ParseOperatorPacketInfo(bitString []string) PacketInfo {
	version, _  := ParsePacketVersionAndId(bitString)
	lengthTypeId := bitString[6]
	if lengthTypeId == "0" {
		packetInfos, length := ParseOperatorSubPacketInfosByLength(bitString)
		versionSumOfSubPackets := sumVersions(packetInfos)
		return PacketInfo{VersionSums: versionSumOfSubPackets + version, Length: length}
	}
	return PacketInfo{VersionSums: version}
}

func sumVersions(packetInfos []PacketInfo) int {
	versionSumOfSubPackets := 0
	for _, info := range packetInfos {
		versionSumOfSubPackets += info.VersionSums
	}
	return versionSumOfSubPackets
}
