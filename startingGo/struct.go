package main

import "fmt"

func main() {
	type Base struct {
		ID    int
		Owner string
	}

	type A struct {
		Base
		Name string
		Area string
	}

	type B struct {
		Base
		Title  string
		Bodies []string
	}

	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	b := B{
		Base:   Base{21, "shaoyu"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	fmt.Printf("%d\n", a.ID)
	fmt.Printf("%d\n", b.ID)

}
