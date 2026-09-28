package model

type City struct {
	TokenID   int64  `json:"tokenId"`
	CityName  string `json:"cityName"`
	Latitude  int64  `json:"latitude"`
	Longitude int64  `json:"longitude"`
	FromDate  int64  `json:"fromDate"`
	ToDate    int64  `json:"toDate"`
}
