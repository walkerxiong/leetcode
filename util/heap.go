package util

type CompareFunc func(i, j int) bool

type IntHeap struct {
	Data []int
	f    CompareFunc
	Cap  int
}

func NewIntHeap(f CompareFunc, cap int) *IntHeap {
	h := &IntHeap{
		f:    f,
		Cap:  cap,
		Data: make([]int, 0, cap),
	}
	return h
}

func (h *IntHeap) Pop() int {
	res := h.Data[0]
	h.Data = h.Data[1:]
	h.down()
	return res
}

func (h *IntHeap) Push(v int) {
	l := len(h.Data)
	if l == h.Cap {
		if !h.f(h.Data[0], v) {
			h.Data[0] = v
			h.down()
		}
	} else {
		h.Data = append(h.Data, v)
		h.down()
	}
}

func (h *IntHeap) down() {
	l := len(h.Data)
	for i := l; i >= 0; i-- {
		left := 2*i + 1
		if left > l {
			continue
		}
		n := i
		if left < l && h.f(h.Data[n], h.Data[left]) {
			n = left
		}

		right := left + 1
		if right < l && h.f(h.Data[n], h.Data[right]) {
			n = right
		}
		if i != n {
			h.Data[i], h.Data[n] = h.Data[n], h.Data[i]
		}

	}
}
