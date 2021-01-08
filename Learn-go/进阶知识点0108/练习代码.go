package main

import (
	"fmt"
	"strings"
)

//CountStr 统计字符串中出现的单词个数
func CountStr(str string) map[string]int {
	m1 := map[string]int{}
	strc := strings.Fields(str)
	fmt.Print(strc)
	for _, data := range strc {
		if _, has := m1[data]; has == true {
			m1[data]++
		} else {
			m1[data] = 1
		}
	}
	return m1
}
func main() {
	str := "i love my work and and and  i love my family too"
	m := CountStr(str)
	fmt.Println(m)
}
