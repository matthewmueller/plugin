package main

import (
	"fmt"
	"net/rpc"

	"github.com/matthewmueller/go-plugin"
)

func main() {
	conn, err := plugin.Start("go", "run", "./plugin/main.go")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClient(conn)
	var result string
	err = client.Call("Svelte.Transform", "<h2>hi world</h2>", &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
