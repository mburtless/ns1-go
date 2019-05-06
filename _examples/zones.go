package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mburtless/ns1-go-v3/ns1"
	"github.com/mburtless/ns1-go-v3/ns1/zones"
)

func main() {
	k := os.Getenv("NS1_APIKEY")

	client, err := ns1.New(ns1.Config{APIKey: k})
	if err != nil {
		log.Panic(err)
	}
	zoneIn := zones.NewZonesByZoneGetParams()
	zoneIn.SetZone("ipsumzone.test")
	zone, err := client.Zones.ZonesByZoneGet(context.Background(), zoneIn)
	if err != nil {
		log.Panic(err)
	}
	for _, record := range zone.Payload.Records {
		fmt.Printf("%#v\n", *record.Domain)
	}
}
