package huggingface

const (
	huggingfaceDownloadFileTemplate = "%s/%s/resolve/%s/%s"
	HuggingfaceURL                  = "https://huggingface.co"
	defaultRevision                 = "main"
)

type ModelInfo struct {
	ID               string           `json:"_id"`
	ModelID          string           `json:"modelId"`
	Author           string           `json:"author"`
	SHA              string           `json:"sha"`
	LastModified     string           `json:"lastModified"`
	Private          bool             `json:"private"`
	Disabled         bool             `json:"disabled"`
	Gated            bool             `json:"gated"`
	PipelineTag      string           `json:"pipeline_tag"`
	Tags             []string         `json:"tags"`
	Downloads        int              `json:"downloads"`
	LibraryName      string           `json:"library_name"`
	WidgetData       *[]WidgetData    `json:"widgetData"`
	Likes            int              `json:"likes"`
	ModelIndex       *[]ModelIndex    `json:"model-index"`
	Config           Config           `json:"config"`
	CardData         CardData         `json:"cardData"`
	TransformersInfo TransformersInfo `json:"transformersInfo"`
	Spaces           *[]string        `json:"spaces"`
	Siblings         *[]Sibling       `json:"siblings"`
}

type WidgetData struct {
	Text         string `json:"text"`
	ExampleTitle string `json:"example_title"`
	Group        string `json:"group"`
}

type ModelIndex struct {
	Name    string    `json:"name"`
	Results []Results `json:"results"`
}

type Results struct {
	Task    Task      `json:"task"`
	Dataset Dataset   `json:"dataset"`
	Metrics []Metrics `json:"metrics"`
}

type Task struct {
	Type string `json:"type"`
}

type Dataset struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Metrics struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
	Verified bool    `json:"verified"`
}

type Config struct {
	Architectures []string          `json:"architectures"`
	ModelType     string            `json:"model_type"`
	AutoMap       map[string]string `json:"auto_map"`
}

type CardData struct {
	PipelineTag         string        `json:"pipeline_tag"`
	License             string        `json:"license"`
	Tags                []string      `json:"tags"`
	ProgrammingLanguage []string      `json:"programming_language"`
	Metrics             []string      `json:"metrics"`
	Inference           bool          `json:"inference"`
	Widget              *[]WidgetData `json:"widget"`
	ModelIndex          *[]ModelIndex `json:"model-index"`
	Datasets            []string      `json:"datasets"`
}

type TransformersInfo struct {
	AutoModel   string `json:"auto_model"`
	PipelineTag string `json:"pipeline_tag"`
	Processor   string `json:"processor"`
}

type Sibling struct {
	Rfilename string `json:"rfilename"`
}
