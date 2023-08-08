package utils

import "fmt"

func ByteString(bytes []byte) string {
	var byteStr string
	for _, b := range bytes {
		byteStr += fmt.Sprintf("%02x ", b)
	}
	return byteStr
}

func Must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
