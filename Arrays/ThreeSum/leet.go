package main

import (
	"fmt"
)



func main(){
	nums := []int{-1,0,1,2,-1,-4}
   result := threeSum(nums)
   fmt.Println(result)
}

func threeSum(nums []int) [][]int {
    var result  [][]int
    for i:=0;i<=len(nums);i++{
        for j:=0;j<=len(nums);j++{
            for k:=0;k<=len(nums);k++{
                if nums[i] + nums[j] + nums[k] == 0{
                    fmt.Print(nums[i],nums[j],nums[k])
               return result

                
                } 
            }
        }
    }
    return result
}

