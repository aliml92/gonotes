// source: https://blog.logrocket.com/advanced-unit-testing-patterns-go/ 
package encode

import (
	"encoding/base64"
)

func getBase64Encoding(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}


