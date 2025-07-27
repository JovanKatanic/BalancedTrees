package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Node struct {
	value int
	left  *Node
	right *Node
	b     int
}

func (n *Node) print(depth int) {
	fmt.Print("|")
	fmt.Print(strings.Repeat("__", depth))
	fmt.Print(" ")
	if n != nil {
		fmt.Println(n.value, "(", n.b, ")")
		depth++
		n.left.print(depth)
		n.right.print(depth)
	} else {
		fmt.Println("x")
	}
}

func (n *Node) validate() (int, bool) {
	if n != nil {
		l, l_valid := n.left.validate()
		if l_valid {
			return -1, true
		}
		r, r_valid := n.right.validate()
		if r_valid || l_valid {
			return -1, true
		} else if n.b != max(r-l) {
			fmt.Printf("[ERROR] Invalid b at node %d\n", n.value)
			return -1, true
		} else if l-r > 1 || r-l > 1 {
			fmt.Printf("[ERROR] Height mismatch at node %d\n", n.value)
			return -1, true
		} else if n.left != nil && n.left.value > n.value {
			fmt.Printf("[ERROR] Left is greater than current at node %d\n", n.value)
			return -1, true
		} else if n.right != nil && n.right.value < n.value {
			fmt.Printf("[ERROR] Right is less than current at node %d\n", n.value)
			return -1, true
		}
		return max(l, r) + 1, (l_valid || r_valid)
	} else {
		return 0, false
	}
}

func (tree *BalancedTree) Validate() bool {
	_, valid := tree.root.validate()
	return valid
}

func (tree *BalancedTree) Print() {
	depth := 0
	tree.root.print(depth)
}

func BuildTree() *Node {
	node7 := &Node{value: 7}
	node6 := &Node{value: 4}
	node5 := &Node{value: 4, left: node6}
	node4 := &Node{value: 4, left: node5}
	node3 := &Node{value: 3}
	node1 := &Node{value: 1, left: node7}
	node2 := &Node{value: 2, left: node3, right: node4}
	root := &Node{value: 0, left: node1, right: node2}

	return root
}

func (node *Node) balance() *Node {
	switch node.right.b {
	case 1:
		newRoot := node.right
		temp := node.right.left
		newRoot.left = node
		node.right = temp

		newRoot.b = 0
		node.b = 0
		return newRoot
	case -1:
		newRoot := node.right.left

		right := newRoot.right
		left := newRoot.left

		newRoot.right = node.right
		newRoot.right.left = right

		newRoot.left = node
		newRoot.left.right = left

		if newRoot.b == 0 {
			newRoot.left.b = 0
			newRoot.right.b = 0
		} else if newRoot.b == 1 {
			newRoot.left.b = -1
			newRoot.right.b = 0
		} else if newRoot.b == -1 {
			newRoot.left.b = 0
			newRoot.right.b = 1
		}

		newRoot.b = 0

		return newRoot
	default:
		return node
	}
}

func (node *Node) balanceLeft() *Node {
	switch node.left.b {
	case -1:
		newRoot := node.left
		temp := node.left.right
		newRoot.right = node
		node.left = temp

		newRoot.b = 0
		node.b = 0

		return newRoot
	case 1:
		newRoot := node.left.right

		right := newRoot.right
		left := newRoot.left

		newRoot.left = node.left
		newRoot.left.right = left

		newRoot.right = node
		newRoot.right.left = right

		if newRoot.b == 0 {
			newRoot.left.b = 0
			newRoot.right.b = 0
		} else if newRoot.b == 1 {
			newRoot.left.b = -1
			newRoot.right.b = 0
		} else if newRoot.b == -1 {
			newRoot.left.b = 0
			newRoot.right.b = 1
		}

		newRoot.b = 0

		return newRoot
	default:
		//node.b--
		return node
	}

}

func (node *Node) add(val int) (int, bool) {
	if val < node.value {
		if node.left == nil {
			node.left = &Node{value: val}
			node.b--
			return node.b, false
		}
		b, valid := node.left.add(val)
		if valid {
			return node.b, true
		}
		if b <= -2 {
			node.left = node.left.balanceLeft()
			return node.b, true
		} else if b >= 2 {
			node.left = node.left.balance()
			return node.b, true
		}
		if b == 0 {
			return node.b, true
		}
		node.b--

	} else {
		if node.right == nil {
			node.right = &Node{value: val}
			node.b++
			return node.b, false
		}
		b, valid := node.right.add(val)
		if valid {
			return node.b, true
		}
		if b >= 2 {
			node.right = node.right.balance()
			return node.b, true
		} else if b <= -2 {
			node.right = node.right.balanceLeft()
			return node.b, true
		}
		if b == 0 {
			return node.b, true
		}
		node.b++
	}
	return node.b, false
}

type BalancedTree struct {
	root *Node
}

func (tree *BalancedTree) Add(val int) {
	node := tree.root
	if node == nil {
		tree.root = &Node{value: val}
		return
	}
	b, _ := node.add(val)
	if b >= 2 {
		tree.root = node.balance()
	} else if b <= -2 {
		tree.root = node.balanceLeft()
	}
	tree.Validate()
}

func main() {
	var tree BalancedTree

	rand.Seed(time.Now().UnixNano())

	n := 10000
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}

	//fmt.Println(nums)
	fmt.Println()

	for i, val := range nums {
		if i == len(nums)-1 {
			fmt.Println("start")
		}
		tree.Add(val)
	}
	//tree.Print()
}
