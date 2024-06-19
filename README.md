# the-golang
A collection of studies about the programming languange Go based on the official book.

## Chapter 1 - Tutorial
### Hello world

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```

### CLI Args
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

