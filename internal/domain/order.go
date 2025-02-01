package domain

type Order struct {
	ID            int    `json:"id"`
	Status        bool   `json:"status"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	DeliveryDate  string `json:"delivery_date"`
	DeliveryTime  string `json:"delivery_time"`
	PlaceDelivery string `json:"place_delivery"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	Comment       string `json:"comment"`
}
