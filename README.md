# explosm-rss

The missing image rss feed for `https://explosm.net`. This service downloads the rss feed from explosm.net and replaces the description of the content with the image instead of a link.

Add <https://explosm-1311.appspot.com> to your rss news feed now!

## Installation

`go get github.com/daniellowtw/explosm-rss`

### On appengine

This software is compatible with [`AppEngine`](https://cloud.google.com/appengine/docs/go/quickstart).

To use with `AppEngine`:

* Clone this repo.
* Follow the instructions on the [quickstart page](https://cloud.google.com/appengine/docs/standard/go/quickstart) and download the gcloud SDK.
* Deploy a version with `gcloud app deploy --no-promote` and check that the new version works. Migrate traffic to new version.

## Usage

```
go run main.go
```

Go to `http://localhost:8080`

### Configuration

* `port` - the port that the server is listening on
* `refresh_interval` - how often to poll the actual feed from `explosm.net`

## License

[MIT License](http://choosealicense.com/licenses/mit/)