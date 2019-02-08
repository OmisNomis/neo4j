package db

// User reperesents a user node
type User struct {
	Name   string
	Mobile string
}

// Group reperesents a group node
type Group struct {
	Name string
}

// Alert represents an alert node
type Alert struct {
	ID string
}
