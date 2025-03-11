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

	// go func() {
	// 	cookie, err := c.LoginToAPI()

	// 	if err != nil {
	// 		return
	// 	}

	// 	c.FetchAndSaveLeads(ctx, cookie)
	// }()

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:

				// taskCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				// defer cancel()

				// cookie, err := c.LoginToAPI()

				// if err != nil {
				// 	return
				// }

				// c.FetchAndSaveLeads(ctx, cookie)

			case <-ctx.Done():
				log.Println("Stopping 120-second cron tasks...")
				return

			}
		}
	}()

}
