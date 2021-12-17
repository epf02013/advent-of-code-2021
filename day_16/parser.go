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
	numberOfParsedBits := 5 * (len(valueBits) / 4)
	//numberOfTrailingZeros := numberOfParsedBits % 4
	endPosition := numberOfParsedBits + 6
	value, _ := strconv.ParseInt(strings.Join(valueBits, ""), 2, 64)
	version, _ := ParsePacketVersionAndId(bitString)
	return PacketInfo{ Value: int(value), Length: endPosition, VersionSums: version }
}

func ParseOperatorSubPacketInfosByLength(bitString []string) ([]PacketInfo, int) {
	lengthOfSubPackets, _ := strconv.ParseInt(strings.Join(bitString[7:22], ""), 2, 64)
	var packetInfos []PacketInfo
	if len(bitString) < 22 {
		fmt.Println(bitString)
	}
	for i := 22; i < 22 + int(lengthOfSubPackets) && i < len(bitString); {
		_, typeId := ParsePacketVersionAndId(bitString[i:i+7])
		var packetInfo PacketInfo
		if typeId == 4 {
			packetInfo = ParsePacketInfoFromLiteralPacket(bitString[i:])
			packetInfos = append(packetInfos, packetInfo)
		} else {
			packetInfo = ParseOperatorPacketInfo(bitString[i:])
			packetInfos = append(packetInfos, packetInfo)
		}
		i += packetInfo.Length
	}
	return packetInfos, int(lengthOfSubPackets) + 22
}

func ParseOperatorSubPacketByNumberOfPackets(bitString []string) ([]PacketInfo, int) {
	numberOfSubPackets, _ := strconv.ParseInt(strings.Join(bitString[7:18], ""), 2, 64)
	var packetInfos []PacketInfo
	numberOfSubPacketsParsed := 0
	i := 18
	for ; i < len(bitString) && numberOfSubPacketsParsed < int(numberOfSubPackets); {
		_, typeId := ParsePacketVersionAndId(bitString[i:i+7])
		var packetInfo PacketInfo
		if typeId == 4 {
			packetInfo = ParsePacketInfoFromLiteralPacket(bitString[i:])
			packetInfos = append(packetInfos, packetInfo)
		} else {
			packetInfo = ParseOperatorPacketInfo(bitString[i:])
			packetInfos = append(packetInfos, packetInfo)
		}
		numberOfSubPacketsParsed += 1
		i += packetInfo.Length
	}
	return packetInfos, i
}

func isAllZeros(bitString []string) bool{
	for _, s := range bitString {
		if s != "0" {
			return false
		}
	}
	return true
}
func ParseOperatorPacketInfo(bitString []string) PacketInfo {
	version, _  := ParsePacketVersionAndId(bitString)
	lengthTypeId := bitString[6]
	if lengthTypeId == "0" {
		packetInfos, length := ParseOperatorSubPacketInfosByLength(bitString)
		versionSumOfSubPackets := sumVersions(packetInfos)
		return PacketInfo{VersionSums: versionSumOfSubPackets + version, Length: length}
	}
	packetInfos, length := ParseOperatorSubPacketByNumberOfPackets(bitString)
	versionSumOfSubPackets := sumVersions(packetInfos)
	return PacketInfo{VersionSums: versionSumOfSubPackets + version, Length: length}
}

func sumVersions(packetInfos []PacketInfo) int {
	versionSumOfSubPackets := 0
	for _, info := range packetInfos {
		versionSumOfSubPackets += info.VersionSums
	}
	return versionSumOfSubPackets
}

func GetPacketInfo(hexString string) PacketInfo {
	bitString := HexToBinary(hexString)
	var packetInfo = ParseOperatorPacketInfo(bitString)
	return packetInfo
}
