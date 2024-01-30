// source: https://blog.logrocket.com/advanced-unit-testing-patterns-go/ 
package encode

import (
	"os"
	"path/filepath"
	"testing"
)


func TestGetBase64Encoding(t *testing.T) {
	cases := []string{"lily01", "lily02", "lily03"}

	for _, v := range cases {
		t.Run(v, func(t *testing.T) {
			b, err := os.ReadFile(filepath.Join("testdata", v+".jpg"))
			if err != nil {
				t.Fatal(err)
			}

			expected, err := os.ReadFile(filepath.Join("testdata", v+"_data.txt"))
			if err != nil {
				t.Fatal(err)
			}

			got := getBase64Encoding(b)

			if string(expected) != got {
				t.Fatalf("Expected output to be: '%s', but got: '%s'", string(expected), got)
			}
		})
	}
}