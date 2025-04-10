package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"top1affiliate/internal/models"
	"top1affiliate/internal/store"
)

type DataService interface {
	Getstatistics(ctx context.Context, id string) ([]models.Statistics, error)
	GetweeklyStatsWithMonthly(ctx context.Context, id string) (*models.WeeklyStatsWithMonthly, error)
	GetTransactions(ctx context.Context, id, from, to string) ([]models.CommissionTxn, error)
	GetDashboardStats(ctx context.Context, id string) (*models.DashboardStats, error)
	GetBalance(ctx context.Context, id string) (float64, error)
	GetNetStats(ctx context.Context, id string) (*models.Stats, error)
	GetLeaderboard(ctx context.Context) ([]models.Leaderboard, error)
	GetSubAffiliates(ctx context.Context, id string) ([]models.User, error)
	GetSubAffiliatePath(ctx context.Context, id string) ([]models.AffiliatePath, error)
	GetSubAffiliateTree(ctx context.Context, id string) (*models.Tree, error)
	GetSubAffiliateList(ctx context.Context, id string) ([]models.TreeNode, error)
}

type dataService struct {
	store store.DataStore
}

func NewDataService(store store.DataStore) DataService {
	return &dataService{store: store}
}

func (s *dataService) Getstatistics(ctx context.Context, id string) ([]models.Statistics, error) {

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

func (s *dataService) GetLeaderboard(ctx context.Context) ([]models.Leaderboard, error) {

	leaderboard, err := s.store.GetLeaderboard(ctx)

	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}

func (s *dataService) GetNetStats(ctx context.Context, id string) (*models.Stats, error) {

	stats, err := s.store.GetNetStats(ctx, id)

	if err != nil {
		return nil, err
	}

	return stats, nil

}

func (s *dataService) GetBalance(ctx context.Context, id string) (float64, error) {

	bal, err := s.store.GetBalance(ctx, id)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return bal, nil
}

func (s *dataService) GetSubAffiliates(ctx context.Context, id string) ([]models.User, error) {

	aff, err := s.store.GetSubAffiliates(ctx, id)

	if err != nil {
		return nil, err
	}

	return aff, nil
}

func (s *dataService) GetSubAffiliatePath(ctx context.Context, id string) ([]models.AffiliatePath, error) {

	path, err := s.store.GetSubAffiliatePath(ctx, id)

	if err != nil {
		return nil, err
	}

	return path, nil
}

func (s *dataService) GetSubAffiliateTree(ctx context.Context, id string) (*models.Tree, error) {
	users, err := s.store.GetAllUsers(ctx, id)
	if err != nil {
		return nil, err
	}

	childrenMap := make(map[string][]models.TreeNode)
	var rootUser *models.TreeNode

	for i := range users {
		u := users[i]
		if u.AddedBy != "" {
			childrenMap[u.AddedBy] = append(childrenMap[u.AddedBy], u)
		}
		if u.ID == id {
			rootUser = &users[i]
		}
	}

	if rootUser == nil {
		return nil, fmt.Errorf("affiliate_id %s not found", id)
	}

	var buildTree func(user models.TreeNode) models.Tree
	buildTree = func(user models.TreeNode) models.Tree {
		children := make([]models.Tree, 0, len(childrenMap[user.ID]))
		for _, child := range childrenMap[user.ID] {
			children = append(children, buildTree(child))
		}
		return models.Tree{
			ID:         user.ID,
			CrmID:      "#" + user.AffiliateID,
			Name:       user.Name,
			Country:    user.Country,
			Commission: user.Commission,
			Recruits:   len(children),
			Children:   children,
		}
	}

	tree := buildTree(*rootUser)
	return &tree, nil
}

func (s *dataService) GetSubAffiliateList(ctx context.Context, id string) ([]models.TreeNode, error) {
	users, err := s.store.GetAllUsers(ctx, id)
	if err != nil {
		return nil, err
	}

	var usersList []models.TreeNode

	for _, user := range users {
		if user.ID != id {
			usersList = append(usersList, user)
		}
	}

	return usersList, nil
}
