package model

type FormSubmission struct {
	GameID   string
	Score    string
	Username string
	Initials string
}

// / Struct for a user searching on the website
type UserSearchForm struct {
	Fields        []string
	FieldErrors   map[string]string
	GeneralErrors []string
	IsUsername    bool
}
