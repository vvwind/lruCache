// cache.go

package main

import "fmt"

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	Size     int
	Nodes    map[string]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Nodes:    make(map[string]*Node),
	}
}

func (c *LRUCache) Get(key string) string {
	if node, exists := c.Nodes[key]; exists {
		c.moveToHead(node)
		return node.Value
	}
	return ""
}

func (c *LRUCache) Set(key, value string) {
	if node, exists := c.Nodes[key]; exists {
		node.Value = value
		c.moveToHead(node)
	} else {
		newNode := &Node{Key: key, Value: value}
		c.Nodes[key] = newNode
		c.addToHead(newNode)
		c.Size++

		if c.Size > c.Capacity {
			removed := c.removeTail()
			delete(c.Nodes, removed.Key)
			c.Size--
		}
	}
}

func (c *LRUCache) Rem(key string) {
	if node, exists := c.Nodes[key]; exists {
		c.removeNode(node)
		delete(c.Nodes, key)
		c.Size--
	}
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeNode(node *Node) {
	prev := node.Prev
	next := node.Next

	if prev != nil {
		prev.Next = next
	} else {
		c.head = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		c.tail = prev
	}

	node.Prev = nil
	node.Next = nil
}

// Добавляет ноду в начало списка
func (c *LRUCache) addToHead(node *Node) {
	node.Prev = nil
	node.Next = c.head

	if c.head != nil {
		c.head.Prev = node
	} else {
		c.tail = node
	}

	c.head = node
}

// Удаляет последний элемент (tail)
func (c *LRUCache) removeTail() *Node {
	tail := c.tail
	if tail == nil {
		return nil
	}
	c.removeNode(tail)
	return tail
}

// main — для тестирования
func main() {
	cache := NewLRUCache(3)
	cache.Set("Jesse", "Pinkman")
	cache.Set("Walter", "White")
	cache.Set("Jesse", "James")

	fmt.Println(cache.Get("Jesse")) // James
	cache.Rem("Walter")
	fmt.Println(cache.Get("Walter")) // ""

	cache.Set("1", "1")
	cache.Set("2", "2")
	cache.Set("3", "3")
	cache.Set("4", "4")
	cache.Set("5", "5")

	fmt.Println(cache.Get("1"), cache.Get("2"), cache.Get("3"), cache.Get("4"), cache.Get("5"))
}
