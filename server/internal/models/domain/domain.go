package domain

import "time"

type User struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Email               string    `gorm:"unique;not null" json:"email"`
	Name                string    `json:"name"`
	Surname             string    `json:"surname"`
	Username            string    `gorm:"unique;not null" json:"username"`
	Description         string    `json:"description"`
	Photo               string    `json:"photo"` // URL to avatar
	PassHash            string    `gorm:"not null" json:"-"`
	RefreshTokenVersion uint      `json:"-"` // скрываем с фронта
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"` // eager load, если нужно
	Title     string    `gorm:"not null" json:"title"`
	Text      string    `json:"text"`
	Image     string    `json:"image"`
	Category  string    `json:"category"`
	Likes     []Like    `gorm:"constraint:OnDelete:CASCADE" json:"likes,omitempty"`
	Comments  []Comment `gorm:"constraint:OnDelete:CASCADE" json:"comments,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"postId"`
	UserID    uint      `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"` // если нужно
	Text      string    `gorm:"not null" json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Like struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `gorm:"not null" json:"postId"`
	UserID    uint      `gorm:"not null" json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
