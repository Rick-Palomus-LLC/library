package types

type PaymentEvent struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Token  string `json:"token"`
}
