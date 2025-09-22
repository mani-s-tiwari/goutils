package deque

type Deque struct {
	data []int
}

func New() *Deque { return &Deque{data: []int{}} }
func (d *Deque) PushBack(x int)   { d.data = append(d.data, x) }
func (d *Deque) PushFront(x int)  { d.data = append([]int{x}, d.data...) }
func (d *Deque) PopBack() int     { x := d.data[len(d.data)-1]; d.data = d.data[:len(d.data)-1]; return x }
func (d *Deque) PopFront() int    { x := d.data[0]; d.data = d.data[1:]; return x }
func (d *Deque) Front() int       { return d.data[0] }
func (d *Deque) Back() int        { return d.data[len(d.data)-1] }
func (d *Deque) Empty() bool      { return len(d.data) == 0 }
func (d *Deque) Len() int         { return len(d.data) }
