# rassemble-go
[![CI Status](https://github.com/itchyny/rassemble-go/workflows/CI/badge.svg)](https://github.com/itchyny/rassemble-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/itchyny/rassemble-go)](https://goreportcard.com/report/github.com/itchyny/rassemble-go)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/itchyny/rassemble-go/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/itchyny/rassemble-go/all.svg)](https://github.com/itchyny/rassemble-go/releases)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/itchyny/rassemble-go)](https://pkg.go.dev/github.com/itchyny/rassemble-go)

**This package is still in its early developing status!**

### Go implementation of [Regexp::Assemble](https://metacpan.org/pod/Regexp::Assemble)
```go
package main

import (
	"fmt"
	"log"

	"github.com/itchyny/rassemble-go"
)

func main() {
	xs, err := rassemble.Join([]string{"abc", "ab", "acbd", "abe"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xs) // a(?:b[ce]?|cbd)
}
```

## Bug Tracker
Report bug at [Issues・itchyny/rassemble-go - GitHub](https://github.com/itchyny/rassemble-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
