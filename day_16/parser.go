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
