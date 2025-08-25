package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	IsProgress  bool      `json:"isProgress"`
	IsDone      bool      `json:"isDone"`
}
