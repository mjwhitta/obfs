package main

import (
	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/log"
	"github.com/mjwhitta/obfs"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			log.ErrX(Exception, r.(error).Error())
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
