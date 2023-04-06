package service

import (
	"github.com/pari-27/HdfcAssignment/store"
	"github.com/pari-27/HdfcAssignment/types"
)

// Service is the interface to call API
type Service interface {
	AddWebsites(req *types.WebsiteRequest) error
	ListWebsitesStatus() ([]*types.WebsiteStatus, error)
	GetWebsiteStatus(queryParams map[string]string) (*types.WebsiteStatus, error)
}

type service struct {
	Store store.Store
}
