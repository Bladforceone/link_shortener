package link

import (
	"go_pro_api/internal/stat"
	"gorm.io/gorm"
	"math/rand/v2"
)

type Link struct {
	gorm.Model
	URL   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint: OnUpdate:CASCADE, OnDelete: SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		URL: url,
	}
	link.GenerateHash()
	return link
}

func (l *Link) GenerateHash() {
	l.Hash = RandStringRunes(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
