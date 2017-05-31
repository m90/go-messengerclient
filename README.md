# go-messengerclient

[![Build Status](https://travis-ci.org/m90/go-messengerclient.svg?branch=master)](https://travis-ci.org/m90/go-messengerclient)
[![godoc](https://godoc.org/github.com/m90/go-messengerclient?status.svg)](http://godoc.org/github.com/m90/go-messengerclient)

> Send messages to messenger like what?!?!

## Installation

Install using `go get`:

```sh
$go get github.com/m90/go-messengerclient
```

## Usage

Calling `New` returns a `Client` that exposes `Send(MessagePayload)`:

```go
client := msngrclient.New("my_access_token")
payload := msngrclient.NewMessagePayload(
	"recipientID",
	msgnrclient.NewTextMessage("Hello Gophers!"),
)
if err := client.Send(payload); err != nil {
	// handle err
}
```

## Tests

Run the tests:

```sh
$ make
```

### License
MIT © [Frederik Ring](http://www.frederikring.com)
