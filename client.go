package laofida

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// notifyAccessTokenExpired do infinite loop with period of time
// to perform auto renew token from LAOFIDA api backend with
// exponential backOff strategy.
func (l *laofida) notifyAccessTokenExpired(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return

		case <-ticker.C:
			// log.Println("notifyAccessTokenExpired: ticker.C => token expired")

		case <-l.toggleTokenRefresher:
			log.Println("notifyAccessTokenExpired: toggleTokenRefresher => token refreshed")
		}

		// reconnect with exponential backOff strategy
		backOff := time.Second

	Retry:
		for {
			token, err := l.connect(ctx)
			switch err {
			case nil:
				l.setAccessToken(token)

				break Retry

			default:
				log.Printf("notifyAccessTokenExpired: %v", err)

				select {
				case <-ctx.Done():
					return

				case <-time.After(backOff):
					backOff *= 2
				}
			}
		}
	}
}

// setAccessToken set access token to client.
func (l *laofida) setAccessToken(accessToken string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.accessToken = accessToken
}

// getAccessToken get access token from client.
func (l *laofida) getAccessToken() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.accessToken
}

// connect makes http call to perform authentication with LAOFIDA api backend.
func (l *laofida) connect(ctx context.Context) (string, error) {
	form := url.Values{}
	form.Set("username", l.username)
	form.Set("password", l.password)

	body := strings.NewReader(form.Encode())

	_baseURL, _ := url.Parse(l.baseUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/v1/login", _baseURL), body)
	if err != nil {
		return "", fmt.Errorf("connectLAOFIDA: http.NewRequestWithContext: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := l.hc.Do(req)
	if err != nil {
		return "", fmt.Errorf("connectLAOFIDA: http.Client.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		l.toggleTokenRefresher <- struct{}{}
		return "", errors.New("connectLAOFIDA: unauthorized (401)")
	}

	if resp.StatusCode != http.StatusOK {
		rbody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("connectLAOFIDA: status %d, body: %s", resp.StatusCode, rbody)
	}

	var reply struct {
		Success bool `json:"success"`
		Data    struct {
			Token string `json:"token"`
			Name  string `json:"name"`
		} `json:"data"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return "", fmt.Errorf("connectLAOFIDA: json.Decode: %w", err)
	}

	if !reply.Success {
		return "", fmt.Errorf("connectLAOFIDA: API returned error: %s", reply.Message)
	}

	return reply.Data.Token, nil
}

func (l *laofida) getSmartTaxs(ctx context.Context, r *ReqFilter) (SmartTaxRecords, error) {
	payload := map[string]string{
		"from_date": r.DateStart,
		"to_date":   r.DateEnd,
		"tin":       r.TIN,
		"type":      r.Type, // e.g., "EX"
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("getSmartTaxs: json.Marshal: %w", err)
	}

	_baseURL, _ := url.Parse(l.baseUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/v1/smarttax", _baseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("getSmartTaxs: http.NewRequestWithContext: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", l.getAccessToken()))

	resp, err := l.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getSmartTaxs: http.Client.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		l.toggleTokenRefresher <- struct{}{}
		return nil, errors.New("connectLAOFIDA: unauthorized (401)")
	}

	if resp.StatusCode != http.StatusOK {
		rbody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("getSmartTaxs: status %d, body: %s", resp.StatusCode, rbody)
	}

	var reply struct {
		Data SmartTaxRecords `json:"data"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		return nil, fmt.Errorf("getSmartTaxs: json.Decode: %w", err)
	}

	return reply.Data, nil
}
