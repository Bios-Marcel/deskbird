package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Bios-Marcel/yagcl"
	yagcl_json "github.com/Bios-Marcel/yagcl-json"
)

var (
	BaseURL                string = "https://app.deskbird.com/api"
	APIVersion                    = "v1.1"
	EndpointMultiDaBooking        = "multipleDayBooking"
)

type PostBookingsBody struct {
	Bookings []*Booking `json:"bookings"`
}

type Booking struct {
	// StartTime represents the start of the reservation in millisecond precision.
	StartTime int64 `json:"bookingStartTime"`
	// EndTime represents the end of the reservation in millisecond precision.
	EndTime int64 `json:"bookingEndTime"`

	Internal  bool `json:"internal"`
	Anonymous bool `json:"isAnonymous"`
	DayPass   bool `json:"isDayPass"`

	// WorkspaceId is a building, for example "Loft"
	WorkspaceId string `json:"workspaceId"`
	// ResourceId is a subpart of a building, for example "Loft Meetingpoint"
	ResourceId string `json:"resourceId"`
	// ZoneId is probably the office location.
	ZoneId int `json:"zoneItemId"`
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

type Config struct {
	BearerToken string `key:"bearer_token"`

	// WorkspaceId is a building, for example "Loft"
	WorkspaceId string `key:"workspace_id"`
	// ResourceId is a subpart of a building, for example "Loft Meetingpoint"
	ResourceId string `key:"resource_id"`
	// ZoneId is probably the office location.
	ZoneId int `key:"zone_id"`
}

func main() {
	var conf Config
	err := yagcl.
		New[Config]().
		Add(yagcl_json.Source().String("config.json")).
		Parse(&conf)
	if err != nil {
		panic(err)
	}

	var body PostBookingsBody

	startDay := must(time.Parse(time.RFC3339, "2022-11-23T08:00:00Z"))
	endDay := must(time.Parse(time.RFC3339, "2022-11-25T08:00:00Z"))

	for nextStart := startDay; !nextStart.After(endDay); nextStart = nextStart.Add(24 * time.Hour) {
		nextEnd := nextStart.Add(12 * time.Hour)
		booking := Booking{
			Internal:  true,
			DayPass:   true,
			Anonymous: false,

			StartTime: nextStart.UnixMilli(),
			EndTime:   nextEnd.UnixMilli(),

			ResourceId:  conf.ResourceId,
			WorkspaceId: conf.WorkspaceId,
			ZoneId:      conf.ZoneId,
		}

		body.Bookings = append(body.Bookings, &booking)
	}

	requestBodyBytes := must(json.Marshal(body))
	fmt.Println("Request body:")
	fmt.Println(string(requestBodyBytes))

	request := must(
		http.NewRequest(
			http.MethodPost,
			must(url.JoinPath(BaseURL, APIVersion, EndpointMultiDaBooking)),
			bytes.NewReader(requestBodyBytes)))

	request.Header.Add("Authorization", conf.BearerToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", fmt.Sprintf("%d", len(requestBodyBytes)))
	request.Header.Add("Accept", "application/json, text/plain, */*")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")
	// FIXME We ain't uncmpressing rn:
	// request.Header.Add("Accept-Encoding", "gzip, deflate, br")

	// User-Agent doesn't seem to relevant for their backend, but might be smart to specify anyway.
	// request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:104.0) Gecko/20100101 Firefox/104.0")

	// Stuff i am using to try'n get the request through:
	request.Header.Add("Host", "app.deskbird.com")
	request.Header.Add("Origin", "https://app.deskbird.com")
	// FIXME Comment
	request.Header.Add("DNT", "1")
	request.Header.Add("Referer", fmt.Sprintf("https://app.deskbird.com/office/%s/bookings/dashboard/details/%s", conf.WorkspaceId, conf.ResourceId))
	request.Header.Add("Sec-Fetch-Dest", "empty")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("TE", "trailers")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response:\n\tstatus: %d\n", response.StatusCode)
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tbody: %s\n", string(responseBytes))
}
