# LAOFIDA Go Client
   This package provides a client for interacting with the LAOFIDA API, including automatic authentication and token refresh, as well as methods for retrieving SmartTax records.
   
# Features:
   Automatic authentication with LAOFIDA API,
   Auto token refresh with exponential backoff strategy,
   Thread-safe token access,
   SmartTax record retrieval via the GetDataSmartTaxs method.

# Installation:
   "go get github.com/yourusername/laofida"

# Usage:
   "package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yourusername/laofida"
)

func main() {
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
}"
