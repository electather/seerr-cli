package mcp

import (
	"context"
	"encoding/json"

	api "seerr-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerRequestTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("request_list",
			mcp.WithDescription("List media requests"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
			mcp.WithString("filter", mcp.Description("Filter by status (all, approved, available, pending, processing, unavailable, failed)")),
			mcp.WithString("sort", mcp.Description("Sort field")),
		),
		RequestListHandler(),
	)

	s.AddTool(
		mcp.NewTool("request_get",
			mcp.WithDescription("Get a specific media request by ID"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestGetHandler(),
	)

	s.AddTool(
		mcp.NewTool("request_create",
			mcp.WithDescription("Create a new media request"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(false),
			mcp.WithString("mediaType", mcp.Required(), mcp.Description("Media type: movie or tv")),
			mcp.WithNumber("mediaId", mcp.Required(), mcp.Description("TMDB media ID")),
			mcp.WithBoolean("is4k", mcp.Description("Request 4K version")),
		),
		RequestCreateHandler(),
	)

	s.AddTool(
		mcp.NewTool("request_approve",
			mcp.WithDescription("Approve a media request"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(false),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestApproveHandler(),
	)

	s.AddTool(
		mcp.NewTool("request_decline",
			mcp.WithDescription("Decline a media request"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(false),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestDeclineHandler(),
	)

	s.AddTool(
		mcp.NewTool("request_delete",
			mcp.WithDescription("Delete a media request"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(false),
			mcp.WithString("requestId", mcp.Required(), mcp.Description("Request ID")),
		),
		RequestDeleteHandler(),
	)

	s.AddTool(
		mcp.NewTool("request_count",
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithDescription("Get request counts by status"),
		),
		RequestCountHandler(),
	)
}

func RequestListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.RequestAPI.RequestGet(callCtx)
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
			return apiToolError("RequestGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.RequestAPI.RequestRequestIdGet(callCtx, requestId).Execute()
		if err != nil {
			return apiToolError("RequestRequestIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestCreateHandler() server.ToolHandlerFunc {
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
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.RequestAPI.RequestPost(callCtx).RequestPostRequest(*body).Execute()
		if err != nil {
			return apiToolError("RequestPost failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestApproveHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.RequestAPI.RequestRequestIdStatusPost(callCtx, requestId, "approve").Execute()
		if err != nil {
			return apiToolError("RequestRequestIdStatusPost(approve) failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestDeclineHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.RequestAPI.RequestRequestIdStatusPost(callCtx, requestId, "decline").Execute()
		if err != nil {
			return apiToolError("RequestRequestIdStatusPost(decline) failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func RequestDeleteHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		requestId, err := req.RequireString("requestId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		_, err = client.RequestAPI.RequestRequestIdDelete(callCtx, requestId).Execute()
		if err != nil {
			return apiToolError("RequestRequestIdDelete failed", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}

func RequestCountHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.RequestAPI.RequestCountGet(callCtx).Execute()
		if err != nil {
			return apiToolError("RequestCountGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
