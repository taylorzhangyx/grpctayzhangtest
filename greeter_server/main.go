/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/taylorzhangyx/grpctayzhangtest"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")

	loadTotalCount int64
	loadMetrics    = make(map[int64]int64)
	loadchan       = make(chan int64, 10000)
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}

func (s server) LoadTest(ctx context.Context, request *pb.LoadTestRequest) (*pb.LoadTestReply, error) {
	t := time.Now()
	t.Unix()
	loadchan <- t.Unix()
	return &pb.LoadTestReply{RetCode: 0}, nil
}

func (s server) LoadTestMetrics(ctx context.Context, request *pb.LoadTestMetricsRequest) (*pb.LoadTestMetricsReply, error) {
	return &pb.LoadTestMetricsReply{Metrics: loadMetrics, Total: loadTotalCount}, nil
}

func (s server) LoadTestClear(ctx context.Context, request *pb.LoadTestClearRequest) (*pb.LoadTestClearReply, error) {
	loadTotalCount = 0
	loadMetrics = make(map[int64]int64)
	return &pb.LoadTestClearReply{RetCode: 0}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	go consumeLoadCount()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func consumeLoadCount() {
	for {
		select {
		case t := <-loadchan:
			loadMetrics[t]++
			loadTotalCount++
		}
	}
}
