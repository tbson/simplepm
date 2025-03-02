package schema

import (
	"encoding/json"
	"src/common/ctype"
	"src/util/dictutil"
	"time"

	"gorm.io/datatypes"
)

type GitCommit struct {
	ID            string    `json:"id" cql:"id"`
	CommitID      string    `json:"commit_id" cql:"commit_id"`
	CommitURL     string    `json:"commit_url" cql:"commit_url"`
	CommitMessage string    `json:"commit_message" cql:"commit_message"`
	CreatedAt     time.Time `json:"created_at" cql:"created_at"`
}

type GitPush struct {
	ID         string      `json:"id" cql:"id"`
	GitBranch  string      `json:"git_branch" cql:"git_branch"`
	GitCommits []GitCommit `json:"git_commits" cql:"git_commits"`
}

func (g *GitPush) New(data ctype.Dict) {
	g.ID = dictutil.GetValue[string](data, "id")
	g.GitBranch = dictutil.GetValue[string](data, "git_branch")
	commits := dictutil.GetValue[[]map[string]interface{}](data, "git_commits")
	for _, commit := range commits {
		g.GitCommits = append(g.GitCommits, GitCommit{
			ID:            dictutil.GetValue[string](commit, "id"),
			CommitID:      dictutil.GetValue[string](commit, "commit_id"),
			CommitURL:     dictutil.GetValue[string](commit, "commit_url"),
			CommitMessage: dictutil.GetValue[string](commit, "commit_message"),
			CreatedAt:     dictutil.GetValue[time.Time](commit, "created_at"),
		})
	}
}

func (g *GitPush) ToDict() ctype.Dict {
	commits := []ctype.Dict{}
	for _, commit := range g.GitCommits {
		commits = append(commits, ctype.Dict{
			"id":             commit.ID,
			"commit_id":      commit.CommitID,
			"commit_url":     commit.CommitURL,
			"commit_message": commit.CommitMessage,
			"created_at":     commit.CreatedAt,
		})
	}
	return ctype.Dict{
		"id":          g.ID,
		"git_branch":  g.GitBranch,
		"git_commits": commits,
	}
}

type GitPR struct {
	ID         string     `json:"id" cql:"id"`
	Title      string     `json:"title" cql:"title"`
	FromBranch string     `json:"from_branch" cql:"from_branch"`
	ToBranch   string     `json:"to_branch" cql:"to_branch"`
	URL        string     `json:"url" cql:"url"`
	MergedAt   *time.Time `json:"merged_at" cql:"merged_at"`
	State      string     `json:"state" cql:"state"`
}

func (g *GitPR) New(data ctype.Dict) {
	g.ID = dictutil.GetValue[string](data, "id")
	g.Title = dictutil.GetValue[string](data, "title")
	g.FromBranch = dictutil.GetValue[string](data, "from_branch")
	g.ToBranch = dictutil.GetValue[string](data, "to_branch")
	g.URL = dictutil.GetValue[string](data, "url")
	g.MergedAt = dictutil.GetValue[*time.Time](data, "merged_at")
	g.State = dictutil.GetValue[string](data, "state")
}

func (g *GitPR) ToDict() ctype.Dict {
	return ctype.Dict{
		"id":          g.ID,
		"title":       g.Title,
		"from_branch": g.FromBranch,
		"to_branch":   g.ToBranch,
		"url":         g.URL,
		"merged_at":   g.MergedAt,
		"state":       g.State,
	}
}

type Message struct {
	ID         string  `json:"id"`
	UserID     uint    `json:"user_id"`
	TaskID     uint    `json:"task_id"`
	ProjectID  uint    `json:"project_id"`
	Content    string  `json:"content"`
	Type       string  `json:"type"`
	GitPush    GitPush `json:"git_push"`
	GitPR      GitPR   `json:"git_pr"`
	UserName   string  `json:"user_name"`
	UserAvatar string  `json:"user_avatar"`
	UserColor  string  `json:"user_color"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"update_at"`
}

type Attachment struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileURL   string `json:"file_url"`
	FileSize  int    `json:"file_size"`
	CreatedAt string `json:"created_at"`
}

type Change struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	TenantID     uint           `gorm:"not null" json:"tenant_id"`
	ProjectID    uint           `gorm:"not null" json:"project_id"`
	TaskID       uint           `gorm:"not null" json:"task_id"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	UserFullName string         `gorm:"type:text;not null;default:''" json:"user_full_name"`
	SourceType   string         `gorm:"type:text;not null" json:"source_type"`
	SourceID     uint           `gorm:"not null" json:"source_id"`
	SourceTitle  string         `gorm:"type:text;not null;default:''" json:"source_title"`
	Action       string         `gorm:"type:text;not null" json:"action"`
	Value        datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"value"`
	CreatedAt    time.Time      `json:"created_at"`
}

func NewChange(data ctype.Dict) *Change {
	valueJSON, err := json.Marshal(data["Value"])
	if err != nil {
		panic("Failed to marshal Value")
	}
	return &Change{
		TenantID:     dictutil.GetValue[uint](data, "TenantID"),
		ProjectID:    dictutil.GetValue[uint](data, "ProjectID"),
		TaskID:       dictutil.GetValue[uint](data, "TaskID"),
		UserID:       dictutil.GetValue[uint](data, "UserID"),
		UserFullName: dictutil.GetValue[string](data, "UserFullName"),
		SourceType:   dictutil.GetValue[string](data, "SourceType"),
		SourceID:     dictutil.GetValue[uint](data, "SourceID"),
		SourceTitle:  dictutil.GetValue[string](data, "SourceTitle"),
		Action:       dictutil.GetValue[string](data, "Action"),
		Value:        datatypes.JSON(valueJSON),
	}
}
