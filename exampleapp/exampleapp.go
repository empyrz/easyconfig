package main

import _ "github.com/empyrz/easyconfig/exampleapp/examplelib"
import "github.com/empyrz/easyconfig/adaptflag"

//import "flag"
import flag "github.com/ogier/pflag"

func main() {
	adaptflag.AdaptWithFunc(func(info adaptflag.Info) {
		flag.Var(info.Value, info.Name, info.Usage)
	})

	flag.Parse()
}
