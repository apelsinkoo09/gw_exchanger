package postgres

import (
	"context"
	"database/sql"
	"log"
	"testing"
)

func TestGetAllExchangeRates(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=war666 dbname=exchange_base sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	storage := StorageConn{DB: db}

	ctx := context.Background()
	rates, err := storage.GetAllExchangeRates(ctx)
	if err != nil {
		t.Errorf("Error GetAllExchangeRates: %v", err)
	}

	expected := map[string]float32{
		"USD->EUR": 1.1,
		"USD->RUB": 0.1,
		"EUR->USD": 0.95,
		"EUR->RUB": 0.08,
		"RUB->USD": 103,
		"RUB->EUR": 108,
	}

	for pair, rate := range expected {
		if rates[pair] != rate {
			t.Errorf("Error")
		}
	}

}

func TestGetExchangeRate(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=war666 dbname=exchange_base sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	storage := StorageConn{DB: db}

	ctx := context.Background()
	fromCurrency := []string{"USD", "USD", "EUR", "EUR", "RUB", "RUB"}
	toCurrency := []string{"EUR", "RUB", "USD", "RUB", "USD", "EUR"}
	expected := []float64{1.1, 0.1, 0.95, 0.08, 103, 108}

	for i := 0; i < 6; i++ {
		rate, err := storage.GetExchangeRate(ctx, fromCurrency[i], toCurrency[i])
		if err != nil {
			t.Errorf("GetExchangeRate returned error: %v", err)
		}
		if expected[i] != rate {
			t.Errorf("Expected rate != tested rate")
		} else {
			log.Printf("Test №%v is passed", i)
		}
	}
}
