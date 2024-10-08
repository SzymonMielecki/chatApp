package types

import (
	pb "github.com/SzymonMielecki/GoGrpcKafkaGormDemo/usersService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PasswordHash string
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:           uint32(u.Model.ID),
		Username:     u.Username,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}

type Message struct {
	gorm.Model
	Content  string
	SenderID uint
}
