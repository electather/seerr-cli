package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerIssueTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("issue_list",
			mcp.WithDescription("List issues"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
		),
		IssueListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("issue_get",
			mcp.WithDescription("Get a specific issue by ID"),
			mcp.WithNumber("issueId", mcp.Required(), mcp.Description("Issue ID")),
		),
		IssueGetHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("issue_create",
			mcp.WithDescription("Create a new issue"),
			mcp.WithNumber("issueType", mcp.Required(), mcp.Description("Issue type (1=Video, 2=Audio, 3=Subtitle, 4=Other)")),
			mcp.WithString("message", mcp.Required(), mcp.Description("Issue message")),
			mcp.WithNumber("mediaId", mcp.Required(), mcp.Description("Media ID")),
		),
		IssueCreateHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("issue_status_update",
			mcp.WithDescription("Update the status of an issue"),
			mcp.WithString("issueId", mcp.Required(), mcp.Description("Issue ID")),
			mcp.WithString("status", mcp.Required(), mcp.Description("New status (open, resolved)")),
		),
		IssueStatusUpdateHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("issue_count",
			mcp.WithDescription("Get issue counts by status"),
		),
		IssueCountHandler(client, ctx),
	)
}

func IssueListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.IssueAPI.IssueGet(ctx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("IssueGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		issueId, err := req.RequireFloat("issueId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.IssueAPI.IssueIssueIdGet(ctx, float32(issueId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("IssueIssueIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueCreateHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
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
		res, _, err := client.IssueAPI.IssuePost(ctx).IssuePostRequest(body).Execute()
		if err != nil {
			return nil, fmt.Errorf("IssuePost failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueStatusUpdateHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		issueId, err := req.RequireString("issueId")
		if err != nil {
			return nil, err
		}
		status, err := req.RequireString("status")
		if err != nil {
			return nil, err
		}
		res, _, err := client.IssueAPI.IssueIssueIdStatusPost(ctx, issueId, status).Execute()
		if err != nil {
			return nil, fmt.Errorf("IssueIssueIdStatusPost failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func IssueCountHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.IssueAPI.IssueCountGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("IssueCountGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
