package design_linked_list

import "testing"

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)      // linked list becomes 1->2->3
	output := obj.Get(1) == 2 // returns 2
	obj.DeleteAtIndex(1)      // now the linked list is 1->3
	output = obj.Get(1) == 3 && output
	obj.AddAtIndex(2, 2) // linked list becomes 1->3->2
	output = obj.Get(5) == -1 && output
	if !output {
		t.Fatalf("output: %v, expected: %v", output, true)
	}
}
