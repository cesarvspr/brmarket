# brmarket
Package to get indexes info from B3

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
	data, _ := brmarket.GetSharesForIndex("IBOV")
	fmt.Println(data.UnderlyingList)
}

```
