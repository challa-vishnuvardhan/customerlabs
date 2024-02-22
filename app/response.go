package app

type Response struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppId           string               `json:"app_id"`
	UserId          string               `json:"user_id"`
	MessageId       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageUrl         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]TypeValue `json:"attributes"`
	Traits          map[string]TypeValue `json:"traits"`
}

type TypeValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}
