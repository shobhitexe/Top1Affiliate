package cron

import (
	"context"
	"fmt"
	"log"
	"top1affiliate/internal/models"
)

func (c *Cron) InactiveAffiliates(ctx context.Context) error {

	users, err := c.store.InactiveAffiliates(ctx)

	if err != nil {
		return err
	}

	for _, u := range users {

		message := fmt.Sprintf(
			`Inactive affiliate
	
	Crm ID:            %s
	Name:              %s
	Country:           %s
	Last succesful deposit of client
	amount:            %.2f
	date:               %s`,
			u.AffiliateID,
			u.Name,
			u.Country,
			u.LastDepositAmount,
			u.LastDepositDate,
		)

		if err := c.utils.SendNotificationToSlack(ctx, models.InactiveAffiliates, message); err != nil {
			log.Println(err)
			continue
		}

	}

	return nil
}
