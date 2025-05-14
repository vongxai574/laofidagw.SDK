# LAOFIDA Go Client
   This package provides a client for interacting with the LAOFIDA API, including automatic authentication and token refresh, as well as methods for retrieving SmartTax records.
   
# Features:
   - Automatic authentication with LAOFIDA API.
   - Auto token refresh with exponential backoff strategy.
   - Thread-safe token access.
   - SmartTax record retrieval via the GetDataSmartTaxs method.

# Installation:
   ```bash 
   go get github.com/vongxai574/laofidagw.SDK
   ```
> Use go get to retrieve the latest version of the client.
# Usage:
   ```bash
   import "github.com/vongxai574/laofidagw.SDK"
   ```
Construct a new laofida client, then use the various services on the client to access different parts of the laofidagw API. For example:
   ```bash
	ctx := context.Background()

	cfg := &laofida.Config{
		BaseURL:  "https://api.laofida.example.com",
		Username: "your-username",
		Password: "your-password",
	}

	client, err := laofida.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create LAOFIDA client: %v", err)
	}
 ```
# Call get GetDataSmartTaxs example:
```bash
	req := &laofida.ReqFilter{
		DateStart: "2024-01-01",
		DateEnd:   "2024-01-31",
		TIN:       "123456789",
		Type:      "EX",
	}

	data, err := client.GetDataSmartTaxs(ctx, req)
	if err != nil {
		log.Fatalf("failed to get SmartTax data: %v", err)
	}

	for _, record := range data {
		fmt.Printf("Record ID: %s, Total Tax: %s\n", record.InstanceID, record.TotalTax.String())
	}
```
# Contributing:
> Contributions are welcome. Please open up an issue or create a pull request if you would like to help out.
   
