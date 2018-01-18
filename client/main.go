package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/leighmcculloch/demo-google-cloud-appengine-twirp/rpc/geo"
)

func main() {
	client := geo.NewGeoProtobufClient(os.Args[1], &http.Client{})

	response, err := client.Get(context.Background(), &geo.GetRequest{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Country: %q\n", response.Country)
}
