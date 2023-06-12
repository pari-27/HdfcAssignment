package store

import (
	"errors"
	"github.com/pari-27/HdfcAssignment/types"
	"github.com/pari-27/HdfcAssignment/utils/constants"
)

type Store interface {
	GetWebsites() ([]types.WebsiteStatus, error)
	AddWebsites(types.WebsiteRequest) ([]types.WebsiteStatus, error)
	GetWebsite(types.WebsiteRequest) (types.WebsiteStatus, error)
	UpdateWebsiteStatus(status types.WebsiteStatus) error
}

type DLStore struct {
	WebsiteStatusMap map[string]string
}

func (dl *DLStore) AddWebsites(websiteReq types.WebsiteRequest) (websiteResponse []types.WebsiteStatus, err error) {
	m := types.WebsiteStatus{}
	for _, website := range websiteReq.Websites {
		if _, ok := dl.WebsiteStatusMap[website]; !ok {
			dl.WebsiteStatusMap[website] = constants.StatusNotChecked
		}
		m.Name = website
		m.Status = dl.WebsiteStatusMap[website]
		websiteResponse = append(websiteResponse, m)
	}
	return
}
func (dl *DLStore) GetWebsites() (response []types.WebsiteStatus, err error) {
	m := types.WebsiteStatus{}
	for key, val := range dl.WebsiteStatusMap {
		m.Name = key
		m.Status = val
		response = append(response, m)
	}
	return
}
func (dl *DLStore) UpdateWebsiteStatus(wbs types.WebsiteStatus) error {
	if _, ok := dl.WebsiteStatusMap[wbs.Name]; !ok {
		return errors.New("failed to find website for status update")
	}
	dl.WebsiteStatusMap[wbs.Name] = wbs.Status
	return nil
}
