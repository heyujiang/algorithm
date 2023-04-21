### 图的逻辑结构和具体实现

```一个图是由节点和边构成的```

<img src="https://labuladong.gitee.io/algo/images/%e5%9b%be/0.jpg"></img>

```go
type Vertex struct {
     id int
     neighbors []*Vertex                    
}
```

上面的实现是逻辑上的，实际上很少用这个Vertex类实现图，而是用常说的**邻接表**和**邻接矩阵**来实现；
<img src="https://labuladong.gitee.io/algo/images/%e5%9b%be/2.jpeg"></img>

#### 邻接表
    每个节点x的邻居都存在一个列表中，然后把x和这个列表关联起来，这样就可以通过一个节点x找到它的所有相邻节点；

```go
    var graph [][]int
```

#### 邻接矩阵
    邻接矩阵是一个而为布尔数组，如果x和y是相连的，那么就把matrix[x][y]设为true。如果想要找到x的邻居，去扫一圈matrix[x][...]就可以了；

```go
    var martix [][]bool
```

#### 两中存储方式的优劣：
    对于邻接表，好处是占用空间少；但是邻接表无法快速判断两个节点是否相邻；

#### 度（degree）
    在无向图中度就是每个节点相连的边的条数；
    由于有向图的边有方向，所有有向图的中每个节点的度又分为入度和出度；
    
#### 图遍历框架
```go
    var visited []bool //记录已经被遍历过的节点
    var onPath []bool //记录从起点到当前节点的路径

    func traverse(graph Grapth,s int){
        if visited[s] {
            return             
        }
        visited[s] = true  // 经过节点s，标记为已遍历
        onPath[s] = trud //做选择： 标记节点s在路径上
        for _,neighbor := range graph.neighbors(s){
            traverse(graph,neighbor)
        }
        onPath[s] = false //撤销选择：节点s离开路径
    }
```

#### 判断是否有环 
##### DFS解法
    使用visited []bool 记录节点是否访问过，onPath []bool记录访问链路，如果一个节点已经存在访问链路中则证明有环，否则无环；visited减少
    遍历次数；
##### BFS解法 
    核心是计算每个节点的入度，入度为0的节点进入遍历队列，遍历过程中实时计算遍历到的节点的入度，当入度为0时进入遍历队列，遍历结束后count值
    等于节点数量则无环，否则有环；

#### 拓扑排序
##### DFS解法
    在判断环的DFS解法的基础上增加后序遍历结构的记录，将后序遍历记录的结构倒过来就是拓扑排序；
##### BFS解法
    在判断是否有环的BFS解法的基础上，遍历的顺序就是拓扑排序的结果；
    
##### 二分图
二分图的定义：
```二分图的定点集可分割为两个互不相交的子集，图中每条边依附的两个顶点都分属于这两个子集，且两个子集内的顶点不相邻。```