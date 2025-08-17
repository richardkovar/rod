// Package main ...
package main

import (
	"fmt"

	"github.com/richardkovar/rod/lib/launcher"
	"github.com/richardkovar/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
