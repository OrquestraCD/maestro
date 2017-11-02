package tables

import (
	"reflect"
	"testing"
)

func TestTable(t *testing.T) {
	row1 := []string{"foo", "foobar"}
	row2 := []string{"bar", "foo"}
	sampleData := [][]string{row1, row2}

	expect := `|foo|foobar|
-------------
|bar|foo   |
`

	got := Table(sampleData, true)
	if got != expect {
		t.Errorf("Expected:\n %s, Got:\n %s", expect, got)
	}
}

func TestTable_withoutHeader(t *testing.T) {
	row1 := []string{"foo", "foobar"}
	row2 := []string{"bar", "foo"}
	sampleData := [][]string{row1, row2}

	expect := `|foo|foobar|
|bar|foo   |
`

	got := Table(sampleData, false)
	if got != expect {
		t.Errorf("Expected:\n %s, Got:\n %s", expect, got)
	}
}

func TestCustomTable(t *testing.T) {
	row1 := []string{"foo", "foobar"}
	row2 := []string{"bar", "foo"}
	sampleData := [][]string{row1, row2}

	expect := `/foo,foobar/
/bar,foo   /
`

	got := CustomTable(sampleData, false, ",", "/")
	if got != expect {
		t.Errorf("Expected:\n %s, Got:\n %s", expect, got)
	}
}

func TestExtractColumn(t *testing.T) {
	row1 := []string{"foo", "foobar"}
	row2 := []string{"bar", "barfoo"}
	sampleData := [][]string{row1, row2}

	res := extractColumn(sampleData, 0)
	expected := []string{"foo", "bar"}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %+v, Got %+v", expected, res)
	}
}

func TestColumnWidth(t *testing.T) {
	column := []string{"foo", "barbarbar", "foobar"}

	res := columnWidth(column)
	if res != 9 {
		t.Errorf("Expected column width to be 9")
	}
}
