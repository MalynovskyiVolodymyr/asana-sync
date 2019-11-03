package models

type (
	// AsanaProject Project - model
	AsanaProject struct {
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		Gid          string `json:"gid"`
		ResourceType string `json:"resource_type"`
	}
	// AsanaTask Task -model
	AsanaTask struct {
		ID             int64           `json:"id"`
		Name           string          `json:"name"`
		Gid            string          `json:"gid"`
		Assignee       AsanaAssignedTo `json:"assignee"`
		ResourceType   string          `json:"resource_type"`
		AssigneeStatus string          `json:"assignee_status"`
		Completed      bool            `json:"completed"`
		CreatedAt      string          `json:"created_at"`
		DueAt          string          `json:"due_at"`
		DueOn          string          `json:"due_on"`
		Notes          string          `json:"notes"`
	}
	// AsanaAssignedTo required
	AsanaAssignedTo struct {
		ID           int64          `json:"id"`
		Gid          int64          `json:"gid"`
		Email        string         `json:"email"`
		Name         string         `json:"name"`
		Photo        string         `json:"photo"`
		ResourceType string         `json:"resource_type"`
		Workspaces   []AsanaProject `json:"workspaces"`
	}
	// AsanaUser asana user object
	AsanaUser struct {
		ID           int64  `json:"id"`
		Gid          string `json:"gid"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		ResourceType string `json:"resource_type"`
	}
)
