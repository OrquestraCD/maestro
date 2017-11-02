package list

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func fieldsFlag(headers []string) cli.StringFlag {
	return cli.StringFlag{
		Name:  "fields,f",
		Usage: "List the fields to include in output",
		Value: strings.Join(headers, ","),
	}
}

func validateFieldNames(c *cli.Context, available []string) error {
	availableMap := make(map[string]interface{})
	for _, val := range available {
		availableMap[val] = nil
	}

	for _, providedVal := range strings.Split(c.String("fields"), ",") {
		if _, ok := availableMap[providedVal]; !ok {
			return fmt.Errorf("%s not an available field\n", providedVal)
		}
	}

	return nil
}
