package models

type PaymentData struct {
	Product   []string  `json:"product"` // di isi dengan contract
	Qty       []int8    `json:"qty"`     // di isi dengan jumlah contract
	Price     []float64 `json:"price"`   // di isi dengan harga contract
	NotifyURL string    `json:"notifyUrl"` // di isi dengan url notify
}
