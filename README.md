# PrimeTrust Go Library

## Running the Example

```
PRIMETRUST_LOGIN=user@example.com PRIMETRUST_PASSWORD=password PRIMETRUST_ACCOUNT_ID=00000000-0000-0000-0000-000000000000 go run examples/account-types/*.go
```

### Environment Variables

Please specify appropriate variables when running examples according to source code.

* PRIMETRUST_LOGIN - email of existing PrimeTrust user
* PRIMETRUST_PASSWORD - password
* PRIMETRUST_ACCOUNT_ID - uuid of custodial account-id
* PRIMETRUST_CONTACT_ID - uuid of KYC contact
* PRIMETRUST_WEBHOOK_ID - uuid of webhook
* PRIMETRUST_WEBHOOK_URL
* PRIMETRUST_WEBHOOK_EMAIL
* PRIMETRUST_WEBHOOK_SECRET

## Usage In Own Application

```go
package main

import (
	"github.com/cloudmode/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	if accountTypes, err := primetrust.GetAccountTypes(); err != nil {
		log.Println("Error getting account types:", err.Error())
	} else {
		log.Printf("Account Types: %d", len(accountTypes.Data))
		if accountType, err := primetrust.GetAccountType(accountTypes.Data[0].ID); err != nil {
			log.Printf("Error getting \"%s\" account type:", accountTypes.Data[0].ID, err.Error())
		} else {
			log.Printf("Account Type \"%s\": %+v", accountTypes.Data[0].ID, accountType)
		}
	}
}
```