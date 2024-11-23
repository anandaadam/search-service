package models

type SearchUsersResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
