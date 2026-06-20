package dto

type SubRequest struct {
	Id string `json:"id_sub" db:"id_sub"`
	NameSub string `json:"service_name" db:"service_name"`
	PriceMonth int32 `json:"price" db:"price_month"`
	UserId string `json:"user_id" db:"user_uuid"`
	DateStart string `json:"start_date" db:"start_date"`
	DateEnd string `json:"end_date" db:"end_date"`
}

type SubUpdateRequest struct {
    ServiceName *string  `json:"service_name,omitempty"`
    PriceMonth  *int32   `json:"price,omitempty"`
    UserUUID    *string  `json:"user_id,omitempty"`
    StartDate   *string  `json:"start_date,omitempty"`
    EndDate     *string  `json:"end_date,omitempty"`
}

type TotalSumRequest struct {
    UserID      string `json:"user_id"`       
    ServiceName string `json:"service_name"`   
    DateStart  string `json:"start_date"`    
    DateEnd    string `json:"end_date"`      
}
