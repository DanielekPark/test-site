package models

import "time"

type EventsErrorMessage struct {
	Data []struct {
		Day,
		DayOfDate,
		MonthYear,
		Time,
		EventName,
		Address string
	}
}

type EventsPageData struct {
	Data []struct {
		// ID         int    `json:"id"`
		// DocumentID string `json:"documentId"`
		Day       string `json:"Day"`
		DayOfDate int    `json:"Day_Of_Date"`
		MonthYear string `json:"Month_Year"`
		Time      string `json:"Time"`
		EventName string `json:"Event_Name"`
		Address   string `json:"Address"`
		// CreatedAt   time.Time `json:"createdAt"`
		// UpdatedAt   time.Time `json:"updatedAt"`
		// PublishedAt time.Time `json:"publishedAt"`
		Details string `json:"Details"`
	} `json:"data"`
	// Meta struct {
	// 	Pagination struct {
	// 		Page      int `json:"page"`
	// 		PageSize  int `json:"pageSize"`
	// 		PageCount int `json:"pageCount"`
	// 		Total     int `json:"total"`
	// 	} `json:"pagination"`
	// } `json:"meta"`
}

// About Page
type TribeMembers struct {
	Data []struct {
		ID          int       `json:"id"`
		DocumentID  string    `json:"documentId"`
		Email       string    `json:"Email"`
		Name        string    `json:"Name"`
		Position    string    `json:"Position"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		PublishedAt time.Time `json:"publishedAt"`
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

// Membership Page
type MembershipDocs struct {
	Data []struct {
		ID          int       `json:"id"`
		DocumentID  string    `json:"documentId"`
		STIPhotoID  string    `json:"STI_Photo_ID"`
		UpdateForm  string    `json:"Update_Form"`
		FeeSchedule string    `json:"Fee_Schedule"`
		Enrollment  string    `json:"Enrollment"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		PublishedAt time.Time `json:"publishedAt"`
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

// Programs Page
type ProgramDocs struct {
	Data []struct {
		ID                  int       `json:"id"`
		DocumentID          string    `json:"documentId"`
		Application         string    `json:"Application"`
		ParticipatingStores string    `json:"Participating_Stores"`
		NonFoodItems        string    `json:"Non_Food_Items"`
		IncomeLimit         string    `json:"Income_Limit"`
		RequiredDocs        string    `json:"Required_Docs"`
		CreatedAt           time.Time `json:"createdAt"`
		UpdatedAt           time.Time `json:"updatedAt"`
		PublishedAt         time.Time `json:"publishedAt"`
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
