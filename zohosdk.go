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

func main() {
	fmt.Println("Hello from the BHI Zoho SDK")

}

// ZohoErrHandler handles all errors
func ZohoErrHandler(e error) {
	fmt.Println("ZOHO SDK ERROR")
	panic(e)
}
