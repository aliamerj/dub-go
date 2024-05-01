package main

import (
	"fmt"

	"github.com/aliamerj/dub-go/pkg/dub"
	"github.com/aliamerj/dub-go/pkg/links"
)

func main() {
	client := dub.NewConfig("DUB_API_KEY", "ws_clugls3...")
	opts, _ := client.Links.Create(links.RequestOptions{URL: "https://www.worknify.com"})
	fmt.Printf("Link created successfully: %+v\n", opts)
}
