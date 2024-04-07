package main

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("hello-grpc-master:8005", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc error, error info: %v", err)
		return
	}
	sampleHash := os.Args[1]
	fmt.Printf("check the data from hello-grpc, the file hash is %s", sampleHash)
	client := breif_info_query_service.NewBriefInfoQueryServiceClient(conn)
}