# zohosdk

## Usage

A quick example to retrieve the IDs of all tickets marked "OPEN" (A Zoho built-in)


```
package main

import (
	"fmt"
	z "github.com/imm-llc/zohosdk"
)

func main() {

	zsdk := z.ZohoHeaders{
		"API_TOKEN",
		"ORGANZATION ID",
	}

	t := zsdk.GetAllTickets([]string{"OPEN"})

	fmt.Println(t)

}
```