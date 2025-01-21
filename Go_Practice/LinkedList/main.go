package main

func main() {
	var nums [4]int = [4]int{3, 5, 7, 8}
	var list List = List{}
	var limit = 0

	for limit < 4 {
		list.push(nums[limit])
		limit++
	}

	list.walkList(list.Head)
}