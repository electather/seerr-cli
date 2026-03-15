package mcp

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPrompts(s *server.MCPServer) {
	s.AddPrompt(
		mcp.NewPrompt("request_media",
			mcp.WithPromptDescription("Request a movie or TV show to be downloaded. Use this when the user says 'I want to watch X' or asks to add something."),
			mcp.WithArgument("title",
				mcp.ArgumentDescription("The title of the movie or TV show to request"),
				mcp.RequiredArgument(),
			),
			mcp.WithArgument("media_type",
				mcp.ArgumentDescription("Media type: 'movie' or 'tv' (optional, will search both if omitted)"),
			),
			mcp.WithArgument("quality",
				mcp.ArgumentDescription("Preferred quality: '4k' or 'hd' (optional)"),
			),
		),
		RequestMediaPromptHandler(),
	)

	s.AddPrompt(
		mcp.NewPrompt("discover_content",
			mcp.WithPromptDescription("Find something to watch based on preferences or trending content."),
			mcp.WithArgument("preferences",
				mcp.ArgumentDescription("Optional mood, genre, or theme description (e.g. 'something funny' or 'sci-fi thriller')"),
			),
			mcp.WithArgument("media_type",
				mcp.ArgumentDescription("Limit to 'movie' or 'tv' (optional)"),
			),
		),
		DiscoverContentPromptHandler(),
	)

	s.AddPrompt(
		mcp.NewPrompt("manage_requests",
			mcp.WithPromptDescription("Admin dashboard for reviewing and actioning pending media requests."),
			mcp.WithArgument("status",
				mcp.ArgumentDescription("Filter by request status: pending, approved, available, failed (optional, defaults to pending)"),
			),
		),
		ManageRequestsPromptHandler(),
	)

	s.AddPrompt(
		mcp.NewPrompt("report_issue",
			mcp.WithPromptDescription("Report a problem with a movie or TV show (e.g. wrong audio, missing subtitles, corrupt file)."),
			mcp.WithArgument("media_title",
				mcp.ArgumentDescription("The title of the media that has a problem"),
				mcp.RequiredArgument(),
			),
			mcp.WithArgument("description",
				mcp.ArgumentDescription("Optional description of the problem"),
			),
		),
		ReportIssuePromptHandler(),
	)

	s.AddPrompt(
		mcp.NewPrompt("my_dashboard",
			mcp.WithPromptDescription("Show a personal overview: your active requests, quota usage, and recent activity."),
		),
		MyDashboardPromptHandler(),
	)

	s.AddPrompt(
		mcp.NewPrompt("admin_overview",
			mcp.WithPromptDescription("System health check: server status, pending requests, open issues, and scheduled jobs."),
		),
		AdminOverviewPromptHandler(),
	)
}

// RequestMediaPromptHandler returns a prompt handler that guides the AI through
// searching for media, checking availability, and creating a download request.
func RequestMediaPromptHandler() server.PromptHandlerFunc {
	return func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		title := req.Params.Arguments["title"]
		mediaType := req.Params.Arguments["media_type"]
		quality := req.Params.Arguments["quality"]

		instructions := fmt.Sprintf(`You are helping the user request media to be downloaded via Seerr.

The user wants to watch: %q

Follow these steps in order:

1. Call search_multi with query=%q to find the title.
2. Present the top 3 results to the user with title, year, and media type. Ask the user to confirm which one they want.
3. Once confirmed, check the mediaInfo.status field on the result:
   - If status is 5 (Available) or 4 (Partially Available), tell the user it is already available and no request is needed.
   - If status is 3 (Processing) or 2 (Pending), tell the user a request is already in progress.
4. If not available, call request_list to check whether an existing request already covers this media (match by mediaId).
5. If no existing request, call request_create with:
   - mediaType: %q (or as determined from search results if not specified)
   - mediaId: the TMDB ID from the search result
   - is4k: %v (true if quality is "4k", false otherwise)
   Use the seerr://services/radarr or seerr://services/sonarr resource to find the appropriate server ID and quality profile if needed.
6. Confirm the request was created and give the user the request ID.`,
			title, title,
			mediaTypeOrDefault(mediaType),
			quality == "4k",
		)

		return &mcp.GetPromptResult{
			Description: "Guide to request a movie or TV show for download",
			Messages: []mcp.PromptMessage{
				{
					Role:    mcp.RoleUser,
					Content: mcp.TextContent{Type: "text", Text: instructions},
				},
			},
		}, nil
	}
}

