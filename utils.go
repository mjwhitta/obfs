package obfs

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/mjwhitta/errors"
)

func bootstrap(size int) (data []byte, e error) {
	var bInt *big.Int
	var inc int

	if bInt, e = rand.Int(rand.Reader, big.NewInt(MaxInc)); e != nil {
		e = errors.Newf("failed to read random int: %w", e)
		return
	}

	inc = int(bInt.Int64()&0xff) + 2 //nolint:mnd // no 0 or 1

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

func generateSrc(function string, data []byte) string {
	var line string
	var sb strings.Builder

	sb.WriteString("obfs." + function + "(")
	sb.WriteString("\n    []byte{")

	for i, b := range data {
		if (i != 0) && ((i % 9) == 0) { //nolint:mnd // wrap every 9
			sb.WriteString("\n       " + line)
			line = fmt.Sprintf(" 0x%02x,", b)
		} else {
			line += fmt.Sprintf(" 0x%02x,", b)
		}
	}

	if len(line) > 0 {
		sb.WriteString("\n       " + line)
	}

	sb.WriteString("\n    },")
	sb.WriteString("\n)")

	return sb.String()
}
