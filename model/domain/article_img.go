package domain

import "time"

type ArticleImg struct {
	Id        int
	ArticleId int
	Filename  string
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
