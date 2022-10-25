package main

import (
	"github.com/Bios-Marcel/deskbird/api"
	"github.com/Bios-Marcel/deskbird/auth"
	"github.com/Bios-Marcel/yagcl"
	yagcl_json "github.com/Bios-Marcel/yagcl-json"
)

type Config struct {
	// WorkspaceId is a building, for example "Loft"
	WorkspaceId string `key:"workspace_id"`
	// ResourceId is a subpart of a building, for example "Loft Meetingpoint"
	ResourceId string `key:"resource_id"`
	// ZoneId is probably the office location.
	ZoneId int `key:"zone_id"`
}

func main() {
	bearerToken, err := auth.RetrieveBearerToken()
	if err != nil {
		panic(err)
	}

	var conf Config
	err = yagcl.
		New[Config]().
		Add(yagcl_json.Source().Path("config.json")).
		Parse(&conf)
	if err != nil {
		panic(err)
	}

	deskbirdApi := api.New(bearerToken)

	// ------------------------------
	// Check in example
	// Requires a valid booking ID

	// responseBody, err := deskbirdApi.CheckIn(&api.CheckIn{
	// 	Internal:    true,
	// 	BookingId:   "1394800",
	// 	WorkspaceId: conf.WorkspaceId,
	// 	ResourceId:  conf.ResourceId,
	// })
	// if err != nil {
	// 	log.Fatalln("error creating bookings", err)
	// }

	// log.Println("Response:\n", string(responseBody))

	// ------------------------------
	// Multi day booking example

	// startDay := must(time.Parse(time.RFC3339, "2022-11-08T09:00:00Z"))
	// endDay := must(time.Parse(time.RFC3339, "2022-11-08T09:00:00Z"))

	// var bookings []*api.Booking
	// for nextStart := startDay; !nextStart.After(endDay); nextStart = nextStart.Add(24 * time.Hour) {
	// 	// We ain't working during weekends.
	// 	if weekday := nextStart.Weekday(); weekday == time.Saturday || weekday == time.Sunday {
	// 		continue
	// 	}

	// 	nextEnd := nextStart.Add(10 * time.Hour)
	// 	booking := api.Booking{
	// 		Internal: true,
	// 		DayPass:  true,

	// 		StartTime: nextStart.UnixMilli(),
	// 		EndTime:   nextEnd.UnixMilli(),

	// 		ResourceId:  conf.ResourceId,
	// 		WorkspaceId: conf.WorkspaceId,
	// 		ZoneId:      conf.ZoneId,
	// 	}

	// 	bookings = append(bookings, &booking)
	// }

	// responseBody, err := deskbirdApi.CreateBookings(bookings)
	// if err != nil {
	// 	log.Fatalln("error creating bookings", err)
	// }

	// log.Println("Response:\n", string(responseBody))
}
