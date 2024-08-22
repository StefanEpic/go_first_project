package schemas

type CreateNerResults struct {
	Event          *string `json:"event"`
	Organization   *string `json:"organization"`
	Person         *string `json:"person"`
	Location       *string `json:"location"`
	ReportId       *int    `json:"report_id"`
	TextDocumentId *int    `json:"text_document_id"`
	PostId         *int    `json:"post_id"`
}

type UpdateNerResults struct {
	CreateNerResults
}

type ReadNerResults struct {
	ID int `json:"id"`
	CreateNerResults
}
