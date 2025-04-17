package cron

import (
	"context"
	"log"
	"time"
	"top1affiliate/internal/store"
	"top1affiliate/pkg/utils"
)

type Cron struct {
	store store.DataStore
	utils utils.Utils
}

func NewCronScheduler(store store.DataStore, utils utils.Utils) *Cron {
	return &Cron{store: store, utils: utils}
}

func (c *Cron) StartCron(ctx context.Context) {

	go func() {
		ticker := time.NewTicker(30 * time.Minute)
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
						log.Println("Error fetching new deposit txns:", err)
					}

					if err := c.FetchAndSaveTransactionsWithdrawals(taskCtx, cookie, lastHour); err != nil {
						log.Println("Error fetching new withdrawal txns:", err)
					}

				}()

			case <-ctx.Done():
				log.Println("Stopping 30 min cron tasks...")
				return
			}
		}
	}()

	go func() {
		var lastRunDate string

		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				now := time.Now().UTC()
				today := now.Format("2006-01-02")

				if now.Hour() == 0 && now.Minute() == 0 && lastRunDate != today {
					log.Println("Running InactiveAffiliates task at 12AM UTC")

					taskCtx, cancel := context.WithTimeout(ctx, 20*time.Minute)
					err := c.InactiveAffiliates(taskCtx)
					cancel()

					if err != nil {
						log.Println("Error running InactiveAffiliates:", err)
					} else {
						log.Println("Successfully ran InactiveAffiliates task")
						lastRunDate = today
					}
				}

			case <-ctx.Done():
				log.Println("Stopping 12AM daily inactive check...")
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

	// if err := c.FetchAndSaveTransactionsWithdrawals(ctx, cookie, lastHour); err != nil {
	// 	log.Println("Error fetching new txns:", err)
	// }

}
