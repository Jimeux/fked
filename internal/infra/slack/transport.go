package slack

// TODO This should be AppMentionEvent
type Event struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Text        string `json:"text"`
	ClientMsgID string `json:"client_msg_id"`
	TS          string `json:"ts"`
	Channel     string `json:"channel"`
	EventTS     string `json:"event_ts"`
}

type EventRequest struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	ApiAppID string `json:"api_app_id"`
	Event    Event  `json:"event"`
	// Determine request type with this
	Type        string   `json:"type" binding:"required"`
	EventId     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`

	// Need to merge structs because Slack sends requests
	// to a single endpoint.
	ChallengeRequest
}

type ChallengeRequest struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type ChallengeResponse struct {
	Challenge string `json:"challenge"`
}

type MessageChannelRequest struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}
