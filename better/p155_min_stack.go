package better

// MinStack is a stack and maintains smallest elements
type MinStack struct {
	// use to store smallest element the corresponding
	// element array sees
	minArray []int
	array    []int
	ptr      int
}

// Constructor initialize your data structure here
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

// Push an element into min stack
func (minstack *MinStack) Push(x int) {
	minstack.ptr++
	if minstack.ptr >= len(minstack.array) {
		newArray := make([]int, len(minstack.array)*2)
		copy(newArray[:len(minstack.array)], minstack.array[:len(minstack.array)])
		minstack.array = newArray

		newMinArray := make([]int, len(minstack.minArray)*2)
		copy(newMinArray[:len(minstack.minArray)], minstack.minArray[:len(minstack.minArray)])
		minstack.minArray = newMinArray
	}

	minstack.array[minstack.ptr] = x
	if minstack.ptr == 0 || x <= minstack.minArray[minstack.ptr-1] {
		minstack.minArray[minstack.ptr] = x
	} else {
		minstack.minArray[minstack.ptr] = minstack.minArray[minstack.ptr-1]
	}
}

// Pop out an element
func (minstack *MinStack) Pop() {
	if minstack.ptr != -1 {
		minstack.ptr--
	}
}

// Top gets the element on top of min stack
func (minstack *MinStack) Top() int {
	if minstack.ptr == -1 {
		return 0
	}
	return minstack.array[minstack.ptr]
}

// GetMin get the min value
func (minstack *MinStack) GetMin() int {
	if minstack.ptr == -1 {
		return 0
	}
	return minstack.minArray[minstack.ptr]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
