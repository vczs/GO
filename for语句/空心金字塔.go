package main

import "fmt"

/**
控制台输出空心金字塔
*/
func main() {
	var num int = 4 //定义金字塔层数
	for i := 1 ; i <= num ; i++ { //控制金字塔层数
		for k := 1 ; k <= num - i ; k++ {
			fmt.Print(" ") //每层开始先打印空格数，打印（总层数-当前层数）个
		}
		//空格打印完打印*,打印个数为（2*当前层数-1）个
		//实心金字塔在for语句里打印*就可以
		for j := 1 ; j <= 2*i-1 ; j++ {
			//因为要打印空心的所以进行判断，只打印第一个和最后一个，其他打印空格
			//(i == num && j % 2 == 1)该条件是为了控制最后一层，当当前层数等于总层数并且位数是奇数时输出*
			//(i == num && j % 2 == 1)中 j % 2 == 1 条件可有可无，但是有的话输出的空心金字塔更好看些
			if j == 1 || j == 2*i-1 || (i == num && j % 2 == 1) {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println() //每层打印完换行
	}
}