package zohosdk

import (
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

type ZohoTicket struct {
	TicketNumber string `json:"ticketNumber"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
}

type ZohoTicketUpdate struct {
	Channel string `json:"channel"`
	Content string `json:"content"`
}

type ZohoHeaders struct {
	Token string
	OrgID string
}

// Get a
func (h *ZohoHeaders) GetAllTickets() {

}
