package middleware

import (
	"github.com/aws/aws-sdk-go/aws"
	maestro "github.com/rackerlabs/maestro/pkg/config"
	"github.com/urfave/cli"
)

const (
	AWSConfigMetadataKey     = "AWS_CONFIG"
	MaestroConfigMetadataKey = "MAESTRO_CONFIG"
)

// Set new AWS session in app metadata
func SetAWSConfig(c *cli.Context) {
	awsConfig := aws.Config{}
	if c.GlobalString("region") != "" {
		awsConfig.Region = aws.String(c.GlobalString("region"))
	}

	c.App.Metadata[AWSConfigMetadataKey] = &awsConfig
}

// Get AWS Config set as CLI Metadata
func GetAWSConfig(c *cli.Context) *aws.Config {
	config := c.App.Metadata[AWSConfigMetadataKey]

	// Type Cast interface back to session and return
	return config.(*aws.Config)
}

// Set new Maestro Config
func SetMaestroConfig(c *cli.Context) error {
	conf, err := maestro.Load(c.GlobalString("config"))
	if err != nil {
		return err
	}

	c.App.Metadata[MaestroConfigMetadataKey] = conf
	return nil
}

// Get Config set as CLI Metadata
func GetMaestroConfig(c *cli.Context) *maestro.Config {
	config := c.App.Metadata[MaestroConfigMetadataKey]
	if config == nil {
		return &maestro.Config{}
	}

	// Type Cast interface back to session and return
	return config.(*maestro.Config)
}
