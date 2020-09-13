package main

import (
	"fmt"

	"github.com/zyhcool/go-demo/morestrings"

	"github.com/google/go-cmp/cmp"
	"github.com/zyhcool/go-demo/sayhello"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	sayhello.SayIt()
}
