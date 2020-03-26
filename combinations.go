package combinations

import (
	"context"
)

func Combinations(ctx context.Context, set []interface{}, repeat int) chan []interface{} {
	ch := make(chan []interface{})

	go func() {
		defer close(ch)

		if repeat > len(set) {
			return
		}

		pos := make([]int, repeat)
		for i := 0; i < repeat; i++ {
			pos[i] = i
		}

		for {
			if ctx != nil {
				select {
				case <-ctx.Done():
					return
				default:
				}
			}

			pair := make([]interface{}, repeat)
			for i, p := range pos {
				pair[i] = set[p]
			}
			ch <- pair

			idx, ok := func() (int, bool) {
				for i := len(pos) - 1; i > -1; i-- {
					if pos[i] != i+len(set)-repeat {
						return i, true
					}
				}
				return 0, false
			}()

			if !ok {
				return
			}

			pos[idx]++
			for i := idx + 1; i < repeat; i++ {
				pos[i] = pos[i-1] + 1
			}
		}
	}()

	return ch
}

func Multicombinations(ctx context.Context, set []interface{}, repeat int) chan []interface{} {
	ch := make(chan []interface{})

	go func() {
		defer close(ch)

		pos := make([]int, repeat)

		for {
			if ctx != nil {
				select {
				case <-ctx.Done():
					return
				default:
				}
			}

			pair := make([]interface{}, repeat)
			for i, p := range pos {
				pair[repeat-i-1] = set[p]
			}
			ch <- pair

			ok := func() bool {
				for i := 0; i < repeat; i++ {
					if pos[i] != len(set)-1 {
						return true
					}
				}
				return false
			}()

			if !ok {
				break
			}

			for i := 0; i < repeat; i++ {
				pos[i]++
				if pos[i] != len(set) {
					break
				}
				for j := 0; j < i+1; j++ {
					pos[j] = pos[i+1] + 1
				}
			}
		}
	}()

	return ch
}