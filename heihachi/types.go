package go_watch

type Site struct {
	URL string `json:"url"`
	Port int `json:"port"`
}

type NotificationEmailSettings struct {
	Smtp string `json:"smtp"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port int `json:"port"`
}

type TimeSettings struct {
	Duration int64 `json:"duration"`
	Unit string `json:"unit`
}

type Config struct {
	Sites []Site `json:"sites"`
	EmailSettings NotificationEmailSettings `json:"notification_email"`
	Every TimeSettings `json:"every"`
}