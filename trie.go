package aho_corasick

/*
   @Author: orbit-w
   @File: trie
   @2023 10月 周二 18:56
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

func (ins *Node) Root() bool {
	return ins.isRoot
}

func (ins *Node) Exist() bool {
	return len(ins.output) > 0
}

func (ins *Node) Code() int {
	return int(ins.code)
}

type Trie struct {
	count int32
	root  *Node
}

func (ins *Trie) Build(ks StrKeySlice) {
	ins.root = &Node{
		isRoot: true,
		depth:  0,
		index:  IndexRoot,
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
				father.children = append(father.children, child)
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

func (ins *Trie) DFS(iter func(node *Node) bool) {
	stack := NodeStack{}
	stack.Push(ins.root)
	for stack.Length() > 0 {
		head := stack.Pop()
		stop := iter(head)
		if stop {
			break
		}
		for i := len(head.children) - 1; i >= 0; i-- {
			child := head.children[i]
			stack.Push(child)
		}
	}
}

func (ins *Trie) Free() {
	ins.root = nil
}
