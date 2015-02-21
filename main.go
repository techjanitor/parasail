package main

import (
	"fmt"
	"net/url"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.Port = 6000

	k.HandleFunc("hello", Hello)

	k.Register(&url.URL{Scheme: "http", Host: "localhost:6000/kite"})
	k.Run()
}

func Hello(r *kite.Request) (interface{}, error) {
	// Print a log on remote Kite.
	// This message will be printed on client's console.
	r.Client.Go(fmt.Sprintf("Hello %s!", r.LocalKite.Kite().Name))

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}
