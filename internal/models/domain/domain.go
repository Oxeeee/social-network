package domain

import "time"

type User struct {
	ID              uint
	Email           string
	Name            string
	Surname         string
	Username        string
	Description     string
	Photo           string // URL или путь к файлу
	PassHash        string
	JWTRefreshToken string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Post struct {
	ID        uint
	UserID    uint // связь с автором
	Title     string
	Text      string
	Image     string    // URL к картинке
	Category  string    // можно enum-ом потом
	Likes     []Like    // отдельная таблица Like
	Comments  []Comment // связь с комментами
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint
	PostID    uint
	UserID    uint
	Text      string
	CreatedAt time.Time
}

type Like struct {
	ID        uint
	PostID    uint
	UserID    uint
	CreatedAt time.Time
}
