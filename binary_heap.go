package go_binary_heap

// 节点接口
type NodeType interface {
	GetPriority() int
	GetData() interface{}
}

type HeapType int

var (
	HeapType_MaxBinaryHeap HeapType = 1
	HeapType_MinBinaryHeap HeapType = 2
)

type BinaryHeap struct {
	Elements []NodeType
	HeapType HeapType
}

// 插入新节点
func (h *BinaryHeap) Push(node NodeType) {
	h.Elements = append(h.Elements, node)
	h.up(len(h.Elements) - 1)
}

// 移除某位置节点并返回
func (h *BinaryHeap) Remove(i int) NodeType {
	n := len(h.Elements) - 1
	if n != i {
		h.swap(i, n)       // 第i个节点与最后一个节点交换
		if !h.down(i, n) { // 下沉原来的最后一个节点，也就是现在第i个节点，形成新的二叉堆
			h.up(i)
		}
	}
	// 删除最后一个元素
	res := h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]
	return res
}

// 返回某位置的节点
func (h *BinaryHeap) Get(i int) NodeType {
	return h.Elements[i]
}

// 重新构建二叉堆
func (h *BinaryHeap) Rebuild() {
	n := len(h.Elements)
	// 最后一个节点下表为n-1，因此其父节点为(n-1-1)/2，即n/2-1
	for i := n/2 - 1; i >= 0; i-- { // 最底下一排不用下沉了
		h.down(i, n)
	}
}

func (h *BinaryHeap) less(i, j int) bool {
	if h.HeapType == HeapType_MaxBinaryHeap {
		return h.Elements[i].GetPriority() > h.Elements[j].GetPriority()
	} else if h.HeapType == HeapType_MinBinaryHeap {
		return h.Elements[i].GetPriority() < h.Elements[j].GetPriority()
	} else {
		panic("heap type error")
	}
}

func (h *BinaryHeap) swap(i, j int) {
	h.Elements[i], h.Elements[j] = h.Elements[j], h.Elements[i]
}

// 节点上浮到最上面
func (h BinaryHeap) up(j int) {
	for {
		i := (j - 1) / 2             // 父节点下标
		if i == j || !h.less(j, i) { // 如果该节点不小于父节点，则无法再上浮，退出
			break
		}
		h.swap(i, j) // 上浮
		j = i        // 继续上浮
	}
}

// 节点下沉到最下面
// i0 要下沉节点的下标
func (h BinaryHeap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1          // 左子节点下标
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		// j2为右子节点下标，j为最小的子节点的下标
		j := j1
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) { // 右子节点优先级小于左子节点
			j = j2 // j设置为右子节点下标
		}
		if !h.less(j, i) { // 如果最小的子节点小于父节点，则交换这两个节点
			break
		}
		h.swap(i, j)
		i = j // 当前节点设置为最小子节点，让其继续下沉
	}
	return i > i0
}

func NewMaxHeap(elements []NodeType, heapType HeapType) BinaryHeap {
	maxHeap := BinaryHeap{
		Elements: elements,
		HeapType: heapType,
	}
	maxHeap.Rebuild()
	return maxHeap
}
