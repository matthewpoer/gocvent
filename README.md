# gocvent [![Go Report Card](https://goreportcard.com/badge/github.com/matthewpoer/gocvent)](https://goreportcard.com/report/github.com/matthewpoer/gocvent)

A Golang wrapper for some common [Cvent SOAP API](https://developers.cvent.com/documentation/soap-api/) functions.

## Thanks to gosoap
[gosoap](https://github.com/tiaguinho/gosoap) gave this wrapper a big head start. There was a lot of hacking on gosoap to make it work with Cvent's API, so you will find it forked and included with `gocvent`.

## Build your own Definitions
This package includes Golang Structures for the out-of-the-box versions of Contacts, Events, Registrations and Users CvObjects in the [`definitions/`](definitions/) directory. If you want to work with other CvObjects and/or use custom fields on these structures, you will need to build your own `struct`s using `cvent.StructGen("filePath", "CvObject")`, then use the resulting `struct`s in your Cvent API retrieve calls, e.g.

```go
var objectDef objDefs.EventRetrieveResult
err := cvent.Retrieve("Event", v, &objectDef)
if err != nil {
    fmt.Printf("TestRetrieveEvent err from cvent.Retrieve: %s", err)
}
fmt.Printf("Event Title: %s\n", objectDef.CvObject.EventTitle)
```

## Check out the Tests
Demonstrations of each of the implemented Cvent API methods has at least one corresponding test in [gocvent_test.go](gocvent_test.go). Current implementation includes:
* [Login](https://developers.cvent.com/documentation/soap-api/call-definitions/authentication/login/)
* [DescribeCvObject](https://developers.cvent.com/documentation/soap-api/call-definitions/object-metadata-calls/describecvobject/)
* [DescribeGlobal](https://developers.cvent.com/documentation/soap-api/call-definitions/object-metadata-calls/describeglobal/)
* [Search](https://developers.cvent.com/documentation/soap-api/call-definitions/search-and-retrieve/search/) (only _AndSearch_ style filtering is currently implemented, see [CvSearchObject](https://developers.cvent.com/documentation/soap-api/object-definitions/cvsearchobject/) and [`gosoap/encode.go`](gosoap/encode.go#L76))
* [Retrieve](https://developers.cvent.com/documentation/soap-api/call-definitions/search-and-retrieve/retrieve/) (gocvent only supports retrieval of one record at a time, which is different than the Cvent's function which would support several. Pull requests welcomed.)
