# Dub.co Go SDK

Welcome to the Dub.co Go SDK! This toolkit is designed to make your life easier when integrating Go applications with Dub.co's APIs. Whether you're building a new project or adding to an existing one, this SDK aims to streamline your workflow with easy setup and comprehensive features.

## Features

- **Easy Integration**: Quick setup and authentication process.
- **Comprehensive Endpoint Coverage**: Access all Dub.co APIs efficiently.
- **Advanced Error Handling**: Gracefully manage errors and API changes.
- **Real-Time Data Synchronization**: Stay updated with the latest data.

## Getting Started

Follow these steps to start using the Dub.co Go SDK in your project:

### Prerequisites

- Go 1.22 or later

### Installation

Install the SDK with Go's package manager:

```bash
go get github.com/aliamerj/dub-go
```
### Usage
Here's a quick example to get you started:
```go
package main

import (
	"fmt"
	"github.com/aliamerj/dub-go/pkg/dub"
	"github.com/aliamerj/dub-go/pkg/links"
)

// The main function demonstrates how to use the Dub.co Go SDK to efficiently create and retrieve shortened links.
func main() {
	// Initialize the Dub client with your API token and workspace ID.
	// Replace "QKZ2..." and "ws_clv..." with your actual API token and workspace ID.
	client := dub.NewConfig("QKZ2...", "ws_clv...")

	// Attempt to create a new link. Specify the URL to be shortened and the domain under which it should be registered.
	createOpts, errCreated := client.Links.Create(links.CreateOptions{
		URL:    "https://www.worknify.com", // URL to shorten
		Domain: "dub.sh",                  // Domain under which the link is registered
	})

	// Handle errors that may occur during the link creation process.
	if errCreated != nil {
		fmt.Printf("Failed to create link: %+v\n", errCreated)
		return // Exit the function if an error occurs to prevent further execution.
	}

	// If the link is created successfully, output the short link.
	fmt.Printf("Link created successfully: %s\n", createOpts.ShortLink)

	// Attempt to retrieve a link by specifying its domain and key.
	getOpts, errGet := client.Links.Get(links.GetOptions{
		Domain: "dub.sh", // Domain of the link
		Key:    "RLTVEzV", // Key associated with the link
	})

	// Handle errors that may occur during the link retrieval process.
	if errGet != nil {
		fmt.Printf("Failed to fetch links: %+v\n", errGet)
		return // Exit the function if an error occurs to prevent further execution.
	}

	// If the link is retrieved successfully, output the URL.
	fmt.Printf("Links fetched successfully: %s\n", getOpts.URL)
}
	// Update link's info by link Id
	res, err := client.Links.Update("linkId", links.RequestOptions{})
	if err != nil {
		fmt.Printf("Failed to update links: %+v\n", err)
		return // Stop further execution if there's an error
	}
	fmt.Printf("Links updated successfully: %+v\n", res)

  // delete link by id 
  deleteRes, err := client.Links.Delete("link Id")
  if err != nil {
  		fmt.Printf("Failed to Delete links: %+v\n", err)
		return // Stop further execution if there's an error
  }
	fmt.Printf("Links updated successfully: %+v\n", deleteRes)
```

