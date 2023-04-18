package main

import (
	"fmt"
	"log"
	"math-service/xmath"
	"net/rpc"
	"time"
	"time-service/xtime"
)

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
		t := new(xtime.NowReply)
		err = timeClient.Call("TimeService.Now", "", &t)
		if err != nil {
			fmt.Printf("err in TimeService: %v\n", err)
		}

		args := &xmath.SumArgs{
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
