package gitlab

type Milestone struct {
	Id          int    `json:"id,omitempty" yaml:"id,omitempty"`
	Iid         int    `json:"iid,omitempty" yaml:"iid,omitempty"`
	GroupId     int    `json:"group_id,omitempty" yaml:"group_id,omitempty"`
	ProjectId   int    `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	State       string `json:"state,omitempty" yaml:"state,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	DueDate     string `json:"due_date,omitempty" yaml:"due_date,omitempty"`
	StartDate   string `json:"start_date,omitempty" yaml:"start_date,omitempty"`
	WebUrl      string `json:"web_url,omitempty" yaml:"web_url,omitempty"`
}
