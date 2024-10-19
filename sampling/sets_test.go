package sampling

import (
	"context"
	"testing"
)

func TestSplitSet(t *testing.T) {
	var source = NewSliceSource([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	var sets, err = SplitSet[int](source, 7, 1.5, 1.5)
	if err != nil {
		t.Fatalf("SplitSet failed: %v", err)
	}
	if len(sets) != 3 {
		t.Fatalf("SplitSet failed: expected 3 sets, got %d", len(sets))
	}
	var counts = make([]int, 3)
	for i, set := range sets {
		counts[i], err = set.Count(context.TODO())
		if err != nil {
			t.Fatalf("SplitSet failed: %v", err)
		}
	}
	if counts[0] != 7 {
		t.Fatalf("SplitSet failed: expected 6 elements in the first set, got %d", counts[0])
	}
	if counts[1] != 1 {
		t.Fatalf("SplitSet failed: expected 2 element in the second set, got %d", counts[1])
	}
	if counts[2] != 2 {
		t.Fatalf("SplitSet failed: expected 2 elements in the third set, got %d", counts[2])
	}
}
