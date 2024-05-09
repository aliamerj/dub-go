package main

import (
	"fmt"
	"os"

	"github.com/aliamerj/dub-go/pkg/analytics"
	"github.com/aliamerj/dub-go/pkg/domains"
	"github.com/aliamerj/dub-go/pkg/dub"
	"github.com/aliamerj/dub-go/pkg/links"
	"github.com/aliamerj/dub-go/pkg/workspace"
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
	clicks, err := client.Analytics.Clicks(&analytics.Options{
		LinkId: "linkId", Interval: analytics.Interval30Day})
	if err != nil {
		fmt.Printf("Error retrieving link clicks: %v\n", err)
		return
	}
	fmt.Printf("Total clicks: %d\n", clicks)

	// Example of retrieving timeseries analytics
	timeseriesData, err := client.Analytics.Timeseries(&analytics.Options{
		Interval: analytics.Interval30Day,
	})
	if err != nil {
		fmt.Printf("Failed to retrieve timeseries data: %v\n", err)
	} else {
		fmt.Println("Timeseries data retrieved successfully:", timeseriesData)
	}

	// Example of retrieving analytics by country
	countryData, err := client.Analytics.Countries(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve country data: %v\n", err)
	} else {
		fmt.Println("Country data retrieved successfully:", countryData)
	}

	// Example of retrieving analytics by city
	cityData, err := client.Analytics.Cities(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve city data: %v\n", err)
	} else {
		fmt.Println("City data retrieved successfully:", cityData)
	}

	// Example of retrieving analytics by device
	deviceData, err := client.Analytics.Devices(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve device data: %v\n", err)
	} else {
		fmt.Println("Device data retrieved successfully:", deviceData)
	}

	// Example of retrieving analytics by browser
	browserData, err := client.Analytics.Browsers(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve browser data: %v\n", err)
	} else {
		fmt.Println("Browser data retrieved successfully:", browserData)
	}

	// Example of retrieving OS analytics
	osData, err := client.Analytics.OS(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve OS data: %v\n", err)
	} else {
		fmt.Println("OS data retrieved successfully:", osData)
	}

	// Example of retrieving referer analytics
	refererData, err := client.Analytics.Referers(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve referer data: %v\n", err)
	} else {
		fmt.Println("Referer data retrieved successfully:", refererData)
	}

	// Example of retrieving top links analytics
	topLinkData, err := client.Analytics.TopLinks(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve top links: %v\n", err)
	} else {
		fmt.Println("Top links retrieved successfully:", topLinkData)
	}

	// Example of retrieving top URLs analytics
	topUrlsData, err := client.Analytics.TopUrls(nil)
	if err != nil {
		fmt.Printf("Failed to retrieve top URLs: %v\n", err)
	} else {
		fmt.Println("Top URLs retrieved successfully:", topUrlsData)
	}

	// Example of creating new workspace
	workspace, err := client.Workspace.Create(workspace.Options{Name: "work", Slug: "dub-slug"})
	if err != nil {
		fmt.Printf("Failed to Create new workspace: %v\n", err)
	}
	fmt.Println("New workspace created : " + workspace.Name)

	// Example of retrieving workspace by ID
	myWorkSpace, err := client.Workspace.Get("Workspace ID")
	if err != nil {
		fmt.Printf("Failed to retrieve workspace: %v\n", err)
		return
	}
	fmt.Println("workspace retrieved successfully : " + myWorkSpace.Name)

	// Example of retrieving all workspaces
	workspaces, err := client.Workspace.List()
	if err != nil {
		fmt.Printf("Failed to retrieve workspaces: %v\n", err)
		return
	}
	for _, workspace := range workspaces {
		fmt.Println("workspace :" + workspace.Name)
	}
	//Add a domain to the authenticated workspace.
	domain, err := client.Domains.Add(domains.Options{Slug: "acme.com"})
	if err != nil {
		fmt.Printf("Failed to Add domain: %v\n", err)
	}
	fmt.Println("New domain added : " + domain.ID)

	// Edit a domain for the authenticated workspace.
	updatedDomain, err := client.Domains.Update("test", domains.Options{Slug: "Xacme.com"})
	if err != nil {
		fmt.Printf("Failed to upadate domain: %v\n", err)
		return
	}
	fmt.Println("new Domain: " + updatedDomain.Slug)

	// Delete a domain from a workspace
	deletedDomain, err := client.Domains.Delete("Xacme.com")
	if err != nil {
		fmt.Printf("Failed to delete domain: %v\n", err)
		return
	}
	fmt.Println("Domain deleted successfully: " + deletedDomain.Slug)

	// Retrieve a list of domains associated with the authenticated workspace.
	domains, err := client.Domains.List()
	if err != nil {
		fmt.Printf("Failed to delete domain: %v\n", err)
		return
	}

	for _, domain := range domains {
		fmt.Println("domain :" + domain.Slug)
	}

	// Set a domain as primary for the authenticated workspace.
	primaryDomain, err := client.Domains.SetPrimary("acme.com")
	if err != nil {
		fmt.Printf("Failed to set domain to be primary: %v\n", err)
		return
	}
	fmt.Println("Domain set to be primary successfully: " + primaryDomain.Slug)

	// Transfer a domain to another workspace
	transferDomain, err := client.Domains.Transfer("acme.com", "New_workspace_ID")
	if err != nil {
		fmt.Printf("Failed to Transfer domain : %v\n", err)
		return
	}
	fmt.Println("Domain Transfer successfully: " + transferDomain.Slug)

}
