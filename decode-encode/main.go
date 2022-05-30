package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "jonh wick"

	var encodingString =  base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded :", encodingString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodingString)
	var decodedString = string(decodeByte)
	fmt.Println("decoded : ", decodedString)
}