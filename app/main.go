package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// I know i could import it from time-service module if i created package in it but it's not the point
type NowReply struct {
	Hours   int
	Minutes int
	Seconds int
}
type SumArgs struct {
	X int
	Y int
}

func main() {
	timeClient, err := rpc.DialHTTP("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}

	mathClient, err := rpc.DialHTTP("tcp", "localhost:5001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		t := new(NowReply)
		err = timeClient.Call("TimeService.Now", "", &t)
		if err != nil {
			fmt.Printf("err in TimeService: %v\n", err)
		}

		args := &SumArgs{
			X: t.Minutes,
			Y: t.Seconds,
		}

		sum := 0
		err = mathClient.Call("MathService.Sum", args, &sum)
		if err != nil {
			fmt.Printf("err in MathService: %v\n", err)
		}
		fmt.Printf("sum of %d and %d: %d\n", args.X, args.Y, sum)

		time.Sleep(1 * time.Second)
	}
}
