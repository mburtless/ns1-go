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
		switch err.(type) {
		case *zones.GetZoneNotFound:
			log.Printf("Error: %s\n", *err.(*zones.GetZoneNotFound).Payload.Message)
		default:
			log.Printf("Something else: %#v\n", err.Error())
		}
		os.Exit(1)
	}
	for _, record := range zone.Payload.Records {
		fmt.Printf("%#v\n", *record.Domain)
	}
}
