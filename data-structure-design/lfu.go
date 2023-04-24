package data_structure_design

type LFU struct {
	keyToValue map[int]int         //key-value 映射
	minFreq    int                 //最小使用数
	keyToFreq  map[int]int         //key-freq 映射
	freqToKeys map[int]*freqToKeys //freq-keys 映射  ，需要使用O(1)的时间复杂度增加和删除key，切实删除最久未使用的key，使用双向链表结合map实现
	cap        int
}

func (l *LFU) Get(key int) int {
	if _, ok := l.keyToValue[key]; !ok {
		return -1
	}
	l.addKeyFreq(key)
	return l.keyToValue[key]
}

func (l *LFU) Put(key, val int) {
	if _, ok := l.keyToValue[key]; ok {
		l.keyToValue[key] = val
		l.addKeyFreq(key)
		return
	}

	if len(l.keyToValue) >= l.cap {

	}

	l.keyToValue[key] = val

}

func (l *LFU) addKeyFreq(key int) {
	curFreq := l.keyToFreq[key]
	nextFreq := curFreq + 1

	l.keyToFreq[key] = nextFreq

	curFkeys := l.freqToKeys[curFreq]
	nextKeys := l.freqToKeys[nextFreq]

	curFkeys.delete(key)
	nextKeys.put(key)
}

//==============
type element struct {
	key, val   int
	prev, next *element
}

type doubleListEle struct {
	head, tail *element
	size       int
}

func newDoubleListEle() *doubleListEle {
	head := &element{}
	tail := &element{}
	head.next = tail
	tail.prev = head
	return &doubleListEle{head: head, tail: tail, size: 0}
}

type freqToKeys struct {
	queue *doubleListEle
	km    map[int]*element
}

func newFreqToKeys() *freqToKeys {
	return &freqToKeys{
		queue: newDoubleListEle(),
		km:    make(map[int]*element, 0),
	}
}

func (f *freqToKeys) put(key int) {
	if _, ok := f.km[key]; ok {
		f.delete(key)
	}

	eld := &element{key: key, val: 0}
	f.km[key] = eld
	f.queue.tail.prev.next = eld
	eld.prev = f.queue.tail.prev
	f.queue.tail.prev = eld
	eld.next = f.queue.tail
	f.queue.size++
}

func (f *freqToKeys) delete(key int) {
	eld := f.km[key]
	delete(f.km, key)
	f.queue.size--
	eld.prev.next = eld.next
	eld.next.prev = eld.prev
}

func (f *freqToKeys) size() int {
	return f.queue.size
}
