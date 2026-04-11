package models

// Registration domain model. Persistence moved to repositories.
type Registration struct {
	ID       int
	Status   string
	Event_id int64
	User_id  int64
}
