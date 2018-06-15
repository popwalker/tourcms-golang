# tourcms-golang

[![Build Status](https://travis-ci.org/popwalker/tourcms-golang.svg?branch=master)](https://travis-ci.org/popwalker/tourcms-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/popwalker/tourcms-golang)](https://goreportcard.com/report/github.com/popwalker/tourcms-golang)
[![codecov](https://codecov.io/gh/popwalker/tourcms-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/popwalker/tourcms-golang)

Go wrapper for the [TourCMS API](http://www.tourcms.com/support/api/mp/)

# Install
```shell
go get github.com/popwalker/tourcms-golang
```

# Usage
```golang
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
	tourID        = "1"                    // tour_id
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
```

> for more usage, please see [examples](https://github.com/popwalker/tourcms-golang/blob/master/example/example.go)

# List of methods included in this package

- APIRateLimit
- ListChannels
- ShowChannel
- ChannelPerformance
- SearcSearchTourhTour
- ShowTour
- ListTourLocation
- ListProductFilters
- ShowTourDatesAndDeals
- UpdateTourUrl
- SearchHotelsSpecific
- SearchRawDepartures
- ShowDeparture
- CreateDeparture
- UpdateDeparture
- DeleteDeparture
- ListTours
- ListTourImages
- ShowTourDepartures

> Other API call methods are to be added

# Dependencies

None. Tested with golang 1.10.2

# Copyright

Copyright (c) 2018 Rick.