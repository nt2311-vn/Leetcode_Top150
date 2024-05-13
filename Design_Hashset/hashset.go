package designhashset

type MyHashSet struct {
	payload [1_000_001]bool
}

func Constructor() MyHashSet {
	return MyHashSet{payload: [1_000_001]bool{}}
}

func (h *MyHashSet) Add(key int) {
	h.payload[key] = true
}

func (h *MyHashSet) Remove(key int) {
	h.payload[key] = false
}

func (h *MyHashSet) Contains(key int) bool {
	return h.payload[key]
}
