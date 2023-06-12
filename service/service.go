package service

import (
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"github.com/pari-27/HdfcAssignment/store"
	"github.com/pari-27/HdfcAssignment/types"
	"github.com/pari-27/HdfcAssignment/utils"
	"github.com/pari-27/HdfcAssignment/utils/constants"
	"net/http"
	"time"
)

// Service is the interface to call API
type Service interface {
	AddWebsites(req types.WebsiteRequest) (response []types.WebsiteStatus, err error)
	ListWebsiteStatus() (res *types.WebsiteStatus, err error)
	GetWebsiteStatus(req *types.WebsiteRequest) (res *types.WebsiteStatus, err error)
}

type service struct {
	Logger log.Logger
	Store  store.Store
}

func NewService(lg log.Logger, store store.Store) Service {
	return &service{
		Logger: lg,
		Store:  store,
	}
}

func (srv *service) AddWebsites(req types.WebsiteRequest) (response []types.WebsiteStatus, err error) {
	// check if any website provided in the request
	if len(req.Websites) == 0 {
		return nil, errors.New("no websites provided in the request")
	}
	// add the websites to the websites map if already not present
	response, err = srv.Store.AddWebsites(req)
	if err != nil {
		return nil, err
	}

	return
}

func (srv *service) ListWebsiteStatus() (res *types.WebsiteStatus, err error) {
	return
}
func (srv *service) GetWebsiteStatus(req *types.WebsiteRequest) (res *types.WebsiteStatus, err error) {
	return
}

func (srv *service) WebsiteStatusChecker() {
	wp := NewWorkerPool(constants.StatusCheckJobCount)
	wp.Run() // start accepting the jobs
	// initialise a ticker for running the status checking task
	ticker := time.NewTicker(time.Duration(constants.StatusCheckTime))

	for {

		select {
		case <-ticker.C:
			wp.Add(func() {
				wbStatusMap, err := srv.Store.GetWebsites()
				if err != nil {

				}
				for _, wbStatus := range wbStatusMap {
					err := utils.SendHTTPRequest(wbStatus.Name, http.MethodGet)
					if err != nil {
						wbStatus.Status = constants.StatusDOWN
					} else {
						wbStatus.Status = constants.StatusUP
					}
					err = srv.Store.UpdateWebsiteStatus(wbStatus)
					if err != nil {
						fmt.Errorf("Failed to update the status for website %s ", wbStatus.Name)
					}
				}

			})
		}

	}
}
