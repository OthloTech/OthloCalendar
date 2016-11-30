package models

import (
	//"errors"
	//"time"
)

type Event struct {
	EventId     int    `json:"event_id"`
	EventURL    string `json:"event_url"`
	EventType   string `json:"event_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Catch       string `json:"catch"`
	HashTag     string `json:"hash_tag"`
	Accepted    int    `json:"accepted"`
	Waiting     int    `json:"waiting"`
	Limit       int    `json:"limit"`
	URL         string `json:"url"`

	OwnerID          int    `json:"owner_id"`
	OwnerNickname    string `json:"owner_nickname"`
	OwnerTwitterID   string `json:"owner_twitter_id"`
	OwnerDisplayName string `json:"owner_display_name"`

	UpdatedAt time.Time `json:"updated_at"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`

	Address string `json:"address"`
	Place   string `json:"place"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
}

type EventModel struct{}
