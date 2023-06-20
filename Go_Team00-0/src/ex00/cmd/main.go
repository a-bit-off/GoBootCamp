package main

import (
	"context"
	"ex00/pkg/proto"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	"time"
)

type TransmitterServer struct {
	proto.UnimplementedTransmitterServer
}

func (ts TransmitterServer) Connection(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	// SessionId
	resp = &proto.Response{}
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	resp.SessionId = id.String()

	// Frequency
	minMean := -10.0
	maxMean := 10.0
	minStd := 0.3
	maxStd := 1.5

	rand.Seed(time.Now().UnixNano())
	mean := rand.Float64()*(maxMean-minMean) + minMean
	std := rand.Float64()*(maxStd-minStd) + minStd
	resp.Frequency = rand.NormFloat64()*std + mean

	// UTC
	resp.UTC = timestamppb.Now()

	// stdout
	_, err = fmt.Printf("SessionId: %s\nFrequency: %f\nUTC: %s\n",
		resp.SessionId, resp.Frequency, resp.UTC.AsTime())
	if err != nil {
		return
	}

	return resp, nil
}

func main() {
	serverListen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTransmitterServer(grpcServer, TransmitterServer{})

	err = grpcServer.Serve(serverListen)
	if err != nil {
		log.Fatal(err)
	}
}
