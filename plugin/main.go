package main

import (
	"net/rpc"
	"strings"

	"github.com/matthewmueller/go-plugin"
)

func main() {
	conn, err := plugin.Serve("Svelte")
	if err != nil {
		panic(err)
	}
	server := rpc.NewServer()
	server.Register(&Svelte{})
	server.ServeConn(conn)
}

// Svelte struct
type Svelte struct {
}

// Transform function
func (p *Svelte) Transform(code string, reply *string) error {
	*reply = strings.Replace(code, "world", "mars", -1)
	return nil
}
