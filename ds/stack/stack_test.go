package stack

import "testing"

func TestStack(t *testing.T) {
	stk := New[int]()
	for K := 0; K < 10; K++ {
		stk.Push(K)
		if stk.Top() != K {
			t.Fatalf("expected %v, got %v", K, stk.Top())
		}
	}

	t.Logf("%v", stk.String())
	if stk.Size() != 10 {
		t.Fatalf("expected 10, got %v", stk.Size())
	}

	expVal := 9
	for !stk.Empty() {
		val := stk.Pop()
		if val != expVal {
			t.Fatalf("expected %v, got %v", expVal, val)
		}

		expVal--
	}

	stk = New[int]()
	stk.Clear()
	if !stk.Empty() {
		t.Fatalf("expected empty")
	}

	stk.Push(5)
	if expVal = stk.Pop(); expVal != 5 {
		t.Fatalf("expected 5, got %v", expVal)
	}
}
