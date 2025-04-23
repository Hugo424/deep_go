package main

type OrderedMap struct {
	root *node
	size int
}

type node struct {
	key   int
	value int
	left  *node
	right *node
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	if m == nil {
		return
	}

	if m.root == nil {
		m.root = &node{
			key:   key,
			value: value,
		}
		m.size++
		return
	}

	currNode := m.root

	for currNode != nil {
		if currNode.key == key {
			return
		}

		if key < currNode.key {
			if currNode.left == nil {
				currNode.left = &node{
					key:   key,
					value: value,
				}
				m.size++
				return
			}
			currNode = currNode.left
		}

		if key > currNode.key {
			if currNode.right == nil {
				currNode.right = &node{
					key:   key,
					value: value,
				}
				m.size++
				return
			}
			currNode = currNode.right
		}
	}
}

func (m *OrderedMap) Erase(key int) {
	if m == nil {
		return
	}

	currNode := m.root
	var prevNode *node

	for currNode != nil && currNode.key != key {
		prevNode = currNode
		if key < currNode.key {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}

	if currNode == nil {
		return
	}

	if currNode.right == nil {
		if prevNode == nil {
			*m.root = *currNode.left
		} else if prevNode.left == currNode {
			prevNode.left = currNode.left
		} else {
			prevNode.right = currNode.left
		}
		m.size--
		return
	}

	leftMostNode := currNode.right
	var leftMostNodePrev *node

	for leftMostNode.left != nil {
		leftMostNodePrev = leftMostNode
		leftMostNode = leftMostNode.left
	}

	currNode.key = leftMostNode.key
	currNode.value = leftMostNode.value

	if leftMostNodePrev != nil {
		leftMostNodePrev.left = leftMostNode.right
	} else {
		currNode.right = leftMostNode.right
	}
	m.size--

}

func (m *OrderedMap) Contains(key int) bool {
	if m == nil {
		return false
	}

	currNode := m.root

	for currNode != nil {
		if currNode.key == key {
			return true
		} else if currNode.key < key {
			currNode = currNode.right
		} else {
			currNode = currNode.left
		}
	}

	return false
}

func (m *OrderedMap) Size() int {
	if m == nil {
		return 0
	}

	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m == nil {
		return
	}

	currNode := m.root

	var forEachNode func(n *node)
	forEachNode = func(n *node) {
		if n == nil {
			return
		}
		forEachNode(n.left)
		action(n.key, n.value)
		forEachNode(n.right)
	}

	forEachNode(currNode)
}
