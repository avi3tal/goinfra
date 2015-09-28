package main

import (
	"fmt"
	"github.com/avi3tal/goinfra/cli"
	"github.com/avi3tal/goinfra/context"
)

func main() {
	fmt.Println("Starting tonat ...")
	x := cli.Parse()
	fmt.Println(*x)

	context.Load(x.Json_file)

}
