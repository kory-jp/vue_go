package controllers

import (
	"time"
)

type Context interface {
	ShouldBindJSON(obj interface{}) error
	Done() <-chan struct{}
	BindJSON(obj interface{}) error
	Deadline() (deadline time.Time, ok bool)
	Err() error
	// Value(key interface{}) interface{}
	Value(key any) any
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
}
