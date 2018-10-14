package types

type PhoneCall struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Number        string `json:"number"`
	Created       string `json:"created"`
	Completed     bool   `json:"completed"`
	CompletedTime string `json:"completedTime"`
}
