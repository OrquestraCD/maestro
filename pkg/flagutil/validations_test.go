package flagutil

import (
	"flag"
	"testing"

	"github.com/urfave/cli"
)

func TestValidateRequiredFlags_notset(t *testing.T) {
	localFlagSet := flag.NewFlagSet("testLocal", 0)
	localFlagSet.String("must-be-set", "", "")
	localContext := cli.NewContext(nil, localFlagSet, nil)

	if err := ValidateRequiredFlags(localContext, []string{"must-be-set"}); err == nil {
		t.Errorf("Expected error when flag is not set")
	}
}

func TestValidateRequiredFlags_set(t *testing.T) {
	localFlagSet := flag.NewFlagSet("testLocal", 0)
	localFlagSet.String("must-be-set", "thisisset", "")
	localContext := cli.NewContext(nil, localFlagSet, nil)

	if err := ValidateRequiredFlags(localContext, []string{"must-be-set"}); err != nil {
		t.Errorf("Got unexpected error when flag --must-be-set is set: %v", err)
	}
}
