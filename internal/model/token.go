package model

import "time"

type Token struct {
	ID     uint32 `json:"id"`
	UserID uint32 `json:"userID"`

	Value      string        `json:"value"`
	CreatedAt  time.Time     `json:"createdAt"`
	Lifetime   time.Duration `json:"lifeTime"`
	Expiration time.Time     `json:"expiration"`
	RevokedAt  time.Time     `json:"revokedAt"`
}
