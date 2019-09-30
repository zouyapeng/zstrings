package zstrings

import (
	"strings"
)

var (
	AllNumbers                    = "0123456789"
	AllUpperCharacters            = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AllLowerCharacters            = strings.ToLower(AllUpperCharacters)
	Underline                     = "_"
	AllCharacters                 = AllUpperCharacters + AllLowerCharacters + AllNumbers + Underline
	AllCharactersOnly             = AllUpperCharacters + AllLowerCharacters
	AllCharactersWithOutNumbers   = AllUpperCharacters + AllLowerCharacters + Underline
	AllCharactersWithOutUnderline = AllUpperCharacters + AllLowerCharacters + AllNumbers
)
