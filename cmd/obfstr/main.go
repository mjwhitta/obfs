package main

import (
	"fmt"

	"github.com/mjwhitta/cli"
	"github.com/mjwhitta/log"
	"github.com/mjwhitta/obfs"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r)
			}

			switch r := r.(type) {
			case error:
				log.ErrX(Exception, r.Error())
			case string:
				log.ErrX(Exception, r)
			}
		}
	}()

	validate()

	var e error
	var str string

	if str, e = obfs.ObfuscateString(cli.Arg(0)); e != nil {
		panic(e)
	}

	fmt.Println(str)
}
