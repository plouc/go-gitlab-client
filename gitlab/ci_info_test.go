package gitlab

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

type ciInfoTestCase struct {
	name   string
	env    map[string]string
	expect func() *CiInfo
}

func TestGetCiInfo(t *testing.T) {
	testCases := []*ciInfoTestCase{
		{
			"no env",
			nil,
			func() *CiInfo {
				return &CiInfo{}
			},
		},
		{
			"chat",
			map[string]string{
				"CHAT_INPUT":   "chat_input",
				"CHAT_CHANNEL": "chat_channel",
			},
			func() *CiInfo {
				i := CiInfo{}
				i.Chat.Input = "chat_input"
				i.Chat.Channel = "chat_channel"

				return &i
			},
		},
		{
			"deploy",
			map[string]string{
				"CI_DEPLOY_USER": "plouc",
			},
			func() *CiInfo {
				i := CiInfo{}
				i.Deploy.User = "plouc"

				return &i
			},
		},
		{
			"project",
			map[string]string{
				"CI_PROJECT_ID":         "11",
				"CI_PROJECT_DIR":        "builds/",
				"CI_PROJECT_NAME":       "sample",
				"CI_PROJECT_NAMESPACE":  "plouc/sample",
				"CI_PROJECT_PATH":       "plouc/sample",
				"CI_PROJECT_PATH_SLUG":  "plouc-sample",
				"CI_PROJECT_URL":        "http://fake.io/plouc/sample",
				"CI_PROJECT_VISIBILITY": "private",
			},
			func() *CiInfo {
				i := CiInfo{}
				i.Project.Id = "11"
				i.Project.Dir = "builds/"
				i.Project.Name = "sample"
				i.Project.Namespace = "plouc/sample"
				i.Project.Path = "plouc/sample"
				i.Project.PathSlug = "plouc-sample"
				i.Project.Url = "http://fake.io/plouc/sample"
				i.Project.Visibility = "private"

				return &i
			},
		},
		{
			"job",
			map[string]string{
				"CI_JOB_ID":    "1",
				"CI_JOB_NAME":  "my-job",
				"CI_JOB_STAGE": "test",
				"CI_JOB_URL":   "http://fake.io/job/1",
			},
			func() *CiInfo {
				i := CiInfo{}
				i.Job.Id = "1"
				i.Job.Name = "my-job"
				i.Job.Stage = "test"
				i.Job.Url = "http://fake.io/job/1"

				return &i
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range tc.env {
					os.Unsetenv(k)
				}
			}()

			info, _ := GetCiInfo()
			assert.Equal(t, info, tc.expect())
		})
	}
}
