package gitlab

import (
	"errors"
	"github.com/iancoleman/strcase"
	"io"
	"os"
	"reflect"
	"strings"
)

var CiEnvMapping map[string]*CiEnvKey

type CiEnvKey struct {
	FromGitlabVersion string
	FromRunnerVersion string
	Hide              bool
}

type CiInfo struct {
	Chat struct {
		Input   string `env:"CHAT_INPUT"   json:"input"   yaml:"input"`
		Channel string `env:"CHAT_CHANNEL" json:"channel" yaml:"channel"`
	} `json:"chat" yaml:"chat"`
	Deploy struct {
		User string `env:"CI_DEPLOY_USER" json:"user" yaml:"user"`
	} `json:"deploy" yaml:"deploy"`
	Project struct {
		Id         string `env:"CI_PROJECT_ID"         json:"id"         yaml:"id"`
		Dir        string `env:"CI_PROJECT_DIR"        json:"dir"        yaml:"dir"`
		Name       string `env:"CI_PROJECT_NAME"       json:"name"       yaml:"name"`
		Namespace  string `env:"CI_PROJECT_NAMESPACE"  json:"namespace"  yaml:"namespace"`
		Path       string `env:"CI_PROJECT_PATH"       json:"path"       yaml:"path"`
		PathSlug   string `env:"CI_PROJECT_PATH_SLUG"  json:"path_slug"  yaml:"path_slug"`
		Url        string `env:"CI_PROJECT_URL"        json:"url"        yaml:"url"`
		Visibility string `env:"CI_PROJECT_VISIBILITY" json:"visibility" yaml:"visibility"`
	} `json:"project" yaml:"project"`
	Commit struct {
		Sha         string `env:"CI_COMMIT_SHA"         json:"sha"         yaml:"sha"`
		RefName     string `env:"CI_COMMIT_REF_NAME"    json:"ref_name"    yaml:"ref_name"`
		RefSlug     string `env:"CI_COMMIT_REF_SLUG"    json:"ref_slug"    yaml:"ref_slug"`
		Tag         string `env:"CI_COMMIT_TAG"         json:"tag"         yaml:"tag"`
		Message     string `env:"CI_COMMIT_MESSAGE"     json:"message"     yaml:"message"`
		Title       string `env:"CI_COMMIT_TITLE"       json:"title"       yaml:"title"`
		Description string `env:"CI_COMMIT_DESCRIPTION" json:"description" yaml:"description"`
	} `json:"commit" yaml:"commit"`
	Job struct {
		Id    string `env:"CI_JOB_ID"    json:"id"    yaml:"id"`
		Name  string `env:"CI_JOB_NAME"  json:"name"  yaml:"name"`
		Stage string `env:"CI_JOB_STAGE" json:"stage" yaml:"stage"`
		Url   string `env:"CI_JOB_URL"   json:"url"   yaml:"url"`
	} `json:"job" yaml:"job"`
	Pipeline struct {
		Id     string `env:"CI_PIPELINE_ID"     json:"id"     yaml:"id"`
		Iid    string `env:"CI_PIPELINE_IID"    json:"iid"    yaml:"iid"`
		Source string `env:"CI_PIPELINE_SOURCE" json:"source" yaml:"source"`
		Url    string `env:"CI_PIPELINE_URL"    json:"url"    yaml:"url"`
	} `json:"pipeline" yaml:"pipeline"`
	Runner struct {
		Id             string `env:"CI_RUNNER_ID"              json:"id"              yaml:"id"`
		Description    string `env:"CI_RUNNER_DESCRIPTION"     json:"description"     yaml:"description"`
		Tags           string `env:"CI_RUNNER_TAGS"            json:"tags"            yaml:"tags"`
		Version        string `env:"CI_RUNNER_VERSION"         json:"version"         yaml:"version"`
		Revision       string `env:"CI_RUNNER_REVISION"        json:"revision"        yaml:"revision"`
		ExecutableArch string `env:"CI_RUNNER_EXECUTABLE_ARCH" json:"executable_arch" yaml:"executable_arch"`
	} `json:"runner" yaml:"runner"`
	User struct {
		Id    string `env:"GITLAB_USER_ID"    json:"id"    yaml:"id"`
		Email string `env:"GITLAB_USER_EMAIL" json:"email" yaml:"email"`
		Login string `env:"GITLAB_USER_LOGIN" json:"login" yaml:"login"`
		Name  string `env:"GITLAB_USER_NAME"  json:"name"  yaml:"name"`
	} `json:"user" yaml:"user"`
	Server struct {
		Name     string `env:"CI_SERVER_NAME"     json:"name"     yaml:"name"`
		Revision string `env:"CI_SERVER_REVISION" json:"revision" yaml:"revision"`
		Version  string `env:"CI_SERVER_VERSION"  json:"version"  yaml:"version"`
	} `json:"server" yaml:"server"`
	Registry struct {
		Registry string `env:"CI_REGISTRY"       json:"registry" yaml:"registry"`
		Image    string `env:"CI_REGISTRY_IMAGE" json:"image"    yaml:"image"`
		User     string `env:"CI_REGISTRY_USER"  json:"user"     yaml:"user"`
	} `json:"registry" yaml:"registry"`
	Environment struct {
		Name string `env:"CI_ENVIRONMENT_NAME" json:"name" yaml:"name"`
		Slug string `env:"CI_ENVIRONMENT_SLUG" json:"slug" yaml:"slug"`
		Url  string `env:"CI_ENVIRONMENT_URL"  json:"url"  yaml:"url"`
	} `json:"environment" yaml:"environment"`
}

