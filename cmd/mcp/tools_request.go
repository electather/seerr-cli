package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerRequestTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("request_list",
			mcp.WithDescription("List media requests"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
			mcp.WithString("filter", mcp.Description("Filter by status (all, approved, available, pending, processing, unavailable, failed)")),
			mcp.WithString("sort", mcp.Description("Sort field")),
		),
		RequestListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("request_get",
			mcp.WithDescription("Get a specific media request by ID"),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestGetHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("request_create",
			mcp.WithDescription("Create a new media request"),
			mcp.WithString("mediaType", mcp.Required(), mcp.Description("Media type: movie or tv")),
			mcp.WithNumber("mediaId", mcp.Required(), mcp.Description("TMDB media ID")),
			mcp.WithBoolean("is4k", mcp.Description("Request 4K version")),
		),
		RequestCreateHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("request_approve",
			mcp.WithDescription("Approve a media request"),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestApproveHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("request_decline",
			mcp.WithDescription("Decline a media request"),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestDeclineHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("request_delete",
			mcp.WithDescription("Delete a media request"),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestDeleteHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("request_count",
			mcp.WithDescription("Get request counts by status"),
		),
		RequestCountHandler(client, ctx),
	)
}

func RequestListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.RequestAPI.RequestGet(ctx)
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
			return nil, fmt.Errorf("RequestGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.RequestAPI.RequestRequestIdGet(ctx, requestId).Execute()
		if err != nil {
			return nil, fmt.Errorf("RequestRequestIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestCreateHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mediaType, err := req.RequireString("mediaType")
		if err != nil {
			return nil, err
		}
		mediaId, err := req.RequireFloat("mediaId")
		if err != nil {
			return nil, err
		}
		body := api.NewRequestPostRequest(mediaType, float32(mediaId))
		if is4k := req.GetBool("is4k", false); is4k {
			body.Is4k = &is4k
		}
		res, _, err := client.RequestAPI.RequestPost(ctx).RequestPostRequest(*body).Execute()
		if err != nil {
			return nil, fmt.Errorf("RequestPost failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestApproveHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.RequestAPI.RequestRequestIdStatusPost(ctx, requestId, "approve").Execute()
		if err != nil {
			return nil, fmt.Errorf("RequestRequestIdStatusPost(approve) failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestDeclineHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.RequestAPI.RequestRequestIdStatusPost(ctx, requestId, "decline").Execute()
		if err != nil {
			return nil, fmt.Errorf("RequestRequestIdStatusPost(decline) failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestDeleteHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		_, err = client.RequestAPI.RequestRequestIdDelete(ctx, requestId).Execute()
		if err != nil {
			return nil, fmt.Errorf("RequestRequestIdDelete failed: %w", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}

func RequestCountHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.RequestAPI.RequestCountGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("RequestCountGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
