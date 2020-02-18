package ordercontroller

import "capsulego/models"

type OrderInteractor struct {
	repo models.OrderRepository
}

func (i OrderInteractor) SaveOrder(order models.Order) error {
	// priceModel := models.Price{
	// 	OrgID:         order.OrgID,
	// 	PriceCategory: "CAPSULE BUS",
	// 	Status:        "ACTIVE",
	// }
	// getPrice, err := getPrice(priceModel)
	// if err != nil {
	// 	return err
	// }
	// existingOrder, err := existingOrder(order)
	// if err != nil || existingOrder > 0 {
	// 	return err
	// }

	err := i.repo.SaveOrder(order)
	if err != nil {
		return err
	}
	return nil
}

// func getPrice(price models.Price) (float64, error) {
//
// 	return
// }
//
// func existingOrder(order models.Order) (int8, error) {
// 	r, err := repo.isAnyOrder(order)
// 	if err != nil {
// 		return false, err
// 	}
// 	return r, nil
// }
//
// func save(order models.Order)
