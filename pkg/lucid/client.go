package lucid

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

const ApiUrl = "https://api.clickup.com/api/v2/"
const MockApiUrl = "https://api.clickup.com/api/v2/"

type Client struct {
	//cfg     *config.Config
	client  *http.Client
	url     string
	ownerId string
	apiKey  string
	//sink    *sink.Sink
}

func NewClient(url string, ownerId, apiKey string) *Client {
	return &Client{
		client:  http.DefaultClient,
		url:     url,
		ownerId: ownerId,
		apiKey:  apiKey,
	}
}

func (c *Client) GetEntities(ctx context.Context) {
	c.sink = sink.New(c.cfg.BufferSize, c.ownerId)
	defer c.sink.Dump()
	defer c.sink.Close()

	log.Println("Searching among documents")
	documents, err := c.SearchDocuments(ctx)
	if err != nil {
		log.Print(err)
		return
	}

	//log.Println("Getting spaces")
	//spaces := c.getSpacesForWorkspaces(ctx, documents)
	//
	//log.Println("Getting folders")
	//folders := c.getFoldersForSpaces(ctx, spaces)
	//
	//log.Println("Getting lists")
	//lists := c.getListsForFolders(folders)
	//
	//log.Println("Getting tasks")
	//_ = c.getTasksForLists(ctx, lists)
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		responseBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("failed to retrieve workspaces: %s", string(responseBytes))
	}

	return resp, nil
}