// DiscoverContentPromptHandler returns a prompt handler that guides the AI through
// finding trending or genre-filtered content for the user.
func DiscoverContentPromptHandler() server.PromptHandlerFunc {
	return func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		preferences := req.Params.Arguments["preferences"]
		mediaType := req.Params.Arguments["media_type"]

		var instructions string
		if preferences != "" {
			instructions = fmt.Sprintf(`You are helping the user discover content to watch.

The user's preferences: %q

Follow these steps:
1. Read the seerr://genres/movies and seerr://genres/tv resources to find genre IDs that match the user's preferences.
2. Based on the media type preference (%q), call search_discover_movies or search_discover_tv with a relevant genre filter.
3. Also call search_trending to supplement with trending content.
4. Present a curated list of 5-8 recommendations, including title, year, genre, and whether each is already Available in Seerr (check mediaInfo.status == 5).
5. Ask the user if they would like to request any of the shown titles.`,
				preferences, mediaTypeOrDefault(mediaType))
		} else {
			instructions = fmt.Sprintf(`You are helping the user discover content to watch.

Follow these steps:
1. Call search_trending to get currently popular content.
2. Based on the media type preference (%q), optionally also call search_discover_movies or search_discover_tv for additional variety.
3. Present a curated list of 5-8 recommendations with title, year, and whether each is Available (mediaInfo.status == 5).
4. Ask the user if they would like to request any of the shown titles.`,
				mediaTypeOrDefault(mediaType))
		}

		return &mcp.GetPromptResult{
			Description: "Guide to discover trending or genre-filtered content",
			Messages: []mcp.PromptMessage{
				{
					Role:    mcp.RoleUser,
					Content: mcp.TextContent{Type: "text", Text: instructions},
				},
			},
		}, nil
	}
}

// ManageRequestsPromptHandler returns a prompt handler that guides an admin through
// reviewing and bulk-actioning pending media requests.
func ManageRequestsPromptHandler() server.PromptHandlerFunc {
	return func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		status := req.Params.Arguments["status"]
		if status == "" {
			status = "pending"
		}

		instructions := fmt.Sprintf(`You are helping an admin review and action media requests in Seerr.

Status filter: %q

Follow these steps:
1. Call request_count to get a summary of all request statuses.
2. Call request_list with filter=%q to retrieve the relevant requests.
3. Group the results by media type (movie vs TV show).
4. For each request, display: title, requester, requested date, and current status.
5. Ask the admin whether they want to approve, decline, or skip each request (or handle them in bulk).
6. For each decision, call request_approve or request_decline with the appropriate requestId.
7. Summarise the actions taken at the end.`,
			status, status)

		return &mcp.GetPromptResult{
			Description: "Admin workflow to review and action pending media requests",
			Messages: []mcp.PromptMessage{
				{
					Role:    mcp.RoleUser,
					Content: mcp.TextContent{Type: "text", Text: instructions},
				},
			},
		}, nil
	}
}

