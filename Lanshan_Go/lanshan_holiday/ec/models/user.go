package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}

//将user.go 用户模型实现 JwtUser 接口，
//后续其他的用户模型都可以通过实现 JwtUser 接口，
//来调用 CreateToken() 颁发 Token
