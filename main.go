package main

import (
	"fmt"
	"net/url"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("parasail", "1.0.0")
	k.SetLogLevel(kite.DEBUG)

	k.Config.Username = "parasail"
	k.Config.KiteKey = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0MjQ0OTI0ODEsImlzcyI6InBhcmFzYWlsIiwianRpIjoiZGMwMWY2NDQtM2QzZC00YmMxLTYzYzgtZTA4MThhMTYwZWRlIiwia29udHJvbEtleSI6Ii0tLS0tQkVHSU4gUFVCTElDIEtFWS0tLS0tXG5NSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXhXSXhJTVI1U3I1MFVXWGZWTU5sXG5UZmR5VWszcXpPM0trc1ptU0dURGw1bXpmb1NOc2xRanBlOUVJc0tINUN5N1IrR2F1MFJaWEdQRTVpSUpjck1UXG5DMWZOckE0WW45ZXVOb0ZCd2k2TEkwdWV1d0ZySGViakt2eG1mSVU3RVZkM3J6bDI3QzhralA0UUZwMjJGQ1l3XG5RdTJRcFR0WFpuNnh5bS9HWEUxUlpXYXh1cVpYZjRrQU5xcGgwNUVFaTBOSi9wTVFSWTFkd1Y1bW8wZndHRk9VXG5OV2RacFYrVWRYRDRUYitYUXVOcHM4dTlPa1d6R0FUbXFLbmFieDhXNWVWODBzK1A5Q3hRRUd6ci9XVWE2cC9SXG5jVE5ZVTd0REtneFE2SzgvaStvcjd5TTBnNlliMXk2T2xremJPaXFWKzhqZDhTczFMODZOOENSeFkyUjBvU0UxXG5Vd0lEQVFBQlxuLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tIiwia29udHJvbFVSTCI6Imh0dHA6Ly9kaXNjb3ZlcnkubW9kbm9kZS5jb206NTU1NS9raXRlIiwic3ViIjoicGFyYXNhaWwifQ.wogrzTMlVf0bl4xgu0YwpmI7X7zkwBqT7Duvc2noXT1794aIMwynT6087RBQJ5tAoNti_ZeQWSv7kGxS_hPRt_qcDBXq4vI6kcnnGvRGGQizzdRtQmrPVm6oEc-629d3dQes5n-UtKajw403Sq99FBUe6I6LEC6zM_UEdNfuNKxzuOykOp8QQqD6vSWbONv4Y55HjLPy1brA1wNyaaxOAeJmxvqXgO5HjQbTJw8rYBHIrX0W7VJU4B3091rWq4oABsMiJHywC4SBX_D12zuZzpbLkPyy72sspAP7YuS6LT8DO4bCTx8ft0NwfqjPBaLDwaTUt2RmVAI67p4ADIJ9hg"

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
	r.Client.Go("kite.log", fmt.Sprintf("Hello %s!", r.LocalKite.Kite().Name))

	// You can return anything as result, as long as it is JSON marshalable.
	return nil, nil
}
