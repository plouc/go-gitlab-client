package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	variablesUrl = "/:type/:id/variables"      // Get list of a project/group variables.
	variableUrl  = "/:type/:id/variables/:key" // Get detail of project/group variable.
)

type Variable struct {
	Key              string `json:"key"`
	Value            string `json:"value"`
	Protected        bool   `json:"protected"`
	EnvironmentScope string `json:"environment_scope,omitempty"`
}

func (g *Gitlab) getVariables(resourceType, id string, o *PaginationOptions) ([]*Variable, *ResponseMeta, error) {
	u := g.ResourceUrl(variablesUrl, map[string]string{
		":type": resourceType,
		":id":   id,
	})
	if o != nil {
		q := u.Query()

		if o.Page != 1 {
			q.Set("page", strconv.Itoa(o.Page))
		}
		if o.PerPage != 0 {
			q.Set("per_page", strconv.Itoa(o.PerPage))
		}

		u.RawQuery = q.Encode()
	}

	var variables []*Variable

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &variables)
	}

	return variables, meta, err
}

func (g *Gitlab) ProjectVariables(projectId string, o *PaginationOptions) ([]*Variable, *ResponseMeta, error) {
	return g.getVariables("projects", projectId, o)
}

func (g *Gitlab) GroupVariables(groupId string, o *PaginationOptions) ([]*Variable, *ResponseMeta, error) {
	return g.getVariables("groups", groupId, o)
}

func (g *Gitlab) getVariable(resourceType, projectId, varKey string) (*Variable, *ResponseMeta, error) {
	u := g.ResourceUrl(variableUrl, map[string]string{
		":type": "projects",
		":id":   projectId,
		":key":  varKey,
	})

	variable := new(Variable)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &variable)
	}

	return variable, meta, err
}

func (g *Gitlab) ProjectVariable(projectId, varKey string) (*Variable, *ResponseMeta, error) {
	return g.getVariable("projects", projectId, varKey)
}

func (g *Gitlab) GroupVariable(groupId, varKey string) (*Variable, *ResponseMeta, error) {
	return g.getVariable("groups", groupId, varKey)
}

func (g *Gitlab) addVariable(resourceType, id string, variable *Variable) (*Variable, *ResponseMeta, error) {
	u := g.ResourceUrl(variablesUrl, map[string]string{
		":type": resourceType,
		":id":   id,
	})

	variableJson, err := json.Marshal(variable)
	if err != nil {
		return nil, nil, err
	}

	var createdVariable *Variable
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), variableJson)
	if err == nil {
		err = json.Unmarshal(contents, &createdVariable)
	}

	return createdVariable, meta, err
}

func (g *Gitlab) AddProjectVariable(projectId string, variable *Variable) (*Variable, *ResponseMeta, error) {
	return g.addVariable("projects", projectId, variable)
}

func (g *Gitlab) AddGroupVariable(groupId string, variable *Variable) (*Variable, *ResponseMeta, error) {
	return g.addVariable("groups", groupId, variable)
}

func (g *Gitlab) removeVariable(resourceType, id, varKey string) (*ResponseMeta, error) {
	u := g.ResourceUrl(variableUrl, map[string]string{
		":type": resourceType,
		":id":   id,
		":key":  varKey,
	})

	_, meta, err := g.buildAndExecRequest("DELETE", u.String(), nil)

	return meta, err
}

func (g *Gitlab) RemoveProjectVariable(projectId, varKey string) (*ResponseMeta, error) {
	return g.removeVariable("projects", projectId, varKey)
}

func (g *Gitlab) RemoveGroupVariable(groupId, varKey string) (*ResponseMeta, error) {
	return g.removeVariable("groups", groupId, varKey)
}
