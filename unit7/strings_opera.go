package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println()

	s := []string{"foo", "bar", "baz"}
	fmt.Println(s)
	fmt.Println(strings.Join(s, ", "))
	fmt.Println()

	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println()

	//相当于python中字符串的乘法
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println()

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println()

	// split返回一个slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Println()

	fmt.Printf("[%q]", strings.Trim("!!! Achtung !!!", "! "))
	fmt.Println()

	fmt.Printf("Fields are: %q", strings.Fields(" foo bar baz"))

}
