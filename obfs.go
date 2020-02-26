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
func ObfuscateByteArray(bArr []byte) (out string, e error) {
	var data []byte
	var increment int

	if data, e = bootstrap(len(bArr)); e != nil {
		return
	}

	increment = int(data[0])

	for i, b := range bArr {
		data[(i*increment)+1] = b
	}

	out = generateSrc("DeobfuscateByteArray", data)
	return
}

// ObfuscateString will generate go source to deobfuscate a []byte
// back into the original string.
func ObfuscateString(str string) (out string, e error) {
	var data []byte
	var increment int

	if data, e = bootstrap(len(str)); e != nil {
		return
	}

	increment = int(data[0])

	for i, c := range []byte(str) {
		data[(i*increment)+1] = c
	}

	out = generateSrc("DeobfuscateString", data)
	return
}
