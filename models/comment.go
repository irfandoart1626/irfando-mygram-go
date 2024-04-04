package models

type Comment struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	PhotoID int    `json:"photo_id"`
	Message string `json:"message"`
}
