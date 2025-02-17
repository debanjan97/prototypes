# Generic Pool

A simple, generic connection pool implementation in Go.

## Features

- Generic type support - can pool any type of resource
- Configurable pool size
- Built-in logging with zap
- Thread-safe connection management

## Usage

```go
import (    
    "github.com/debanjan97/pool"
    "go.uber.org/zap"
)

func main() {
    logger, err := zap.NewDevelopment()
    if err != nil {
        panic(err)
    }
    factory := func() *Resource {
        return &Resource{}
    }
    pool := pool.NewPool(10, logger, factory) // factory is a function that returns a new resource
    resource := pool.Get() // Get a resource from the pool
    defer pool.Put(resource) // Put the resource back in the pool
}
```

## Example
Check out the [DB connections pool example](../examples/db_connections_pool/main.go) to see how this pool can be used to manage database connections efficiently.







