package main

func main() {

}

func postorderTraversal(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		path := []int{}
		for ; node != nil; node = node.Right {
			path = append(path, node.Val)
		}
		for i := len(path) - 1; i >= 0; i-- {
			res = append(res, path[i])
		}
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}
