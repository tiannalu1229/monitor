package main

import (
	"fmt"
	"security/core"
)

func main() {
	c := core.GetTokenSecurityTokenSniffer("0xac7945a20c3a6629749f40ad57aAC77aD9C60fEd")
	fmt.Println(c.Score)
}
