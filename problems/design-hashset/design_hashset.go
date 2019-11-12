package problem705

type MyHashSet struct {
	data map[int]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make(map[int]bool, 0)}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = true
}

func (this *MyHashSet) Remove(key int) {
	delete(this.data, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.data[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
