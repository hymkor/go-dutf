go-dutf
=======

Implementation of DUTF encoder/decoder by [the programming language Go](https://go.dev)

- [DUTF, a Dynamic Unicode Transformation Format](https://www.ietf.org/id/draft-yaoyang-dutf-01.html) 
- [yaoyang-ietf/dutf](https://github.com/yaoyang-ietf/dutf)

```example.go
package main

import (
    "fmt"

    "github.com/hymkor/go-dutf"
)

func main() {
    sourceString := "your string to be encoded"
    bytes := dutf.EncodeString(sourceString)
    decodedString, _ := dutf.DecodeString(bytes)
    fmt.Println(decodedString)
}
```
