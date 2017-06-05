# Variable byte codes

vbc is a golang library for [Variable byte codes](https://nlp.stanford.edu/IR-book/html/htmledition/variable-byte-codes-1.html).

## Example

```go
package main

import (
	"fmt"

	"github.com/bluele/vbc"
)

func main() {
	encoded := vbc.Encode32([]uint32{
		38, 103, 157, 363, 364, 383, 480, 506, 572,
	})
	values := vbc.Decode32(encoded)

	// [38 103 157 363 364 383 480 506 572]
	fmt.Println(values)
}
```

## Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
