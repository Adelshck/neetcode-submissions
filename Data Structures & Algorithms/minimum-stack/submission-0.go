type Stack struct {
    vals []int
}
func NewStack() Stack {
	return Stack{}
}
func (s *Stack) Push(val int) {
    s.vals = append(s.vals, val)
}

func (s *Stack) Pop() int {
    n := len(s.vals)
    val := s.vals[n-1]
    s.vals = s.vals[:n-1]
    return val
}

func (s *Stack) Top() int {
    return s.vals[len(s.vals)-1]
}

func (s *Stack) IsEmpty() bool {
    return len(s.vals) == 0
}

type MinStack struct {
	val Stack
	min Stack
}

func Constructor() MinStack {
	val := NewStack()
	min := NewStack()
	return MinStack{
		val : val,
		min : min,
	}
}

func (this *MinStack) Push(val int) {
    this.val.Push(val)

    if this.min.IsEmpty() {
        this.min.Push(val)
    } else {
        this.min.Push(min(val, this.min.Top()))
    }
}

func (this *MinStack) Pop() {
	this.val.Pop()
	this.min.Pop()
}

func (this *MinStack) Top() int {
	return this.val.Top()
}

func (this *MinStack) GetMin() int {
	return this.min.Top()
}
