package obfs

// DeobfuscateByteArray will return the unobfuscated []byte.
func DeobfuscateByteArray(obfuscated []byte) []byte {
	var b []byte
	var increment int = int(obfuscated[0])

	for i := 1; i < len(obfuscated); i += increment {
		b = append(b, obfuscated[i])
	}

	return b
}

// DeobfuscateString will return the unobfuscated string.
func DeobfuscateString(obfuscated []byte) string {
	return string(DeobfuscateByteArray(obfuscated))
}

// ObfuscateByteArray will generate go source to deobfuscate a []byte
// back into the original []byte.
func ObfuscateByteArray(b []byte) (string, error) {
	var data []byte
	var e error
	var increment int
	var obfuscated string

	if data, e = bootstrap(len(b)); e != nil {
		return "", e
	}

	increment = int(data[0])

	for i, c := range b {
		data[(i*increment)+1] = c
	}

	obfuscated = generateSrc("DeobfuscateByteArray", data)

	return obfuscated, nil
}

// ObfuscateString will generate go source to deobfuscate a []byte
// back into the original string.
func ObfuscateString(s string) (string, error) {
	var data []byte
	var e error
	var increment int
	var obfuscated string

	if data, e = bootstrap(len(s)); e != nil {
		return "", e
	}

	increment = int(data[0])

	for i, c := range []byte(s) {
		data[(i*increment)+1] = c
	}

	obfuscated = generateSrc("DeobfuscateString", data)

	return obfuscated, nil
}
