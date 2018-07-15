package output

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"gopkg.in/yaml.v2"
	"io"
)

func Meta(meta *gitlab.ResponseMeta, withPagination bool) {
	color.Yellow("\nResponse meta")

	fmt.Println("")
	fmt.Printf("  Method          %s\n", meta.Method)
	fmt.Printf("  Url             %s\n", meta.Url)
	fmt.Printf("  StatusCode      %d\n", meta.StatusCode)
	fmt.Printf("  Request id      %s\n", meta.RequestId)
	fmt.Printf("  Runtime         %f\n", meta.Runtime)

	if withPagination {
		fmt.Printf("  Page            %d\n", meta.Page)
		fmt.Printf("  Items per page  %d\n", meta.PerPage)
		fmt.Printf("  Previous page   %d\n", meta.PrevPage)
		fmt.Printf("  Next page       %d\n", meta.NextPage)
		fmt.Printf("  Total pages     %d\n", meta.TotalPages)
		fmt.Printf("  Total           %d\n", meta.Total)
	}

	fmt.Println("")
}

func Json(w io.Writer, v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	var indented bytes.Buffer
	json.Indent(&indented, j, "", "  ")

	indented.WriteTo(w)
}

func Yaml(w io.Writer, v interface{}) {
	j, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Write(j)
}
