package db

import (
	"time"
)

type Client struct{
	ID int64 `json:"ClientID"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	PhoneNum string `json:"Phone_Number"`
	PassWord string `json:"Password"`
	Address string `json:"Address"`
}

type Booker struct{
	ID int64 `json:"BookerID"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	PhoneNum string `json:"Phone_Number"`
	PassWord string `json:"Password"`
	Address string `json:"Address"`
}

type Booking struct{
	ID int64 `json:"BookingID"`
	ClientID int64 `json:"ClientID"`
	BookerID int64 `json:"BookerID"`
	DateAndTime time.Time `json:"Date_and_Time"`
	Status string `json:"Status"`
}

type Restaurant struct{
	ID int64 `json:"RestaurantID"`
	Name string `json:"Name"`
	Location string `json:"Location"`
	Contact string `json:"Contact_Number"`
}

type ChatMessage struct{
	ID int64 `json:"MessageID"`
	SenderID int64 `json:"SenderID"`
	ReceiverID int64 `json:"ReceiverID"`
	MessageText string `json:"MessageText"`
	Timestamp time.Time `json:"Timestamp"`
}