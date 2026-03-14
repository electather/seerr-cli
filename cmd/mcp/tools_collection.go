package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerCollectionTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("collection_get",
			mcp.WithDescription("Get collection details by TMDB collection ID"),
			mcp.WithNumber("collectionId", mcp.Required(), mcp.Description("TMDB collection ID")),
		),
		CollectionGetHandler(client, ctx),
	)
}

func CollectionGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		collectionId, err := req.RequireFloat("collectionId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.CollectionAPI.CollectionCollectionIdGet(ctx, float32(collectionId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("CollectionCollectionIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
