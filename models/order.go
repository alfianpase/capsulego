package models

type Order struct {
	OrgID           int64  `json:"org_id"`
	MemberCellPhone string `json:"member_cell_phone"`
}

type Price struct {
	PriceValue    float64
	OrgID         int64
	PriceCategory string
	Status        string
}

type OrderContract interface {
	SaveOrder(Order) error
	getPrice(orgID int64, priceCategory string, status string) (float64, error)
	existingOrder(Order) (int8, error)
}
type OrderRepository interface {
	SaveOrder(Order) error
	// getPrice(orgID int64, priceCategory string, status string) (float64, error)
	// existingOrder(Order) (int8, error)
}

// type OrderRepo interface {
// 	SaveOrder(Order) error
// 	getPrice(orgID int64, priceCategory string, status string) (float64, error)
// 	existingOrder(Order) (int8, error)
// }
