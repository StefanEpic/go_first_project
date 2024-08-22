package schemas

import (
	"github.com/google/uuid"
)

type CreateReports struct {
	Uuid                 uuid.UUID                `json:"uuid" binding:"required"`
	MimeTypeFull         string                   `json:"mime_type_full"`
	MimeTypeShort        string                   `json:"mime_type_short"`
	ReportName           string                   `json:"report_name" binding:"required"`
	ReportTitle          string                   `json:"report_title" binding:"required"`
	MessageCode          int                      `json:"message_code" binding:"required"`
	Source               string                   `json:"source" binding:"required"`
	Author               string                   `json:"author" binding:"required"`
	PathToFile           string                   `json:"path_to_file"`
	TextLanguage         string                   `json:"text_language"`
	Alpha3Code           string                   `json:"alpha3_code"`
	LanguageProbability  float64                  `json:"language_probability"`
	LanguageRatio        float64                  `json:"language_ratio"`
	TextTelevancy        string                   `json:"text_relevancy"`
	RelevancyProbability string                   `json:"relevancy_probability"`
	Keywords             string                   `json:"keywords"`
	Morphology           string                   `json:"morphology"`
	CntMorphology        string                   `json:"cnt_morphology"`
	PathToSyntaxResult   string                   `json:"path_to_syntax_result"`
	Annotation           string                   `json:"annotation"`
	TargetTask           string                   `json:"target_task"`
	Route                string                   `json:"route"`
	Translate            string                   `json:"translate"`
	ProjectName          string                   `json:"project_name" binding:"required"`
	StatusClass          [][]string               `json:"status_class"`
	KnowledgeElement     []map[string]interface{} `json:"knowledge_element"`
	Images               []string                 `json:"images"`
	SendingDataStatus    string                   `json:"sending_data_status"`
	ChainedData          string                   `json:"chained_data"`
}

type UpdateReports struct {
	Uuid                 uuid.UUID                `json:"uuid"`
	MimeTypeFull         string                   `json:"mime_type_full"`
	MimeTypeShort        string                   `json:"mime_type_short"`
	ReportName           string                   `json:"report_name"`
	ReportTitle          string                   `json:"report_title"`
	MessageCode          int                      `json:"message_code"`
	Source               string                   `json:"source"`
	Author               string                   `json:"author"`
	PathToFile           string                   `json:"path_to_file"`
	TextLanguage         string                   `json:"text_language"`
	Alpha3Code           string                   `json:"alpha3_code"`
	LanguageProbability  float64                  `json:"language_probability"`
	LanguageRatio        float64                  `json:"language_ratio"`
	TextTelevancy        string                   `json:"text_relevancy"`
	RelevancyProbability string                   `json:"relevancy_probability"`
	Keywords             string                   `json:"keywords"`
	Morphology           string                   `json:"morphology"`
	CntMorphology        string                   `json:"cnt_morphology"`
	PathToSyntaxResult   string                   `json:"path_to_syntax_result"`
	Annotation           string                   `json:"annotation"`
	TargetTask           string                   `json:"target_task"`
	Route                string                   `json:"route"`
	Translate            string                   `json:"translate"`
	ProjectName          string                   `json:"project_name"`
	StatusClass          [][]string               `json:"status_class"`
	KnowledgeElement     []map[string]interface{} `json:"knowledge_element"`
	Images               []string                 `json:"images"`
	SendingDataStatus    string                   `json:"sending_data_status"`
	ChainedData          string                   `json:"chained_data"`
}

type ReadReports struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
	CreateReports
}
