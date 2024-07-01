package main

import (
	"fmt"
	"strings"
)

var pf = fmt.Println

func main() {
	word1 := "Hi Nagaraj"
	pf(word1)

	pf(len(word1)) // 10

	// split

	pf("Split :", strings.Split("N-a-g-a-r-a-j", "-")) // Split : [N a g a r a j]

	pf("Lower:", strings.ToLower(word1)) // Lower: hi nagaraj

	pf("Higher:", strings.ToUpper(word1)) // Higher: HI NAGARAJ

}
