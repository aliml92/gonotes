package main

import (
	"encoding/base64"
)

func getBase64Encoding(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}


