package tables

import (
	"reflect"
	"testing"
)

func newTestTable() (OrderedTable, [][]string) {
	return OrderedTable{
			data: map[string][]string{
				"header1": []string{"foo", "bar"},
				"header2": []string{"foobar", "barfoo"},
			},
			order: []string{
				"header1",
				"header2",
			},
		}, [][]string{
			[]string{"header1", "header2"},
			[]string{"foo", "foobar"},
			[]string{"bar", "barfoo"},
		}
}

// Create a new ordered table given a two dimensional slice.
func TestNewOrderedTableFromMatrix(t *testing.T) {
	testMatrix := make([][]string, 3)
	testMatrix[0] = []string{"header1", "header2"}
	testMatrix[1] = []string{"firstItem", "secondItem"}
	testMatrix[2] = []string{"thirdItem", "fourthItem"}

	tbl := NewOrderedTableFromMatrix(testMatrix)

	hdr1Expect := []string{"firstItem", "thirdItem"}
	hdr1Got := tbl.data["header1"]
	if !reflect.DeepEqual(hdr1Got, hdr1Expect) {
		t.Errorf("Expected %v, Got %v", hdr1Expect, hdr1Got)
	}

	hdr2Expect := []string{"secondItem", "fourthItem"}
	hdr2Got := tbl.data["header2"]
	if !reflect.DeepEqual(hdr2Got, hdr2Expect) {
		t.Errorf("Expected %v, Got %v", hdr2Expect, hdr2Got)
	}
}

func TestOrderedTableAddRow(t *testing.T) {
	tbl, _ := newTestTable()
	tbl.AddRow([]string{"baz", "barbaz"})

	got := tbl.data["header1"][2]
	want := "baz"
	if got != want {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, got)
	}
}

func TestOrderedTableAddColumn(t *testing.T) {
	tbl, _ := newTestTable()
	tbl.AddColumn([]string{"header3", "itemOne", "itemTwo"})

	got := tbl.data["header3"][1]
	want := "itemTwo"
	if got != want {
		t.Errorf("Expected \"%s\", Got \"%s\"", want, got)
	}
}

func TestOrderedTableMatrix(t *testing.T) {
	tbl, matrix := newTestTable()

	got := tbl.Matrix()
	want := matrix
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}

func TestOrderedTableMatrixWithOrder(t *testing.T) {
	tbl := OrderedTable{
		data: map[string][]string{
			"header1": []string{"foo", "bar"},
			"header2": []string{"foobar", "barfoo"},
		},
		order: []string{
			"header1",
			"header2",
		},
	}

	want := [][]string{
		[]string{"header2"},
		[]string{"foobar"},
		[]string{"barfoo"},
	}
	got := tbl.MatrixWithOrder([]string{"header2"})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}

func TestOrderedTableString(t *testing.T) {
	tbl := OrderedTable{
		data: map[string][]string{
			"header1": []string{"foo", "bar"},
			"header2": []string{"foobar", "barfoo"},
		},
		order: []string{
			"header1",
			"header2",
		},
	}

	want := `|header1|header2|
------------------
|foo    |foobar |
|bar    |barfoo |
`
	got := tbl.String()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}

func TestOrderedTableStringWithOrder(t *testing.T) {
	tbl := OrderedTable{
		data: map[string][]string{
			"header1": []string{"foo", "bar"},
			"header2": []string{"foobar", "barfoo"},
		},
		order: []string{
			"header1",
			"header2",
		},
	}

	want := `|header2|
----------
|foobar |
|barfoo |
`
	got := tbl.StringWithOrder([]string{"header2"})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}

func TestOrderedTableCustomString(t *testing.T) {
	tbl := OrderedTable{
		data: map[string][]string{
			"header1": []string{"foo", "bar"},
			"header2": []string{"foobar", "barfoo"},
		},
		order: []string{
			"header1",
			"header2",
		},
	}

	want := `|header1,header2|
------------------
|foo    ,foobar |
|bar    ,barfoo |
`
	got := tbl.CustomString(",", "|")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}

func TestOrderedTableCustomStringWithOrder(t *testing.T) {
	tbl := OrderedTable{
		data: map[string][]string{
			"header1": []string{"foo", "bar"},
			"header2": []string{"foobar", "barfoo"},
		},
		order: []string{
			"header1",
			"header2",
		},
	}

	want := `header2
--------
foobar 
barfoo 
`
	got := tbl.CustomStringWithOrder([]string{"header2"}, ",", "")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}
