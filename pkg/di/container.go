package di

import (
	"context"
	"top1affiliate/internal/cron"
	"top1affiliate/internal/handlers"
	"top1affiliate/internal/service"
	"top1affiliate/internal/store"
	"top1affiliate/pkg/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	DataHandler *handlers.DataHandler
	UserHandler *handlers.UserHandler
}

func NewContainer(db *pgxpool.Pool) *Container {

	utils := utils.NewUtils()

	//data
	datastore := store.NewDataStore(db)
	dataservice := service.NewDataService(datastore)

	//user
	userstore := store.NewUserStore(db)
	userservice := service.NewUserService(userstore)

	//cron
	ctx := context.Background()
	c := cron.NewCronScheduler(datastore)
	c.StartCron(ctx)

	return &Container{
		DataHandler: handlers.NewDataHandler(dataservice, utils),
		UserHandler: handlers.NewUserHandler(userservice, utils),
	}

}
