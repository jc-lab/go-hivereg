package hiveutil

import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

func stringToUtf16LE(buffer *bytes.Buffer, s string) error {
	utf16Encoded := utf16.Encode([]rune(s))
	return binary.Write(buffer, binary.LittleEndian, utf16Encoded)
}

func StringToUtf16LE(s string) ([]byte, error) {
	var buffer bytes.Buffer
	err := stringToUtf16LE(&buffer, s)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func StringsToMultiUtf16LE(list []string) ([]byte, error) {
	var buffer bytes.Buffer
	for _, s := range list {
		err := stringToUtf16LE(&buffer, s)
		if err != nil {
			return nil, err
		}
		buffer.WriteByte(0)
		buffer.WriteByte(0)
	}
	buffer.WriteByte(0)
	buffer.WriteByte(0)
	return buffer.Bytes(), nil
}
