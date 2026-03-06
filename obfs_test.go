//nolint:godoclint // These are tests
package obfs_test

import (
	"bufio"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/mjwhitta/obfs"
	assert "github.com/stretchr/testify/require"
)

func srcToBytes(src string) []byte {
	var b []byte
	var br *bufio.Reader
	var e error
	var line string
	var sb strings.Builder

	br = bufio.NewReader(strings.NewReader(src))

	// Don't need first two lines
	for e = nil; e == nil; line, e = br.ReadString('\n') {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "0x") {
			continue
		}

		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "0x", "")

		sb.WriteString(line)
	}

	b, _ = hex.DecodeString(sb.String())

	return b
}

func TestObfuscateByteArray(t *testing.T) {
	var actual []byte
	var e error
	var expected []byte = []byte("this is a test")
	var src string

	src, e = obfs.ObfuscateByteArray(expected)
	assert.NoError(t, e)
	assert.NotEmpty(t, src)

	actual = obfs.DeobfuscateByteArray(srcToBytes(src))
	assert.Equal(t, expected, actual)
}

func TestObfuscateString(t *testing.T) {
	var actual string
	var e error
	var expected string = "this is a test"
	var src string

	src, e = obfs.ObfuscateString(expected)
	assert.NoError(t, e)
	assert.NotEmpty(t, src)

	actual = obfs.DeobfuscateString(srcToBytes(src))
	assert.Equal(t, expected, actual)
}
