package contentpage

type ContentPageData struct {
	ImgURL string
	Name   string
	Author string
}

type ParsedURL struct {
	Type *string
	ID   *string
	Name string
}
