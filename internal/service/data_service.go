package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"top1affiliate/internal/models"
	"top1affiliate/internal/store"
)

type DataService interface {
	Getstatistics(ctx context.Context, id string) ([]models.Leads, error)
	GetweeklyStatsWithMonthly(ctx context.Context, id string) (*models.WeeklyStatsWithMonthly, error)
	GetTransactions(ctx context.Context, id, from, to string) ([]models.CommissionTxn, error)
	GetDashboardStats(ctx context.Context, id string) (*models.DashboardStats, error)
}

type dataService struct {
	store store.DataStore
}

func NewDataService(store store.DataStore) DataService {
	return &dataService{store: store}
}

func (s *dataService) Getstatistics(ctx context.Context, id string) ([]models.Leads, error) {

	leads, err := s.store.Getstatistics(ctx, id)

	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (s *dataService) FetchAndSaveLeads(cookie string, username string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://publicapi.fxlvls.com/management/leads?limit=100&minRegistrationDate=2025-05-01", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Cookie", cookie)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var responseData []models.Leads

	if err := json.Unmarshal(body, &responseData); err != nil {
		log.Println(err)
		return err
	}

	ctx := context.Background()

	for _, data := range responseData {

		if err := s.store.SaveLeadsData(ctx, data); err != nil {
			log.Println(err)
			return err
		}

	}

	return nil

}

func (s *dataService) GetweeklyStats(ctx context.Context, id string) (*models.Stats, error) {

	d, err := s.store.GetweeklyStats(ctx, id)

	if err != nil {
		return nil, err
	}

	return d, err

}

func (s *dataService) GetweeklyStatsWithMonthly(ctx context.Context, id string) (*models.WeeklyStatsWithMonthly, error) {

	var weekly, monthly *models.Stats
	var weeklyErr, monthlyErr error

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		weekly, weeklyErr = s.store.GetweeklyStats(ctx, id)
	}()

	go func() {
		defer wg.Done()
		monthly, monthlyErr = s.store.GetMonthlyStats(ctx, id)
	}()

	wg.Wait()

	if weeklyErr != nil {
		return nil, weeklyErr
	}

	if monthlyErr != nil {
		return nil, monthlyErr
	}

	stats := models.WeeklyStatsWithMonthly{
		Registrations: weekly.Registrations,
		Deposits:      weekly.Deposits,
		Withdrawals:   weekly.Withdrawals,
		Commissions:   weekly.Commissions,

		RegistrationsMonthly: monthly.Registrations,
		DepositsMonthly:      monthly.Deposits,
		WithdrawalsMonthly:   monthly.Withdrawals,
		CommissionsMonthly:   monthly.Commissions,
	}

	return &stats, nil

}

func (s *dataService) GetTransactions(ctx context.Context, id, from, to string) ([]models.CommissionTxn, error) {

	txn, err := s.store.GetTransactions(ctx, id, from, to)

	if err != nil {
		return nil, err
	}

	return txn, err

}

func (s *dataService) GetDashboardStats(ctx context.Context, id string) (*models.DashboardStats, error) {

	var weekly *models.Stats
	var weeklyErr error

	var txns []models.CommissionTxn
	var txnErr error

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		weekly, weeklyErr = s.store.GetweeklyStats(ctx, id)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		txns, txnErr = s.store.GetLatestFiveTransactions(ctx, id)
	}()

	wg.Wait()

	if weeklyErr != nil {
		log.Println(weeklyErr)
		return nil, weeklyErr
	}

	if txnErr != nil {
		return nil, txnErr
	}

	return &models.DashboardStats{Weekly: *weekly, Commissions: txns}, nil
}
