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
> Construct a new Telbiz client, then use the various services on the client to access different parts of the laofida API. For example:
# Usage:
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

   
