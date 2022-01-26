package set

import (
	"fmt"
	"testing"
)

func populateSet(count int, start int) *ItemSet {
	set := ItemSet{}
	for i := start; i < (start + count); i++ {
		set.Add(fmt.Sprintf("item%d", i))
	}
	return &set
}

func TestAdd(t *testing.T) {
	set := populateSet(3, 0)
	fmt.Println(set)
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	set.Add("item1") //should not add it, already there
	if size := set.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	set.Add("item4") //should not add it, already there
	if size := set.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
	fmt.Println(set)
}

func TestClear(t *testing.T) {
	set := populateSet(3, 0)
	set.Clear()
	if size := set.Size(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
	fmt.Println(set)
}

func TestDelete(t *testing.T) {
	set := populateSet(3, 0)
	set.Delete("item2")
	if size := set.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestHas(t *testing.T) {
	set := populateSet(3, 0)
	has := set.Has("item2")
	if !has {
		t.Errorf("expected item2 to be there")
	}
	set.Delete("item2")
	has = set.Has("item2")
	if has {
		t.Errorf("expected item2 to be removed")
	}
	set.Delete("item1")
	has = set.Has("item1")
	if has {
		t.Errorf("expected item1 to be removed")
	}
}

func TestItems(t *testing.T) {
	set := populateSet(3, 0)
	items := set.Items()
	if len(items) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(items))
	}
	set = populateSet(520, 0)
	items = set.Items()
	if len(items) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(items))
	}
}

func TestSize(t *testing.T) {
	set := populateSet(3, 0)
	items := set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
	set = populateSet(0, 0)
	items = set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
	set = populateSet(10000, 0)
	items = set.Items()
	if len(items) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(items))
	}
}