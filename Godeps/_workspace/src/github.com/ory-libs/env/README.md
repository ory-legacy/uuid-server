# ory-libs/env

[![Build Status](https://travis-ci.org/ory-libs/env.svg)](https://travis-ci.org/ory-libs/env)

A very handy extension which adds defaults to `os.GetEnv()`.

```go
import "github.com/ory-libs/env"

func main() {
  port := env.Getenv("PORT", "80")
}
```

versus

```go
import "os"

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "80"
  }
}
```
