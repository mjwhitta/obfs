package obfs

// DeobfuscateByteArray will return the unobfuscated []byte.
func DeobfuscateByteArray(data []byte) []byte {
	return deobfs(data)
}

// DeobfuscateString will return the unobfuscated string.
func DeobfuscateString(data []byte) string {
	return string(deobfs(data))
}

// ObfuscateByteArray will generate go source to deobfuscate a []byte
// back into the original []byte.
func ObfuscateByteArray(bArr []byte) (string, error) {
	var data []byte
	var e error
	var increment int

	if data, e = bootstrap(len(bArr)); e != nil {
		return "", e
	}

	increment = int(data[0])

	for i, b := range bArr {
		data[(i*increment)+1] = b
	}

	return generateSrc("DeobfuscateByteArray", data), nil
}

// ObfuscateString will generate go source to deobfuscate a []byte
// back into the original string.
func ObfuscateString(str string) (string, error) {
	var data []byte
	var e error
	var increment int

	if data, e = bootstrap(len(str)); e != nil {
		return "", e
	}

	increment = int(data[0])

	for i, c := range []byte(str) {
		data[(i*increment)+1] = c
	}

	return generateSrc("DeobfuscateString", data), nil
}
