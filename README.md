## Usage

Create a `config.json`, for example:

```json
{
    "bearer_token": "Your bearer token",
    "resource_id": "1234",
    "workspace_id": "1234",
    "zone_Id": 1234
}
```

The bearer token can be extracted from any request that contains the
`Authorization` request header.

The rest needs to be configured depending on where you want to book.
However, right now, you still have to find these out manually as well.
I also don't want to publish internal data here ;)

Then run:

```sh
go run main.go
```

Note that golang 1.19 is required.
