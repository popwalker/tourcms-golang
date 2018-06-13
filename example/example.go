package example

import (
	"fmt"
	"encoding/xml"
	"github.com/popwalker/tourcms-golang/client"
)

var (
	marketPlaceID = "your_marketplace_id"  // please change this field to your real marketPlaceID before test
	privateAPIKey = "your_private_api_key" // please change this field to your read APIKey before test
	channelID     = "3930"                 // this is a test channel
	tourID        = "1"                    // tourID
)

// ListChannels Check remaining limit for API requests
func ListChannels() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	body, err := c.ListChannels("/p/channels/list.xml", channelID)
	if err != nil {
		return fmt.Errorf("request list channel API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// APIRateLimit Check remaining limit for API requests
func APIRateLimit() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	body, err := c.APIRateLimit("/api/rate_limit_status.xml", channelID)
	if err != nil {
		return fmt.Errorf("request rate limit API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ShowChannel Return channel information
func ShowChannel() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	body, err := c.ShowChannel("/c/channel/show.xml", channelID)
	if err != nil {
		return fmt.Errorf("request show channel API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ChannelPerformance List top 50 channels by number of unique visitor clicks
func ChannelPerformance() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	var path = ""
	if channelID == "0" {
		path = "/p/channels/performance.xml"
	} else {
		path = "/c/channel/performance.xml"
	}
	body, err := c.ChannelPerformance(path, channelID)
	if err != nil {
		return fmt.Errorf("request channel performance API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// SearchTour Create a search results page, a page listing all tours or a category page.
func SearchTour() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	var path string
	if channelID == "0" {
		path = "/p/tours/search.xml"
	} else {
		path = "/c/tours/search.xml"
	}

	params := "lat=56.82127&long=-6.09139&k=walking"

	path += "?" + params

	body, err := c.SearchTour(path, channelID)
	if err != nil {
		return fmt.Errorf("request search tour API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ShowTour Return full information about a specific Tour.
func ShowTour() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	path := "/c/tour/show.xml?id=" + tourID
	var params = "show_offers=1"
	if params != "" {
		path += "&" + params
	} else {
		path += "&show_options=1"
	}

	body, err := c.ShowTour(path, channelID)
	if err != nil {
		return fmt.Errorf("request show tour API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ListTourLocation Retrieve a list of tour locations.
func ListTourLocation() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	var params = "lang=en"
	var path string
	if channelID == "0" {
		path = "/p/tours/locations.xml?" + params
	} else {
		path = "/c/tours/locations.xml?" + params
	}

	body, err := c.ListTourLocation(path, channelID)
	if err != nil {
		return fmt.Errorf("request list tour location API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ListProductFilters Product Filters are used to group tours. Use this API to find out which tours are in which Product Filters.
func ListProductFilters() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	body, err := c.ListProductFilters("/c/tours/filters.xml", channelID)
	if err != nil {
		return fmt.Errorf("request list product filters API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ShowTourDatesAndDeals Retrieve dates and deals on a particular tour
func ShowTourDatesAndDeals() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	var queryString = ""
	path := "/c/tour/datesprices/datesndeals/search.xml?id=" + tourID + "&" + queryString

	body, err := c.ShowTourDatesAndDeals(path, channelID)
	if err != nil {
		return fmt.Errorf("request tour_datesdeals_show API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// UpdateTourUrl Update the tour url on an existing Tour.
func UpdateTourUrl() error {
	type tour struct {
		TourID  string `xml:"tour_id"`
		TourUrl string `xml:"tour_url"`
	}

	reqData := tour{TourID: tourID, TourUrl: "/tours/1_example_tour/"}
	b, err := xml.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal tour data failed, err:%+v", err)
	}

	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	body, err := c.UpdateTourUrl("/c/tour/update.xml", channelID, b)
	if err != nil {
		return fmt.Errorf("request update_tour_url API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// SearchHotelsSpecific Create a hotel/chalet/villa search results page based on availability for specific dates
func SearchHotelsSpecific() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	parameters := "startdate_yyyymmdd=2018-11-01&hdur=7&ad=2&ch=0"

	var path string
	if channelID == "0" {
		path = "/p/hotels/search_avail.xml?" + parameters + "&single_tour_id=" + tourID
	} else {
		path = "/c/hotels/search_avail.xml?" + parameters + "&single_tour_id=" + tourID
	}

	body, err := c.SearchHotelsSpecific(path, channelID)
	if err != nil {
		return fmt.Errorf("request update_tour_url API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// SearchRawDepartures Get a raw departure data list for a particular Tour
func SearchRawDepartures() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	// set some queryString. like:start_date_start=2019-01-01
	var queryString = ""
	path := "/c/tour/datesprices/dep/manage/search.xml?id=" + tourID + "&" + queryString

	body, err := c.SearchRawDepartures(path, channelID)
	if err != nil {
		return fmt.Errorf("request search_raw_departures API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ShowDeparture Show details of a specific tour departure, including bookings on that departure
func ShowDeparture() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	departureID := "3191"
	path := "/c/tour/datesprices/dep/manage/show.xml?id=" + tourID + "&departure_id=" + departureID

	body, err := c.ShowDeparture(path, channelID)
	if err != nil {
		return fmt.Errorf("request show_departure API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// CreateDeparture Create a new departure on a Tour
func CreateDeparture() error {
	type departure struct {
		TourID    string `xml:"tour_id"`
		StartDate string `xml:"start_date"`
		EndDate   string `xml:"end_date"`
		Code      string `xml:"code"`
	}

	reqData := departure{
		TourID:    tourID,
		StartDate: "2019-11-22",
		EndDate:   "2019-11-22",
		Code:      "13:00-16:00",
	}

	b, err := xml.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal request data failed, err:%+v", err)
	}

	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	body, err := c.CreateDeparture("/c/tour/datesprices/dep/manage/new.xml", channelID, b)
	if err != nil {
		return fmt.Errorf("request create_departure API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// UpdateDeparture Create a new departure on a Tour
func UpdateDeparture() error {
	type departure struct {
		TourID    string `xml:"tour_id"`
		StartDate string `xml:"start_date"`
		EndDate   string `xml:"end_date"`
		Code      string `xml:"code"`
	}

	reqData := departure{
		TourID:    tourID,
		StartDate: "2019-11-22",
		EndDate:   "2019-11-22",
		Code:      "13:30-16:30",
	}

	b, err := xml.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal request data failed, err:%+v", err)
	}

	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	body, err := c.UpdateDeparture("/c/tour/datesprices/dep/manage/update.xml", channelID, b)
	if err != nil {
		return fmt.Errorf("request update_departure API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// DeleteDeparture Delete a departure from a Tour
func DeleteDeparture() error {
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	departureID := "3191"
	path := "/c/tour/datesprices/dep/manage/delete.xml?id=" + tourID + "&departure_id=" + departureID

	body, err := c.DeleteDeparture(path, channelID)
	if err != nil {
		return fmt.Errorf("request delete_departure API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ListTours List of tours
func ListTours() error {
	var path string
	var queryString string
	if channelID == "0" {
		path = "/p/tours/list.xml?" + queryString
	} else {
		path = "/c/tours/list.xml?" + queryString
	}
	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	body, err := c.ListTours(path, channelID)
	if err != nil {
		return fmt.Errorf("request list_tours API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ListTourImages List of tours including their image URLs.
func ListTourImages() error {
	var path, queryString string
	if channelID == "0" {
		path = "/p/tours/images/list.xml?" + queryString
	} else {
		path = "/c/tours/images/list.xml?" + queryString
	}

	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")

	body, err := c.ListTourImages(path, channelID)
	if err != nil {
		return fmt.Errorf("request list_tour_images API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}

// ShowTourDepartures Retrieve Departures for a particular tour
func ShowTourDepartures() error {
	var queryString string
	path := "/c/tour/datesprices/dep/show.xml?id=" + tourID + "&" + queryString

	c := client.NewTourCmsClient(marketPlaceID, privateAPIKey, "")
	body, err := c.ShowTourDepartures(path, channelID)
	if err != nil {
		return fmt.Errorf("request show_tour_departures API failed, err:%+v", err)
	}
	fmt.Println(string(body))
	return nil
}
