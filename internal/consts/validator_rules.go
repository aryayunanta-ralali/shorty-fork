package consts

var (
	// RulesUserID rules
	RulesUserID = []string{"between:1,40", "uuid_v4"}

	// RulesURL rules
	RulesURL = []string{"required", "url", "min:3", "max:255"}

	// RulesShortCodeURL rules
	RulesShortCodeURL = []string{"required", "min:3", "max:255"}
)
