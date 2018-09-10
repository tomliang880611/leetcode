package better

type MinStack struct {
	// use to store smallest element the corresponding
	// element array sees
	minArray []int
	array    []int
	ptr      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	array := make([]int, 10)
	minArray := make([]int, 10)
	ptr := -1
	return MinStack{
		array:    array,
		minArray: minArray,
		ptr:      ptr,
	}
}

func (this *MinStack) Push(x int) {
	this.ptr++
	if this.ptr >= len(this.array) {
		newArray := make([]int, len(this.array)*2)
		copy(newArray[:len(this.array)], this.array[:len(this.array)])
		this.array = newArray

		newMinArray := make([]int, len(this.minArray)*2)
		copy(newMinArray[:len(this.minArray)], this.minArray[:len(this.minArray)])
		this.minArray = newMinArray
	}

	this.array[this.ptr] = x
	if this.ptr == 0 || x <= this.minArray[this.ptr-1] {
		this.minArray[this.ptr] = x
	} else {
		this.minArray[this.ptr] = this.minArray[this.ptr-1]
	}
}

func (this *MinStack) Pop() {
	if this.ptr != -1 {
		this.ptr--
	}
}

func (this *MinStack) Top() int {
	if this.ptr == -1 {
		return 0
	}
	return this.array[this.ptr]
}

func (this *MinStack) GetMin() int {
	if this.ptr == -1 {
		return 0
	}
	return this.minArray[this.ptr]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
