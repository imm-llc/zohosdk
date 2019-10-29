package zohosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//  ZohoThread returns the summary of a ticket thread: https://desk.zoho.com/DeskAPIDocument#Threads#Threads_Getathread
type ZohoThread struct {
	Summary string `json:"summary"`
}

// GetTicketThreadSummary returns the summary of a ticket thread
func (h *ZohoHeaders) GetTicketThreadSummary(id string) (string, error) {
	url := fmt.Sprintf("%s/%s/latestThread", ZohoBaseURL, id)

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

	responseBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Ticket thread response: ", responseBody)
	fmt.Println("HTTP Status code: ", resp.StatusCode)

	if err != nil {
		fmt.Println("Error reading Zoho response")
		return "", err
	}

	r := ZohoThread{}

	err = json.Unmarshal(responseBody, &r)

	if err != nil {
		fmt.Println("Error unmarshalling JSON")
		return "", err
	}

	return r.Summary, nil

}
