
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	quitCh := make(chan struct{}, 0)

}



func cat(count uint64, catCh, DogCh chan) {
	for{
		if count >= 100{
			quitCh <- struct{}
		}
		atomic.AddUint64(&count, 1)
	}
	fmt.Println("cat")
	DogCh <- struct{}
}

func dog() {
	fmt.Println("dog")
}

func fish() {
	fmt.Println("fish")
}