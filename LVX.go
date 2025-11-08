package main

import (
	"errors"
	"fmt"
)

type myInt func() error
type begin_computer func() float64

// cyjj vs kqjj
var set_string string
var set_cy []string

// 作业要求的方法
func kq_cy(cy []string) map[string]int {
	cyjj_zd := make(map[string]int)
	for i := 0; i < len(cy); i++ {
		cyjj_zd[cy[i]]++
	}
	return cyjj_zd
}

// 输入
func set_in() {
	fmt.Println("输入一串数字,输入cy结束，回车输入下一个：")
	for i := 0; ; i++ {
		fmt.Println("输入数字：")
		fmt.Scanln(&set_string)
		if set_string == "cy" {
			return
		}
		set_cy = append(set_cy, set_string)
	}
}

// 功能选择
func choose_op() error {
	num := 0
	fmt.Println("选择功能：\n------------------\n1-数字出现\n2-计算机")
	fmt.Scanln(&num)
	switch num {
	case 1:
		set_in()
		fmt.Println(kq_cy(set_cy))
		return errors.New("cyjj game win")
	case 2:
		var ccyy begin_computer = number_cy()
		fmt.Println(ccyy())
		return errors.New("cyjj game win")
	default:
		return errors.New("cyjj game over")
	}
}

// 计算机
func number_cy() func() float64 {
	var num1, num2 float64
	var con string
	fmt.Println("输入计算式(数字和符号间须有空格：")
	fmt.Scanln(&num1, &con, &num2)
	return func() float64 {
		switch con {
		case "+":
			return num1 + num2
		case "-":
			return num1 - num2
		case "*":
			return num1 * num2
		case "/":
			return num1 / num2
		default:
			return 0
		}
	}
}
func main() {
	var p myInt
	p = choose_op
	fmt.Println(p())
}
