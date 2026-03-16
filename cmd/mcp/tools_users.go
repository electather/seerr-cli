package mcp

import (
	"context"
	"encoding/json"

	api "seerr-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerUsersTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("users_list",
			mcp.WithDescription("List all users"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
		),
		UsersListHandler(),
	)

	s.AddTool(
		mcp.NewTool("users_get",
			mcp.WithDescription("Get a specific user by ID"),
			mcp.WithNumber("userId", mcp.Required(), mcp.Description("User ID")),
		),
		UsersGetHandler(),
	)

	s.AddTool(
		mcp.NewTool("users_quota",
			mcp.WithDescription("Get request quota for a specific user"),
			mcp.WithNumber("userId", mcp.Required(), mcp.Description("User ID")),
		),
		UsersQuotaHandler(),
	)

	s.AddTool(
		mcp.NewTool("users_update",
			mcp.WithDescription("Update a user's information by user ID"),
			mcp.WithNumber("userId", mcp.Required(), mcp.Description("User ID")),
			mcp.WithString("email", mcp.Description("New email address")),
			mcp.WithString("username", mcp.Description("New username")),
			mcp.WithNumber("permissions", mcp.Description("Permission bitmask")),
		),
		UsersUpdateHandler(),
	)
}

func UsersListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.UsersAPI.UserGet(callCtx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("UserGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func UsersGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		userId, err := req.RequireFloat("userId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.UsersAPI.UserUserIdGet(callCtx, float32(userId)).Execute()
		if err != nil {
			return apiToolError("UserUserIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func UsersUpdateHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		userId, err := req.RequireFloat("userId")
		if err != nil {
			return nil, err
		}
		body := api.UserUpdatePayload{}
		changed := false
		if v := req.GetString("email", ""); v != "" {
			body.SetEmail(v)
			changed = true
		}
		if v := req.GetString("username", ""); v != "" {
			body.SetUsername(v)
			changed = true
		}
		if v := req.GetFloat("permissions", -1); v >= 0 {
			f := float32(v)
			body.SetPermissions(f)
			changed = true
		}
		if !changed {
			return mcp.NewToolResultError("at least one field must be provided"), nil
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.UsersAPI.UserUserIdPut(callCtx, float32(userId)).UserUpdatePayload(body).Execute()
		if err != nil {
			return apiToolError("UserUserIdPut failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func UsersQuotaHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		userId, err := req.RequireFloat("userId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.UsersAPI.UserUserIdQuotaGet(callCtx, float32(userId)).Execute()
		if err != nil {
			return apiToolError("UserUserIdQuotaGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
