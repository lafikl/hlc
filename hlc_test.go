package hlc

import "testing"

func TestNew(t *testing.T) {
	h := NewWithPT(PT{Seconds: true})
	ts := h.Now()
	if ts.count != 0 {
		t.Fatal("expected count to be zero")
	}

	h.Update(ts)
	if h.t.count != 1 {
		t.Fatal("expected count to be one")
	}

	// test ordering
	if ts.Less(h.Now()) != true {
		t.Fatal("ordering test failed")
	}
}
