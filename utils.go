package obfs

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"
)

func bootstrap(size int) ([]byte, error) {
	var bInt *big.Int
	var data []byte
	var e error
	var inc int

	if bInt, e = rand.Int(rand.Reader, big.NewInt(MaxInc)); e != nil {
		return []byte{}, e
	}
	inc = int(bInt.Int64()&0xff) + 2 // Don't allow 0 or 1

	data = make([]byte, (inc*size)+1)

	if _, e = io.ReadFull(rand.Reader, data); e != nil {
		return []byte{}, e
	}

	data[0] = byte(inc)

	return data, nil
}

func deobfs(data []byte) []byte {
	var deobfs []byte
	var increment = int(data[0])

	for i := 1; i < len(data); i += increment {
		deobfs = append(deobfs, data[i])
	}

	return deobfs
}

func generateSrc(function string, data []byte) string {
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
	src = append(src, ")")

	return strings.Join(src, "\n")
}
