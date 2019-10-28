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

// GetAllContacts returns a ZohoContactsAll struct
func (h *ZohoHeaders) GetAllContacts() ZohoContactsAllResponse {

	url := fmt.Sprintf("%s/contacts", ZohoBaseURL)

	tokenHeaderString := fmt.Sprintf("Zoho-authtoken %s", h.Token)

	c := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request to GetAllTickets")
		ZohoErrHandler(err)
	}

	// Limit defaults to 10
	q := req.URL.Query()
	q.Add("limit", "100")

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

	r := ZohoContactsAllResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling GetAllContacts JSON")
		ZohoErrHandler(err)
	}

	return r

}
