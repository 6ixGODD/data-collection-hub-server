package models

type CacheList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}
