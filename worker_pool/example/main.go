package main

import (
	"fmt"

	wpool "github.com/rasoro/go-concurrency-patterns/worker_pool"
)

func main() {
	cancel, in, out := wpool.Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		in <- wpool.WorkRequest{Op: wpool.Hash, Text: []byte(fmt.Sprintf("messages %d", i))}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- wpool.WorkRequest{Op: wpool.Compare, Text: res.Wr.Text, Compare: res.Result}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string: \"%s\"; matched: %v\n", string(res.Wr.Text), res.Matched)
	}
}
