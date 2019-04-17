package main

import (
	"fmt"
	"net/http/httputil"
	"os"
	"reflect"

	"github.com/mburtless/ns1-go/ns1"
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

	out, err := r.Send()
	t := reflect.TypeOf(out)
	fmt.Printf("\nOut: %v\nErr: %v\nType: %v\n", out, err, t)
	fmt.Printf("Resp: %v\n", r.HTTPResponse)
	//fmt.Printf("Ipsum: %v\n", (*out)[0])
	fmt.Printf("Ipsum: %v\n", out.Zones[0])

	zn := "ipsumzone.test"
	zoneIn := ns1.GetZoneInput{ZoneName: &zn}
	fmt.Printf("\nInput: %v\n", *zoneIn.ZoneName)

	zr := client.Zones.GetZoneRequest(&zoneIn)
	byts, _ = httputil.DumpRequest(zr.HTTPRequest, true)
	fmt.Println("\n\nRequest:", string(byts))

	zOut, err := zr.Send()
	t = reflect.TypeOf(zOut)
	fmt.Printf("\nOut: %v\nErr: %v\nType: %v\n", zOut, err, t)
}
