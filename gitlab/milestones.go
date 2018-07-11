package gitlab

type Milestone struct {
	Id          int    `json:"id,omitempty" yaml:"id,omitempty"`
	Iid         int    `json:"iid,omitempty" yaml:"iid,omitempty"`
	ProjectID   int    `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	State       string `json:"state,omitempty" yaml:"state,omitempty"`
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	StartDate   string `json:"start_date,omitempty" yaml:"start_date,omitempty"`
	DueDate     string `json:"due_date,omitempty" yaml:"due_date,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
}
