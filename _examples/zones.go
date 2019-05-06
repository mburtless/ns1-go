package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mburtless/ns1-go/ns1"
	"github.com/mburtless/ns1-go/ns1/zones"
)

func main() {
	// Create client
	k := os.Getenv("NS1_APIKEY")
	client, err := ns1.New(ns1.Config{APIKey: k})
	if err != nil {
		log.Panic(err)
	}

	// Create input
	zoneIn := zones.NewGetZoneParams()
	zoneIn.SetZone("ipsumzone.test")

	// Get zone
	zone, err := client.Zones.GetZone(context.Background(), zoneIn)
	if err != nil {
		log.Panic(err)
	}
	for _, record := range zone.Payload.Records {
		fmt.Printf("%#v\n", *record.Domain)
	}
}
