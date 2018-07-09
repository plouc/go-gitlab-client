package cmd

var resources = map[string][]string{
	"group":            {"id"},
	"project":          {"id"},
	"user":             {"id"},
	"runner":           {"id"},
	"namespace":        {"id"},
	"project-branch":   {"project_id", "branch_name"},
	"project-pipeline": {"project_id", "pipeline_id"},
	"group-var":        {"group_id", "var_key"},
	"project-var":      {"project_id", "var_key"},
	"project-hook":     {"project_id", "hook_id"},
}

var resourceTypes []string

func init() {
	for key := range resources {
		resourceTypes = append(resourceTypes, key)
	}
}

func isValidResourceType(resourceType string) bool {
	for _, t := range resourceTypes {
		if t == resourceType {
			return true
		}
	}

	return false
}
