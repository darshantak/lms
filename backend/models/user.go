package models

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	EmpType  string `json:"emp_type"`
	Role     string `json:"role"` //admin or user
}
