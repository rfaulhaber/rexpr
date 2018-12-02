package expr

type nodeStack struct {
	arr []*Node
}
func (ns *nodeStack) Push(n *Node) {
	ns.arr = append(ns.arr, n)
}

func (ns *nodeStack) Pop() *Node {
	if len(ns.arr) > 0 {
		lastIndex := len(ns.arr) - 1

		ret := ns.arr[lastIndex]
		newArr := ns.arr[:lastIndex]

		ns.arr = newArr

		return ret
	}

	return &Node{}
}

func (ns *nodeStack) Empty() bool {
	return len(ns.arr) == 0
}

func (ns *nodeStack) Size() int {
	return len(ns.arr)
}

func (ns *nodeStack) Peek() *Node {
	return ns.arr[len(ns.arr) - 1]
}
