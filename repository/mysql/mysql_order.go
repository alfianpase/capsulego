package repository

import (
	"capsulego/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// getPrice(orgID int64, priceCategory string, status string) (float64, error)
// existingOrder(Order) (int8, error)
type orderRepository struct {
	conn *sql.DB
}

// NewUserRepository returns initialized UserRepositoryImpl
func NewOrderRepository(conn *sql.DB) models.OrderRepository {
	return &orderRepository{conn: conn}
}

func (r *orderRepository) SaveOrder(order models.Order) error {
	stmt, err := r.conn.Prepare("INSERT INTO mt_transaction (org_id, member_cell_phone) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.OrgID, order.MemberCellPhone)
	return err
}
