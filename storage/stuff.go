package storage

type Request struct {
	Body int `json:"body"`
}

type RequestChild struct {
	Request
	ChildID string `json:"childID`
}
