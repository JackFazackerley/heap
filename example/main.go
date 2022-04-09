package main

import (
	"fmt"
	"time"

	"github.com/JackFazackerley/heap"
)

type Photo struct {
	timestamp time.Time
}

func (p Photo) Compare() int64 {
	return p.timestamp.Unix()
}

func main() {
	h := heap.New[Photo, int64]()

	h.Push(Photo{
		timestamp: time.Date(2022, 12, 12, 12, 12, 24, 0, time.UTC),
	})
	h.Push(Photo{
		timestamp: time.Date(2022, 12, 12, 12, 12, 20, 0, time.UTC),
	})
	h.Push(Photo{
		timestamp: time.Date(2022, 12, 12, 12, 12, 16, 0, time.UTC),
	})
	h.Push(Photo{
		timestamp: time.Date(2022, 12, 12, 12, 12, 15, 0, time.UTC),
	})
	h.Push(Photo{
		timestamp: time.Date(2022, 12, 12, 12, 12, 12, 0, time.UTC),
	})

	for {
		t, err := h.Pop()
		if err != nil {
			break
		}
		fmt.Println(t.timestamp)
	}
}
