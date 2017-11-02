package middleware

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/urfave/cli"
)

const (
	SessionMetadataKey = "AWS_SESSION"
)

// Set new AWS session in app metadata
func SetSession(c *cli.Context) error {
	sess, err := session.NewSession(GetAWSConfig(c))
	if err != nil {
		return err
	}
	c.App.Metadata[SessionMetadataKey] = sess

	return nil
}

// Get Session set as CLI Metadata
func GetSession(c *cli.Context) *session.Session {
	sess := c.App.Metadata[SessionMetadataKey]

	// Type Cast interface back to session and return
	return sess.(*session.Session)
}
