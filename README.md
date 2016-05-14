# explosm-rss

The missing image rss feed for `https://explosm.net`. This service downloads the rss feed from explosm.net and replaces the description of the content with the image instead of a link.

Add <https://explosm-1311.appspot.com> to your rss news feed now!

## Installation

`go get github.com/daniellowtw/explosm-rss`

## Usage

```
go install
./explosm-rss
```

Go to `http://localhost:20480`

### Configuration

* `port` - the port that the server is listening on
* `refresh_interval` - how often to poll the actual feed from `explosm.net`

## License

[MIT License](http://choosealicense.com/licenses/mit/)