package obfs

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"
)

// Deobfuscate will return the unobfuscated string.
func Deobfuscate(data []byte, increment int) string {
	var str string

	for i := 0; i < len(data); i += increment {
		str += string(data[i])
	}

	return str
}

// Obfuscate will generate go source to deobfuscate a []byte back into
// a string.
func Obfuscate(str string) (string, error) {
	var data []byte
	var e error
	var inc *big.Int
	var increment int
	var line string
	var src []string

	if inc, e = rand.Int(rand.Reader, big.NewInt(16)); e != nil {
		return "", e
	}
	increment = int(inc.Int64()&0xff) + 2 // Don't allow 0 or 1

	data = make([]byte, increment*len(str))

	if _, e = io.ReadFull(rand.Reader, data); e != nil {
		return "", e
	}

	for i, c := range []byte(str) {
		data[i*increment] = c
	}

	src = append(src, "// "+str)
	src = append(src, "var str string = obfs.Deobfuscate(")
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
	src = append(src, fmt.Sprintf("    %d,", increment))
	src = append(src, ")")

	return strings.Join(src, "\n"), nil
}
