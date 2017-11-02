package stringset

import (
	"reflect"
	"testing"
)

func TestNewFromSlice(t *testing.T) {
	testSlice := []string{"this", "is", "a", "test"}

	gotSet := NewFromSlice(testSlice)

	for _, item := range testSlice {
		if !gotSet.data[item] {
			t.Errorf("expected \"%s\" to be in the new set", item)
		}
	}
}

func TestAdd(t *testing.T) {
	set := New()

	set.Add("foo")

	if !set.data["foo"] {
		t.Errorf("expected key \"foo\" to be set to true")
	}
}

func TestRemove(t *testing.T) {
	set := StringSet{
		data: map[string]bool{
			"foo": true,
			"bar": true,
		},
	}

	set.Remove("bar")
	if set.data["bar"] {
		t.Errorf("expected bar to be removed")
	}
}

func TestSlice(t *testing.T) {
	set := StringSet{
		data: map[string]bool{
			"foo": true,
			"bar": true,
		},
	}

	expect := []string{"foo", "bar"}
	got := set.Slice()

	notFound := true
	for _, expectItem := range expect {
		for _, gotItem := range got {
			if expectItem == gotItem {
				notFound = false
				break
			}
		}

		if notFound {
			t.Errorf("missing item %s in returned slice", expect)
		}
	}
}

func TestUnion(t *testing.T) {
	firstSet := StringSet{
		data: map[string]bool{
			"foo": true,
			"bar": true,
			"baz": true,
		},
	}

	secondSet := &StringSet{
		data: map[string]bool{
			"baz":    true,
			"foobar": true,
		},
	}

	expect := &StringSet{
		data: map[string]bool{
			"foo":    true,
			"bar":    true,
			"baz":    true,
			"foobar": true,
		},
	}

	got := firstSet.Union(secondSet)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected: %+v\n, got: %+v", expect, got)
	}
}

func TestIntersection(t *testing.T) {
	firstSet := StringSet{
		data: map[string]bool{
			"foo": true,
			"bar": true,
			"baz": true,
		},
	}

	secondSet := &StringSet{
		data: map[string]bool{
			"foo":    true,
			"baz":    true,
			"foobar": true,
		},
	}

	expect := &StringSet{
		data: map[string]bool{
			"foo": true,
			"baz": true,
		},
	}

	got := firstSet.Intersection(secondSet)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected: %+v\n, got: %+v", expect, got)
	}
}

func TestDifference(t *testing.T) {
	firstSet := &StringSet{
		data: map[string]bool{
			"foo":    true,
			"bar":    true,
			"baz":    true,
			"foobar": true,
		},
	}

	secondSet := &StringSet{
		data: map[string]bool{
			"foo": true,
			"baz": true,
		},
	}

	expect := &StringSet{
		data: map[string]bool{
			"bar":    true,
			"foobar": true,
		},
	}

	got := firstSet.Difference(secondSet)
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected: %+v\n, got: %+v", expect, got)
	}
}
