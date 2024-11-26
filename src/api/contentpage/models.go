package contentpage

type ContentPageData struct {
	Summary ContentSummary
	Content []string
}

type ContentSummary struct {
	CurrentPath string
	ImgURL      string
	Name        string
	Published   string
	Authors     []string
	Genres      []string
}
