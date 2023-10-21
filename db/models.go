package db

import (
	"time"
)

type Client struct{
	ID int64 `json:"id"`
	Name string `json:"name_client"`
	Email string `json:"email"`
	PhoneNum string `json:"phone_number"`
	Password string `json:"pwd"`
	Address string `json:"address_client"`
	CreatedAt time.Time `json:"created_at"`
}

type Booker struct{
	ID int64 `json:"id"`
	Name string `json:"name_booker"`
	Email string `json:"email"`
	PhoneNum string `json:"phone_number"`
	Password string `json:"pwd"`
	Address string `json:"address_client"`
	CreatedAt time.Time `json:"created_at"`
}

type Booking struct{
	ID int64 `json:"id"`
	ClientID int64 `json:"client_id"`
	BookerID int64 `json:"booker_id"`
	RestaurantID int64 `json:"restaurant_id"`
	DateAndTime time.Time `json:"date_and_time"`
	Status string `json:"status_booking"`
}

type Restaurant struct{
	ID int64 `json:"id"`
	Name string `json:"name_restaurant"`
	Location string `json:"location_restaurant"`
	Contact string `json:"contact_number"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatMessage struct{
	ID int64 `json:"messageID"`
	SenderID int64 `json:"senderID"`
	ReceiverID int64 `json:"receiverID"`
	MessageText string `json:"messageText"`
	Timestamp time.Time `json:"timestamp"`
}