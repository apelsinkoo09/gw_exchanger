package service

import (
	"context"
	"fmt"
	"gw_exchanger/internal/storages/postgres"

	_ "github.com/apelsinkoo09/proto-exchange"
)

type ExchangeService struct {
	db *postgres.StorageConn
}

func NewExchangeService(db *postgres.StorageConn) *ExchangeService {
	return &ExchangeService{db: db}
}

// Реализация метода GetExchangeRates
func (s *ExchangeService) GetExchangeRates(ctx context.Context, req *exchange.Empty) (*exchange.ExchangeRatesResponse, error) {
	// Получаем курсы валют из базы данных
	rates, err := s.db.GetAllExchangeRates(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rates: %v", err)
	}

	// Формируем ответ
	response := &exchange.ExchangeRatesResponse{Rates: make(map[string]float32)}
	for pair, rate := range rates {
		response.Rates[pair] = rate
	}

	return response, nil
}

// Реализация метода GetExchangeRateForCurrency
func (s *ExchangeService) GetExchangeRateForCurrency(ctx context.Context, req *exchange.CurrencyRequest) (*exchange.ExchangeRateResponse, error) {
	// Получаем курс конкретной валютной пары из базы данных
	rate, err := s.db.GetExchangeRate(ctx, req.FromCurrency, req.ToCurrency)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate for %s -> %s: %v", req.FromCurrency, req.ToCurrency, err)
	}

	// Формируем ответ
	return &exchange.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         float32(rate),
	}, nil
}
