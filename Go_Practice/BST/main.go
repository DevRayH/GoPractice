package main

func main(){
	//var intNum int = 1
	//fmt.Println(intNum)

	var nums[4]int = [4]int{5, 3, 9, 2}

	var tree = Tree{}
	var limit int = 0
	for limit < 4{
		tree.addNode(nums[limit])
		limit++
	}
	
	tree.inorder()
}