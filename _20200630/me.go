package main

/**
剑指 Offer 09. 用两个栈实现队列

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例1：
```
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
```

示例2：
```
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
```

提示：
- `1 <= values <= 10000`
- `最多会对 appendTail、deleteHead 进行 10000 次调用`
*/

/**
Go中没有队列和栈，都是通过数组来模拟的，
*/
type MyCQueue struct {
	arr    []int
	length int
}

func MyConstructor() CQueue {
	return CQueue{}
}

func (this *MyCQueue) AppendTail(value int) {
	this.arr = append(this.arr, value)
	this.length++
}

func (this *MyCQueue) DeleteHead() int {
	if this.length == 0 {
		return -1
	}
	first := this.arr[0]
	this.arr = this.arr[1:len(this.arr)]
	this.length--
	return first
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
