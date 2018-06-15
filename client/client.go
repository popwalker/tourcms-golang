package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/popwalker/tourcms-golang/utils"
)

// Client is a tourCms client which include some method to connect to tourCms API
type Client struct {
	MarketPlaceId string
	PrivateApiKey string
}

// BaseURI is a base URI for tourCms
var BaseURI = "https://api.tourcms.com"

// NewTourCmsClient return a tourCms client.
// Can be used to connect to tourCms API
func NewTourCmsClient(marketPlaceID, apiKey, baseURI string) *Client {
	if baseURI != "" {
		BaseURI = baseURI
	}

	return &Client{
		MarketPlaceId: marketPlaceID,
		PrivateApiKey: apiKey,
	}
}

// Request will create a http client and request to tourCms API
func (c *Client) Request(path string, method string, channelID string, postData []byte) ([]byte, error) {
	if channelID == "" {
		channelID = "0"
	}

	outboundTime := fmt.Sprintf("%d", time.Now().Unix())
	signature := c.GenerateSignature(path, method, channelID, outboundTime)

	requestUrl := BaseURI + path

	client := &http.Client{}
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, requestUrl, nil)
	} else {
		req, err = http.NewRequest(method, requestUrl, bytes.NewReader(postData))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml;charset=\"utf-8\"")
	req.Header.Set("Date", utils.GetHTTPDateTime())
	req.Header.Set("Authorization", "TourCMS "+channelID+":"+c.MarketPlaceId+":"+signature)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateSignature generate a signature.will be used for every request
func (c *Client) GenerateSignature(path, method, channel, outboundTime string) string {
	strToSign := strings.Trim(channel+"/"+c.MarketPlaceId+"/"+method+"/"+outboundTime+path, " ")
	dig := utils.Hmac(c.PrivateApiKey, strToSign)
	res := base64.StdEncoding.EncodeToString(dig)
	return url.QueryEscape(res)
}

// APIRateLimit Check remaining limit for API requests
func (c *Client) APIRateLimit(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ListChannels List all connected channels
func (c *Client) ListChannels(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ShowChannel Return channel information
func (c *Client) ShowChannel(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ChannelPerformance List top 50 channels by number of unique visitor clicks
func (c *Client) ChannelPerformance(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// SearchTour Create a search results page, a page listing all tours or a category page.
func (c *Client) SearchTour(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ShowTour Return full information about a specific Tour.
func (c *Client) ShowTour(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ListTourLocation Retrieve a list of tour locations.
func (c *Client) ListTourLocation(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ListProductFilters Product Filters are used to group tours. Use this API to find out which tours are in which Product Filters.
func (c *Client) ListProductFilters(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ShowTourDatesAndDeals Retrieve dates and deals on a particular tour
func (c *Client) ShowTourDatesAndDeals(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// UpdateTourUrl Update the tour url on an existing Tour.
func (c *Client) UpdateTourUrl(path, channelID string, reqData []byte) ([]byte, error) {
	return c.UpdateTour(path, channelID, reqData)
}

// UpdateTour Update the details on an existing Tour.
func (c *Client) UpdateTour(path, channelID string, reqData []byte) ([]byte, error) {
	return c.Request(path, http.MethodPost, channelID, reqData)
}

// SearchHotelsSpecific Create a hotel/chalet/villa search results page based on availability for specific dates
func (c *Client) SearchHotelsSpecific(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// SearchRawDepartures Get a raw departure data list for a particular Tour
func (c *Client) SearchRawDepartures(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ShowDeparture Show details of a specific tour departure, including bookings on that departure
func (c *Client) ShowDeparture(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// CreateDeparture Create a new departure on a Tour
func (c *Client) CreateDeparture(path, channelID string, data []byte) ([]byte, error) {
	return c.Request(path, http.MethodPost, channelID, data)
}

// UpdateDeparture Create a new departure on a Tour
func (c *Client) UpdateDeparture(path, channelID string, data []byte) ([]byte, error) {
	return c.Request(path, http.MethodPost, channelID, data)
}

// DeleteDeparture Delete a departure from a Tour.
func (c *Client) DeleteDeparture(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ListTours List of tours
func (c *Client) ListTours(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ListTourImages List of tours including their image URLs.
func (c *Client) ListTourImages(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}

// ShowTourDepartures Retrieve Departures for a particular tour
func (c *Client) ShowTourDepartures(path, channelID string) ([]byte, error) {
	return c.Request(path, http.MethodGet, channelID, nil)
}
