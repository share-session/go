# Share Session

A go client wrapper for creating and retrieving share sessions

## Installation

```shell
go get github.com/share-session/go
```

## Usage

### Create a session

```go
package main

import (
	"fmt"
	"github.com/share-session/go"
)

func main() {
	session, err := sharesession.Create(&sharesession.Session{
		Cookies: &sharesession.Cookies{
			"cookie-name": "cookie-value",
		},
	})

	fmt.Println(session.Id)
}
```

### Retrieve a session

```go
package main

import (
	"fmt"
	"github.com/share-session/go"
)

func main() {
	session, err := sharesession.Get("some-session-id")

	fmt.Println(session.Id)
}
```