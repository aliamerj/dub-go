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
	"os"

	"github.com/aliamerj/dub-go/pkg/analytics"
	"github.com/aliamerj/dub-go/pkg/dub"
	"github.com/aliamerj/dub-go/pkg/links"
)

func main() {
	// Initialize the Dub client configuration with your API token and workspace ID.
	client := dub.NewConfig("YOUR_API_TOKEN", "YOUR_WORKSPACE_ID")

	// Create a new link.
	createOpts, err := client.Links.Create(links.Options{
		URL:    "https://www.google.com",
		Domain: "dub.sh",
	})
	if err != nil {
		fmt.Printf("Failed to create link: %v\n", err)
		return
	}
	fmt.Printf("Link created successfully: %s\n", createOpts.ShortLink)

	// Retrieve a link by domain and key.
	getOpts, err := client.Links.Get(links.GetOptions{
		Domain: "dub.sh",
		Key:    "RLTVEzV",
	})
	if err != nil {
		fmt.Printf("Failed to fetch links: %v\n", err)
		return
	}
	fmt.Printf("Link fetched successfully: %s\n", getOpts.URL)

	// Update a link's information.
	updateOpts, err := client.Links.Update("linkId", links.Options{Domain: "example.com"})
	if err != nil {
		fmt.Printf("Failed to update link: %v\n", err)
		return
	}
	fmt.Printf("Link updated successfully: %v\n", updateOpts)

	// Delete a link by ID.
	if _, err := client.Links.Delete("linkId"); err != nil {
		fmt.Printf("Failed to delete link: %v\n", err)
		return
	}
	fmt.Println("Link deleted successfully.")

	// List links with sorting options.
	listOpts, err := client.Links.List(&links.GetListOptions{Sort: links.Clicks})
	if err != nil {
		fmt.Printf("Failed to list links: %v\n", err)
		return
	}
	fmt.Printf("Links listed: %v\n", listOpts)

	// Count the number of links.
	count, err := client.Links.Count(nil)
	if err != nil {
		fmt.Printf("Failed to count links: %v\n", err)
		return
	}
	fmt.Printf("Total links: %d\n", count)

	// Create multiple links in bulk.
	bulkLinks := []links.Options{
		{URL: "https://github.com/aliamerj/dub-go"},
		{URL: "https://github.com/aliamerj"},
	}
	bulk, err := client.Links.CreateMany(bulkLinks)
	if err != nil {
		fmt.Printf("Failed to create bulk links: %v\n", err)
		return
	}
	fmt.Printf("Bulk links created: %v\n", bulk)

	// Retrieve QR code.
	qr, err := client.QR.Get(nil)
	if err != nil {
		fmt.Printf("Error retrieving QR code: %v\n", err)
		return
	}
	if err := os.WriteFile("qr_code.png", qr, 0644); err != nil {
		fmt.Printf("Error saving QR code image: %v\n", err)
		return
	}
	fmt.Println("QR code saved successfully.")

	// Retrieve link click analytics.
	clicks, err := client.Analytics.Clicks(&analytics.Options{LinkId: "linkId"})
	if err != nil {
		fmt.Printf("Error retrieving link clicks: %v\n", err)
		return
	}
	fmt.Printf("Total clicks: %d\n", clicks)
}
```

