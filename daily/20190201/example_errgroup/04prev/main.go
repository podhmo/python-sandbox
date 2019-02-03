package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	type R struct {
		N  int
		NN int
	}
	var r R

	g, _ := errgroup.WithContext(context.Background())
	for k := 0; k < 1000; k++ {
		k := k
		g.Go(func() error {
			for i := 0; i < 10; i++ {
				fmt.Println("n", k, i)
				r.N++
				// time.Sleep(100 * time.Millisecond)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
