package main

import (
	"fmt"
	"os"
	"github.com/mburtless/ns1-go/ns1"
	"net/http/httputil"
)

func main() {
	k := os.Getenv("NS1_APIKEY")
	if k == "" {
		fmt.Println("NS1_APIKEY environment variable is not set, giving up")
	}
	config, _ := ns1.NewConfig(&ns1.NewConfigInput{APIKey: k})
	fmt.Printf("Config: %v\n", config)

	client := ns1.NewClient(config)
	fmt.Printf("\nClient: %v\n", client)

	listIn := ns1.ListZonesInput{}
	fmt.Printf("\nInput: %v", listIn)

	r := client.Zones.ListZonesRequest(&listIn)
	byts, _ := httputil.DumpRequest(r.HTTPRequest, true)
	fmt.Println("\n\nRequest:", string(byts))

	out, _ := r.Send()
	fmt.Printf("\nOut: %v\n", out)
}
