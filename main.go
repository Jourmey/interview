package main

type node struct {
	left  *node
	right *node
}

//func main() {
//	var a, b, c, d node
//
//	a.left = &b
//	a.right = &c
//	b.left = &d
//
//	result := make(map[int]int, 0)
//	traverse(&a, 1, result)
//
//	fmt.Println(result)
//}

func traverse(node *node, level int, result map[int]int) {
	if node == nil {
		return
	}
	result[level] ++

	traverse(node.left, level+1, result)
	traverse(node.right, level+1, result)
}
