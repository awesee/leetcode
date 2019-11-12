package problem933

import "container/list"

type RecentCounter struct {
	queue *list.List
}

func Constructor() RecentCounter {
	return RecentCounter{list.New()}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue.PushBack(t)
	for this.queue.Front().Value.(int) < t-3000 {
		this.queue.Remove(this.queue.Front())
	}
	return this.queue.Len()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
