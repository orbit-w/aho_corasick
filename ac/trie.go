package ac

import "fmt"

/*
   @Time: 2023/8/22 00:06
   @Author: david
   @File: trie
*/

type Node struct {
	isRoot   bool
	isLeaf   bool
	code     rune
	depth    int
	index    int
	father   *Node
	fail     *Node
	children []*Node
	output   []rune
}

func (ins *Node) findChild(r rune) *Node {
	for _, child := range ins.children {
		if child.code == r {
			return child
		}
	}
	return nil
}

func (ins *Node) addChild(child *Node) {
	ins.children = append(ins.children, child)
}

func (ins *Node) Root() bool {
	return ins.isRoot
}

func (ins *Node) Exist() bool {
	return len(ins.output) > 0
}

type Trie struct {
	count int32
	root  *Node
}

func (ins *Trie) Build(ks strKeySlice) {
	ins.root = &Node{
		isRoot: true,
		depth:  0,
	}

	for _, key := range ks {
		father := ins.root
		for i, r := range key {
			child := father.findChild(r)
			if child == nil {
				child = &Node{
					code:   r,
					isLeaf: i == len(key)-1,
					depth:  father.depth + 1,
					father: father,
				}
				if child.isLeaf {
					child.output = append(child.output, rune(len(key)))
				}
				ins.count++
				father.addChild(child)
			}
			father = child
		}
	}
}

func (ins *Trie) f(node *Node) {
	fail := ins.g(node)
	if len(fail.output) > 0 {
		node.output = append(node.output, fail.output...)
	}
	node.fail = fail
}

func (ins *Trie) g(node *Node) (fail *Node) {
	for fn := node.father.fail; fn != nil; fn = fn.fail {
		if fail = fn.findChild(node.code); fail != nil {
			return
		}
	}
	fail = ins.root
	return
}

func (ins *Trie) BFS(iter func(node *Node) bool) {
	queue := make([]*Node, 0, ins.count+1)
	queue = append(queue, ins.root)
	for len(queue) > 0 {
		//pop
		head := queue[0]
		queue = queue[1:]
		stop := iter(head)
		if stop {
			break
		}

		for i := range head.children {
			child := head.children[i]
			queue = append(queue, child)
		}
	}
}

func (ins *Trie) Print() {
	ins.BFS(func(node *Node) bool {
		if node.Root() {
			return false
		}
		fmt.Print("code: ", node.code)
		fmt.Println("fail code: ", node.fail.code)
		return false
	})
}

func printTree(node *Node, lv int) {
	for i := 0; i < lv; i++ {
		fmt.Print("  ")
	}
	fmt.Println(string(node.code))
	for _, child := range node.children {
		printTree(child, lv+1)
	}
}
