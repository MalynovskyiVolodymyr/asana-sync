package models

// TractionTaskModel - models
type (
	// TractionTaskModel - models
	TractionTaskModel struct {
		ID           int64              `json:"Id"`
		Type         string             `json:"Type"`
		Name         string             `json:"Name"`
		DetailsURL   string             `json:"DetailsUrl"`
		Origin       string             `json:"Origin"`
		OriginID     int64              `json:"OriginId"`
		DueDate      string             `json:"DueDate"`
		Owner        TractionOwnerModel `json:"Owner"`
		CompleteTime string             `json:"CompleteTime"`
		CreateTime   string             `json:"CreateTime"`
		Complete     bool               `json:"Complete"`
		TodoType     string             `json:"TodoType"`
		Ordering     int64              `json:"Ordering"`
		Permission   string             `json:"Permission"`
	}

	// TractionOwnerModel - model
	TractionOwnerModel struct {
		ID       int64  `json:"Id"`
		Type     string `json:"Type"`
		Name     string `json:"Name"`
		Key      string `json:"Key"`
		ImageURL string `json:"ImageUrl"`
	}

	// TractionUserModel - model
	TractionUserModel struct {
		UserID          int64                `json:"Id"`
		Name            string               `json:"Name"`
		Description     string               `json:"Description"`
		OrganizationID  int64                `json:"OrganizationId"`
		Organization    TractionOrganization `json:"Organization"`
		Email           string               `json:"Email"`
		ImageURL        string               `json:"ImageUrl"`
		AltIcon         TractionAltIcon      `json:"AltIcon"`
		ResultType      string               `json:"ResultType"`
		ItemName        string               `json:"ItemName"`
		ItemImageURL    string               `json:"ItemImageUrl"`
		ItemValue       string               `json:"ItemValue"`
		ItemDescription string               `json:"ItemDescription"`
	}

	TractionOrganization struct {
		OrganizationID int64 `json:"Id"`
	}

	TractionAltIcon struct {
	}

	TractionMittingModel struct {
		ID   int64  `json:"Id"`
		Name string `json:"Name"`
	}
)
