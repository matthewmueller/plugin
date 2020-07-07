# Plugin

Plugin is a generic plugin system for Go. It was originally built for extending CLIs with custom commands.

## Install

```
go get github.com/matthewmueller/plugin
```

## Example

In this example, the host binary spawns a plugin binary, then communicates with that plugin over RPC via `FD=3` and `FD=4`.

### Host

**host/main.go**

```go
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
```

### Plugin

**plugin/main.go**

```go
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
```

### Usage

```sh
go run host/main.go
<h1>hi mars</h1>
```

## TODO

- [ ] Add some tests

## License

MIT
