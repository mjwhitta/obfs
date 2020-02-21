package obfs

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"
)

func bootstrap(size int) ([]byte, int, error) {
	var bInt *big.Int
	var data []byte
	var e error
	var inc int

	if bInt, e = rand.Int(rand.Reader, big.NewInt(16)); e != nil {
		return []byte{}, 0, e
	}
	inc = int(bInt.Int64()&0xff) + 2 // Don't allow 0 or 1

	data = make([]byte, inc*size)

	if _, e = io.ReadFull(rand.Reader, data); e != nil {
		return []byte{}, 0, e
	}

	return data, inc, nil
}

func generateSrc(function string, data []byte, inc int) string {
	var line string
	var src []string

	src = append(src, "var deobfs = obfs."+function+"(")
	src = append(src, "    []byte{")

	for i, b := range data {
		if (i != 0) && ((i % 9) == 0) {
			src = append(src, "       "+line)
			line = fmt.Sprintf(" 0x%02x,", b)
		} else {
			line += fmt.Sprintf(" 0x%02x,", b)
		}
	}

	if len(line) > 0 {
		src = append(src, "       "+line)
	}

	src = append(src, "    },")
	src = append(src, fmt.Sprintf("    %d,", inc))
	src = append(src, ")")

	return strings.Join(src, "\n")
}
