package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/wattx-backend/model"
	pb "github.com/wattx-backend/proto"
	"google.golang.org/grpc"
)

const (
	topAssetsUrl = "https://min-api.cryptocompare.com/data/top/volumes?tsym=USD&limit=500"
	currencyUrl  = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=%s"

	assetUpdateTime = 60 * time.Second
	batchSize       = 20
	numTopAssets    = 200
)

var grpcServerIp = flag.String("serverIp", "localhost", "grpc server ip")
var grpcPort = flag.String("port", "50051", "grpc server port")
var headers = map[string]string{
	"X-CMC_PRO_API_KEY": "4701ef03-3006-4d0a-8564-eb27b7715c00",
}

var DefaultClient = &http.Client{}

func httpGet(url string, headers map[string]string) ([]byte, error) {
	// Init new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Fail to create request, %v\n", err)
		return nil, err
	}

	// Add header
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	// Send out request
	resp, err := DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Fail to send request %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func getTopAssets(topAssetsChan chan []string, ticker *time.Ticker, url string) {
	t := time.Time{}
	for ; true; t = <-ticker.C {
		// Fetching top 500
		fmt.Printf("Fetching top 500 assets %v\n", t)
		resp, err := httpGet(url, nil)
		if err != nil {
			fmt.Printf("Fail to fetch top asset: %v\n", err)
			continue
		}
		// Unmarshal Asset data
		crypto := model.Crypto{}
		if err := json.Unmarshal(resp, &crypto); err != nil {
			fmt.Printf("Fail to unmarshal assets: %v\n", err)
			continue
		}

		// Sort by volume
		sort.Slice(crypto.Data,
			func(i, j int) bool {
				return crypto.Data[j].Volume24hourto < crypto.Data[i].Volume24hourto
			},
		)

		// Make sure assets are whitelisted and select top 200
		var topAssets = make([]string, 0, numTopAssets)
		var j = 0
		for i := 0; i < len(crypto.Data) && j < cap(topAssets); i += 1 {
			if isWhiteListed(crypto.Data[i].Symbol) {
				topAssets = append(topAssets, crypto.Data[i].Symbol)
				j += 1
			}
		}
		fmt.Printf("Top assets size: %v\n", topAssets)
		topAssetsChan <- topAssets
	}
}

func getAssetValue(topAssetsChan chan []string, assetDoneChan chan model.Conversion, currencyUrl string) {
	for topAssets := range topAssetsChan {
		for i := 0; i < len(topAssets); i += batchSize {
			assets := topAssets[i : i+batchSize]
			assetsStr := strings.Join(assets, ",")
			fmt.Printf("url: %s\n", assetsStr)
			go func(symbols string, done chan model.Conversion) {
				var url = fmt.Sprintf(currencyUrl, symbols)
				var resp, err = httpGet(url, headers)
				if err != nil {
					fmt.Println("error:", err)
					return
				}
				conversion, err := model.UnmarshalConversion(resp)
				if err != nil {
					return
				}
				done <- conversion
			}(assetsStr, assetDoneChan)
		}
	}
}

func main() {
	flag.Parse()
	topAssetsChan := make(chan []string)

	assetValueDoneChan := make(chan model.Conversion)
	go getAssetValue(topAssetsChan, assetValueDoneChan, currencyUrl)

	topAssetTicker := time.NewTicker(assetUpdateTime)
	go getTopAssets(topAssetsChan, topAssetTicker, topAssetsUrl)

	// Set up a connection to the server.
	var grpcAddr = fmt.Sprintf("%s:%s", *grpcServerIp, *grpcPort)
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for value := range assetValueDoneChan {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		pbData := model.ToProtoConversion(value)
		fmt.Printf("...... %v\n", pbData)
		r, err := c.UpdateAsset(ctx, &pbData)
		if err != nil {
			fmt.Printf("Fail to send udpated asset to ranking service: %v\n", err)
			continue
		}
		fmt.Printf("Successfully update asset to ranking service: %v\n", r.Status)
		cancel()
	}
}
