package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"top1affiliate/internal/models"
	"top1affiliate/internal/store"
)

type DataService interface {
	Getstatistics(ctx context.Context, id string) ([]models.Leads, error)
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
