package main

import (
	"fmt"
)

func main(){
	nums := []int{2,7,11,15} 
	target := 9
	result := TwoSum(nums,target)
	fmt.Println(result)
}

func TwoSum(nums []int, target int) []int {
 result := make(map[int]int)
 for i :=0;i<=len(nums);i++{
	element := nums[i]
	if value,ok := result[target-element];ok{
		return []int{value,i}
	}
	result[element] = i
 }
 return nil
}

