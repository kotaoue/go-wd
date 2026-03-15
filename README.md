# go-wd

[![Go](https://github.com/kotaoue/go-wd/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kotaoue/go-wd/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kotaoue/go-wd)](https://goreportcard.com/report/github.com/kotaoue/go-wd)
[![License](https://img.shields.io/github/license/kotaoue/go-wd)](https://github.com/kotaoue/go-wd/blob/main/LICENSE)

Get the same working directory path at `go run` and after `go build`.

## Installation

```bash
go get github.com/kotaoue/go-wd
```

## Import

```go
import "github.com/kotaoue/go-wd"
```

## Usage

| Function | Description |
| ---- | ---- |
| `wd.Get() (string, error)` | Return the working directory path consistently between `go run` and compiled binaries |
| `wd.FullPath(path string) (string, error)` | Return the full path by joining the working directory with the given relative path |

## Examples

### Get the working directory

```go
package main

import (
	"fmt"

	"github.com/kotaoue/go-wd"
)

func main() {
	dir, err := wd.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(dir)
}
```

Output:

```
/path/to/your/project
```

### Get the full path of a file

```go
package main

import (
	"fmt"

	"github.com/kotaoue/go-wd"
)

func main() {
	path, err := wd.FullPath("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(path)
}
```

Output:

```
/path/to/your/project/config.json
```
