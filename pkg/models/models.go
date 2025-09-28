package models

import "time"

type EventsPageData struct {
	Data []struct {
		ID          int       `json:"id"`
		DocumentID  string    `json:"documentId"`
		Day         string    `json:"Day"`
		DayOfDate   int       `json:"Day_Of_Date"`
		MonthYear   string    `json:"Month_Year"`
		Time        string    `json:"Time"`
		EventName   string    `json:"Event_Name"`
		Address     string    `json:"Address"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		PublishedAt time.Time `json:"publishedAt"`
		Details     string    `json:"Details"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Page      int `json:"page"`
			PageSize  int `json:"pageSize"`
			PageCount int `json:"pageCount"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}
