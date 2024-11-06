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

type CreatePeopleRequest struct {
	// Name of this person. Mandatory, max 150 characters.
	Name string `json:"name,omitempty"`

	// Legal document number (CPF), must follow format /\d{11}/.
	Cpf string `json:"cpf,omitempty"`

	// Organization ID, integer that links this person to an organization.
	Organization int `json:"organization,omitempty"`

	// Person's role in their organization, optional.
	Role string `json:"role,omitempty"`

	// Ranking displayed as stars on the person's page. Enum values: 0-5.
	Ranking int `json:"ranking,omitempty"`

	// Description of this person, optional, max 700 characters.
	Description string `json:"description,omitempty"`

	// Person's birthdate with format YYYY-MM-DD or MM-DD, optional.
	Birthday string `json:"birthday,omitempty"`

	// User ID or email of the owner of this person.
	OwnerUser int `json:"ownerUser,omitempty"`

	// Contact information, stored as a nested object.
	Contact Contact `json:"contact,omitempty"`

	// Address information, stored as a nested object.
	Address Address `json:"address,omitempty"`

	// Lead origin ID or name that identifies the source of the lead.
	LeadOrigin int `json:"leadOrigin,omitempty"`

	// Category ID or name for categorizing this person, optional.
	Category int `json:"category,omitempty"`

	// Array of product IDs associated with this person.
	Products []int `json:"products,omitempty"`

	// Array of IDs of users allowed to view this person, optional.
	AllowedUsers []int `json:"allowedUsers,omitempty"`

	// Set to true if this person should be visible to all users, optional.
	AllowToAllUsers bool `json:"allowToAllUsers,omitempty"`

	// Custom fields information, stored as a key-value object, optional.
	CustomFields map[string]any `json:"customFields,omitempty"`
}

func (s *Client) CreatePeople(ctx context.Context, people CreatePeopleRequest) (*People, error) {
	resp, err := s.request(ctx, people, http.MethodPost, peopleEndpoint)
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

	var toReturn Response[People]
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn.Data, nil
}

// FilterPeoplesRequest is a filter to filter peoples by fields
type FilterPeoplesRequest struct {
	Email       string `query:"email"`
	MobilePhone string `query:"mobile_phone"`
	CPF         string `query:"cpf"`
	Name        string `query:"name"`
}

func (s *Client) ListPeoples(ctx context.Context, filter FilterPeoplesRequest) ([]People, error) {
	query, err := StructToQueryString(filter)
	if err != nil {
		return nil, err
	}

	resp, err := s.request(ctx, nil, http.MethodGet, peopleEndpoint+"?"+query)
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
		return nil, fmt.Errorf("failed to list people with status code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn Response[[]People]
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return toReturn.Data, nil
}

type CreatePeopleDealRequest struct {
	// Title to identify this deal. Mandatory, max 150 characters.
	Title string `json:"title,omitempty"`

	// Deal status name, allowed values: "ongoing", "won", "lost".
	DealStatusText string `json:"dealStatusText,omitempty"`

	// Description of the deal, optional, max 2000 characters.
	Description string `json:"description,omitempty"`

	// When this deal was finished. Optional, in date-time format.
	EndTime time.Time `json:"endTime,omitempty"`

	// List of product entities associated with this deal, optional.
	ProductsEntities []CreatePeopleDealRequestProduct `json:"products_entities,omitempty"`

	// List of product names associated with this deal, optional.
	Products []string `json:"products,omitempty"`

	// Ranking displayed as stars on the deal page. Enum values: 0-5.
	Ranking int `json:"ranking,omitempty"`

	// When this deal started, in date-time format. Optional.
	StartTime time.Time `json:"startTime,omitempty"`

	// Owner user ID or email. Mandatory.
	OwnerUser int `json:"ownerUser,omitempty"`

	// Funnel ID, optional.
	Funnel int `json:"funnel,omitempty"`

	// Deal stage sequence number, optional.
	DealStage int `json:"dealStage,omitempty"`

	// Monetary value of this deal, optional.
	Value int `json:"value,omitempty"`

	// Array of user IDs that can view this deal, optional.
	AllowedUsers []int `json:"allowedUsers,omitempty"`

	// Set to true if this deal is visible to all users, optional.
	AllowToAllUsers bool `json:"allowToAllUsers,omitempty"`

	// Custom fields information, stored as a key-value object, optional.
	CustomFields map[string]any `json:"customFields,omitempty"`
}

type CreatePeopleDealRequestProduct struct {
	Id        int `json:"id,omitempty"`
	UnitValue int `json:"unitValue,omitempty"`
	Discount  int `json:"discount,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}

func (s *Client) CreatePeopleDeal(ctx context.Context, peopleID int, peopleDeal CreatePeopleDealRequest) (*Deal, error) {
	resp, err := s.request(ctx, peopleDeal, http.MethodPost, fmt.Sprintf(peopleDealsEndpoint, peopleID))
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
