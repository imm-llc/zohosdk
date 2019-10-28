package zohosdk

import (
	"fmt"
)

const (
	ZohoBaseURL = "https://desk.zoho.com/api/v1"
)

// ZohoHeaders Headers for API requests
type ZohoHeaders struct {
	Token string
	OrgID string
}

// You must initialize this package by providing your API token and Organization ID
func (h *ZohoHeaders) main() {

}

func ZohoErrHandler(e error) {
	fmt.Println("ZOHO SDK ERROR")
	panic(e)
}
