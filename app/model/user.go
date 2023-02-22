package model

// User 用户模型
type User struct {
	BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	CommonTimestampsField
}
