package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/koding/kite"
	"github.com/koding/kite/kitekey"
	"github.com/koding/kite/protocol"
)

var (
	discoveryUrl = "http://discovery.modnode.com:6000/kite"
	k            = kite.New("parasol", "1.0.0")
	username     = "prim"
)

func init() {

	k.Config.Username = username
	k.Config.Environment = "digitalocean"
	k.Config.Region = "nyc"

}

func main() {

	methodFlag := flag.String("method", "", "method to be used on the hosts")
	commandFlag := flag.String("command", "", "the command to be executed on the hosts")
	initialFlag := flag.Bool("initial", false, "register with kontrol")
	flag.Parse()

	if *methodFlag == "" {
		fmt.Println("method required")
		return
	}

	if *initialFlag {
		initial()
	}

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

	k.Config.KiteKey = key

	k.Config.KontrolURL = discoveryUrl
	k.Config.KontrolUser = "discovery"
	k.Config.KontrolKey = token.Claims["kontrolKey"].(string)

	kites, err := k.GetKites(&protocol.KontrolQuery{
		Username: "pram",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	var wg sync.WaitGroup

	wg.Add(len(kites))

	for _, client := range kites {

		go func(client *kite.Client) {
			defer wg.Done()

			client.Dial()

			response, err := client.TellWithTimeout(*methodFlag, 5*time.Minute, *commandFlag)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintln(f, client.Hostname)

			output := response.MustString()

			fmt.Fprint(f, output)

			client.Close()

		}(client)

	}

	wg.Wait()

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
