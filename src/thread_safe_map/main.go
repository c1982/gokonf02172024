package main

import (
	"fmt"
	"sync"
)

func main() {
	var list = struct {
		sync.RWMutex
		m map[string]int
	}{m: map[string]int{"a": 1, "b": 2, "c": 3}}

	go func() {
		list.RLock()
		println("rlock a: ", list.m["a"])
		list.RUnlock()
	}()

	go func() {
		list.Lock()
		list.m["a"] = 5
		list.Unlock()
	}()

	fmt.Scanf("...")
}
