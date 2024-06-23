package main

import "fmt"


func binarySearch(nums []int, target int) int {
  leftIndex := 0
  rightIndex := len(nums)
  var mid int
  var midIndex int

  for leftIndex < rightIndex {
    fmt.Println(leftIndex, mid, rightIndex)
    midIndex = (leftIndex + rightIndex) / 2
    mid = nums[midIndex]

    if mid == target {
      return midIndex
    } else if target < mid {
      rightIndex = midIndex
    } else if target > mid {
      leftIndex = midIndex + 1
    }
  }
  return -1
}

func main() {
  // var testArr []int = []int{11, 234, 454, 656, 7658, 8797, 10000}
  var testArr []int = []int{-1, 0, 3, 5, 9, 12}
  target := 9
  fmt.Printf("Hellop, %v, looking for %v\n", testArr, target)
  result := binarySearch(testArr, target)
  fmt.Printf("Here's the index: %v", result)
}
