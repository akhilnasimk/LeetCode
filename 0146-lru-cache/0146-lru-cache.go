type LinkedNode struct {
	Key  int
	Val  int
	Prev *LinkedNode
	Next *LinkedNode
}

type LRUCache struct {
	Storage map[int]*LinkedNode
	Cap     int
	Head    *LinkedNode
	Tail    *LinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Storage: make(map[int]*LinkedNode),
		Cap:     capacity,
	}
}

func (this *LRUCache) remove(node *LinkedNode) {

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		this.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		this.Tail = node.Prev
	}

	node.Next = nil
	node.Prev = nil
}


func (this *LRUCache) insertFront(node *LinkedNode) {

	node.Next = this.Head
	node.Prev = nil

	if this.Head != nil {
		this.Head.Prev = node
	}

	this.Head = node

	if this.Tail == nil {
		this.Tail = node
	}
}

func (this *LRUCache) Get(key int) int {

	node, exists := this.Storage[key]

	if !exists {
		return -1
	}

	this.remove(node)
	this.insertFront(node)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {

	if node, exists := this.Storage[key]; exists {

		node.Val = value

		this.remove(node)
		this.insertFront(node)

		return
	}

	node := &LinkedNode{
		Key: key,
		Val: value,
	}

	this.Storage[key] = node

	this.insertFront(node)

	if len(this.Storage) > this.Cap {
		lru := this.Tail
		this.remove(lru)
		delete(this.Storage, lru.Key)
	}
}