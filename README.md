![pkg-go-dev](https://pkg.go.dev/badge/mod/github.com/nate-trojian/goregistry)
![Go Report Card](https://goreportcard.com/badge/github.com/nate-trojian/goregistry)
![go-version](https://img.shields.io/github/go-mod/go-version/nate-trojian/goregistry) 
# goregistry
goregistry is a Go implementation of the [Registry](https://martinfowler.com/eaaCatalog/registry.html) design pattern.

## Install
```shell
go get github.com/nate-trojian/goregistry
```

## Why use a Registry?
The original problem I was looking to solve when I first made this package was trying to decode a JSON message into one of many possible message structs based on a key field.

For example:
```json
{
    "type": "key1",
    "from": "PersonA",
    "foo": 1
}
```

and

```json
{
    "type": "key2",
    "from": "PersonA",
    "bar": "baz"
}
```

are both messages a service could receive, would have to decode, and then do some operation with.  By first deserializing each into a base structure to retrieve their unique "key", it is then possible to determine which struct it should end up as.

For the above example, the base struct would look like:
```go
type base struct {
    Type string
}
```

A registry provides a centralized data structure where you can save each of the two structs under a key, which can then be accessed based on the key of in each message.


## Alternatives
### [mapstructure](https://github.com/mitchellh/mapstructure)
When I was researching different solutions for this, I came across the [mapstructure](https://github.com/mitchellh/mapstructure) package, which was also looking to solve a similar issue.  The problem I saw with this package though was that it was heavily reliant on using reflect for decoding the final structures, which is generally not performant.  I made a quick [benchmark](https://github.com/nate-trojian/mapstructure-benchmark) testing this out, and it showed that this approach was indeed slower, and more memory intensive, than just doing a two-pass JSON parsing.