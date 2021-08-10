package tree

import (
	"reflect"
	"testing"
)

func TestGenerateTree(t *testing.T) {
	rootNode1 := Node{
		probability: 1,
		Parent:      nil,
		Left:        nil,
		Right:       nil,
		Value:       rune("a"[0]),
		Name:        "'a'",
	}
	testTree1 := Tree{
		Root: &rootNode1,
		CharNodeMap: map[rune]*Node{
			rune("a"[0]): &rootNode1,
		},
	}

	leafA := Node{
		probability: 0.5,
		Parent:      nil,
		Left:        nil,
		Right:       nil,
		Value:       rune("a"[0]),
		Name:        "'a'",
	}
	leafB := Node{
		probability: 0.5,
		Parent:      nil,
		Left:        nil,
		Right:       nil,
		Value:       rune("b"[0]),
		Name:        "'b'",
	}
	rootNode2 := Node{
		probability: 1.0,
		Parent:      nil,
		Left:        &leafA,
		Right:       &leafB,
		Value:       0,
		Name:        "'a' + 'b'",
	}
	leafA.Parent = &rootNode2
	leafB.Parent = &rootNode2

	testTree2 := Tree{
		Root: &rootNode2,
		CharNodeMap: map[rune]*Node{
			rune("a"[0]): &leafA,
			rune("b"[0]): &leafB,
		},
	}
	type args struct {
		strToEncode string
	}
	tests := []struct {
		name string
		args args
		want Tree
	}{
		{
			name: "empty string",
			args: args{
				strToEncode: "",
			},
			want: Tree{
				Root:        nil,
				CharNodeMap: map[rune]*Node{},
			},
		},
		{
			name: "one char string",
			args: args{
				strToEncode: "a",
			},
			want: testTree1,
		},
		{
			name: "two char string",
			args: args{
				strToEncode: "ab",
			},
			want: testTree2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateTree(tt.args.strToEncode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateBranch(t *testing.T) {
	node1 := Node{
		probability: 0.3,
		Parent:      nil,
		Left:        nil,
		Right:       nil,
		Value:       0,
		Name:        "a",
	}
	node2 := Node{
		probability: 0.7,
		Parent:      nil,
		Left:        nil,
		Right:       nil,
		Value:       0,
		Name:        "b",
	}
	type args struct {
		node1 *Node
		node2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Generate Branch",
			args: args{
				node1: &node1,
				node2: &node2,
			},
			want: &Node{
				probability: 1.0,
				Parent:      nil,
				Left:        &node1,
				Right:       &node2,
				Value:       0,
				Name:        "a + b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateBranch(tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Encode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		tree Tree
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Encode(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_getOutput(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		tree Tree
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.getOutput(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.getOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Decode(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		tree Tree
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.Decode(tt.args.bytes); got != tt.want {
				t.Errorf("Tree.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_DFS(t *testing.T) {
	tests := []struct {
		name string
		tree Tree
		want []Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tree.DFS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
