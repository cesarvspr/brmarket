# brmarket

Package to get indexes info from B3

### Import

```go
go get github.com/cesarvspr/brmarket@V1.1.0
```

### Usage

```go
import "go get github.com/cesarvspr/brmarket"
```

#### Indexes available

  IBOV
  IBXX
  IBXL
  IBRA
  IGCX
  ITAG
  IGNM
  IGCT
  IDIV
  MLCX
  IVBX
  ICO2
  ISEE
  ICON
  IEEX
  IFNC
  IMOB
  INDX
  IMAT
  UTIL
  IFIX
  BDRX

- ### main.go

```go
package main

import (
 "fmt"

 "github.com/cesarvspr/brmarket"
)

func main() {
 data, _ := brmarket.GetSharesForIndex(brmarket.IBOV)
 fmt.Println(data.UnderlyingList)
}

```
