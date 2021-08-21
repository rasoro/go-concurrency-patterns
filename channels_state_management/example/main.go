package main

import (
	"context"
	"fmt"

	state "github.com/rasoro/go-concurrency-patterns/channels_state_management"
)

func main() {
	in := make(chan *state.WorkRequest, 6)
	out := make(chan *state.WorkResponse, 6)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processor(ctx, in, out)

	req1 := state.WorkRequest{state.Add, 3, 4}
	in <- &req1

	req2 := state.WorkRequest{state.Subtract, 5, 3}
	in <- &req2

	req3 := state.WorkRequest{state.Multiply, 3, 3}
	in <- &req3

	req4 := state.WorkRequest{state.Divide, 8, 2}
	in <- &req4

	req5 := state.WorkRequest{state.Divide, 5, 0}
	in <- &req5

	for i := 0; i < 5; i++ {
		resp := <-out
		fmt.Printf("Request: %v; Result: %v; Error: %v\n", resp.Wr, resp.Result, resp.Err)

	}

}
