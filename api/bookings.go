package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	EndpointMultiDayBooking = "multipleDayBooking"
	EndpointEarlyRelease    = "user/booking/%d/early-release"
	EndpointCheckIn         = "workspaces/%s/checkIn"
)

type postBookingsBody struct {
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

type CheckIn struct {
	Internal bool `json:"isInternal"`

	BookingId string `json:"bookingId"`
	// WorkspaceId is a building, for example "Loft"
	WorkspaceId string `json:"workspaceId"`
	// ResourceId is a subpart of a building, for example "Loft Meetingpoint"
	ResourceId string `json:"resourceId"`
}

func (api *API) CreateBooking(booking *Booking) ([]byte, error) {
	return api.CreateBookings([]*Booking{booking})
}

func (api *API) CreateBookings(bookings []*Booking) ([]byte, error) {
	requestBodyBytes, err := json.Marshal(postBookingsBody{Bookings: bookings})
	if err != nil {
		return nil, err
	}

	url, err := url.JoinPath(BaseURL, APIVersion, EndpointMultiDayBooking)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", api.bearerToken)
	request.Header.Add("Content-Length", fmt.Sprintf("%d", len(requestBodyBytes)))
	injectDefaultHeaders(request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func (api *API) CheckIn(checkIn *CheckIn) ([]byte, error) {
	requestBodyBytes, err := json.Marshal(checkIn)
	if err != nil {
		return nil, err
	}

	url, err := url.JoinPath(BaseURL, APIVersion, fmt.Sprintf(EndpointCheckIn, checkIn.WorkspaceId))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewReader(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", api.bearerToken)
	request.Header.Add("Content-Length", fmt.Sprintf("%d", len(requestBodyBytes)))
	injectDefaultHeaders(request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
