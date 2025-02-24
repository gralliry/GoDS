# Array

## Usage

```go
package main

import (
	"fmt"
	"github.com/gralliry/gods/array/array"
)

func main() {
	arr := array.New[int]()
	arr.Append(1)
	arr.Append(2)
	arr.Append(3)
	fmt.Println(arr.Get(0), arr.Get(1), arr.Get(2))
}

```