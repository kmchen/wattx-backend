package model

import (
	"encoding/json"
	"fmt"

	pb "github.com/wattx-backend/proto"
)

type Data struct {
	Symbol         string  `json:"symbol"`
	Volume24hourto float32 `json:"Volume24hourto"`
}

type Crypto struct {
	Data []Data
}

type Conversion struct {
	Data map[string]Currency
}

type Currency struct {
	Quote  map[string]Price `json: "quote"`
	Symbol string           `json: "symbol"`
}

type Price struct {
	Price float32 `json: "price"`
}

func UnmarshalConversion(resp []byte) (Conversion, error) {
	conversion := Conversion{}
	err := json.Unmarshal(resp, &conversion)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conversion, err
}

type AssetValue struct {
	Key   string
	Value float32
}

func ToProtoConversion(conversion Conversion) pb.Data {
	var pbAssetValue = make([]*pb.AssetValue, 0)
	for currencyName, currency := range conversion.Data {
		for _, Price := range currency.Quote {
			assetValue := &pb.AssetValue{
				Key:   currencyName,
				Value: Price.Price,
			}
			pbAssetValue = append(pbAssetValue, assetValue)
		}
	}
	return pb.Data{Data: pbAssetValue}
}
