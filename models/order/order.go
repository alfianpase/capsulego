package models

type Order struct {
	OrgID           int64  `json:"org_id"`
	MemberCellPhone string `json:"member_cell_phone"`
}

type Price struct {
	PriceValue float64
}

type OrderContract interface {
	SaveOrder(Order) error
	GetPrice(orgID int64, priceCategory string, status string) (Price, error)
	IsAnyOrder(orgID int64, memberCellPhone string, valueOrder string) bool
}
