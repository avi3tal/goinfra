package cli

import (
	"flag"
)

type Args struct {
	Json_file string
}

func Parse() (x *Args) {
	x = new(Args)
	flag.StringVar(&x.Json_file, "json", "", "path to json file")

	flag.Parse()
	return
}
