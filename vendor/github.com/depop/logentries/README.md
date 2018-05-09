# Logentries client (golang) 

### Install

```sh
$ go get -d github.com/depop/logentries
```

### Usage

Set LOGENTRIES_TOKEN environment variable with your token

```go
package main

import (
   "fmt"
   logentries "github.com/depop/logentries"
)

func main() {
  client := logentries.New{}
  resp, err := client.LogSet.Create(&logentries.LogSetCreateRequest{
  	logentries.LogSetFields{
  		Name: "foobar",
  	},
  })
  if err != nil {
  	panic(err)
  }

  fmt.Println(resp)
}
```

