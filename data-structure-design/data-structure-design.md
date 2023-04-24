### 数据结构设计
#### LRU  （Least Recently Used）
    最近最少使用，认为最近使用过的数据是有用的，很久没有使用的数据应该是无用的，内存满了就优先删除哪些很久没有使用过的数据；
    
Get(key) 和 PUT(key,value) 方法都是O(1)的时间复杂度；
所以LRU cache的数据结构必要条件为：
- cache中元素必须有时序，以区分最近使用和久未使用的数据，当容量满了优先删除久未使用的数据；  （双向链表：记住head和tail）
- 要在cache中快速找到key是否存在并得到对应的val；  （哈希表可以实现）
- 每次访问cache中某个key，需要将这个元素变为最近使用的，也就是说cache要支持任意位置快速插入和删除元素； （双向链表） 

所以LRU cache的数据结构是双向链表和哈希表的结合体；
    
#### LFU （Least Frequently Used）
    淘汰哪些最少使用的数据；如果由多个数据最少使用，淘汰最久未使用的数据；
```go
type LFU struct {
	keyToValue map[int]int  //key-value 映射
	minFreq int  //最小使用数
	keyToFreq map[int]int  //key-freq 映射
	freqToKeys map[int]*freqToKeys //freq-keys 映射  ，需要使用O(1)的时间复杂度增加和删除key，切实删除最久未使用的key，使用双向链表结合map实现
	cap int
}
```
    
#### 前缀树 trie
#### 单调栈 （monotone stack）
    单调栈是一种特殊的栈，顾名思义就是栈内元素按照单调递增（或递减）排序的栈；
    单调栈主要用于以时间复杂度为O(n)来解决NGE（next greater element）问题，通俗的说就是解决类似【找出第一个大于x或者第一个小于x的元素】的问题
#### 单调队列 （monotone queue）
    队列中的元素是单调递增或递减的；
    可以解决滑动窗口内极值的问题；
    
「单调队列」的核心思路和「单调栈」类似，push 方法依然在队尾添加元素，但是要把前面比自己小或者比自己大的元素都删掉；
```go
type MonotonicQueue interface {
    func push(x interface{})  //在队尾加入x
    func pop(x interfaces{})  //队头如果是x，删除它
    func max() int            //返回当前队列中最大的值
}
```

#### 二叉堆实现优先队列 
#### 队列实现栈以及栈实现队列