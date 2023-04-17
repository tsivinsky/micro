package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type TimeService struct{}

type NowReply struct {
	Hours   int
	Minutes int
	Seconds int
}

func (ts *TimeService) Now(arg string, reply *NowReply) error {
	t := time.Now()

	*reply = NowReply{
		Hours:   t.Hour(),
		Minutes: t.Minute(),
		Seconds: t.Second(),
	}

	return nil
}

func main() {
	rpc.RegisterName("TimeService", &TimeService{})

	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(l, nil)
}
