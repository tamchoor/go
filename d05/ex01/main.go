package main

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left   *BinaryNode
	right  *BinaryNode
	hasToy bool
}

type BinaryTree struct {
	root *BinaryNode
}

type Trunk struct {
	prev *Trunk
	str  string
}

func addTrunk(prev *Trunk, str string) Trunk {
	var newT Trunk

	newT.prev = prev
	newT.str = str
	return newT
}

func showTrunks(p *Trunk) {
	if p == nil {
		return
	}
	showTrunks(p.prev)
	fmt.Printf(p.str)
}

func printTree(root *BinaryNode, prev *Trunk, isLeft bool) {
	if root == nil {
		return
	}
	prev_str := "    "
	trunk := addTrunk(prev, prev_str)
	printTree(root.right, &trunk, true)
	if prev == nil {
		trunk.str = "———"
	} else if isLeft == true {
		trunk.str = ".———"
		prev_str = "   |"
	} else {
		trunk.str = "`———"
		prev.str = prev_str
	}
	showTrunks(&trunk)
	if root.hasToy == true {
		fmt.Println(" 1")
	} else {
		fmt.Println(" 0")
	}
	if prev != nil {
		prev.str = prev_str
	}
	trunk.str = "   |"
	printTree(root.left, &trunk, false)
}

func (t *BinaryTree) insert(hasToy bool) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{hasToy: hasToy, left: nil, right: nil}
	} else {
		t.root.insert(hasToy)
	}
	return t
}

var iR int

func (n *BinaryNode) insert(hasToy bool) {
	if n == nil {
		return
	} else if n.left == nil {
		n.left = &BinaryNode{hasToy: hasToy, left: nil, right: nil}
	} else if n.right == nil {
		n.right = &BinaryNode{hasToy: hasToy, left: nil, right: nil}
	} else if iR == 0 {
		n.left.insert(hasToy)
		iR = 1
	} else {
		n.right.insert(hasToy)
		iR = 0
	}
}

func main() {

	var btree0 BinaryTree
	btree0.root = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.left.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.left.left.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.left.left.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.left.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.left.right.left = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.left.right.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.right.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.right.left.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.right.left.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.right.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree0.root.right.right.left = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree0.root.right.right.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	fmt.Println("Example 0")
	fmt.Println(countLevel(btree0.root, 0))
	printTree(btree0.root, nil, false)
	fmt.Println("__________")
	res := unrollGarland(btree0.root)
	fmt.Println(res)
	fmt.Println("__________")

	var btree1 BinaryTree
	btree1.root = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree1.root.left = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree1.root.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree1.root.left.left = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree1.root.left.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	fmt.Println("Example 1")
	fmt.Println(countLevel(btree1.root, 0))
	printTree(btree1.root, nil, false)
	fmt.Println("__________")
	res = unrollGarland(btree1.root)
	fmt.Println(res)
	fmt.Println("__________")

	var btree2 BinaryTree
	btree2.root = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree2.root.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree2.root.left.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree2.root.left.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree2.root.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree2.root.right.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree2.root.right.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	fmt.Println("Example 2")
	printTree(btree2.root, nil, false)
	fmt.Println("__________")
	res = unrollGarland(btree2.root)
	fmt.Println(res)
	fmt.Println("__________")

	var btree3 BinaryTree
	btree3.root = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree3.root.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree3.root.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	fmt.Println("Example 3")
	printTree(btree3.root, nil, false)
	fmt.Println("__________")
	res = unrollGarland(btree3.root)
	fmt.Println(res)
	fmt.Println("__________")

	var btree4 BinaryTree
	btree4.root = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree4.root.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree4.root.left.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree4.root.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree4.root.right.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	fmt.Println("Example 4")
	printTree(btree4.root, nil, false)
	fmt.Println("__________")
	res = unrollGarland(btree4.root)
	fmt.Println(res)
	fmt.Println("__________")

	var btree5 BinaryTree
	btree5.root = &BinaryNode{hasToy: false, left: nil, right: nil}
	fmt.Println("Example 5")
	printTree(btree5.root, nil, false)
	fmt.Println("__________")
	res = unrollGarland(btree5.root)
	fmt.Println(res)
	fmt.Println("__________")

	var btree BinaryTree
	var tmp int
	for _, err := fmt.Scanln(&tmp); err != io.EOF; _, err = fmt.Scanln(&tmp) {
		if err == nil {
			if tmp == 1 {
				btree.insert(true)
			} else if tmp == 0 {
				btree.insert(false)
			} else {
				fmt.Println("Error: it is not 1 and 0")
				os.Exit(1)
			}
		} else {
			fmt.Println("Error: it is not int\n")
			os.Exit(1)
		}
		printTree(btree.root, nil, false)
	}
	printTree(btree.root, nil, false)
	fmt.Println("__________")
	res = unrollGarland(btree.root)
	fmt.Println(res)
	fmt.Println("__________")

}

