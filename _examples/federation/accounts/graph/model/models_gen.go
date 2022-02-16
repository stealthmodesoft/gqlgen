// Code generated by github.com/stealthmodesoft/gqlgen, DO NOT EDIT.

package model

type EmailHost struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (EmailHost) IsEntity() {}

type User struct {
	ID       string     `json:"id"`
	Host     *EmailHost `json:"host"`
	Email    string     `json:"email"`
	Username string     `json:"username"`
}

func (User) IsEntity() {}
