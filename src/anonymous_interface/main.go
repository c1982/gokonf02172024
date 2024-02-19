package main

import "encoding/base64"

func main() {
	var data = []byte("hello")

	var encoder interface {
		EncodeToString([]byte) string
	} = base64.StdEncoding

	encoded := encoder.EncodeToString(data)
	print(encoded)
}
