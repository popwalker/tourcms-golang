package main

import (
	"fmt"
	"github.com/popwalker/tourcms-golang/client"
	"log"
)

var (
	marketPlaceID = "your_marketplace_id"  // please change this field to your real marketPlaceID before test
	privateAPIKey = "your_private_api_key" // please change this field to your read APIKey before test
	channelID     = "3930"                 // this is a test channel
	tourID        = "1"                    // tourID
)

func main() {
	// new a TourCms Client to connect TourCms API
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	// call the APIRateLimit method to check remaining limit for API requests
	body, err := c.APIRateLimit("/api/rate_limit_status.xml", channelID)
	if err != nil {
		log.Fatalf("request rate limit API failed, err:%+v", err)
	}

	// when get the response body, do whatever you want~
	fmt.Println(string(body))
}
