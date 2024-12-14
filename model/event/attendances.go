package event

import "time"

type Attendances struct {
	IsArrival  bool `json:"is_arrival"`
	ArrivalTime time.Time `json:"arrival_time"`
}

type AttendancesResponse struct{
	Message string `json:"message"`
}

type RequestVote struct{
	UserId int `json:"user_id"`
	Option string `json:"option"`
}