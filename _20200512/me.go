package main

/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

- push(x) —— 将元素 x 推入栈中。
- pop() —— 删除栈顶的元素。
- top() —— 获取栈顶元素。
- getMin() —— 检索栈中的最小元素。
*/

/**
OK
*/
type MinStackWithMe struct {
	arr []int
	len int
}

/** initialize your data structure here. */
func ConstructorWithMe() MinStackWithMe {
	var stack MinStackWithMe
	stack.len = 0
	return stack
}

func (this *MinStackWithMe) Push(x int) {
	this.arr = append(this.arr, x)
	this.len += 1
}

func (this *MinStackWithMe) Pop() {
	if this.len > 0 {
		this.arr = this.arr[0 : this.len-1]
		this.len -= 1
	}
}

func (this *MinStackWithMe) Top() int {
	if this.len > 0 {
		return this.arr[this.len-1]
	}
	return 0
}

func (this *MinStackWithMe) GetMin() int {
	if this.len == 0 {
		return 0
	}
	min := this.arr[0]
	for i := 1; i < this.len; i++ {
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