// ReportIssuePromptHandler returns a prompt handler that guides the user through
// searching for media and filing an issue report.
func ReportIssuePromptHandler() server.PromptHandlerFunc {
	return func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		mediaTitle := req.Params.Arguments["media_title"]
		description := req.Params.Arguments["description"]

		descriptionLine := ""
		if description != "" {
			descriptionLine = fmt.Sprintf("\nThe user has described the problem as: %q\n", description)
		}

		instructions := fmt.Sprintf(`You are helping the user report a problem with media in Seerr.

The user has an issue with: %q%s
Follow these steps:
1. Call search_multi with query=%q to find the media.
2. Present the top matches and ask the user to confirm which one has the problem.
3. Ask the user what type of problem they are experiencing:
   - 1: Video quality issue
   - 2: Audio quality issue
   - 3: Subtitle issue
   - 4: Wrong content (wrong episode, wrong cut, etc.)
   - 5: Other
4. If no description was provided, ask the user to describe the problem briefly.
5. Call issue_create with:
   - mediaType: the type from search results
   - mediaId: the TMDB ID
   - issueType: the number selected above (1-5)
   - message: the problem description
6. Confirm the issue was filed and provide the issue ID.`,
			mediaTitle, descriptionLine, mediaTitle)

		return &mcp.GetPromptResult{
			Description: "Guide to report a playback or content issue with media",
			Messages: []mcp.PromptMessage{
				{
					Role:    mcp.RoleUser,
					Content: mcp.TextContent{Type: "text", Text: instructions},
				},
			},
		}, nil
	}
}

// MyDashboardPromptHandler returns a prompt handler that presents a personalised
// overview of the current user's requests and quota.
func MyDashboardPromptHandler() server.PromptHandlerFunc {
	return func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		instructions := `You are presenting a personal Seerr dashboard for the current user.

Follow these steps:
1. Call auth_me to get the current user's profile, ID, and permissions.
2. In parallel:
   a. Call request_list filtered to the current user's requests (match requestedBy.id from auth_me).
   b. Call users_quota with the user's ID to check their remaining request quota.
3. Present a summary including:
   - User name, email, and role
   - Number of active requests (pending/processing) vs completed
   - Quota remaining (movie requests and TV show requests)
   - Any requests that are in a failed or unavailable state, with a note to retry or report an issue
4. Ask if the user would like to take any action (e.g. make a new request, report an issue, or retry a failed request).`

		return &mcp.GetPromptResult{
			Description: "Personal Seerr dashboard showing requests and quota",
			Messages: []mcp.PromptMessage{
				{
					Role:    mcp.RoleUser,
					Content: mcp.TextContent{Type: "text", Text: instructions},
				},
			},
		}, nil
	}
}

// AdminOverviewPromptHandler returns a prompt handler that presents a system-wide
// health check for Seerr administrators.
func AdminOverviewPromptHandler() server.PromptHandlerFunc {
	return func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		instructions := `You are presenting a system health overview for a Seerr administrator.

Follow these steps (run all calls in parallel where possible):
1. Call status_system to get system health and version information.
2. Call request_count to get the breakdown of requests by status.
3. Call issue_count to get the count of open issues.
4. Call settings_jobs_list to see scheduled job statuses.

Present a structured summary:
- System Status: version, whether an update is available, any warnings
- Requests: total pending, processing, failed — highlight any that need attention
- Issues: total open issues — note if there are unresolved issues older than 7 days
- Scheduled Jobs: list any jobs that are currently running or recently failed

Finally, ask the admin if they want to take any action, such as:
- Approving or declining pending requests (use manage_requests prompt)
- Reviewing open issues
- Triggering a scheduled job manually (settings_jobs_run)`

		return &mcp.GetPromptResult{
			Description: "System health check for Seerr administrators",
			Messages: []mcp.PromptMessage{
				{
					Role:    mcp.RoleUser,
					Content: mcp.TextContent{Type: "text", Text: instructions},
				},
			},
		}, nil
	}
}

// mediaTypeOrDefault returns the provided media type string, or "movie or tv" if empty.
func mediaTypeOrDefault(mediaType string) string {
	if mediaType == "" {
		return "movie or tv"
	}
	return mediaType
}
