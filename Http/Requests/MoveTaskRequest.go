package Requests

type MoveTaskRequest struct {
	ColumnID int `json:"column_id" binding:"required"`
}
