package demo

import (
	"fmt"
	"strconv"
	"time"
)

type Item struct {
	key string `json:"key"`

	ID    int    `json:"ID"`
	Label string `json:"Label"`

	prev, next *Item
}

type LRU struct {
	tail     *Item //tail pointers to the last item,so when list capacity is full, it's next will be nil
	head     *Item //head doesn't store any data,next Item is the first data for use
	capacity int
	biteMap  map[string]*Item //key for query,value(*Item) for underly linklist operation
}

func NewItem(id int, name, label string) *Item {
	return &Item{
		ID:    id,
		key:   name,
		Label: label,
	}
}

func (l *LRU) move2Head(item *Item) {
	item.prev.next = item.next
	item.next.prev = item.prev //connect prev and next node

	item.next = l.head.next
	l.head.next.prev = item
	item.prev = l.head

	l.head.next = item
}

func (l *LRU) add2Tail(item *Item) {
	item.prev = l.tail
	l.tail.next = item
	l.tail = item
}

func (l *LRU) push2HeadPopTail(item *Item) {
	l.tail.prev.next = l.tail.next
	l.tail = l.tail.prev

	item.next = l.head.next
	l.head.next.prev = item
	item.prev = l.head
	l.head.next = item
}

func (l *LRU) Get(name string) *Item {
	if v, ok := l.biteMap[name]; ok {
		l.move2Head(v)
		return v
	}
	return nil
}

func (l *LRU) Traverl() []string {
	item := l.head.next
	s := make([]string, 0, l.capacity)
	for item != l.tail.next {
		s = append(s, item.key)
		item = item.next
	}

	fmt.Printf("s: %+v\n", s)

	return s
}

func (l *LRU) Set(name string, item *Item) {
	if v := l.Get(name); v != nil {
		v.Label = item.Label
		v.ID = item.ID
		l.move2Head(v)
		return
	}

	if len(l.biteMap) < l.capacity {
		l.add2Tail(item)
		l.biteMap[name] = item
		return
	}
	l.push2HeadPopTail(item)
	l.biteMap[name] = item
	delete(l.biteMap, name)
}

func InitLRU(capacity int, firstItem *Item) *LRU {
	l := &LRU{
		capacity: capacity,
		head:     NewItem(0, "head", "head"),
		tail:     NewItem(-1, "tail", "tail"),
		biteMap:  make(map[string]*Item, capacity),
	}
	firstItem.prev = l.head
	l.head.next = firstItem
	l.tail = firstItem
	l.biteMap[firstItem.key] = firstItem
	return l
}

func LRUDemo(i int) int {
	println("//<<-------------------------LRUDemo start-----------")
	start := time.Now()
	lru := InitLRU(10, NewItem(0, "0", "0"))
	const count = 10
	for i := 1; i < count; i++ {
		n := strconv.Itoa(i)
		lru.Traverl()
		lru.Set(n, NewItem(i, n, n))
	}
	lru.Traverl()
	lru.Set("3", NewItem(3, "a", "aa"))
	lru.Set("4", NewItem(3, "a", "aa"))
	lru.Set("5", NewItem(3, "a", "aa"))
	lru.Traverl()
	lru.Get("3")
	lru.Traverl()
	lru.Get("5")
	lru.Traverl()
	lru.Set("a", NewItem(13, "a", "aa"))
	lru.Traverl()
	lru.Set("b", NewItem(14, "b", "bb"))
	lru.Traverl()
	fmt.Printf("LRUDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------LRUDemo end----------->>")
	return 1
}
