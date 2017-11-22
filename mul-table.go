package main

import "fmt"

func main() {
	for i = 1; i <= 100; i++  {
		if i % 3 == 0 {
			fmt.Println ("foo")}
		if i % 5 == 0 {
			fmt.Println ("bar") }
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println ("foobar") }


		fmt.Println(i)
	}
}



//fmt.Println(i, "is odd")