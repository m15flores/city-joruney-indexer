package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type City struct {
	TokenID   int64  `json:"tokenId"`
	CityName  string `json:"cityName"`
	Latitude  int64  `json:"latitude"`
	Longitude int64  `json:"longitude"`
	FromDate  int64  `json:"fromDate"`
	ToDate    int64  `json:"toDate"`
}

func main() {
	client, err := connectClient("https://arb1.arbitrum.io/rpc")
	if err != nil {
		log.Fatal(err)
	}

	contract, err := NewCityJourneyNFT(common.HexToAddress("0xA066b02716BEFaAb59B370224Af1c8C0bBEA1eDd"), client)
	if err != nil {
		log.Fatal(err)
	}

	cities, err := readMintedCities(client, contract)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loaded", len(cities), "cities")

	db, err := openDB("cities.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := insertCities(db, cities); err != nil {
		log.Fatal(err)
	}

	server := NewServer(db)
	log.Fatal(server.Start(":8080"))
}

func connectClient(link string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(link)
	return client, err
}

func readMintedCities(client *ethclient.Client, contract *CityJourneyNFT) ([]City, error) {
	iterator, err := contract.FilterMintNFT(&bind.FilterOpts{Start: 508821437})
	if err != nil {
		log.Fatal(err)
	}

	var cities []City

	for iterator.Next() {
		cityData, err := contract.CityData(nil, iterator.Event.TokenId)
		if err != nil {
			log.Fatal(err)
		}

		var city City
		city.TokenID = iterator.Event.TokenId.Int64()
		city.CityName = cityData.CityName
		city.Latitude = cityData.Latitude.Int64()
		city.Longitude = cityData.Longitude.Int64()
		city.FromDate = cityData.FromDate.Int64()
		city.ToDate = cityData.ToDate.Int64()

		fmt.Printf(
			"For token: %d => City name: %s, Latitude: %d, Longitude: %d, From Date: %d, To Date: %d\n",
			city.TokenID, city.CityName, city.Latitude, city.Longitude, city.FromDate, city.ToDate,
		)
		cities = append(cities, city)
	}
	return cities, nil
}
