package insert_delete_getrandom

import "math/rand"

type RandomizedSet struct {
	set  map[int]int
	arr  []int
	size int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set:  make(map[int]int),
		arr:  make([]int, 0),
		size: 0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exist := this.set[val]

	if exist {
		this.arr[index] = this.arr[len(this.arr)-1]
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.set))
}
