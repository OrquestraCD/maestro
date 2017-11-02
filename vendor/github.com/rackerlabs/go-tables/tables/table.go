package tables

import (
	"fmt"
	"strconv"
	"strings"
)

// Given the data in the format [row][column] return a
// string with a grid using | for both the delimiter and outer characters.
func Table(data [][]string, header bool) string {
	return CustomTable(data, header, "|", "|")
}

// Given the data in the format [row][column] return a
// string with a grid using specified delimter and frame characters.
func CustomTable(data [][]string, header bool, delimiter, frame string) string {
	columnEsc := make([]string, 0)
	var result string

	var width int
	for i := 0; i < len(data[0]); i++ {
		width = columnWidth(extractColumn(data, i))
		columnEsc = append(columnEsc, "%-"+strconv.Itoa(width)+"s")
	}

	// Generate a format string for the row
	rowFmt := fmt.Sprintf(
		"%s%s%s\n", frame,
		strings.Join(columnEsc, delimiter),
		frame)

	var rowStr string
	for i, row := range data {
		rowStr = fmt.Sprintf(rowFmt, iface(row)...)
		if i == 0 && header {
			rowStr = rowStr + strings.Repeat("-", len(rowStr)) + "\n"
		}
		result = result + rowStr
	}

	return result
}

// Convert String Slice into interface slice
func iface(list []string) []interface{} {
	ifaceList := make([]interface{}, len(list))
	for i, item := range list {
		ifaceList[i] = item
	}

	return ifaceList
}

// Extract a single column from a matrix
func extractColumn(data [][]string, columnID int) []string {
	column := make([]string, 0)
	for _, row := range data {
		column = append(column, row[columnID])
	}

	return column
}

// Given a column return the column width
func columnWidth(column []string) int {
	width := 0
	for _, row := range column {
		if len(row) > width {
			width = len(row)
		}
	}

	return width
}