func (i *CiInfo) RenderJson(w io.Writer) error {
	return renderJson(w, i)
}

func (i *CiInfo) RenderYaml(w io.Writer) error {
	return renderYaml(w, i)
}

func getStructInfo(i interface{}) {
	structType := reflect.TypeOf(i)
	var structValue reflect.Value

	if structType.Kind() != reflect.Ptr {
		panic(errors.New("can only deal with pointers"))
	}

	structValue = reflect.ValueOf(i).Elem()
	structType = reflect.TypeOf(structValue.Interface())

	if structType.Kind() != reflect.Struct {
		panic(errors.New("can only deal with structs"))
	}

	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldValue := structValue.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			if fieldValue.CanAddr() {
				getStructInfo(fieldValue.Addr().Interface())
			}
			continue
		}

		envKey := ""

		tag := fieldType.Tag.Get("env")
		if tag != "" {
			directives := strings.Split(tag, ",")
			if len(directives) > 0 {
				envKey = directives[0]
			}
		}

		if envKey == "-" {
			continue
		}

		if envKey == "" {
			envKey = strings.ToUpper(strcase.ToSnake(fieldType.Name))
		}

		_, ok := CiEnvMapping[envKey]
		if !ok {
			continue
		}

		envVal := os.Getenv(envKey)
		if envVal != "" {
			fieldValue.SetString(envVal)
		}
	}
}

func GetCiInfo() (*CiInfo, error) {
	i := CiInfo{}
	getStructInfo(&i)

	return &i, nil
}

func PopulateCiInfo(out interface{}) error {
	getStructInfo(&out)

	return nil
}

