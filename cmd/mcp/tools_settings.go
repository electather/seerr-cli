package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerSettingsTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("settings_about",
			mcp.WithDescription("Get Seer system information and settings"),
		),
		SettingsAboutHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("settings_jobs_list",
			mcp.WithDescription("List all scheduled jobs"),
		),
		SettingsJobsListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("settings_jobs_run",
			mcp.WithDescription("Trigger a scheduled job to run immediately"),
			mcp.WithString("jobId", mcp.Required(), mcp.Description("Job ID")),
		),
		SettingsJobsRunHandler(client, ctx),
	)
}

func SettingsAboutHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.SettingsAPI.SettingsAboutGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("SettingsAboutGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SettingsJobsListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.SettingsAPI.SettingsJobsGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("SettingsJobsGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SettingsJobsRunHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		jobId, err := req.RequireString("jobId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.SettingsAPI.SettingsJobsJobIdRunPost(ctx, jobId).Execute()
		if err != nil {
			return nil, fmt.Errorf("SettingsJobsJobIdRunPost failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
