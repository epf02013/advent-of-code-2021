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

func ParseValueAndEndFromLiteralPacket(bitString []string) (int64, int) {
	literalValueLength := len(bitString) - ((len(bitString) - 6)%5)
	valueBits := []string{}

	for i := 6; i < literalValueLength-4; i+=5 {
		elems := bitString[i+1 : i+5]
		valueBits = append(valueBits, elems...)
		if bitString[0] == "0" {
			break
		}
	}
	numberOfParsedBits := (5 * len(valueBits)) / 4
	numberOfTrailingZeros := numberOfParsedBits % 4
	endPosition := numberOfParsedBits + numberOfTrailingZeros - 1 + 6
	value, _ := strconv.ParseInt(strings.Join(valueBits, ""), 2, 64)
	return value, endPosition
}
