package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/koding/kite"
	"github.com/koding/kite/kitekey"
)

var (
	discoveryUrl = "http://discovery.modnode.com:6000/kite"
	k            = kite.New("parasail", "1.0.0")
	username     = "pram"
)

func init() {

	kontrol := k.NewClient(discoveryUrl)
	err := kontrol.Dial()
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := kontrol.TellWithTimeout("registerMachine", 5*time.Minute, username)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = kitekey.Write(result.MustString())
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {

	key, err := kitekey.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	token, err := kitekey.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	k.Config.Username = username
	k.Config.Environment = "digitalocean"
	k.Config.Region = "nyc"

	k.Config.Port = 6000
	k.Config.KiteKey = key

	k.Config.KontrolURL = discoveryUrl
	k.Config.KontrolUser = "discovery"
	k.Config.KontrolKey = token.Claims["kontrolKey"].(string)

	discovery := &url.URL{
		Scheme: "http",
		Host:   "discovery.modnode.com:6000",
		Path:   "kite",
	}

	k.Register(discovery)

	k.HandleFunc("hello", Hello)

	k.Run()

}

func Hello(r *kite.Request) (interface{}, error) {
	// Print a log on remote Kite.
	// This message will be printed on client's console.
	r.Client.Go("kite.log", fmt.Sprintf("Hello %s!", r.LocalKite.Kite().Name))

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}
