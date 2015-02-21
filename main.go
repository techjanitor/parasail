package main

import (
	"fmt"
	"net/url"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.Port = 6000
	k.HandleFunc("hello", func(r *kite.Request) {
		fmt.Println("hello!")
		return
	})

	k.Register(&url.URL{Scheme: "http", Host: "localhost:6000/kite"})
	k.Run()
}
