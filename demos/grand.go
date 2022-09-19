package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/grand"
)

func main() {

	for a := 1; a < 20; a++ {
		j := grand.Intn(15)
		m := grand.Intn(15 - j)
		s := grand.Intn(15 - j - m)
		h := grand.Intn(15 - j - m - s)
		t := 15 - j - m - s - h
		fmt.Println(j, m, s, h, t)

	}

}