func init() {
	// ARTIFACT_DOWNLOAD_ATTEMPTS	8.15	1.9	Number of attempts to download artifacts running a job
	// CI	all	0.4	Mark that job is executed in CI environment
	// CI_CONFIG_PATH	9.4	0.5	The path to CI config file. Defaults to .gitlab-ci.yml
	// CI_DEBUG_TRACE	all	1.7	Whether debug tracing is enabled
	// CI_REPOSITORY_URL	9.0	all	The URL to clone the Git repository
	// CI_SHARED_ENVIRONMENT	all	10.1	Marks that the job is executed in a shared environment (something that is persisted across CI invocations like shell or ssh executor). If the environment is shared, it is set to true, otherwise it is not defined at all.
	// GET_SOURCES_ATTEMPTS	8.15	1.9	Number of attempts to fetch sources running a job
	// GITLAB_CI	all	all	Mark that job is executed in GitLab CI environment
	// GITLAB_FEATURES	10.6	all	The comma separated list of licensed features available for your instance and plan
	// RESTORE_CACHE_ATTEMPTS	8.15	1.9	Number of attempts to restore the cache running a job
	CiEnvMapping = map[string]*CiEnvKey{
		// Additional arguments passed in the ChatOps command
		"CHAT_INPUT": {
			"10.6",
			"*",
			false,
		},

		// Source chat channel which triggered the ChatOps command
		"CHAT_CHANNEL": {
			"10.6",
			"*",
			false,
		},

		// Authentication username of the GitLab Deploy Token,
		// only present if the Project has one related.
		"CI_DEPLOY_USER": {
			"10.8",
			"*",
			false,
		},

		// Authentication password of the GitLab Deploy Token,
		// only present if the Project has one related.
		"CI_DEPLOY_PASSWORD": {
			"10.8",
			"*",
			true,
		},

		// The unique id of the current project that GitLab CI uses internally
		"CI_PROJECT_ID": {
			"*",
			"*",
			false,
		},

		// The full path where the repository is cloned and where the job is run
		"CI_PROJECT_DIR": {
			"*",
			"*",
			false,
		},

		// The project name that is currently being built (actually it is project folder name)
		"CI_PROJECT_NAME": {
			"8.10",
			"0.5",
			false,
		},

		// The project namespace (username or groupname) that is currently being built
		"CI_PROJECT_NAMESPACE": {
			"8.10",
			"0.5",
			false,
		},

		// The namespace with project name
		"CI_PROJECT_PATH": {
			"8.10",
			"0.5",
			false,
		},

		// $CI_PROJECT_PATH lowercased and with everything
		// except 0-9 and a-z replaced with -. Use in URLs and domain names.
		"CI_PROJECT_PATH_SLUG": {
			"9.3",
			"*",
			false,
		},

		// The HTTP address to access project
		"CI_PROJECT_URL": {
			"8.10",
			"0.5",
			false,
		},

		// The project visibility (internal, private, public)
		"CI_PROJECT_VISIBILITY": {
			"10.3",
			"*",
			false,
		},

		// The commit revision for which project is built
		"CI_COMMIT_SHA": {
			"9.0",
			"*",
			false,
		},

		// The branch or tag name for which project is built
		"CI_COMMIT_REF_NAME": {
			"9.0",
			"*",
			false,
		},

		// $CI_COMMIT_REF_NAME lowercased, shortened to 63 bytes,
		// and with everything except 0-9 and a-z replaced with -. No leading / trailing -.
		// Use in URLs, host names and domain names.
		"CI_COMMIT_REF_SLUG": {
			"9.0",
			"*",
			false,
		},

		// The commit tag name. Present only when building tags.
		"CI_COMMIT_TAG": {
			"9.0",
			"0.5",
			false,
		},

		// The full commit message.
		"CI_COMMIT_MESSAGE": {
			"10.8",
			"*",
			false,
		},

		// The title of the commit - the full first line of the message
		"CI_COMMIT_TITLE": {
			"10.8",
			"*",
			false,
		},

		// The description of the commit: the message without first line,
		// if the title is shorter than 100 characters; full message in other case.
		"CI_COMMIT_DESCRIPTION": {
			"10.8",
			"*",
			false,
		},

		// The unique id of the current job that GitLab CI uses internally
		"CI_JOB_ID": {
			"9.0",
			"*",
			false,
		},

		// CI_JOB_MANUAL	8.12	all	The flag to indicate that job was manually started

		// The name of the job as defined in .gitlab-ci.yml
		"CI_JOB_NAME": {
			"9.0",
			"0.5",
			false,
		},

		// The name of the stage as defined in .gitlab-ci.yml
		"CI_JOB_STAGE": {
			"9.0",
			"0.5",
			false,
		},

		// Token used for authenticating with the GitLab Container Registry.
		// Also used to authenticate with multi-project pipelines when triggers are involved.
		"CI_JOB_TOKEN": {
			"9.0",
			"1.2",
			true,
		},

		// Job details URL
		"CI_JOB_URL": {
			"11.1",
			"0.5",
			false,
		},

		// The unique id of the current pipeline that GitLab CI uses internally
		"CI_PIPELINE_ID": {
			"8.10",
			"0.5",
			false,
		},

		// The unique id of the current pipeline scoped to project
		"CI_PIPELINE_IID": {
			"11.0",
			"*",
			false,
		},

		// Indicates how the pipeline was triggered.
		// Possible options are: push, web, trigger, schedule, api, and pipeline.
		// For pipelines created before GitLab 9.5, this will show as unknown
		"CI_PIPELINE_SOURCE": {
			"10.0",
			"*",
			false,
		},

		// Pipeline details URL
		"CI_PIPELINE_URL": {
			"11.1",
			"0.5",
			false,
		},

		// CI_PIPELINE_TRIGGERED	all	all	The flag to indicate that job was triggered

		// The unique id of runner being used
		"CI_RUNNER_ID": {
			"8.10",
			"0.5",
			false,
		},

		// The description of the runner as saved in GitLab
		"CI_RUNNER_DESCRIPTION": {
			"8.10",
			"0.5",
			false,
		},

		// The defined runner tags
		"CI_RUNNER_TAGS": {
			"8.10",
			"0.5",
			false,
		},

		// GitLab Runner version that is executing the current job
		"CI_RUNNER_VERSION": {
			"*",
			"10.6",
			false,
		},

		// GitLab Runner revision that is executing the current job
		"CI_RUNNER_REVISION": {
			"*",
			"10.6",
			false,
		},

		// The OS/architecture of the GitLab Runner executable
		// (note that this is not necessarily the same as the environment of the executor)
		"CI_RUNNER_EXECUTABLE_ARCH": {
			"*",
			"10.6",
			false,
		},

		// The id of the user who started the job
		"GITLAB_USER_ID": {
			"8.12",
			"*",
			false,
		},

		// The email of the user who started the job
		"GITLAB_USER_EMAIL": {
			"8.12",
			"*",
			false,
		},

		// The login username of the user who started the job
		"GITLAB_USER_LOGIN": {
			"10.0",
			"*",
			false,
		},

		// The real name of the user who started the job
		"GITLAB_USER_NAME": {
			"10.0",
			"*",
			false,
		},

		// The name of CI server that is used to coordinate jobs
		"CI_SERVER_NAME": {
			"*",
			"*",
			false,
		},

		// GitLab revision that is used to schedule jobs
		"CI_SERVER_REVISION": {
			"*",
			"*",
			false,
		},

		// GitLab version that is used to schedule jobs
		"CI_SERVER_VERSION": {
			"*",
			"*",
			false,
		},

		// CI_SERVER all	all	Mark that job is executed in CI environment

		// If the Container Registry is enabled it returns
		// the address of GitLab's Container Registry
		"CI_REGISTRY": {
			"8.10",
			"0.5",
			false,
		},

		// If the Container Registry is enabled for the project
		// it returns the address of the registry tied to the specific project
		"CI_REGISTRY_IMAGE": {
			"8.10",
			"0.5",
			false,
		},

		// The password to use to push containers to the GitLab Container Registry
		"CI_REGISTRY_PASSWORD": {
			"9.0",
			"*",
			true,
		},

		// The username to use to push containers to the GitLab Container Registry
		"CI_REGISTRY_USER": {
			"9.0",
			"*",
			false,
		},

		// CI_DISPOSABLE_ENVIRONMENT	all	10.1	Marks that the job is executed in a disposable environment (something that is created only for this job and disposed of/destroyed after the execution - all executors except shell and ssh). If the environment is disposable, it is set to true, otherwise it is not defined at all.

		// The name of the environment for this job
		"CI_ENVIRONMENT_NAME": {
			"8.15",
			"*",
			false,
		},

		// A simplified version of the environment name,
		// suitable for inclusion in DNS, URLs, Kubernetes labels, etc.
		"CI_ENVIRONMENT_SLUG": {
			"8.15",
			"*",
			false,
		},

		// The URL of the environment for this job
		"CI_ENVIRONMENT_URL": {
			"9.3",
			"*",
			false,
		},
	}
}
