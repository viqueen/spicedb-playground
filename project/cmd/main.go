package main

import (
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const spicedbEndpoint = `localhost:50051`

func main() {
	client, err := authzed.NewClient(spicedbEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	setupOrg(client)
}

func setupOrg(client *authzed.Client) {

}
