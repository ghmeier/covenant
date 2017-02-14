package models

import (
	"database/sql"
	"errors"

	"github.com/pborman/uuid"
)

/*Subscription is the representation of subscription entries in covenant*/
type Subscription struct {
	ID uuid.UUID `json: "id"`
	UserID uuid.UUID `json: "userId"` 
	Status SubscriptionStatus `json:"status"` 
	CreatedAt string `json:"createdAt"` 
	StartAt string `json:"startAt"` 
	RoasterID uuid.UUID `json: "roasterID"`
	ItemID uuid.UUID `json: "itemID"`
}

/*NewSubscription creates a new subscription with a new uuid*/
func NewSubscription(userID uuid.UUID, createdAt string, startAt string, roasterID uuid.UUID, itemID uuid.UUID) *Subscription {
	return &Subscription{ 
		ID: uuid.NewUUID(), 
		UserID: userID, 
		Status: ACTIVE, 
		CreatedAt: createdAt, 
		StartAt: startAt, 
		RoasterID: roasterID,
		ItemID: itemID, 
	}
}

/*SubscriptionFromSql returns a new subscription slice from a group of sql rows*/
func SubscriptionFromSql(rows *sql.Rows) ([]*Subscription, error) {
	subscription := make([]*Subscription,0)

	for rows.Next() {
		s := &Subscription{}
		var sStatus string
		rows.Scan(&s.ID, &s.UserID, &sStatus, &s.CreatedAt, &s.StartAt, &s.RoasterID, &s.ItemID)

		var ok bool
		s.Status, ok = toSubscriptionType(sStatus)
		if !ok {
			return nil, errors.New("invalid subscriptionStatus string")
		}
		subscription = append(subscription, s)
	}


	return subscription, nil
}

/*toSubscriptionType checks whether a given status is valid*/
func toSubscriptionType(s string) (SubscriptionStatus, bool) {
	switch s {
	case ACTIVE:
		return ACTIVE, true
	case PENDING:
		return PENDING, true
	case CANCELLED:
		return CANCELLED, true
	case INACTIVE:
		return INACTIVE, true
	default:
		return "", false
	}
}

/*SubscriptionStatus is an enum wrapper for valid subscription type*/
type SubscriptionStatus string 

/*Valid Subscription Statuses*/
const (
	ACTIVE = "ACTIVE"
	PENDING = "PENDING"
	CANCELLED = "CANCELLED"
	INACTIVE = "INACTIVE"
)
