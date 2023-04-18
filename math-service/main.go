package main

import (
	"log"
	"math-service/xmath"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct{}

func (ms *MathService) Sum(args *xmath.SumArgs, reply *int) error {
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
