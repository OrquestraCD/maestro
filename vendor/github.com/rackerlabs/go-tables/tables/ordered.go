package tables

type OrderedTable struct {
	data  map[string][]string
	order []string
}

func NewOrderedTable() OrderedTable {
	return OrderedTable{
		data:  make(map[string][]string, 0),
		order: make([]string, 0),
	}
}

// Create a new ordered table given a two dimensional slice.
func NewOrderedTableFromMatrix(data [][]string) OrderedTable {
	header := data[0]
	tbl := OrderedTable{
		order: make([]string, len(header)),
		data:  make(map[string][]string),
	}

	for i, item := range header {
		tbl.order[i] = item
		tbl.data[item] = extractColumn(data[1:], i)
	}

	return tbl
}

// Add a row to the table object
func (t *OrderedTable) AddRow(row []string) {
	for i, key := range t.order {
		t.data[key] = append(t.data[key], row[i])
	}
}

// Add a new column to the table object, expects the first
// item in the slice to be the header name.
func (t *OrderedTable) AddColumn(column []string) {
	header := column[0]
	t.data[header] = column[1:]
}

// Return the ordered table as a two dimensional slice [][]string
func (t *OrderedTable) Matrix() [][]string {
	return t.MatrixWithOrder(t.order)
}

// Return the ordered table as a two dimensional slice [][]string
// given a specified order
func (t *OrderedTable) MatrixWithOrder(order []string) [][]string {
	matrix := make([][]string, t.rowCount()+1)
	matrix[0] = order

	for i, _ := range matrix[1:] {
		row := make([]string, len(matrix[0]))
		for k, key := range matrix[0] {
			row[k] = t.data[key][i]
		}
		matrix[i+1] = row
	}

	return matrix
}

// Return the table as a pretty string
func (t *OrderedTable) String() string {
	return Table(t.Matrix(), true)
}

// Return the table as a string with the given order
func (t *OrderedTable) StringWithOrder(order []string) string {
	return Table(t.MatrixWithOrder(order), true)
}

// Return the table as a string using a provided delimiter and frame
// character.
func (t *OrderedTable) CustomString(delimiter, frame string) string {
	return CustomTable(t.Matrix(), true, delimiter, frame)
}

// Return the table as a string using the provided delimiter and frame
// characters in the specified order.
func (t *OrderedTable) CustomStringWithOrder(order []string, delimiter, frame string) string {
	return CustomTable(t.MatrixWithOrder(order), true, delimiter, frame)
}

func (t *OrderedTable) rowCount() int {
	if len(t.data) != 0 {
		for _, item := range t.data {
			return len(item)
		}
	}

	return 0
}
