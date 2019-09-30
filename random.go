package zstrings

import "math/rand"

func RandomCustom(n int64, sourceString string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = sourceString[rand.Intn(len(sourceString))]
	}
	return string(b)
}

func RandomString(n int64) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AllCharacters[rand.Intn(len(AllCharacters))]
	}
	return string(b)
}

func RandomStringWithOutUnderline(n int64) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AllCharactersWithOutUnderline[rand.Intn(len(AllCharactersWithOutUnderline))]
	}
	return string(b)
}

func RandomStringWithOutNumber(n int64) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AllCharactersWithOutNumbers[rand.Intn(len(AllCharactersWithOutNumbers))]
	}
	return string(b)
}

func RandomStringOnlyCharacters(n int64) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AllCharactersOnly[rand.Intn(len(AllCharactersOnly))]
	}
	return string(b)
}
