package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/plouc/go-gitlab-client/gitlab"
	"gopkg.in/yaml.v2"
)

func metaOutput(meta *gitlab.ResponseMeta, withPagination bool) {
	if !verbose {
		return
	}

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

func jsonOutput(v interface{}) {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	var indented bytes.Buffer
	json.Indent(&indented, j, "", "  ")

	indented.WriteTo(output)
}

func yamlOutput(v interface{}) {
	j, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}

	output.Write(j)
}
