package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerSettingsTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("settings_about",
			mcp.WithDescription("Get Seer system information and settings"),
		),
		SettingsAboutHandler(),
	)

	s.AddTool(
		mcp.NewTool("settings_jobs_list",
			mcp.WithDescription("List all scheduled jobs"),
		),
		SettingsJobsListHandler(),
	)

	s.AddTool(
		mcp.NewTool("settings_jobs_run",
			mcp.WithDescription("Trigger a scheduled job to run immediately"),
			mcp.WithString("jobId", mcp.Required(), mcp.Description("Job ID")),
		),
		SettingsJobsRunHandler(),
	)
}

func SettingsAboutHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.SettingsAPI.SettingsAboutGet(callCtx).Execute()
		if err != nil {
			return apiToolError("SettingsAboutGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SettingsJobsListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.SettingsAPI.SettingsJobsGet(callCtx).Execute()
		if err != nil {
			return apiToolError("SettingsJobsGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SettingsJobsRunHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		jobId, err := req.RequireString("jobId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.SettingsAPI.SettingsJobsJobIdRunPost(callCtx, jobId).Execute()
		if err != nil {
			return apiToolError("SettingsJobsJobIdRunPost failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
