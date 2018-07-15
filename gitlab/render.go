package gitlab

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
)

type Renderable interface {
	RenderJson(w io.Writer) error
	RenderYaml(w io.Writer) error
}

func renderJson(w io.Writer, v interface{}) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var indented bytes.Buffer
	err = json.Indent(&indented, j, "", "  ")
	if err != nil {
		return err
	}

	_, err = indented.WriteTo(w)

	return err
}

func renderYaml(w io.Writer, v interface{}) error {
	j, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	w.Write(j)

	return nil
}
