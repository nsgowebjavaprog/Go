package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pf = fmt.Println

func main() {
	AB := 20000
	BC := strconv.Itoa(AB)
	pf(reflect.TypeOf(BC)) // string
}
