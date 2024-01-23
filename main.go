package main

import "fmt"

// I've never made a base64 encoder before or written something in go :)

const b64str string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const b64pad string = "="

func b64encode(bytes []byte) string {
	i := 0
	var ret string = ""
	var byte1 byte
	var byte2 byte
	var byte3 byte
	for {
		if i >= len(bytes) {
			break
		}
		byte1 = bytes[i]
		i++

		if i >= len(bytes) {
			byte2 = 0
		} else {
			byte2 = bytes[i]
			i++
		}

		if i >= len(bytes) {
			byte3 = 0
		} else {
			byte3 = bytes[i]
			i++
		}

		// Byte 1 is guaranteed to exist
		ret += string(b64str[byte1>>2])
		ret += string(b64str[(byte1&0b00000011<<4)|byte2>>4])

		// Byte 2 is not
		if byte2 == 0 {
			ret += string(b64pad)
		} else {
			ret += string(b64str[(byte2&0b00001111<<2)|((byte3&0b11000000)>>6)])
		}

		// Byte 3 is not either
		if byte3 == 0 {
			ret += string(b64pad)
		} else {
			ret += string(b64str[byte3&0b00111111])
		}
	}

	return ret
}

func b64encode_str(str string) string {
	return b64encode([]byte(str))
}

func main() {
	data := "Hello, World!"
	fmt.Println(data)
	fmt.Println(b64encode_str(data))
}
