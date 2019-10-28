package zohosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ZohoGetAllTicketsResponse struct {
	Data []struct {
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		Anon           bool   `json:"isAnonymous"`
		ModifiedTime   string `json:"modifiedTime"`
		Country        string `json:"country"`
		SecondaryEmail string `json:"secondaryEmail"`
		City           string `json:"city"`
		//Description    string    `json:"secondaryEmail"`
		OwnerID        string    `json:"ownerId"`
		Type           string    `json:"type"`
		Title          string    `json:"title"`
		PhotoURL       string    `json:"photoURL"`
		Twitter        string    `json:"twitter"`
		Deleted        string    `json:"isDeleted"`
		Trashed        string    `json:"isTrashed"`
		Street         string    `json:"street"`
		CreatedTime    time.Time `json:"createdTime"`
		ZohoCRMContact struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		}
		State        string `json:"state"`
		HappyMetrics struct {
			BadPercent  string `json:"badPercentage"`
			OkPercent   string `json:"okPercentage"`
			GoodPercent string `json:"goodPercentage"`
		}
		ID    string `json:"id"`
		Email string `json:"email"`
		Zip   string `json:"zip"`
		CF    struct {
			PermAddr    string `json:"cf_permanentaddress"`
			LastContact string `json:"cf_lastcontactedon"`
		}
		FB        string `json:"facebook"`
		Mobile    string `json:"mobile"`
		AccountID string `json:"accountId"`
		Phone     string `json:"phone"`
		WebURL    string `json:"webUrl"`
	} `json:"data"`
}

type ZohoGetSingleTicketResponse struct {
	TicketNumber string `json:"ticketNumber"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
}

// GetAllTickets Get all tickets of the given status, returning a list of Ticket IDs
func (h *ZohoHeaders) GetAllTickets(statuses string) []string {

	url := fmt.Sprintf("%s/tickets", ZohoBaseURL)

	//zh := &ZohoHeaders{}

	tokenHeaderString := fmt.Sprintf("Zoho-authtoken %s", h.Token)

	c := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating HTTP request to GetAllTickets")
		ZohoErrHandler(err)
	}

	q := req.URL.Query()

	// statuses should be e.g. "ON HOLD,OPEN,WAITING"
	q.Add("status", statuses)

	req.URL.RawQuery = q.Encode()

	req.Header.Set("orgId", h.OrgID)
	req.Header.Set("Authorization", tokenHeaderString)

	resp, err := c.Do(req)

	if err != nil {
		fmt.Println("Error making request to Zoho API to GetAllTickets")
		ZohoErrHandler(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading Zoho response")
		panic(err)
	}

	r := ZohoGetAllTicketsResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling GetAllTickets JSON")
		ZohoErrHandler(err)
	}

	fmt.Println("Unmarshalled JSON")

	//var ticketStringSlice []string

	ticketStringSlice := make([]string, len(r.Data))

	for i, t := range r.Data {
		ticketStringSlice[i] = t.ID
	}

	return ticketStringSlice

}

// GetSingleTicket returns a struct of the ticket subject, description, and ticket number
func (h *ZohoHeaders) GetSingleTicket(id string) ZohoGetSingleTicketResponse {

	url := fmt.Sprintf("%s/tickets/%s", ZohoBaseURL, id)

	//zh := &ZohoHeaders{}

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
		fmt.Println("Error making request to Zoho API")
		ZohoErrHandler(err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading Zoho response")
		panic(err)
	}

	r := ZohoGetSingleTicketResponse{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling JSON")
		panic(err)
	}

	return r

}
