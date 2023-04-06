package store

import "github.com/pari-27/HdfcAssignment/types"

type Store interface {
	GetWebsites() (types.WebsiteStatus, error)
}
