package hard

/*
https://leetcode.com/problems/lru-cache
*/

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	head     *Node
	tail     *Node
	size     int
	capacity int
	indexMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{nil, nil, 0, capacity, map[int]*Node{}}
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.indexMap[key]; ok {
		this.delete(node)
		this.addToFront(node)
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.indexMap[key]; ok {
		node.Val = value
		this.delete(node)
		this.addToFront(node)
	} else {
		node = &Node{key, value, nil, nil}
		this.addToFront(node)
		this.indexMap[key] = node
		this.size++

		if this.size > this.capacity {
			delete(this.indexMap, this.tail.Key)
			this.removeFromTail()
			this.size--
		}
	}

	return
}

func (this *LRUCache) delete(node *Node) {
	prev := node.Prev
	next := node.Next
	if prev != nil {
		prev.Next = next
	} else {
		this.head = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		this.tail = prev
	}
}

func (this *LRUCache) addToFront(node *Node) {
	node.Prev = nil
	if this.head == nil {
		node.Next = nil
		this.head = node
		this.tail = node
		return
	}
	node.Next = this.head
	this.head.Prev = node
	this.head = node
}

func (this *LRUCache) removeFromTail() {
	if this.tail == nil {
		return
	}
	this.delete(this.tail)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
