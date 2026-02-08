package repositories

import (
	"database/sql"
	"fmt"

	"github.com/heru-oktafian/tugas-02/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	var (
		res *models.Transaction
	)

	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0

	details := make([]models.TransactionDetails, 0)

	for _, item := range items {
		var productName string
		var productID, productPrice, productStock int
		err := tx.QueryRow("SELECT id, name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productID, &productName, &productPrice, &productStock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %d not found", item.ProductID)
		}

		if err != nil {
			return nil, err
		}

		subtotal := item.Quantity * productPrice
		totalAmount += subtotal

		if productStock < item.Quantity {
			return nil, fmt.Errorf("product with ID %d has insufficient stock", item.ProductID)
		}

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetails{
			ProductID:   productID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)", transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	res = &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}

	return res, nil
}

func (repo *TransactionRepository) GetReport(startDate, endDate string) (*models.ReportResponse, error) {
	var report models.ReportResponse

	// Base query condition
	dateCondition := ""
	args := []interface{}{}
	if startDate != "" && endDate != "" {
		dateCondition = " WHERE created_at BETWEEN $1 AND $2"
		args = append(args, startDate, endDate)
	}

	// 1. Total Revenue
	queryRevenue := "SELECT COALESCE(SUM(total_amount), 0) FROM transactions" + dateCondition
	err := repo.db.QueryRow(queryRevenue, args...).Scan(&report.TotalRevenue)
	if err != nil {
		return nil, err
	}

	// 2. Total Transactions
	queryCount := "SELECT COUNT(*) FROM transactions" + dateCondition
	err = repo.db.QueryRow(queryCount, args...).Scan(&report.TotalTransactions)
	if err != nil {
		return nil, err
	}

	// 3. Best Selling Product
	// Need to join with transactions table to filter by date if needed
	queryBestSelling := `
		SELECT p.name, SUM(td.quantity) as total_qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
	` + dateCondition + `
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`

	err = repo.db.QueryRow(queryBestSelling, args...).Scan(&report.BestSelling.Name, &report.BestSelling.QtySold)
	if err != nil {
		if err == sql.ErrNoRows {
			report.BestSelling = models.BestSellingProduct{Name: "", QtySold: 0}
		} else {
			return nil, err
		}
	}

	return &report, nil
}
