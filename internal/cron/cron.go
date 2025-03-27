package cron

import (
	"context"
	"log"
	"time"
	"top1affiliate/internal/store"
)

type Cron struct {
	store store.DataStore
}

func NewCronScheduler(store store.DataStore) *Cron {
	return &Cron{store: store}
}

func (c *Cron) StartCron(ctx context.Context) {

	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:

				cookie, err := c.LoginToAPI()
				if err != nil {
					log.Println("Error logging in to API:", err)
					return
				}

				taskCtx, cancel := context.WithTimeout(ctx, 20*time.Minute)
				defer cancel()

				go func() {

					currentTime := time.Now().UTC()
					lastHour := currentTime.Add(-2 * time.Hour).Format("2006-01-02 15:04")

					log.Println("Fetching leads from:", lastHour, "to:", currentTime.Format("2006-01-02 15:04"))

					if err := c.FetchAndSaveLeads(taskCtx, cookie, lastHour, currentTime.Format("2006-01-02 15:04")); err != nil {
						log.Println("Error fetching new leads:", err)
					}

					if err := c.FetchAndSaveTransactionsDeposit(taskCtx, cookie, lastHour); err != nil {
						log.Println("Error fetching new txns:", err)
					}

					if err := c.FetchAndSaveTransactionsWithdrawals(taskCtx, cookie, lastHour); err != nil {
						log.Println("Error fetching new txns:", err)
					}

				}()

			case <-ctx.Done():
				log.Println("Stopping hourly cron tasks...")
				return
			}
		}
	}()

	// cookie, err := c.LoginToAPI()
	// if err != nil {
	// 	log.Println("Error logging in to API:", err)
	// 	return
	// }

	// currentTime := time.Now().UTC()
	// lastHour := currentTime.Add(-2610 * time.Hour).Format("2006-01-02 15:04")

	// if err := c.FetchAndSaveTransactionsDeposit(ctx, cookie, lastHour); err != nil {
	// 	log.Println("Error fetching new txns:", err)
	// }

}
