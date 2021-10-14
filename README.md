# Go Collections for generics worlds

## Usages
### `Sets`
```go
package main

import "github.com/hienduyph/collections/sets"

func main() {
	setStr := sets.NewWithCap[string](10)
	setStr.Add("my")
	setStr.Has("my") // true
}
```

## Development

Install `gotip` first: https://pkg.go.dev/golang.org/dl/gotip
```bash 
go install golang.org/dl/gotip@latest
gotip download
```

**Run Test**
```bash 
gotip test -v ./...
```

