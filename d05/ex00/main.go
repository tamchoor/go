package main

import (
	// "bufio"
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

	var btree1 BinaryTree
	btree1.root = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree1.root.left = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree1.root.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree1.root.left.left = &BinaryNode{hasToy: false, left: nil, right: nil}
	btree1.root.left.right = &BinaryNode{hasToy: true, left: nil, right: nil}
	fmt.Println("Example 1")
	printTree(btree1.root, nil, false)
	fmt.Println("__________")
	fmt.Println("areToysBalanced ", areToysBalanced(btree1.root))
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
	fmt.Println("areToysBalanced ", areToysBalanced(btree2.root))
	fmt.Println("__________")

	var btree3 BinaryTree
	btree3.root = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree3.root.left = &BinaryNode{hasToy: true, left: nil, right: nil}
	btree3.root.right = &BinaryNode{hasToy: false, left: nil, right: nil}
	fmt.Println("Example 3")
	printTree(btree3.root, nil, false)
	fmt.Println("__________")
	fmt.Println("areToysBalanced ", areToysBalanced(btree3.root))
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
	fmt.Println("areToysBalanced ", areToysBalanced(btree4.root))
	fmt.Println("__________")

	var btree5 BinaryTree
	btree5.root = &BinaryNode{hasToy: false, left: nil, right: nil}
	fmt.Println("Example 5")
	printTree(btree5.root, nil, false)
	fmt.Println("__________")
	fmt.Println("areToysBalanced ", areToysBalanced(btree5.root))
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
	fmt.Println("areToysBalanced ", areToysBalanced(btree.root))
	fmt.Println("__________")

}

func countToysInNode(node *BinaryNode) int {
	var count, toy int

	if node == nil {
		return 0
	}

	if node.hasToy == true {
		toy = 1
	}
	count = toy + countToysInNode(node.left) + countToysInNode(node.right)
	return count
}

func areToysBalanced(root *BinaryNode) bool {
	if root == nil {
		fmt.Println("root == nil")
		return false
	}
	var count1, count2 int
	count1 = countToysInNode(root.left)
	count2 = countToysInNode(root.right)
	if count1 == count2 {
		return true
	} else {
		return false
	}
}
