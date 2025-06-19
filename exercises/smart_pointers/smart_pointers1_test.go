package main

import "testing"

func TestRefCounterBasic(t *testing.T) {
	rc := NewRefCounter(42)
	if rc == nil {
		t.Errorf("NewRefCounter should not return nil")
		return
	}

	if rc.GetValue() != 42 {
		t.Errorf("GetValue() = %d; want 42", rc.GetValue())
	}

	if rc.GetCount() != 1 {
		t.Errorf("Initial count should be 1, got %d", rc.GetCount())
	}
}

func TestRefCounterClone(t *testing.T) {
	rc1 := NewRefCounter(100)
	rc2 := rc1.Clone()

	if rc2 == nil {
		t.Errorf("Clone should not return nil")
		return
	}

	// Both should have same value
	if rc1.GetValue() != rc2.GetValue() {
		t.Errorf("Cloned RefCounter should have same value")
	}

	// Count should be 2
	if rc1.GetCount() != 2 {
		t.Errorf("After clone, count should be 2, got %d", rc1.GetCount())
	}

	if rc2.GetCount() != 2 {
		t.Errorf("Cloned RefCounter should also show count 2, got %d", rc2.GetCount())
	}
}

func TestRefCounterRelease(t *testing.T) {
	rc1 := NewRefCounter(50)
	rc2 := rc1.Clone()

	// Release one reference
	rc1.Release()

	// Count should be 1
	if rc2.GetCount() != 1 {
		t.Errorf("After one release, count should be 1, got %d", rc2.GetCount())
	}

	// Value should still be accessible
	if rc2.GetValue() != 50 {
		t.Errorf("Value should still be 50 after partial release, got %d", rc2.GetValue())
	}
}

func TestSwapPointers(t *testing.T) {
	a, b := 10, 20
	SwapPointers(&a, &b)

	if a != 20 {
		t.Errorf("After swap, a should be 20, got %d", a)
	}

	if b != 10 {
		t.Errorf("After swap, b should be 10, got %d", b)
	}
}

func TestDetectCycleNoCycle(t *testing.T) {
	// Create list: 1 -> 2 -> 3 -> nil
	head := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}

	if DetectCycle(head) {
		t.Errorf("DetectCycle should return false for list without cycle")
	}
}

func TestDetectCycleWithCycle(t *testing.T) {
	// Create list with cycle: 1 -> 2 -> 3 -> 2 (cycle)
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node2 // Creates cycle

	if !DetectCycle(node1) {
		t.Errorf("DetectCycle should return true for list with cycle")
	}
}

func TestDeepCopyList(t *testing.T) {
	// Create original list: 1 -> 2 -> 3
	original := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}

	copied := DeepCopyList(original)

	if copied == nil {
		t.Errorf("DeepCopyList should not return nil")
		return
	}

	// Check values are same
	if copied.Value != 1 {
		t.Errorf("First node value should be 1, got %d", copied.Value)
	}

	// Check that it's actually a deep copy (different memory addresses)
	if copied == original {
		t.Errorf("DeepCopyList should create new nodes, not copy references")
	}

	if copied.Next == original.Next {
		t.Errorf("Deep copy should create new nodes for all elements")
	}
}

func TestFindMiddle(t *testing.T) {
	// Test odd length: 1 -> 2 -> 3 -> 4 -> 5 (middle is 3)
	head := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 4,
					Next: &Node{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}

	middle := FindMiddle(head)
	if middle == nil {
		t.Errorf("FindMiddle should not return nil")
		return
	}

	if middle.Value != 3 {
		t.Errorf("Middle of 5-node list should be 3, got %d", middle.Value)
	}
}

func TestReverseList(t *testing.T) {
	// Create list: 1 -> 2 -> 3
	head := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}

	reversed := ReverseList(head)

	if reversed == nil {
		t.Errorf("ReverseList should not return nil")
		return
	}

	// Should be: 3 -> 2 -> 1
	if reversed.Value != 3 {
		t.Errorf("First node after reverse should be 3, got %d", reversed.Value)
	}

	if reversed.Next == nil || reversed.Next.Value != 2 {
		t.Errorf("Second node after reverse should be 2")
	}

	if reversed.Next.Next == nil || reversed.Next.Next.Value != 1 {
		t.Errorf("Third node after reverse should be 1")
	}
}

func TestObjectPool(t *testing.T) {
	pool := NewObjectPool()
	if pool == nil {
		t.Errorf("NewObjectPool should not return nil")
		return
	}

	// Get object from empty pool
	obj1 := pool.Get()
	if obj1 == nil {
		t.Errorf("ObjectPool.Get() should not return nil")
		return
	}

	// Put it back
	pool.Put(obj1)

	// Get again - should reuse the same object
	obj2 := pool.Get()
	if obj1 != obj2 {
		t.Errorf("ObjectPool should reuse objects")
	}
}
