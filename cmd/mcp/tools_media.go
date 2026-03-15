package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// statusLabels maps numeric Seer media status codes to human-readable names.
var statusLabels = map[float64]string{
	1: "UNKNOWN",
	2: "PENDING",
	3: "PROCESSING",
	4: "PARTIALLY_AVAILABLE",
	5: "AVAILABLE",
	6: "DELETED",
}

func registerMediaTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("media_list",
			mcp.WithDescription(
				"List movies and TV shows tracked by Seer (i.e. requested or added to the library), "+
					"including their download/availability status and associated requests. "+
					"Supports pagination via take/skip. "+
					"Each result includes statusLabel and status4kLabel string fields alongside the numeric status codes.",
			),
			mcp.WithNumber("take", mcp.Description("Max number of results to return per page (omit for all).")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip, for pagination.")),
			mcp.WithString("filter", mcp.Description(
				"Filter results by availability status. Valid values: "+
					"all (default, no filter), "+
					"available (fully available), "+
					"partial (partially available), "+
					"allavailable (available or partially available), "+
					"processing (currently being downloaded), "+
					"pending (requested but not yet downloading), "+
					"deleted (removed from library).",
			)),
			mcp.WithString("sort", mcp.Description(
				"Field to sort results by. Valid values: "+
					"added (default, sort by date added to Seer), "+
					"modified (sort by last modified date), "+
					"mediaAdded (sort by date media was added to the library).",
			)),
		),
		MediaListHandler(),
	)

	s.AddTool(
		mcp.NewTool("media_status_update",
			mcp.WithDescription("Update the status of a media item"),
			mcp.WithString("mediaId", mcp.Required(), mcp.Description("Media ID")),
			mcp.WithString("status", mcp.Required(), mcp.Description("New status (available, partial, processing, pending, unknown)")),
		),
		MediaStatusUpdateHandler(),
	)
}

func MediaListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.MediaAPI.MediaGet(callCtx)
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
			return apiToolError("MediaGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}

		// Enrich the response with human-readable status labels so agents can
		// interpret numeric codes without a separate lookup table.
		var envelope map[string]interface{}
		if err := json.Unmarshal(b, &envelope); err != nil {
			return mcp.NewToolResultText(string(b)), nil
		}
		if results, ok := envelope["results"].([]interface{}); ok {
			for _, item := range results {
				m, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if v, ok := m["status"].(float64); ok {
					m["statusLabel"] = statusLabels[v]
				}
				if v, ok := m["status4k"].(float64); ok {
					m["status4kLabel"] = statusLabels[v]
				}
			}
		}
		envelope["_statusLegend"] = map[string]string{
			"1": "UNKNOWN",
			"2": "PENDING",
			"3": "PROCESSING",
			"4": "PARTIALLY_AVAILABLE",
			"5": "AVAILABLE",
			"6": "DELETED",
		}
		enriched, err := json.Marshal(envelope)
		if err != nil {
			return mcp.NewToolResultText(string(b)), nil
		}
		return mcp.NewToolResultText(string(enriched)), nil
	}
}

func MediaStatusUpdateHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		mediaId, err := req.RequireString("mediaId")
		if err != nil {
			return nil, err
		}
		status, err := req.RequireString("status")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.MediaAPI.MediaMediaIdStatusPost(callCtx, mediaId, status).Execute()
		if err != nil {
			return apiToolError("MediaMediaIdStatusPost failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
