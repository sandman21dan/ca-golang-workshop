package demo

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestRuneDemo(t *testing.T) {
	t.Run("Rune demo", func(t *testing.T) {
		name := "ダニエル"
		fmt.Println("Length of ", name, " is ", len(name))
		fmt.Println("Rune length of ", name, " is ", utf8.RuneCountInString(name))
		i := 0
		for _, r := range name {
			fmt.Println("Rune as value: ", r)
			fmt.Printf("Rune as Unicode point %U\n", r)
			fmt.Println("Rune as string", string(r))
			i++
		}
		fmt.Println("Number of iterations", i)
		t.Fail()
	})
}

func TestMap(t *testing.T) {
	t.Run("Map demo", func(t *testing.T) {
		codeMap := map[string]string{}

		codeMap["a"] = "234sdgsdfjb"
		codeMap["b"] = "123483sfgjkgsd"

		fmt.Println(codeMap)
		fmt.Println(codeMap["a"])
		fmt.Printf(`"%v"\n`, codeMap["z"])

		t.Fail()
	})
}
