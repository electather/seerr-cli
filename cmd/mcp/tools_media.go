package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerMediaTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("media_list",
			mcp.WithDescription("List media items"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
			mcp.WithString("filter", mcp.Description("Filter by status")),
			mcp.WithString("sort", mcp.Description("Sort field")),
		),
		MediaListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("media_status_update",
			mcp.WithDescription("Update the status of a media item"),
			mcp.WithString("mediaId", mcp.Required(), mcp.Description("Media ID")),
			mcp.WithString("status", mcp.Required(), mcp.Description("New status (available, partial, processing, pending, unknown)")),
		),
		MediaStatusUpdateHandler(client, ctx),
	)
}

func MediaListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.MediaAPI.MediaGet(ctx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		if filter := req.GetString("filter", ""); filter != "" {
			r = r.Filter(filter)
		}
		if sort := req.GetString("sort", ""); sort != "" {
			r = r.Sort(sort)
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("MediaGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MediaStatusUpdateHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mediaId, err := req.RequireString("mediaId")
		if err != nil {
			return nil, err
		}
		status, err := req.RequireString("status")
		if err != nil {
			return nil, err
		}
		res, _, err := client.MediaAPI.MediaMediaIdStatusPost(ctx, mediaId, status).Execute()
		if err != nil {
			return nil, fmt.Errorf("MediaMediaIdStatusPost failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
