package obfs

import (
	"crypto/rand"
	"io"
	"math/big"

	"github.com/mjwhitta/errors"
	hl "github.com/mjwhitta/hilighter"
)

func bootstrap(size int) (data []byte, e error) {
	var bInt *big.Int
	var inc int

	if bInt, e = rand.Int(rand.Reader, big.NewInt(MaxInc)); e != nil {
		e = errors.Newf("failed to read random int: %w", e)
		return
	}
	inc = int(bInt.Int64()&0xff) + 2 // Don't allow 0 or 1

	data = make([]byte, (inc*size)+1)

	if _, e = io.ReadFull(rand.Reader, data); e != nil {
		e = errors.Newf("failed to generate random data: %w", e)
		return
	}

	data[0] = byte(inc)

	return
}

func deobfs(data []byte) (deobfs []byte) {
	var increment int = int(data[0])

	for i := 1; i < len(data); i += increment {
		deobfs = append(deobfs, data[i])
	}

	return
}

func generateSrc(function string, data []byte) (src string) {
	var line string

	src = "obfs." + function + "("
	src += "\n    []byte{"

	for i, b := range data {
		if (i != 0) && ((i % 9) == 0) {
			src += "\n       " + line
			line = hl.Sprintf(" 0x%02x,", b)
		} else {
			line += hl.Sprintf(" 0x%02x,", b)
		}
	}

	if len(line) > 0 {
		src += "\n       " + line
	}

	src += "\n    },"
	src += "\n)"

	return
}
