package main

import (
	"flag"
	"fmt"
	"os/exec"
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

	k.Config.Port = 6000
	k.Config.Username = username
	k.Config.Environment = "digitalocean"
	k.Config.Region = "nyc"

}

func main() {

	initialFlag := flag.Bool("initial", false, "register with kontrol")
	flag.Parse()

	if *initialFlag {
		initial()
	}

	k.HandleFunc("hello", Hello)
	k.HandleFunc("exec", Exec)

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

	k.SetLogLevel(kite.INFO)

	k.Config.KiteKey = key

	k.Config.KontrolURL = discoveryUrl
	k.Config.KontrolUser = "discovery"
	k.Config.KontrolKey = token.Claims["kontrolKey"].(string)

	k.Register(k.RegisterURL(false))

	k.Run()

}

func initial() {
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

	return
}

func Hello(r *kite.Request) (interface{}, error) {
	// Print a log on remote Kite.
	// This message will be printed on client's console.
	r.Client.Go("kite.log", fmt.Sprintf("Hello from %s!", r.LocalKite.Kite().Hostname))

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}

func Exec(r *kite.Request) (interface{}, error) {
	r.Client.Go("kite.log", fmt.Sprintf("Executing on %s!", r.LocalKite.Kite().Hostname))

	command := r.Args.One().MustString()

	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return nil, err
	}

	response := string(out[:])

	return response, nil

}
