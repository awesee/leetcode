package design_hashset

import "testing"

func TestDesignHashset(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	output := obj.Contains(1)
	output = !obj.Contains(3) && output
	obj.Add(2)
	output = obj.Contains(2) && output
	obj.Remove(2)
	output = !obj.Contains(2) && output
	if !output {
		t.Fatalf("output: %v, expected: %v", output, true)
	}
}
