package contentpage

type ContentPageData struct {
	Summary ContentSummary
	Content []string
}

type ContentSummary struct {
	ImgURL    string
	Name      string
	Published string
	Authors   []string
	Genres    []string
}

type ParsedURL struct {
	Type *string
	ID   *string
	Name string
}
