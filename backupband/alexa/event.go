package alexa

type Application struct {
	ApplicationID string
}

type User struct {
	UserID string
}

type Session struct {
	New       bool ``
	SessionID string
	Application
	Attributes map[string]interface{}
	User
}

type AudioPlayer struct {
	PlayerActivity string
}

type Device struct {
	SupportedInterfaces map[string]interface{}
}

type System struct {
	Application
	User
}

type Context struct {
	AudioPlayer
	System
}

type Event struct {
	Version string
	Session
	Request map[string]interface{}
}
