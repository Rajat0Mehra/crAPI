// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"CreatedAt"`
	Author    *User     `json:"author"`
}

type CommentInput struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type Coupon struct {
	CouponCode string    `json:"coupon_code"`
	Amount     string    `json:"amount"`
	CreatedAt  time.Time `json:"CreatedAt"`
}

type CouponInput struct {
	CouponCode string `json:"coupon_code"`
	Amount     string `json:"amount"`
}

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Author    *User      `json:"author"`
	Comments  []*Comment `json:"comments"`
	Authorid  string     `json:"authorid"`
	CreatedAt time.Time  `json:"CreatedAt"`
}

type PostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Nickname      string    `json:"nickname"`
	Email         string    `json:"email"`
	Vehicleid     string    `json:"vehicleid"`
	ProfilePicURL string    `json:"profile_pic_url"`
	CreatedAt     time.Time `json:"created_at"`
}

type UserInput struct {
	Nickname      string     `json:"nickname"`
	Email         string     `json:"email"`
	VehicleID     string     `json:"vehicle_id"`
	ProfilePicURL string     `json:"profile_pic_url"`
	CreatedAt     *time.Time `json:"created_at"`
}
