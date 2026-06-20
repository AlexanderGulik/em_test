package dto

type SubResponse struct {
	Id        string `json:"id_sub" db:"id_sub" example:"1"`
	NameSub   string `json:"service_name" db:"service_name" example:"Netflix"`
	PriceMonth int32 `json:"price" db:"price_month" example:"1000"`
	UserId    string `json:"user_id" db:"user_uuid" example:"d2f6cf95-ea12-4d74-b7fb-4215c94475a7"`
	DateStart string `json:"start_date" db:"start_date" example:"05-2026"`
	DateEnd   string `json:"end_date" db:"end_date" example:"06-2026"`
}


type SubRequest struct {
	NameSub   string `json:"service_name" db:"service_name" example:"Netflix"`
	PriceMonth int32 `json:"price" db:"price_month" example:"1000"`
	UserId    string `json:"user_id" db:"user_uuid" example:"d2f6cf95-ea12-4d74-b7fb-4215c94475a7"`
	DateStart string `json:"start_date" db:"start_date" example:"05-2026"`
	DateEnd   string `json:"end_date" db:"end_date" example:"06-2026"`
}

type SubUpdateRequest struct {
	ServiceName *string `json:"service_name,omitempty" example:"Netflix"`
	PriceMonth  *int32  `json:"price,omitempty" example:"1400"`
	UserUUID    *string `json:"user_id,omitempty" example:"d2f6cf95-ea12-4d74-b7fb-4215c94475a7"`
	StartDate   *string `json:"start_date,omitempty" example:"05-2026"`
	EndDate     *string `json:"end_date,omitempty" example:"06-2026"`
}

type TotalSumRequest struct {
	UserID      string `json:"user_id"`
	ServiceName string `json:"service_name"`
	DateStart   string `json:"start_date" example:"05-2026"`
	DateEnd     string `json:"end_date" example:"08-2026"`
}
