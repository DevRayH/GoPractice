package main

import "fmt"

func main(){
	// Entry point
	var nums [4]int = [4]int{3, 6, 7, 9}
	var list List = List{nil, nil}

	for x := 0; x < 4; x++{
		list.push(nums[x])
	}

	for x := 0; x < 4; x++{
		var node Node = list.pop()
		fmt.Println(node.val)
	}
}