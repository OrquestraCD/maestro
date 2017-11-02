package list

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/rackerlabs/go-tables/tables"
	"github.com/urfave/cli"

	"github.com/rackerlabs/maestro/pkg/middleware"
)

var defaultDocTblHeader = []string{
	"Name",
	"Document Type",
	"Platform Types",
	"Schema Version",
}

func listSSMDocuments(c *cli.Context) error {
	sess := middleware.GetSession(c)
	svc := ssm.New(sess, middleware.GetAWSConfig(c))

	documents := [][]string{
		defaultDocTblHeader,
	}

	input := &ssm.ListDocumentsInput{}
	err := svc.ListDocumentsPages(input,
		func(out *ssm.ListDocumentsOutput, lastPage bool) bool {
			for _, doc := range out.DocumentIdentifiers {
				documents = append(documents, []string{
					aws.StringValue(doc.Name),
					aws.StringValue(doc.DocumentType),
					strings.Join(aws.StringValueSlice(doc.PlatformTypes), ", "),
					aws.StringValue(doc.SchemaVersion),
				})
			}

			return true
		},
	)
	if err != nil {
		return err
	}

	tbl := tables.NewOrderedTableFromMatrix(documents)
	printOutput(tbl, strings.Split(c.String("fields"), ","))

	return nil
}
