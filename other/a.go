package other

type MinStack struct {
	arr []int
	len int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack MinStack
	stack.len = 0
	return stack
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
	this.len += 1
}

func (this *MinStack) Pop() {
	if this.len > 0 {
		this.arr = this.arr[:]
		this.len -= 1
	}
}

func (this *MinStack) Top() int {
	if this.len > 0 {
		return this.arr[0]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	min := 0
	for i := 0; i < this.len; i++ {
		if this.arr[i] < min {
			min = this.arr[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
