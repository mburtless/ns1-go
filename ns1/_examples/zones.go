package main

import (
	"fmt"
	"os"

	"github.com/mburtless/ns1-go/ns1"
)

func main() {
	k := os.Getenv("NS1_APIKEY")
	if k == "" {
		fmt.Println("NS1_APIKEY environment variable is not set, giving up")
		os.Exit(1)
	}
	// Create config
	config, _ := ns1.NewConfig(&ns1.NewConfigInput{APIKey: k})
	// Create client
	client := ns1.NewClient(config)
	// Create input
	listIn := ns1.ListZonesInput{}
	// Create request
	r := client.Zones.ListZonesRequest(&listIn)
	// Send request
	out, _ := r.Send()
	fmt.Printf("\nAll Zones: %v\n", out.Zones)
	fmt.Printf("First Zone's ID: %v\n", out.Zones[0].ID)

	// Create input
	zoneIn := ns1.GetZoneInput{Zone: "ipsumzone.test"}
	// Create request
	zr := client.Zones.GetZoneRequest(&zoneIn)
	fmt.Printf("Body: %v\n", zr.Body)
	// Send request
	zOut, _ := zr.Send()
	fmt.Printf("\nZone: %v\n", zOut)

	// Create input
	createIn := ns1.CreateZoneInput{Zone: "threepnettest5.test", TTL: 3600}
	// Create request
	cr := client.Zones.CreateZoneRequest(&createIn)
	fmt.Printf("Body: %v\n", cr.Body)
	// Send request
	_, err := cr.Send()
	if err != nil {
		fmt.Printf("\nErr: %v\n", err)
	}

	/*if crOut != nil {
		fmt.Printf("\nZone: %v\n", crOut)
	}*/

	// Create input
	updateIn := ns1.UpdateZoneInput{Zone: "threepnettest5.test", Refresh: 41000}
	// Create request
	ur := client.Zones.UpdateZoneRequest(&updateIn)
	fmt.Printf("Body: %v\n", ur.Body)
	// Send request
	urOut, err := ur.Send()
	if err != nil {
		fmt.Printf("\nErr: %v\n", err)
	}
	fmt.Printf("Out: %v\n", urOut)

	// Create input
	deleteIn := ns1.DeleteZoneInput{Zone: "garbagefakezone.test"}
	// Create request
	dr := client.Zones.DeleteZoneRequest(&deleteIn)
	//Send request
	drOut, err := dr.Send()
	if err != nil {
		fmt.Printf("\nErr: %v\n", err)
	}
	fmt.Printf("Out: %v\n", drOut)
}
