package main

import (
	"fmt"
	"math/rand"
	"time"
)
type point struct {
	x, y int
}
func main()  {
	var ent int
	rand.Seed(int64(time.Now().Nanosecond()))
	a := rand.Int() % 100
	fmt.Println("来玩猜数游戏吧 guessTheNumber go-version author:Leo Ma")
	win := false
	for win == false {
		fmt.Scan(&ent)
		if ent == a {
			fmt.Println("你赢了")
			win = true
		} else if (ent > a) {
			fmt.Println("太大了")
		} else {
			fmt.Println("太小了")
		}
	}
	fmt.Print("恭喜!游戏结束 ^__^")
}