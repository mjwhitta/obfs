package obfs

// DeobfuscateByteArray will return the unobfuscated []byte.
func DeobfuscateByteArray(data []byte, increment int) []byte {
	var deobfs []byte

	for i := 0; i < len(data); i += increment {
		deobfs = append(deobfs, data[i])
	}

	return deobfs
}

// DeobfuscateString will return the unobfuscated string.
func DeobfuscateString(data []byte, increment int) string {
	var deobfs string

	for i := 0; i < len(data); i += increment {
		deobfs += string(data[i])
	}

	return deobfs
}

// ObfuscateByteArray will generate go source to deobfuscate a []byte
// back into the original []byte.
func ObfuscateByteArray(bArr []byte) (string, error) {
	var data []byte
	var e error
	var increment int

	if data, increment, e = bootstrap(len(bArr)); e != nil {
		return "", e
	}

	for i, b := range bArr {
		data[i*increment] = b
	}

	return generateSrc("DeobfuscateByteArray", data, increment), nil
}

// ObfuscateString will generate go source to deobfuscate a []byte
// back into the original string.
func ObfuscateString(str string) (string, error) {
	var data []byte
	var e error
	var increment int

	if data, increment, e = bootstrap(len(str)); e != nil {
		return "", e
	}

	for i, c := range []byte(str) {
		data[i*increment] = c
	}

	return generateSrc("DeobfuscateString", data, increment), nil
}
