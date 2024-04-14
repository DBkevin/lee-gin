package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Email    string `json:"email" gorm:"null;comment:用户邮箱可选"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}
