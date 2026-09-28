package main

import (
	"city-journey-indexer/internal/api"
	"city-journey-indexer/internal/onchain"
	"city-journey-indexer/internal/store"
	"fmt"
	"log"
)

const (
	rpcURL          = "https://arb1.arbitrum.io/rpc"
	contractAddress = "0xA066b02716BEFaAb59B370224Af1c8C0bBEA1eDd"
	deployBlock     = 508821437
	dbPath          = "cities.db"
	listenAddr      = ":8080"
)

func main() {
	client, err := onchain.Connect(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	cities, err := onchain.ReadMintedCities(client, contractAddress, deployBlock)
	if err != nil {
		log.Fatal(err)
	}
	client.Close()

	fmt.Println("Loaded", len(cities), "cities")

	db, err := store.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.InsertCities(db, cities); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(db)
	log.Fatal(server.Start(listenAddr))
}
