package types

type PaymentResponse struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}
