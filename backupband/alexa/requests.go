package alexa

type Intent struct {
	Name  string
	Slots map[string]interface{}
}

type IntentRequest struct {
	Type      string
	RequestID string
	Locale    string
	Timestamp string
	Intent
}
