package data_structure_design

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type doubleList struct {
	head *node
	tail *node
	size int
}

func newDoubleList() *doubleList {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return &doubleList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (d *doubleList) addLast(x *node) {
	d.tail.prev.next = x
	x.prev = d.tail.prev
	x.next = d.tail
	d.tail.prev = x
	d.size++
}

func (d *doubleList) remove(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	d.size--
}

func (d *doubleList) removeFirst() *node {
	if d.head.next == d.tail {
		return nil
	}
	first := d.head.next
	d.remove(first)
	return first
}

func (d *doubleList) getSize() int {
	return d.size
}

type LRUCache struct {
	cache *doubleList
	km    map[int]*node
	cap   int
}

func NewLRU(cap int) *LRUCache {
	return &LRUCache{
		cache: newDoubleList(),
		km:    make(map[int]*node),
		cap:   cap,
	}
}

func (l *LRUCache) Get(key int) (int, bool) {
	//判断key是否存在缓存中
	if _, ok := l.km[key]; !ok { //if false 返回flase
		return 0, false
	}
	//if true get val 并将 给node 提升为最近使用
	l.makeRecently(key)
	return l.km[key].val, true
}

func (l *LRUCache) Put(key, val int) {
	//判断key是否存在缓存中
	if _, ok := l.km[key]; ok { //if true 先删除key 再添加
		l.deleteKey(key)
	} else {
		//if false
		if l.cache.size >= l.cap { //判断缓存容量是否充足，如果容量以达到限制，先淘汰一个key 最后在添加
			l.removeLeastRecently()
		}
	}
	l.addRecently(key, val) //添加缓存的key-val
}

func (l *LRUCache) Size() int {
	return l.cache.getSize()
}

func (l *LRUCache) makeRecently(key int) { //提升为最近使用
	x := l.km[key]
	l.cache.remove(x)
	l.cache.addLast(x)
}

func (l *LRUCache) addRecently(key, val int) { //添加最近使用
	node := &node{key: key, val: val}
	l.km[key] = node
	l.cache.addLast(node)
}

func (l *LRUCache) removeLeastRecently() { //删除最近未使用
	dx := l.cache.removeFirst()
	delete(l.km, dx.key)
}

func (l *LRUCache) deleteKey(key int) { //删除key
	x := l.km[key]
	delete(l.km, key)
	l.cache.remove(x)
}
