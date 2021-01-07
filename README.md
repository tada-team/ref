# ref
Shortcuts for pass by reference

```go
package main

import "github.com/tada-team/ref"

type MyStruct struct {
    X *string
}

func main()  {
    print(MyStruct{X: ref.String("yyy")})
}
```
