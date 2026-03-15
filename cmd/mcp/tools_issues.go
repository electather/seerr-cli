package mcp

import (
	"context"
	"encoding/json"

	api "seerr-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerIssueTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("issue_list",
			mcp.WithDescription("List issues"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
		),
		IssueListHandler(),
	)

	s.AddTool(
		mcp.NewTool("issue_get",
			mcp.WithDescription("Get a specific issue by ID"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("issueId", mcp.Required(), mcp.Description("Issue ID")),
		),
		IssueGetHandler(),
	)

	s.AddTool(
		mcp.NewTool("issue_create",
			mcp.WithDescription("Create a new issue"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(false),
			mcp.WithNumber("issueType", mcp.Required(), mcp.Description("Issue type (1=Video, 2=Audio, 3=Subtitle, 4=Other)")),
			mcp.WithString("message", mcp.Required(), mcp.Description("Issue message")),
			mcp.WithNumber("mediaId", mcp.Required(), mcp.Description("Media ID")),
		),
		IssueCreateHandler(),
	)

	s.AddTool(
		mcp.NewTool("issue_status_update",
			mcp.WithDescription("Update the status of an issue"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(false),
			mcp.WithString("issueId", mcp.Required(), mcp.Description("Issue ID")),
			mcp.WithString("status", mcp.Required(), mcp.Enum("open", "resolved")),
		),
		IssueStatusUpdateHandler(),
	)

	s.AddTool(
		mcp.NewTool("issue_count",
			mcp.WithDescription("Get issue counts by status"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		IssueCountHandler(),
	)
}

func IssueListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.IssueAPI.IssueGet(callCtx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("IssueGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		issueId, err := req.RequireFloat("issueId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.IssueAPI.IssueIssueIdGet(callCtx, float32(issueId)).Execute()
		if err != nil {
			return apiToolError("IssueIssueIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueCreateHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		issueType, err := req.RequireFloat("issueType")
		if err != nil {
			return nil, err
		}
		message, err := req.RequireString("message")
		if err != nil {
			return nil, err
		}
		mediaId, err := req.RequireFloat("mediaId")
		if err != nil {
			return nil, err
		}
		issueTypeFloat := float32(issueType)
		mediaIdFloat := float32(mediaId)
		body := api.IssuePostRequest{
			IssueType: &issueTypeFloat,
			Message:   &message,
			MediaId:   &mediaIdFloat,
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.IssueAPI.IssuePost(callCtx).IssuePostRequest(body).Execute()
		if err != nil {
			return apiToolError("IssuePost failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueStatusUpdateHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		issueId, err := req.RequireString("issueId")
		if err != nil {
			return nil, err
		}
		status, err := req.RequireString("status")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.IssueAPI.IssueIssueIdStatusPost(callCtx, issueId, status).Execute()
		if err != nil {
			return apiToolError("IssueIssueIdStatusPost failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueCountHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.IssueAPI.IssueCountGet(callCtx).Execute()
		if err != nil {
			return apiToolError("IssueCountGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
