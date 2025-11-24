package obfs

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/mjwhitta/errors"
)

func bootstrap(size int) ([]byte, error) {
	var b []byte
	var bInt *big.Int
	var e error
	var increment int
	var n int

	if bInt, e = rand.Int(rand.Reader, big.NewInt(MaxInc)); e != nil {
		e = errors.Newf("failed to read random int: %w", e)
		return nil, e
	}

	increment = int(bInt.Int64()&0xff) + 2 //nolint:mnd // no 0 or 1
	b = make([]byte, (increment*size)+1)

	if n, e = rand.Read(b); e != nil {
		e = errors.Newf("failed to generate random data: %w", e)
		return nil, e
	} else if n != (increment*size)+1 {
		e = errors.New("failed to generate random data")
		return nil, e
	}

	b[0] = byte(increment)

	return b, nil
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
