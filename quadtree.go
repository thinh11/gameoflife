package main



type Node struct {
	nw, ne, sw, se *Node
	level int
}

var BaseNodes [2]*Node = [2]*Node{&Node{nil, nil, nil, nil, 0}, &Node{nil, nil, nil, nil, 0}}

var Node func(nw, ne, sw, se *Node) *Node = MemoizeNode()

func Live(n *Node) int {return if n==BaseNodes[1] { return 1} return 0}

func MemoizeNode() func(nw, ne, sw, se *Node) *Node {
	NodeMemoized := make(map[[4]*Node]*Node)
	for _, n1 := range BaseNodes {
		for _, n2 := range BaseNodes {
			for _, n3 := range BaseNodes {
				for _, n4 := range BaseNodes {
					NodeMemoized[[4]*Node{n1,n2,n3,n4}] = &Node{n1,n2,n3,n4,1}
				}
			}
		}
	}
	func Node(nw, ne, sw, se *Node) *Node {
		key := [4]*Node{nw,ne,sw,se}
		if n, ok := NodeMemoized[key]; ok {
			return n
		}
		n := &Node{nw,ne,sw,se,nw.level+1}
		NodeMemoized[key] = n
		return n
	}
}

func CenterSubnode(n *Node) *Node {
	return Node(n.nw.se, n.ne.sw, n.sw.ne, n.se.nw)
}

func CenterHorizontalSubnode(w, e *Node) *Node {
	return Node(w.ne.se, e.nw.sw, w.se.ne, e.sw.nw)
}

func CenterVerticalSubnode(n, s *Node) *Node {
	return Node(n.sw.se, n.se.sw, s.nw.ne, s.ne.nw)
}

func CenterSubSubnode(n *Node) *Node {
   return Node(n.nw.se.se, n.ne.sw.sw, n.sw.ne.ne, n.se.nw.nw)
}

func NextGenSubnode(n *Node) *Node {
	if n.level == 2 { //4-by-4
		ns := []*Node{
			n.nw.nw, n.nw.ne, n.ne.nw, n.ne.ne,
			n.nw.sw, n.nw.se, n.ne.sw, n.ne.se,
			n.sw.nw, n.sw.ne, n.se.nw, n.se.ne,
			n.sw.sw, n.sw.se, n.se.sw, n.se.se }
		
	}
	n00 := CenterSubnode(n.nw)
	n01 := CenterHorizontalSubnode(n.nw, n.ne)
	n02 := CenterSubnode(n.ne)
	n10 := CenterVerticalSubnode(n.nw, n.sw)
	n11 := CenterSubSubnode(n)
	n12 := CenterVerticalSubnode(n.ne, n.se)
	n20 := CenterSubnode(n.sw)
	n21 := CenterHorizontalSubnode(n.sw, n.se)
	n22 := CenterSubnode(n.se)
	nextlevel := level-2
	return Node(
		NextGenSubnode(Node(n00, n01, n10, n11))
		NextGenSubnode(Node(n01, n02, n11, n12)),
		NextGenSubnode(Node(n10, n11, n20, n21)),
		NextGenSubnode(Node(n11, n12, n21, n22)))
	
}