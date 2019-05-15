package design_hashmap

import "testing"

func TestDesignHashmap(t *testing.T) {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	output := obj.Get(1) == 1
	output = obj.Get(3) == -1 && output
	obj.Put(2, 1)
	output = obj.Get(2) == 1 && output
	obj.Remove(2)
	output = obj.Get(2) == -1 && output
	if !output {
		t.Fatalf("output: %v, expected: %v", output, true)
	}
}
