package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string        `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password     string        `gorm:"not null" json:"-"`
	Nickname     string        `gorm:"type:varchar(50)" json:"nickname"`
	Categories   []Category    `gorm:"foreignKey:UserID" json:"categories,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions,omitempty"`
}

func (User) TableName() string {
	return "users"
}

// BeforeCreate хеширует пароль перед созданием записи
func (u *User) BeforeCreate(*gorm.DB) error {
	return u.HashPassword()
}

// BeforeUpdate хеширует пароль перед обновлением записи,
// но только если пароль был изменен
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// Проверяем, изменилось ли поле password
	if tx.Statement.Changed("Password") {
		return u.HashPassword()
	}
	return nil
}

// hashPassword хеширует пароль пользователя
func (u *User) HashPassword() error {
	if u.Password == "" {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// ComparePassword проверяет соответствие пароля хешу
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	)
}
