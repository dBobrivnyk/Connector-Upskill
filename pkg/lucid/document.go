package lucid

import (
	"connectors/pkg/entities"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const EntityTypeDocument = "document"

type Document struct {
	Id                 uuid.UUID         `json:"documentId"`
	Title              string            `json:"title"`
	EditUrl            string            `json:"editUrl"`
	ViewUrl            string            `json:"viewUrl"`
	Version            int64             `json:"version"`
	PageCount          int64             `json:"pageCount"`
	CanEdit            bool              `json:"canEdit"`
	Created            time.Time         `json:"created"`
	CreatorId          int64             `json:"creatorId"`
	LastModified       time.Time         `json:"lastModified"`
	LastModifiedUserId int64             `json:"lastModifiedUserId"`
	CustomAttributes   []CustomAttribute `json:"customAttributes"`
	CustomTags         []string          `json:"customTags"`
	Product            Product           `json:"product"`
	Status             string            `json:"status"`
	Trashed            time.Time         `json:"trashed"`
	Parent             int64             `json:"parent"`
}

func (c *Client) SearchDocuments(ctx context.Context) ([]Document, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url+"documents", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Add("Authorization", c.apiKey)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve document data: %w", err)
	}
	defer resp.Body.Close()

	var documents []Document
	err = json.NewDecoder(resp.Body).Decode(&documents)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return documents, nil
}

func (d Document) ToEntity(ownerId string) entities.Entity {
	return entities.Entity{
		Name:         d.Title,
		EntityUrl:    d.ViewUrl,
		ExternalId:   d.Id.String(),
		Type:         EntityTypeDocument,
		ContentUrl:   d.EditUrl,
		OwnerId:      ownerId,
		LastModified: time.Time{},
		Data:         d,
	}
}
