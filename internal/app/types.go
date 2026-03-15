package app

type SkillSpec struct {
	Name        string         `json:"name"`
	Slug        string         `json:"slug,omitempty"`
	Description string         `json:"description"`
	Audience    string         `json:"audience,omitempty"`
	Category    string         `json:"category,omitempty"`
	Tags        []string       `json:"tags,omitempty"`
	Triggers    []string       `json:"triggers,omitempty"`
	Constraints []string       `json:"constraints,omitempty"`
	Checks      []string       `json:"checks,omitempty"`
	Examples    []string       `json:"examples,omitempty"`
	Workflow    []WorkflowStep `json:"workflow,omitempty"`
	Tools       []ToolSpec     `json:"tools"`
}

type WorkflowStep struct {
	Name string `json:"name"`
	Goal string `json:"goal"`
}

type ToolSpec struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Description string   `json:"description"`
	Example     string   `json:"example,omitempty"`
	Inputs      []string `json:"inputs,omitempty"`
	Outputs     []string `json:"outputs,omitempty"`
}
