package loopia

import (
	"errors"
)

// Record descired Loopia record_obj
type Record struct {
	ID       int64  `xmlrpc:"record_id"`
	TTL      int    `xmlrpc:"ttl"`
	Type     string `xmlrpc:"type"`
	Value    string `xmlrpc:"rdata"`
	Priority int    `xmlrpc:"priority"`
}

// Status - operation status wrapper
type Status struct {
	Status string
	Cause  string
}

// AddZoneRecord - Create zone record
func (api *API) AddZoneRecord(domain string, subdomain string, record *Record) error {
	var result string
	args := []interface{}{
		domain,
		subdomain,
		record,
	}

	if err := api.Call("addZoneRecord", args, &result); err != nil || result != "OK" {
		return err
	}

	// Try figuring out ID of our new zoneRecord.
	// Loopia does not return any kind of identification on the created object.
	results, err := api.GetZoneRecords(domain, subdomain)
	if err != nil {
		return err
	}

	for _, element := range results {
		// Exclude ID before equality check
		id := element.ID
		element.ID = 0

		// Compare by value
		if element == *record {
			// Found our new record, assigning ID
			record.ID = id
			return nil
		}
	}
	return errors.New("Record saved but unable to query for ID")
}

// GetZoneRecords - fetch subdomains records
func (api *API) GetZoneRecords(domain string, subdomain string) ([]Record, error) {
	result := []Record{}
	args := []interface{}{
		domain,
		subdomain,
	}

	if err := api.Call("getZoneRecords", args, &result); err != nil {
		return []Record{}, err
	}

	return result, nil
}

// GetZoneRecord - fetch specific zone record
func (api *API) GetZoneRecord(domain string, subdomain string, id int64) (*Record, error) {
	results, err := api.GetZoneRecords(domain, subdomain)
	if err != nil {
		return &Record{}, err
	}

	for _, element := range results {
		if id == element.ID {
			return &element, nil
		}
	}
	return &Record{}, errors.New("ID Not found")
}

// RemoveZoneRecord - remove zone record
func (api *API) RemoveZoneRecord(domain string, subdomain string, id int64) (*Status, error) {
	var result string
	args := []interface{}{
		domain,
		subdomain,
		id,
	}

	if err := api.Call("removeZoneRecord", args, &result); err != nil {
		return &Status{
			Status: "failed",
			Cause:  result,
		}, err
	}

	return &Status{
		Status: "success",
	}, nil
}

// UpdateZoneRecord -
func (api *API) UpdateZoneRecord(domain string, subdomain string, record Record) (*Status, error) {
	var result string
	args := []interface{}{
		domain,
		subdomain,
		record,
	}

	if err := api.Call("updateZoneRecord", args, &result); err != nil || result != "OK" {
		return &Status{
			Status: "failed",
			Cause:  result,
		}, err
	}

	return &Status{
		Status: "success",
	}, nil
}
