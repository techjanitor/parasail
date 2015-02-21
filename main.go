package main

import (
	"fmt"
	"net/url"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("parasail", "1.0.0")
	k.SetLogLevel(kite.DEBUG)

	c.Config.Username = "parasail"
	c.Config.KiteKey = "/root/.kite/kite.key"

	k.Config.Port = 6000
	k.Config.Environment = "digitalocean"
	k.Config.Region = "nyc"

	k.Config.KontrolURL = "http://discovery.modnode.com:6000/kite"
	k.Config.KontrolUser = "parasail"
	k.Config.KontrolKey = "/root/key.pem"

	discovery := &url.URL{
		Scheme: "http",
		Host:   "discovery.modnode.com:6000",
		Path:   "kite",
	}

	fmt.Println(discovery.String())

	k.RegisterForever(discovery)

	k.HandleFunc("hello", Hello)

	k.Run()
}

func Hello(r *kite.Request) (interface{}, error) {
	// Print a log on remote Kite.
	// This message will be printed on client's console.
	r.Client.Go(fmt.Sprintf("Hello %s!", r.LocalKite.Kite().Name))

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}
