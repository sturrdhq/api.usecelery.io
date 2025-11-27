package waitlist

import (
	"strings"

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

	err = as.db.Create(&newSubscription).Error
	if err != nil && strings.Contains(strings.ToLower(err.Error()), "duplicate") {
		return nil
	}
	return err
}

func (as *Service) getDefaultWaitList() (*models.WaitList, error) {
	const DEFAULT_WAITLIST_NAME = "USECELERY_WAITLIST"

	var defaultWaitlist = models.WaitList{Name: DEFAULT_WAITLIST_NAME}

	tx := as.db.Model(models.WaitList{}).Where("name = ?", DEFAULT_WAITLIST_NAME).FirstOrCreate(&defaultWaitlist)

	if tx.Error == nil {
		return &defaultWaitlist, nil
	}

	return nil, tx.Error
}
