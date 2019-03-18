package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/wattx-backend/model"
)

type TestSuite struct {
	suite.Suite
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupSuite() {
}

func (ts *TestSuite) TestHttpGet() {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	body, err := httpGet(server.URL, nil)
	assert.Equal(ts.T(), []byte("OK"), body, "httpGet should send a Get request and return []byte. [expect: %v, got %v]", []byte("OK"), body)
	assert.Nil(ts.T(), err)
}

func (ts *TestSuite) TestGetTopAsset() {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		data := model.Crypto{
			Data: []model.Data{
				model.Data{
					Symbol:         "EOS",
					Volume24hourto: 1,
				},
				model.Data{
					Symbol:         "XRP",
					Volume24hourto: 3,
				},
				model.Data{
					Symbol:         "LTC",
					Volume24hourto: 2,
				},
				model.Data{
					Symbol:         "NotWhiteListed",
					Volume24hourto: 2,
				},
			},
		}
		dataByte, err := json.Marshal(data)
		assert.Nil(ts.T(), err)
		rw.Write(dataByte)
	}))
	defer server.Close()

	topAssetsChan := make(chan []string)
	topAssetTicker := time.NewTicker(assetUpdateTime)
	go getTopAssets(topAssetsChan, topAssetTicker, server.URL)
	result := <-topAssetsChan

	assert.Equal(ts.T(), []string{"XRP", "LTC", "EOS"}, result)
}

func (ts *TestSuite) TestGetAssetValue() {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		dataByte, err := json.Marshal(conversions)
		assert.Nil(ts.T(), err)
		rw.Write(dataByte)
	}))
	defer server.Close()

	topAssetsChan := make(chan []string)
	assetValueDoneChan := make(chan model.Conversion)
	go getAssetValue(topAssetsChan, assetValueDoneChan, server.URL+"?symbol=%s")
	topAssetsChan <- whiteList[:10]
	result := <-assetValueDoneChan

	assert.Equal(ts.T(), len(result.Data), 10)
}

var conversions = model.Conversion{
	Data: map[string]model.Currency{
		"EOS10": model.Currency{
			Symbol: "EOS10",
			Quote: map[string]model.Price{
				"USD": {
					Price: 10,
				},
			},
		},
		"EOS9": model.Currency{
			Symbol: "EOS9",
			Quote: map[string]model.Price{
				"USD": {
					Price: 9,
				},
			},
		},
		"EOS8": model.Currency{
			Symbol: "EOS8",
			Quote: map[string]model.Price{
				"USD": {
					Price: 8,
				},
			},
		},
		"EOS7": model.Currency{
			Symbol: "EOS7",
			Quote: map[string]model.Price{
				"USD": {
					Price: 7,
				},
			},
		},
		"EOS6": model.Currency{
			Symbol: "EOS6",
			Quote: map[string]model.Price{
				"USD": {
					Price: 6,
				},
			},
		},
		"EOS5": model.Currency{
			Symbol: "EOS5",
			Quote: map[string]model.Price{
				"USD": {
					Price: 5,
				},
			},
		},
		"EOS4": model.Currency{
			Symbol: "EOS4",
			Quote: map[string]model.Price{
				"USD": {
					Price: 4,
				},
			},
		},
		"EOS3": model.Currency{
			Symbol: "EOS3",
			Quote: map[string]model.Price{
				"USD": {
					Price: 3,
				},
			},
		},
		"EOS2": model.Currency{
			Symbol: "EOS2",
			Quote: map[string]model.Price{
				"USD": {
					Price: 2,
				},
			},
		},
		"EOS1": model.Currency{
			Symbol: "EOS1",
			Quote: map[string]model.Price{
				"USD": {
					Price: 1,
				},
			},
		},
	},
}
