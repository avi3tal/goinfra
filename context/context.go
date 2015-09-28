package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonobject struct {
	dd     []NodeType
	ds     []NodeType
	client []NodeType
	// dp     []NodeType
}

type NodeType struct {
	address  string
	username string
	password string
}

// type Context struct {
// 	dd, ds, dp, client map[string]string
// }

// func (self *Context) load(args map[string]string) {
// 	fmt.Println("Start loading context")

// }

func Load(file string) {
	fmt.Println("Start loading context")
	data, e := ioutil.ReadFile(file)
	if e != nil {
		fmt.Printf("Failed to read file: %s\nError: %v\n", file, e)
		os.Exit(1)
	}
	var jsontype jsonobject
	if err := json.Unmarshal(data, &jsontype); err != nil {
		panic(err)
	}
	fmt.Printf("Result: %v\n", jsontype)

	fmt.Println(jsontype.client)
}
