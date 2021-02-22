package main

import (
	"fmt"
	"time"
)

func main() {
	mining := [7]string{"ore", "slv", "aum", "ore", "ore", "x", "ore"}
	found := make(chan string)
	collected := make(chan string)
	go func(mined [len(mining)]string) {
		for _, m := range mined {
			if m == "ore" {
				found <- m
			}
		}
	}(mining)

	go func() {
		for collecting := range found {
			fmt.Println("from finder, found: ", collecting)
			collected <- "minedore"
		}
	}()

	go func() {
		for melting := range collected {
			fmt.Println("from miner: ", melting)
			fmt.Println("from meltor: melting")
		}
	}()
	<-time.After(time.Second * 5)
}
