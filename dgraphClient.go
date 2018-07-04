package main

import (
	"log"
	// "os"
	// "io/ioutil"
	"context"

	"google.golang.org/grpc"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)




func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
	    log.Fatal(err)
	}

	clientDir, err := ioutil.TempDir("", "client_")
	if err != nil {
	    log.Fatal(err)
	}
	defer os.RemoveAll(clientDir)

	defer conn.Close()

	req := client.Req{}

	// dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	// op := &api.Operation{}

	// ctx := context.Background()

	// txn := dgraphClient.NewTxn()
	// pb, err := json.Marshal(p)


}