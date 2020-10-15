package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
fmt.Println(isValid("()[]{}"))
}
func twoSum(nums []int, target int) []int {
	sort.Sort(sort.IntSlice(nums[0:]))
	var lenth = len(nums)
	for ; lenth >= 0; lenth-- {
		if nums[lenth-1] < target {
			break
		}
	}
	arr := nums[0:lenth]
	a := 0
	b := len(arr) - 1
	for a < b {
		if arr[a]+arr[b] == target {
			break
		} else if arr[a]+arr[b] < target {
			a++
		} else {
			b--
		}
	}
	res := make([]int, 2)
	for i, num := range nums {
		if num == arr[a] {
			res[0] = i
		}
		if num == arr[b] {
			res[1] = i
		}
	}
	return res
}

func isValid(s string) bool {
	var stack1 []string
	var stack2 []string
	arr := strings.Split(s, "")
	for _, s2 := range arr {
		if s2 == "(" || s2 == "[" || s2 == "{"{
			stack1 = append(stack1, s2)
		}else {
			stack2 = append(stack2, s2)
		}
		if len(stack1)>0&&len(stack2)>0 {
			if isCorrect(stack1[len(stack1)-1],stack2[len(stack2)-1]) {
				stack1 = stack1[0:len(stack1)-1]
				stack2 = stack2[0:len(stack2)-1]
			}
		}
	}
	if len(stack1)==0&&len(stack2)==0 {
		return true
	}
	return false
}
func isCorrect(a,b string) bool {
	if a =="("&&b ==")"|| a =="["&&b =="]"||a =="{"&&b =="}"{
		return true
	}
	return false
}
