package onchain

import (
	"city-journey-indexer/internal/model"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Connect(link string) (*ethclient.Client, error) {
	return ethclient.Dial(link)
}

func ReadMintedCities(client *ethclient.Client, address string, initBlock uint64) ([]model.City, error) {
	contract, err := NewCityJourneyNFT(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	iterator, err := contract.FilterMintNFT(&bind.FilterOpts{Start: initBlock})
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var cities []model.City

	for iterator.Next() {
		cityData, err := contract.CityData(nil, iterator.Event.TokenId)
		if err != nil {
			return nil, err
		}

		var city model.City
		city.TokenID = iterator.Event.TokenId.Int64()
		city.CityName = cityData.CityName
		city.Latitude = cityData.Latitude.Int64()
		city.Longitude = cityData.Longitude.Int64()
		city.FromDate = cityData.FromDate.Int64()
		city.ToDate = cityData.ToDate.Int64()

		cities = append(cities, city)
	}
	if err := iterator.Error(); err != nil {
		return nil, err
	}

	return cities, nil
}
