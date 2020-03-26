# go-combinations

This package provides a method to generate combinations and multicombinations (like python's itertools.combinations and itertools.combinations_with_replacement) out of given slice.

## Installation

```sh
go get github.com/kyo1/go-combinations
```

## Usage

### `func Combinations(ctx context.Context, set []interface{}, repeat int) chan []interface{}`

`Combinations` function generates elements of k-combinations.

```go
package main

import (
	"context"
	"fmt"

	"github.com/kyo1/go-combinations"
)

func main() {
	chars := []interface{}{0, 1, 2, 3, 4}

	// Generate a k-combinations
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for s := range combinations.Combinations(ctx, chars, 3) {
		fmt.Println(s)

		// The condition for terminating the generator is not required
		// if condition {
		// 	cancel()
		// 	continue
		// }
	}

	// Output
	// [0 1 2]
	// [0 1 3]
	// [0 1 4]
	// [0 2 3]
	// [0 2 4]
	// [0 3 4]
	// [1 2 3]
	// [1 2 4]
	// [1 3 4]
	// [2 3 4]
}
```

### `func Multicombinations(ctx context.Context, set []interface{}, repeat int) chan []interface{}`

`Multicombinations` function generates elements of k-multicombinations.

```go
package main

import (
	"context"
	"fmt"

	"github.com/kyo1/go-combinations"
)

func main() {
	chars := []interface{}{0, 1, 2}

	// Generate a k-multicombinations
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for s := range combinations.Multicombinations(ctx, chars, 3) {
		fmt.Println(s)

		// The condition for terminating the generator is not required
		// if condition {
		// 	cancel()
		// 	continue
		// }
	}

	// Output
	// [0 0 0]
	// [0 0 1]
	// [0 0 2]
	// [0 1 1]
	// [0 1 2]
	// [0 2 2]
	// [1 1 1]
	// [1 1 2]
	// [1 2 2]
	// [2 2 2]
}
```
