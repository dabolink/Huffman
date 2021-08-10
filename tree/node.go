package tree

import "fmt"

type Node struct {
	probability float64
	Parent      *Node
	Left        *Node
	Right       *Node
	Value       rune
	Name        string
}

func (node Node) String() string {
	return fmt.Sprintf("%s : %.2f", node.Name, node.probability)
	//return fmt.Sprintf("%s", node.Name)
}

func (node Node) getProbability() float64 {
	if node.Left == nil && node.Right == nil {
		return node.probability
	} else {
		return node.Left.getProbability() + node.Right.getProbability()
	}
}

func has(nodes []Node, current Node) bool {
	for _, node := range nodes {
		if node.Name == current.Name {
			return true
		}
	}
	return false
}
