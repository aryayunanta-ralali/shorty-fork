package consts

var (
	// RulesUserID rules
	RulesUserID = []string{"between:1,40", "uuid_v4"}

	// RulesURL rules
	RulesURL = []string{"required", "url", "min:3", "max:255"}

	// RulesShortCodeURL rules
	// Example: RalaliBackend-2
	RulesShortCodeURL = []string{"required", "regex:^[a-zA-Z0-9-]{1,255}$", "min:1", "max:255"}
)
