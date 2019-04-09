package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jonathabp/grpc-leak/proto"
	"google.golang.org/grpc"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			conn, err := grpc.Dial(
				":8000",
				grpc.WithInsecure(),
				grpc.WithMaxMsgSize(1024))
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			c := proto.NewDataClient(conn)

			for i := 0; i < 1000; i++ {
				_, err := c.Get(context.Background(), &proto.DataRequest{})
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(50 * time.Millisecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
