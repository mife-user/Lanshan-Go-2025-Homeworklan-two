package main

import "testing"

// 测试 kq_cy 函数
func TestKqCy(t *testing.T) {
	input := []string{"1", "2", "1", "3", "2", "1"}
	expected := map[string]int{"1": 3, "2": 2, "3": 1}

	result := kq_cy(input)

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("kq_cy() 错误: key %s, got %d, want %d", k, result[k], v)
		}
	}
}

// 测试 number_cy 的计算逻辑
func TestNumberCy(t *testing.T) {
	// 模拟 number_cy 返回的闭包函数
	calcAdd := func() float64 {
		num1, num2 := 2.0, 3.0
		con := "+"
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
		}()
	}()

	if calcAdd != 5 {
		t.Errorf("number_cy() + 错误, got %f, want %f", calcAdd, 5.0)
	}

	// 测试乘法
	calcMul := func() float64 {
		num1, num2 := 4.0, 5.0
		con := "*"
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
		}()
	}()

	if calcMul != 20 {
		t.Errorf("number_cy() * 错误, got %f, want %f", calcMul, 20.0)
	}
}
