package postgres

import (
	"database/sql"
)

type StorageConn struct {
	db *sql.DB
}

func (db *StorageConn) GetAllExchangeRates() (map[string]float32, error) {
	rows, err := db.db.Query("SELECT from_currency, rate FROM exchange_rates")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rates := make(map[string]float32)
	for rows.Next() {
		var currency string
		var rate float32
		if err := rows.Scan(&currency, &rate); err != nil {
			return nil, err
		}
		rates[currency] = rate
	}
	return rates, nil
}

func (db *StorageConn) GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	var rate float64
	err := db.db.QueryRow("SELECT rate FROM exchange_rates WHERE from_currency = $1 AND to_currency = $2", fromCurrency, toCurrency).Scan(&rate)
	return rate, err
}
