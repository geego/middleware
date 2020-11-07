# middleware

Here you'll find middleware ready to use with `geen` framework, maintained by geego team.

## Install

`go get -u github.com/geego/geen`

## Examples

See [examples](https://github.com/geego/example) for a variety of examples.

```go
package main

import (
  "net/http"

  "github.com/geego/geen"
  "github.com/geego/middleware"
)

func main() {
  r := geen.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome"))
  })
  http.ListenAndServe(":3000", r)
}
```

## License

Inspired by `chi` and under [MIT License](./LICENSE)