const (
	toleft  = 1
	toright = 2
)

func checkOneLevel(node1 *BinaryNode, node2 *BinaryNode, direct int) []bool {
	var res1 []bool

	if direct == toleft {
		if node2.right != nil {
			res1 = append(res1, node2.right.hasToy)
		}
		if node2.left != nil {
			res1 = append(res1, node2.left.hasToy)
		}
		if node1.right != nil {
			res1 = append(res1, node1.right.hasToy)
		}
		if node1.left != nil {
			res1 = append(res1, node1.left.hasToy)
		}
	} else if direct == toright {
		if node1.left != nil {
			res1 = append(res1, node1.left.hasToy)
		}
		if node1.right != nil {
			res1 = append(res1, node1.right.hasToy)
		}
		if node2.left != nil {
			res1 = append(res1, node2.left.hasToy)
		}
		if node2.right != nil {
			res1 = append(res1, node2.right.hasToy)
		}
	}
	return res1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func addOnelevel(node *BinaryNode, lvl, cur int) []bool {
	var res []bool
	if cur == lvl {
		res = append(res, node.hasToy)
	} else {
		if lvl%2 == 0 {
			if node.right != nil {
				res = append(res, addOnelevel(node.right, lvl, cur+1)...)
			}
			if node.left != nil {
				res = append(res, addOnelevel(node.left, lvl, cur+1)...)
			}

		} else {
			if node.left != nil {
				res = append(res, addOnelevel(node.left, lvl, cur+1)...)
			}
			if node.right != nil {
				res = append(res, addOnelevel(node.right, lvl, cur+1)...)
			}
		}
	}
	return res
}

func countLevel(node *BinaryNode, depth int) int {

	if node == nil {
		return depth
	}
	return Max(countLevel(node.left, depth+1), countLevel(node.right, depth+1))
}

func unrollGarland(root *BinaryNode) []bool {
	var res []bool
	if root == nil {
		fmt.Println("root == nil")
		return nil
	}
	countLvls := countLevel(root, 0)
	for i := 0; i < countLvls; i++ {
		res = append(res, addOnelevel(root, i, 0)...)
	}
	fmt.Println(res)
	return res
}

type myqueue struct {
	elem *BinaryNode
	next *myqueue
}

func addLast(que **myqueue, node *BinaryNode) {
	tmp := *que

	if tmp != nil {
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = &myqueue{elem: node, next: nil}
	} else {
		tmp = &myqueue{elem: node, next: nil}
	}
}

func unrollGarland1(root *BinaryNode) []bool {
	var res []bool
	if root == nil {
		fmt.Println("root == nil")
		return nil
	}
	que := &myqueue{elem: root, next: nil}
	for que != nil {
		res = append(res, que.elem.hasToy)
		if que.elem.left != nil {
			addLast(&que, que.elem.left)
		}
		if que.elem.right != nil {
			addLast(&que, que.elem.right)
		}
		que = que.next
	}
	fmt.Println(res)
	return res
}
