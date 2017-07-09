package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/julienschmidt/httprouter"
	"github.com/wejick/poc_tego/src/random"
	"github.com/wejick/tego/config"
)

func main() {
	err := config.LoadConfigFromFile("./files/etc/poc_tego/config.json")
	if err != nil {
		log.Panicln("couldn't load config file")
	}

	router := httprouter.New()

	router.GET("/random", random.GetRandomHTTP)

	go goRPCServer()

	listenTo := config.Get().HTTP.Listen + ":" + config.Get().HTTP.Port
	log.Fatal(http.ListenAndServe(listenTo, router))
}

func goRPCServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	random.RegisterRandomServer(s, &random.RandomS{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
