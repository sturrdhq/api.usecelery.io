package waitlist

import (
	"github.com/sturrdhq/celery-server/internal/database"
	"github.com/sturrdhq/celery-server/internal/database/models"
)

type Service struct {
	db *database.DBClient
}

func NewWaitListService(db *database.DBClient) *Service {
	return &Service{
		db,
	}
}

func (as *Service) Subscribe(email string) error {
	defaultWaitlist, err := as.getDefaultWaitList()
	if err != nil {
		return err
	}

	var newSubscription = models.Subscription{Email: email, WaitListID: defaultWaitlist.ID}

	return as.db.Create(&newSubscription).Error
}

func (as *Service) getDefaultWaitList() (*models.WaitList, error) {
	const DEFAULT_WAITLIST_NAME = "USECELERY_WAITLIST"

	var defaultWaitlist models.WaitList

	tx := as.db.Model(models.WaitList{}).Where("name = ?", DEFAULT_WAITLIST_NAME).FirstOrCreate(&defaultWaitlist)

	if tx.Error == nil {
		return &defaultWaitlist, nil
	}

	return nil, tx.Error
}
