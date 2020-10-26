package storage

type Request struct {
	Body int `json:"body"`
}

type fan struct {
	Links map[string]string `json:"links"`
}

type RequestChild struct {
	Request
	fan
	ChildID string `json:"childID"`
}

type ArrayOne []string
