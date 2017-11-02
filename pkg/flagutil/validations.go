package flagutil

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

// Validates string flags
func ValidateRequiredFlags(c *cli.Context, requiredFlags []string) error {
	for _, flag := range requiredFlags {
		if c.String(flag) == "" && c.GlobalString(flag) == "" {
			return fmt.Errorf("Flag --%s must be set", flag)
		}
	}

	return nil
}

type ExclusiveFlagGroup []string

func ValidateExclusiveFlagGroups(c *cli.Context, groups []ExclusiveFlagGroup) error {
	setGroups := make([]int, 0)

	for i, group := range groups {
		for _, flg := range group {
			if c.String(flg) != "" {
				setGroups = append(setGroups, i)
				break
			}
		}
	}

	if len(setGroups) == 0 {
		return fmt.Errorf("missing one of the required flags: %q\n", groups)
	}

	if len(setGroups) > 1 {
		firstGroup := groups[setGroups[0]]
		for i, item := range firstGroup {
			firstGroup[i] = "--" + item
		}

		secondGroup := groups[setGroups[1]]
		for i, item := range secondGroup {
			secondGroup[i] = "--" + item
		}

		return fmt.Errorf(
			"Can not set mutually exclusive flags %s and %s\n",
			strings.Join(firstGroup, ","),
			strings.Join(secondGroup, ","),
		)
	}

	// Validate the required flags are set for the group
	return ValidateRequiredFlags(c, groups[setGroups[0]])
}
