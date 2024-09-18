package domain

import "time"

type Department struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	ParentID   *int      `json:"parent_id"`
	Flags      int8      `json:"flags"`
	Hierarchhy string    `json:"hierarchy"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
