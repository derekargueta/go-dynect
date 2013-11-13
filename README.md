# go-dynect

A DynECT REST client for the Go programming language.

## Installation

	$ go get bitbucket.org/_nesv/go-dynect/dynect

## Usage

	package main

	import (
		"bitbucket.org/_nesv/go-dynect/dynect"
		"log"
	)

	func main() {
		client := dynect.NewClient("my-dyn-customer-name")
		err := client.Login("my-dyn-username", "my-dyn-password")
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			err := client.Logout()
			if err != nil {
				log.Fatal(err)
			}
		}()

		// Make a request to the API, to get a list of all, managed DNS zones
		var response ZonesResponse
		if err := client.Do("GET", "Zone", nil, &response); err != nil {
			log.Println(err)
		}

		for _, zone := range response.Data {
			log.Printf("Zone %q Serial %d", zone.Zone, zone.Serial)
		}
	}

More to come!
