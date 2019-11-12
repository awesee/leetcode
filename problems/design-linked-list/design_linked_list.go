package problem707

import "container/list"

type MyLinkedList struct {
	data *list.List
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{list.New()}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.data.Len() {
		return -1
	}
	node := this.data.Front()
	for i := 0; i < index; i++ {
		node = node.Next()
	}
	return node.Value.(int)
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.data.PushFront(val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.data.PushBack(val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.data.Len() {
		this.AddAtTail(val)
	} else if index >= 0 && index < this.data.Len() {
		node := this.data.Front()
		for i := 0; i < index; i++ {
			node = node.Next()
		}
		this.data.InsertBefore(val, node)
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < this.data.Len() {
		node := this.data.Front()
		for i := 0; i < index; i++ {
			node = node.Next()
		}
		this.data.Remove(node)
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
