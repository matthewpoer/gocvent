# gocvent
A Golang wrapper for some common [Cvent SOAP API](https://developers.cvent.com/documentation/soap-api/) functions.

## Thanks to gosoap
[gosoap](https://github.com/tiaguinho/gosoap) gave this wrapper a big head start. There was a lot of hacking on gosoap to make it work with Cvent's API, so you will find it forked and included with `gocvent`.

## Check out the Tests
Demonstrations of each of the implemented Cvent API methods has at least one corresponding test in [gocvent_test.go](gocvent_test.go). Current implementation includes:
* [Login](https://developers.cvent.com/documentation/soap-api/call-definitions/authentication/login/)
* [DescribeCvObject](https://developers.cvent.com/documentation/soap-api/call-definitions/object-metadata-calls/describecvobject/)
* [DescribeGlobal](https://developers.cvent.com/documentation/soap-api/call-definitions/object-metadata-calls/describeglobal/)
* [Search](https://developers.cvent.com/documentation/soap-api/call-definitions/search-and-retrieve/search/) (only _AndSearch_ style filtering is currently implemented, see [CvSearchObject](https://developers.cvent.com/documentation/soap-api/object-definitions/cvsearchobject/) and [`gosoap/encode.go`](gosoap/encode.go#L76))