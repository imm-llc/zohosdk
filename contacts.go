package zohosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ZohoContactsAllResponse provides a struct for a list of contacts
type ZohoContactsAllResponse struct {
	Data []struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		ID        string `json:"id"`
		Email     string `json:"email"`
	} `json:"data"`
}

type ZohoContactsCountResponse struct {
	Count string `json:"count"`
}

type ZohoContactsViewsResponse struct {
	Data []struct {
		Custom   bool   `json:"isCustomView"`
		ViewName string `json:"name"`
		ID       string `json:"id"`
	} `json:"data"`
}

type ZohoSingleContactResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ID        string `json:"id"`
	Email     string `json:"email"`
}

// GetAllContacts returns a ZohoContactsAll struct
func (h *ZohoHeaders) GetAllContacts() ZohoContactsAllResponse {

	url := fmt.Sprintf("%s/contacts?limit=99", ZohoBaseURL)

	tokenHeaderString := fmt.Sprintf("Zoho-authtoken %s", h.Token)

	c := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request to GetAllTickets")
		ZohoErrHandler(err)
	}

	// Limit defaults to 10
	//q := req.URL.Query()
	//q.Add("limit", "15")

	req.Header.Set("orgId", h.OrgID)
	req.Header.Set("Authorization", tokenHeaderString)

	resp, err := c.Do(req)

	if err != nil {
		fmt.Println("Error making request to Zoho API to GetAllContacts")
		ZohoErrHandler(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(responseBody))

	if err != nil {
		fmt.Println("Error reading GetAllContacts response")
		fmt.Println("Dumping response body for debugging")
		fmt.Println(string(responseBody))
		ZohoErrHandler(err)
	}

	r := ZohoContactsAllResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling GetAllContacts JSON")
		fmt.Println("Dumping response body for debugging")
		fmt.Println(string(responseBody))
		ZohoErrHandler(err)
	}

	return r
}

func (h *ZohoHeaders) ListContactViews() ZohoContactsViewsResponse {
	// For some reason, adding the query with URL.Query() isn't working
	url := fmt.Sprintf("%s/views?module=contacts", ZohoBaseURL)

	tokenHeaderString := fmt.Sprintf("Zoho-authtoken %s", h.Token)

	c := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request to GetAllTickets")
		ZohoErrHandler(err)
	}

	req.Header.Set("orgId", h.OrgID)
	req.Header.Set("Authorization", tokenHeaderString)

	//q := req.URL.Query()
	//q.Add("module", "Contacts")

	resp, err := c.Do(req)

	if err != nil {
		fmt.Println("Error making request to Zoho API to GetAllContacts")
		ZohoErrHandler(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading GetAllContacts response")
		ZohoErrHandler(err)
	}

	r := ZohoContactsViewsResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling ListContactViews JSON")
		fmt.Println("Dumping response body for debugging")
		fmt.Println(string(responseBody))
		ZohoErrHandler(err)
	}

	return r

}

func (h *ZohoHeaders) GetContactCount(v string) ZohoContactsCountResponse {

	url := fmt.Sprintf("%s/contacts/count?viewId=%s", ZohoBaseURL, v)

	tokenHeaderString := fmt.Sprintf("Zoho-authtoken %s", h.Token)

	c := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request to GetAllTickets")
		ZohoErrHandler(err)
	}

	req.Header.Set("orgId", h.OrgID)
	req.Header.Set("Authorization", tokenHeaderString)

	resp, err := c.Do(req)

	if err != nil {
		fmt.Println("Error making request to Zoho API to GetAllContacts")
		ZohoErrHandler(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading GetAllContacts response")
		ZohoErrHandler(err)
	}

	r := ZohoContactsCountResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling GetAllContacts JSON")
		fmt.Println("Dumping response body for debugging")
		fmt.Println(string(responseBody))
		ZohoErrHandler(err)
	}

	return r
}

func (h *ZohoHeaders) GetSingleContact(id string) ZohoSingleContactResponse {

	url := fmt.Sprintf("%s/contacts/%s", ZohoBaseURL, id)

	tokenHeaderString := fmt.Sprintf("Zoho-authtoken %s", h.Token)

	c := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request to GetAllTickets")
		ZohoErrHandler(err)
	}

	req.Header.Set("orgId", h.OrgID)
	req.Header.Set("Authorization", tokenHeaderString)

	resp, err := c.Do(req)

	if err != nil {
		fmt.Println("Error making request to Zoho API to GetAllContacts")
		ZohoErrHandler(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading GetAllContacts response")
		ZohoErrHandler(err)
	}

	r := ZohoSingleContactResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling GetSingleContact JSON")
		fmt.Println("Dumping response body for debugging")
		fmt.Println(string(responseBody))
		ZohoErrHandler(err)
	}

	return r

}
