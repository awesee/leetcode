package problem706

import "testing"

func TestDesignHashmap(t *testing.T) {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	got := obj.Get(1) == 1
	got = obj.Get(3) == -1 && got
	obj.Put(2, 1)
	got = obj.Get(2) == 1 && got
	obj.Remove(2)
	got = obj.Get(2) == -1 && got
	if !got {
		t.Fatalf("got: %v, want: %v", got, true)
	}
}
