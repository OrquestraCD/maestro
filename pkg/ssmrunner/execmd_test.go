package ssmrunner

import (
	"regexp"
	"testing"
)

func TestS3Regex(t *testing.T) {
	bucketURLs := map[string]string{
		"regionDomain":    "https://s3-us-west-2.amazonaws.com/maestro-0r5f141055052e5/6daa402c-2242-491a-9272-9287ffd9f492/i-0234sd0sdfd0/awsrunShellScript/0.runShellScript/stdout",
		"regionSubDomain": "https://s3.eu-central-1.amazonaws.com/maestro-0r5f141055052e5/6daa402c-2242-491a-9272-9287ffd9f492/i-0234sd0sdfd0/awsrunShellScript/0.runShellScript/stdout",
		"usStandard":      "https://s3.amazonaws.com/maestro-0r5f141055052e5/6daa402c-2242-491a-9272-9287ffd9f492/i-0234sd0sdfd0/awsrunShellScript/0.runShellScript/stdout",
	}

	for testName, url := range bucketURLs {
		t.Run(testName, func(t *testing.T) {
			re := regexp.MustCompile(S3UrlRegex)
			result := re.FindStringSubmatch(url)

			if len(result) != 5 {
				t.Errorf("expected regex to find 5 groups: %d found", len(result))
			}
		})
	}
}
