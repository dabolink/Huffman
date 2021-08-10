package tree

import (
	"fmt"
	"sort"
)

type Tree struct {
	Root        *Node
	CharNodeMap map[rune]*Node
}

func GenerateTree(strToEncode string) Tree {
	//generate heatmap of values
	numOfChar := map[rune]int{}
	for _, char := range strToEncode {
		numOfChar[char]++
	}

	//generate list of leaf nodes
	leafs := []*Node{}
	charNodeMap := map[rune]*Node{}

	for key, value := range numOfChar {
		node := &Node{
			probability: float64(value) / float64(len(strToEncode)),
			Value:       key,
			Name:        fmt.Sprintf("'%s'", string(key)),
		}
		leafs = append(leafs, node)
		charNodeMap[key] = node
	}
	for len(leafs) > 1 {
		sort.SliceStable(leafs,
			func(i, j int) bool {
				return numOfChar[leafs[i].Value] < numOfChar[leafs[j].Value]
			},
		)
		newNode := generateBranch(leafs[0], leafs[1])
		leafs = leafs[2:]
		leafs = append(leafs, newNode)
	}
	if len(leafs) == 1 {
		return Tree{
			Root:        leafs[0],
			CharNodeMap: charNodeMap,
		}
	} else {
		return Tree{
			Root:        nil,
			CharNodeMap: charNodeMap,
		}
	}

}

func generateBranch(node1 *Node, node2 *Node) *Node {
	// returns parent node with the two nodes attached
	parent := &Node{
		probability: node1.getProbability() + node2.getProbability(),
		Name:        fmt.Sprintf("%s + %s", node1.Name, node2.Name),
		Left:        node1,
		Right:       node2,
	}
	node1.Parent = parent
	node2.Parent = parent
	return parent
}

func (tree Tree) Encode(str string) []byte {
	bytes := []byte{}
	for _, r := range str {
		b := tree.getOutput(r)
		bytes = append(bytes, b...)
	}
	return bytes
}

func (tree Tree) getOutput(r rune) []byte {
	bytes := []byte{}

	currentNode := tree.CharNodeMap[r]
	path := []*Node{}
	for currentNode != nil {
		path = append(path, currentNode)
		currentNode = currentNode.Parent
	}

	//iterate reverse of list
	for i := len(path) - 1; i >= 1; i-- {
		currentNode := path[i]
		nextNode := path[i-1]
		if currentNode.Left == nextNode {
			bytes = append(bytes, 0)
		} else if currentNode.Right == nextNode {
			bytes = append(bytes, 1)
		} else {
			fmt.Println("ERROR")
			break
		}
	}
	return bytes
}

func (tree Tree) Decode(bytes []byte) string {
	outString := ""
	currentNode := tree.Root
	var nextNode *Node
	for _, b := range bytes {
		if b == 0 {
			nextNode = currentNode.Left
		} else {
			nextNode = currentNode.Right
		}

		if nextNode.Value != 0 {
			outString += string(nextNode.Value)
			currentNode = tree.Root
		} else {
			currentNode = nextNode
		}
	}
	return outString
}

func (tree Tree) DFS() []Node {
	stack := []Node{*tree.Root}
	nodes := []Node{*tree.Root}
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]
		if !has(nodes, current) {
			nodes = append(nodes, current)
		}
		if current.Left != nil {
			stack = append(stack, *current.Left)
		}
		if current.Right != nil {
			stack = append(stack, *current.Right)
		}

	}
	return nodes
}
