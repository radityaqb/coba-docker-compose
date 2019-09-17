package main

type (
	RequestData struct {
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	ResponseData struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Status int    `json:"status"`
	}
)
