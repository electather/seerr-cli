package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCollectionTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("collection_get",
			mcp.WithDescription("Get collection details by TMDB collection ID"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("collectionId", mcp.Required(), mcp.Description("TMDB collection ID")),
		),
		CollectionGetHandler(),
	)
}

func CollectionGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		collectionId, err := req.RequireInt("collectionId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.CollectionAPI.CollectionCollectionIdGet(callCtx, float32(collectionId)).Execute()
		if err != nil {
			return apiToolError("CollectionCollectionIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
