package aitutor

type Prompt struct {
	Init    string  `json:"init"`
	AITutor AITutor `json:"ai_tutor"`
}

type AITutor struct {
	Author             string             `json:"Author"`
	Name               string             `json:"name"`
	Version            string             `json:"version"`
	Commands           Commands           `json:"commands"`
	Features           Features           `json:"features"`
	Formats            Formats            `json:"formats"`
	Rules              []string           `json:"rules"`
	StudentPreferences StudentPreferences `json:"student preferences"`
}

type Commands struct {
	Prefix   string            `json:"prefix"`
	Commands map[string]string `json:"commands"`
}

type Features struct {
	Personalization Personalization `json:"personalization"`
	Plugins         bool            `json:"plugins"`        // false
	Internet        bool            `json:"internet"`       // false
	UseEmojis       bool            `json:"use_emojis"`     // true
	PythonEnabled   bool            `json:"python_enabled"` // false
}

type Formats struct {
	Description           string   `json:"Description"`
	Configuration         []string `json:"configuration"`
	ConfigurationReminder []string `json:"configuration_reminder"`
	SelfEvaluation        []string `json:"self-evaluation"`
	Planning              []string `json:"Planning"`
}

type StudentPreferences struct {
	Description        string   `json:"Description"`
	Depth              int      `json:"depth"`
	CommunicationStyle []string `json:"communication_style"`
	// FeedbackType       []string `json:"feedback_type"`
	LearningStyle      []string `json:"learning_style"`
	ReasoningFramework []string `json:"reasoning_framework"`
	ToneStyle          []string `json:"tone_style"`
}
