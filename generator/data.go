package generator

type Question struct {
	Name         string
	AskText      string
	DefaultValue string
}

type Stub struct {
	SourceFiles []string
	DestDir     string
	SourceDir   string
}

var mapStub = map[string]Stub{
	"ucase": Stub{
		SourceDir: "generator/templates/ucase",
		DestDir:   "internal/ucase",
		SourceFiles: []string{
			"ucase_get.go.exp",
			"ucase_post.go.exp",
			"ucase_put.go.exp",
		},
	},
	"presentations": Stub{
		SourceDir: "generator/templates/presentations",
		DestDir:   "internal/presentations",
		SourceFiles: []string{
			"presentations_get.go.exp",
			"presentations_post.go.exp",
			"presentations_put.go.exp",
		},
	},
	"repository": Stub{
		SourceDir: "generator/templates",
		DestDir:   "internal/repositories",
		SourceFiles: []string{
			"repository.go.exp",
		},
	},
	"entity": Stub{
		SourceDir: "generator/templates",
		DestDir:   "internal/entity",
		SourceFiles: []string{
			"entity.go.exp",
		},
	},
}

var ucaseQuestion = []Question{
	{
		Name:         "package-name",
		AskText:      "What is the package name? [ex: order, cart, etc]: ",
		DefaultValue: "",
	},
	{
		Name:         "function-method",
		AskText:      "What is the function method? [enum: get, post, put]: ",
		DefaultValue: "",
	},
	{
		Name:         "function-name",
		AskText:      "What is the file name / function name? (ex: get_list, insert_order) : ",
		DefaultValue: "",
	},
	{
		Name:         "is-create-presentation",
		AskText:      "Do you also want to generate the presentation? (yes | no) : ",
		DefaultValue: "",
	},
}

var repositoryQuestion = []Question{
	{
		Name:         "file-name",
		AskText:      "What is the file name? (ex: digital_good_orders) : ",
		DefaultValue: "",
	},
	{
		Name:         "create-entity",
		AskText:      "Do you also want to generate the entity file? (yes | no) : ",
		DefaultValue: "",
	},
}

var mappedQuestion = map[string][]Question{
	"ucase":      ucaseQuestion,
	"repository": repositoryQuestion,
}
