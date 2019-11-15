package problem707

import "testing"

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)   // linked list becomes 1->2->3
	got := obj.Get(1) == 2 // returns 2
	obj.DeleteAtIndex(1)   // now the linked list is 1->3
	got = obj.Get(1) == 3 && got
	obj.AddAtIndex(2, 2) // linked list becomes 1->3->2
	got = obj.Get(5) == -1 && got
	if !got {
		t.Fatalf("got: %v, want: %v", got, true)
	}
}
