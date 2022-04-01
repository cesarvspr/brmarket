# brmarket
Package to get indexes info from B3
# Models TCStreamsCollector

## Models used on Price Feeds projects

### Import

```go
go get github.com/cesarvspr/brmarket@V1.0.0
```

### Usage

```go
import "go get github.com/cesarvspr/brmarket"
```

- ### main.go
```go
package main

import (
	"fmt"

	"github.com/cesarvspr/brmarket"
)

func main() {
	text, err := brmarket.GetSharesForIndex("IOV")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(text)
}

```
