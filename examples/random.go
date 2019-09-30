package main

import (
	"fmt"
	"github.com/zouyapeng/zstrings"
)

func main() {
	fmt.Println(zstrings.RandomString(64))
	fmt.Println(zstrings.RandomStringOnlyCharacters(64))
	fmt.Println(zstrings.RandomStringWithOutNumber(64))
	fmt.Println(zstrings.RandomStringWithOutUnderline(64))
	fmt.Println(zstrings.RandomCustom(64, zstrings.AllUpperCharacters+zstrings.AllNumbers+"!@#$%^&*("))
}
