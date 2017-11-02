package list

import (
	"github.com/rackerlabs/go-tables/tables"
	. "github.com/rackerlabs/maestro/ui"
)

// Print fields in output in a table
func printOutput(output tables.OrderedTable, order []string) {
	UI.Print(output.CustomStringWithOrder(order, "|", ""))
}
