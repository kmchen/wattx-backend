package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sort"

	"github.com/wattx-backend/model"
	pb "github.com/wattx-backend/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

var ranking = map[string]model.AssetValue{}

func (s *server) UpdateAsset(ctx context.Context, data *pb.Data) (*pb.Reply, error) {
	log.Printf("Received: %v", data)
	for _, value := range data.Data {
		ranking[value.Key] = model.AssetValue{
			Key:   value.Key,
			Value: value.Value,
		}
	}
	return &pb.Reply{Status: "OK"}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	assetValues := make([]model.AssetValue, len(ranking))
	for _, value := range ranking {
		assetValues = append(assetValues, value)
	}
	sort.Slice(assetValues,
		func(i, j int) bool {
			return assetValues[j].Value < assetValues[i].Value
		},
	)

	rankingJson, err := json.MarshalIndent(assetValues[:200], "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rankingJson)
}

func main() {

	http.HandleFunc("/data", handler)
	go http.ListenAndServe(":8888", nil)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
