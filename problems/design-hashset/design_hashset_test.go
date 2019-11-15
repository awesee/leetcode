package problem705

import "testing"

func TestDesignHashset(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	got := obj.Contains(1)
	got = !obj.Contains(3) && got
	obj.Add(2)
	got = obj.Contains(2) && got
	obj.Remove(2)
	got = !obj.Contains(2) && got
	if !got {
		t.Fatalf("got: %v, want: %v", got, true)
	}
}
