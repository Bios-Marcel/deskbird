## Features

* Book a certain time range outside of the allowed timeframe (30 days into the future)
  > Currently requiers to manually input the start and endtime; weekends and off-days aren't treated properly

## Planned Features

* Automatically Accept booking
* Proper CLI for easy of use

## Requirements

* Chrome / Chromium
* Initial manual login into deskbird within the Chromium instance

## Usage

Create a `config.json`, for example:

```json
{
    "resource_id": "1234",
    "workspace_id": "1234",
    "zone_Id": 1234
}
```

The rest needs to be configured depending on where you want to book.
However, right now, you still have to find these out manually as well.
I also don't want to publish internal data here ;)

Then run:

```sh
go run main.go
```

Note that golang 1.19 is required.
