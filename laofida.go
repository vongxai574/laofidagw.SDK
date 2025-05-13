package laofida

import (
	"context"
	"net/http"
	"sync"
	"time"
)

var _ LAOFIDA = (*laofida)(nil)

type (
	Config struct {
		BaseURL string

		Username string
		Password string
	}

	laofida struct {
		baseUrl string

		// clientID is the client id of laofida api backend.
		username string
		password string

		// accessToken is used to authenticate with LDB backend.
		accessToken string

		// mu is used to lock access token.
		mu sync.Mutex

		// toggleTokenRefresher is used to notify token refresher to refresh token.
		toggleTokenRefresher chan struct{}

		// hc is the http client.
		hc *http.Client
	}
)

type LAOFIDA interface {
	GetDataSmartTaxs(ctx context.Context, req *ReqFilter) (SmartTaxRecords, error)
}

// New creates new instance of LAOFIDA client.
func New(ctx context.Context, cfg *Config) (LAOFIDA, error) {
	client := &laofida{
		baseUrl:  cfg.BaseURL,
		username: cfg.Username,
		password: cfg.Password,

		// make a buffered channel to avoid blocking.
		toggleTokenRefresher: make(chan struct{}, 1),

		// set http client with timeout.
		hc: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	token, err := client.connect(ctx)
	if err != nil {
		return nil, err
	}
	client.setAccessToken(token)

	go client.notifyAccessTokenExpired(ctx)

	return client, nil
}

func (l *laofida) GetDataSmartTaxs(ctx context.Context, req *ReqFilter) (SmartTaxRecords, error) {
	return l.getSmartTaxs(ctx, req)
}
