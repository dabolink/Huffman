package tree

import (
	"math"
	"testing"
)

func TestNode_String(t *testing.T) {
	tests := []struct {
		name string
		node Node
		want string
	}{
		{
			name: "single character",
			node: Node{
				probability: 0.1,
				Parent:      nil,
				Left:        nil,
				Right:       nil,
				Value:       rune("e"[0]),
				Name:        "e",
			},
			want: "e : 0.10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_getProbability(t *testing.T) {
	tree := GenerateTree("1223334444")
	tests := []struct {
		name string
		node Node
		want float64
	}{
		{
			name: "leaf node (1)",
			node: *tree.Root.Left.Left.Left,
			want: 0.1,
		},
		{
			name: "stem (1,2)",
			node: *tree.Root.Left.Left,
			want: 0.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.getProbability(); math.Abs(got-tt.want) > 0.0000001 {
				t.Errorf("Node.getProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_has(t *testing.T) {
	nodeToFind := Node{
		probability: 0.1,
		Parent:      &Node{},
		Left:        &Node{},
		Right:       &Node{},
		Value:       rune("e"[0]),
		Name:        "e : 0.10",
	}
	type args struct {
		nodes   []Node
		current Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "does not have node",
			args: args{
				nodes: []Node{
					{
						probability: 0.2,
						Parent:      &Node{},
						Left:        &Node{},
						Right:       &Node{},
						Value:       rune("e"[0]),
						Name:        "e : 0.20",
					},
				},
				current: nodeToFind,
			},
			want: false,
		},
		{
			name: "has node",
			args: args{
				nodes: []Node{
					{
						probability: 0.2,
						Parent:      &Node{},
						Left:        &Node{},
						Right:       &Node{},
						Value:       rune("e"[0]),
						Name:        "e : 0.20",
					},
					nodeToFind,
				},
				current: nodeToFind,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := has(tt.args.nodes, tt.args.current); got != tt.want {
				t.Errorf("has() = %v, want %v", got, tt.want)
			}
		})
	}
}
