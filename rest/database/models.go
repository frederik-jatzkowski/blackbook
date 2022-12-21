package database

import "time"

type Group struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	// associations
	Invitations []Invitation
	Users       []User `gorm:"many2many:user_groups;"`
	// timestamps
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Invitation struct {
	ID      uint   `gorm:"primarykey"`
	Message string `gorm:"not null"`
	// associations
	SenderID   uint  `gorm:"not null"`
	Sender     User  `gorm:"foreignkey:SenderID"`
	ReceiverID uint  `gorm:"not null"`
	Receiver   User  `gorm:"foreignkey:ReceiverID"`
	GroupID    uint  `gorm:"not null"`
	Group      Group `gorm:"foreignkey:GroupID"`
	// timestamps
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	ID        uint   `gorm:"primarykey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	// associations
	Groups              []Group      `gorm:"many2many:user_groups;"`
	SentInvitations     []Invitation `gorm:"foreignkey:SenderID"`
	ReceivedInvitations []Invitation `gorm:"foreignkey:ReceiverID"`
	// authentication
	Salt              string `gorm:"not null"`
	Hash              string `gorm:"not null"`
	Active            bool   `gorm:"not null"`
	ActivationCode    string
	Session           string
	SessionExpiration time.Time
	// timestamps
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
