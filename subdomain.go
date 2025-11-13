package loopia

import "errors"

// Subdomain ...
type Subdomain struct {
	Name string
}

// GetSubdomains - Method for fetching all subdomains
func (api *API) GetSubdomains(domain string) ([]Subdomain, error) {
	result := []string{}
	args := []interface{}{
		domain,
	}

	if err := api.Call("getSubdomains", args, &result); err != nil {
		return []Subdomain{}, err
	}

	subdomains := []Subdomain{}
	for _, value := range result {
		subdomains = append(subdomains, Subdomain{
			Name: value,
		})
	}

	return subdomains, nil
}

// GetSubdomain ...
func (api *API) GetSubdomain(domain string, subdomain string) (*Subdomain, error) {
	results, err := api.GetSubdomains(domain)
	if err != nil {
		return &Subdomain{}, err
	}
	for _, element := range results {
		if subdomain == element.Name {
			return &element, nil
		}
	}
	return &Subdomain{}, errors.New("ID Not found")
}

// AddSubdomain - method for creating subdomain
func (api *API) AddSubdomain(domain string, subdomain string) (*Status, error) {
	var result string
	args := []interface{}{
		domain,
		subdomain,
	}

	if err := api.Call("addSubdomain", args, &result); err != nil || result != "OK" {
		return &Status{
			Status: "failed",
			Cause:  result,
		}, err
	}

	return &Status{
		Status: "success",
	}, nil
}

// RemoveSubDomain - Removes a subdomain
func (api *API) RemoveSubDomain(domain string, subdomain string) (*Status, error) {
	var result string
	args := []interface{}{
		domain,
		subdomain,
	}

	if err := api.Call("removeSubdomain", args, &result); err != nil || result != "OK" {
		return &Status{
			Status: "failed",
			Cause:  result,
		}, err
	}

	return &Status{
		Status: "success",
	}, nil
}
