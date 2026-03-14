package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerUsersTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("users_list",
			mcp.WithDescription("List all users"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
		),
		UsersListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("users_get",
			mcp.WithDescription("Get a specific user by ID"),
			mcp.WithNumber("userId", mcp.Required(), mcp.Description("User ID")),
		),
		UsersGetHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("users_quota",
			mcp.WithDescription("Get request quota for a specific user"),
			mcp.WithNumber("userId", mcp.Required(), mcp.Description("User ID")),
		),
		UsersQuotaHandler(client, ctx),
	)
}

func UsersListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.UsersAPI.UserGet(ctx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("UserGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func UsersGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		userId, err := req.RequireFloat("userId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.UsersAPI.UserUserIdGet(ctx, float32(userId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("UserUserIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func UsersQuotaHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		userId, err := req.RequireFloat("userId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.UsersAPI.UserUserIdQuotaGet(ctx, float32(userId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("UserUserIdQuotaGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
