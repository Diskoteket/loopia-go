package loopia

import "errors"

// Domain type
type Domain struct {
	Name            string `xmlrpc:"domain"`
	Paid            bool   `xmlrpc:"paid"`
	Registered      bool   `xmlrpc:"registered"`
	RenewalStatus   string `xmlrpc:"renewal_status"`
	ExpirationDate  string `xmlrpc:"expiration_date"`
	ReferenceNumber int32  `xmlrpc:"reference_number"`
}

// GetDomains - Method for fetching all domains
func (api *API) GetDomains() ([]Domain, error) {
	result := []Domain{}
	if err := api.Call("getDomains", []interface{}{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDomain
func (api *API) GetDomain(domain string) (*Domain, error) {
	results, err := api.GetDomains()
	if err != nil {
		return &Domain{}, err
	}
	for _, element := range results {
		if domain == element.Name {
			return &element, nil
		}
	}
	return &Domain{}, errors.New("ID Not found")
}
