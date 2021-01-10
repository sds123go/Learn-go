package main

import (
	"fmt"
	"time"
	//"tempconv"
)

//并发练习
func Fab(a int) int {
	if a <= 1 {
		return a
	} else {
		return Fab(a-1) + Fab(a-2)
	}
}
func Warning() {
	for {
		for _,r:=range `--\|/`{
			fmt.Printf("\r%c",r)
			time.Sleep(100 * time.Millisecond)
		}
		
	}
}
func main() {
	go Warning()
	result := Fab(45)

	fmt.Println(result)
}
