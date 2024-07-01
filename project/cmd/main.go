package main

import (
	"context"
	"fmt"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"path/filepath"
)

const spicedbEndpoint = `localhost:50051`

func main() {
	client, err := authzed.NewClient(spicedbEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	setupSchema(client)
	setupOrg(client)
}

func setupSchema(client *authzed.Client) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}

	schemaFile := filepath.Join(wd, "cmd", "schema.zed")
	content, err := os.ReadFile(schemaFile)

	if err != nil {
		log.Fatalf("failed to read schema file: %v", err)
	}

	schema := string(content)
	fmt.Printf("schema: %s\n", schema)

	request := &pb.WriteSchemaRequest{
		Schema: schema,
	}
	response, err := client.WriteSchema(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to write schema: %v", err)
	}
	fmt.Printf("response: %v\n", response)
}

func setupOrg(client *authzed.Client) {

}
