package main

import (
	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/obfs"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			errx(Exception, r.(error).Error())
		}
	}()

	validate()

	var e error
	var str string

	if str, e = obfs.ObfuscateString(cli.Arg(0)); e != nil {
		panic(e)
	}

	hl.Println(str)
}
