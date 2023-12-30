package api

import (
	"fmt"

	"github.com/sauravkuila/portfolio-worth/external"
)

func ReachZerodha() {
	fmt.Println("call to zerodha from here")
	external.GetMargin()
}

func ReachAngel() {
	fmt.Println("call to angel from here")
	external.GetMarginAngel()
}
