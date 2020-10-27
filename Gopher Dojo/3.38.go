package main

import "fmt"

// Score is ユーザ定義型
type Score struct {
	UserID int
	GameID int
	Point  int
}

func main() {
	// 変数に割当して一部値入力
	var x Score
	x.UserID = 2
	fmt.Println(x.UserID)

}
