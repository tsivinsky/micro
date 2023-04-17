package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct{}

type SumArgs struct {
	X int
	Y int
}

func (ms *MathService) Sum(args *SumArgs, reply *int) error {
	*reply = args.X + args.Y

	return nil
}

func main() {
	rpc.RegisterName("MathService", &MathService{})

	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}
