package models

type User struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
