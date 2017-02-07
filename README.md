# explosm-rss

The missing image rss feed for `https://explosm.net`. This service downloads the rss feed from explosm.net and replaces the description of the content with the image instead of a link.

Add <https://explosm-1311.appspot.com> to your rss news feed now!

## Installation

`go get github.com/daniellowtw/explosm-rss`

### On appengine

This software is compatible with [`AppEngine`](https://cloud.google.com/appengine/docs/go/quickstart).

To use with `AppEngine`:

* Make a directory with necessary `app.yaml` file (or use the one provided)
* Clone this repo inside that directory
* Download the [Go SDK](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go) and run `goapp serve` to make sure it is working
* Upload it with `appcfg.py`
** `python appcfg.py update "<path-to-appengine-folder>" -A explosm-1311 -V <version-number>``

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