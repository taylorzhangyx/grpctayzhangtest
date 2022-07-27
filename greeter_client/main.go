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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/taylorzhangyx/grpctayzhangtest/taylorzhtestpb/taylorzhtestpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "ap-shanghai.test.ti.tencentcs.com:443", "the address to connect to")
	// addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	ts := time.Now()
	defer func() {
		log.Print("cost", time.Now().Sub(ts))
	}()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := taylorzhtestpb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.SayHello(ctx, &taylorzhtestpb.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	loadRet, _ := c.LoadTest(ctx, &taylorzhtestpb.LoadTestRequest{Body: "bbb"})
	log.Printf("LoadTest: %v", loadRet.GetRetCode())

	metrics, _ := c.LoadTestMetrics(ctx, &taylorzhtestpb.LoadTestMetricsRequest{})
	jsonStr, _ := json.Marshal(metrics)
	log.Printf("%+v", string(jsonStr))

}
