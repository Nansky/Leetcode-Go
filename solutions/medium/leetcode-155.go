// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

// Example 1:

// Input
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// Output
// [null,null,null,null,-3,null,0,-2]

// Explanation
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -2

// Constraints:

// -231 <= val <= 231 - 1
// Methods pop, top and getMin operations will always be called on non-empty stacks.
// At most 3 * 104 calls will be made to push, pop, top, and getMin.

type MinStack struct {
	MinVal int
	Node   *StackNode
}

type StackNode struct {
	Val     int
	prevMin int
	Next    *StackNode
}

var top *StackNode

func Constructor() MinStack {
	root := MinStack{
		MinVal: math.MaxInt,
	}

	return root
}

func (this *MinStack) Push(val int) {
	if this.Node == nil {
		top = &StackNode{
			Val:  val,
			Next: nil,
		}

		this.MinVal = val
	} else {
		top = &StackNode{
			Val:     val,
			Next:    this.Node,
			prevMin: this.MinVal,
		}

		if this.MinVal > val {
			this.MinVal = val
		}
	}

	this.Node = top
}

func (this *MinStack) Pop() {
	this.MinVal = this.Node.prevMin
	this.Node = this.Node.Next

}

func (this *MinStack) Top() int {
	return this.Node.Val
}

func (this *MinStack) GetMin() int {
	return this.MinVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */