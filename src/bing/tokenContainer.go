package ding

import "time"

type tokenContainer struct {
	accessToken string
	expiresIn   int64
}

func (t *tokenContainer) SetAccessToken(token string) *tokenContainer {
	t.accessToken = token
	return t
}

func (t *tokenContainer) SetExpiresIn(timestamp int64) *tokenContainer {
	t.expiresIn = timestamp
	return t
}

func (t *tokenContainer) CheckToken() bool {
	if t.tokenIsEmpty() || t.timeIsExpired() {
		return false
	}
	return true
}

func (t *tokenContainer) tokenIsEmpty() bool {
	return t.accessToken == ""
}

func (t *tokenContainer) timeIsExpired() bool {
	return time.Now().Unix() >= t.expiresIn
}
