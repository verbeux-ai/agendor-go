package agendor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UpdateDealRequest struct {
	// Monetary value of this deal, optional.
	Value *float32 `json:"value,omitempty"`

	// Your own description of this deal, optional, max 2000 characters.
	Description string `json:"description,omitempty"`

	// When this deal was started, optional, in date-time format.
	StartTime time.Time `json:"startTime,omitempty"`

	// Array of product entities associated with this deal, optional.
	ProductsEntities []ProductEntity `json:"products_entities,omitempty"`

	// Array of product names associated with this deal, optional.
	Products []string `json:"products,omitempty"`

	// Owner user ID or email, optional.
	OwnerUser int `json:"ownerUser,omitempty"`

	// Array of IDs of users allowed to see this deal, optional.
	AllowedUsers []int `json:"allowedUsers,omitempty"`

	// Set to true if this deal should be visible to all users, optional.
	AllowToAllUsers *bool `json:"allowToAllUsers,omitempty"`

	// Custom fields information as a key-value object, optional.
	CustomFields map[string]any `json:"customFields,omitempty"`
}

func (s *Client) UpdateDeal(ctx context.Context, dealID int, req UpdateDealRequest) (*Deal, error) {
	resp, err := s.request(ctx, req, http.MethodPut, fmt.Sprintf(dealsOneEndpoint, dealID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyErr := errors.New(string(body))
		return nil, fmt.Errorf("failed to create people with status code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn Response[Deal]
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}

type UpdateDealStageRequest struct {
	// DealStage Sequence number of new deal stage.
	DealStage int `json:"dealStage,omitempty"`
	// Funnel id the id
	Funnel int `json:"funnel,omitempty"`
}

func (s *Client) UpdateDealStage(ctx context.Context, dealID int, req UpdateDealStageRequest) (*Deal, error) {
	resp, err := s.request(ctx, req, http.MethodPut, fmt.Sprintf(dealsOneStageEndpoint, dealID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyErr := errors.New(string(body))
		return nil, fmt.Errorf("failed to create people with status code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn Response[Deal]
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}

type DealStatusType string

const (
	DealStatusOngoing DealStatusType = "ongoing"
	DealStatusWon     DealStatusType = "won"
	DealStatusLost    DealStatusType = "lost"
)

type UpdateDealStatusRequest struct {
	// DealStage Status that the deal should transition to.
	// Enum: "ongoing" "won" "lost"
	DealStatusText DealStatusType `json:"dealStatusText,omitempty"`
	// OwnerUser Owner user ID or email.
	// Just can send id
	OwnerUser int `json:"ownerUser,omitempty"`
	// Endtime  When the deal was finished.
	EndTime time.Time `json:"endTime,omitempty"`
}

func (s *Client) UpdateDealStatus(ctx context.Context, dealID int, req UpdateDealStatusRequest) (*Deal, error) {
	resp, err := s.request(ctx, req, http.MethodPut, fmt.Sprintf(dealsOneStatusEndpoint, dealID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyErr := errors.New(string(body))
		return nil, fmt.Errorf("failed to create people with status code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn Response[Deal]
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}
