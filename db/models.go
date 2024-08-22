package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Reports struct {
	ID                   int              `json:"id" gorm:"primaryKey"`
	Uuid                 uuid.UUID        `json:"uuid" gorm:"type: uuid; not null"`
	MimeTypeFull         *string          `json:"mime_type_full" gorm:"type:varchar(100); default:''"`
	MimeTypeShort        *string          `json:"mime_type_short" gorm:"type:varchar(30); default:''"`
	Date                 time.Time        `json:"date" gorm:"autoCreateTime"`
	ReportName           string           `json:"report_name" gorm:"type:varchar(255); not null"`
	ReportTitle          string           `json:"report_title" gorm:"type:varchar(1000); index; not null"`
	MessageCode          int              `json:"message_code" gorm:"index; not null"`
	Source               string           `json:"source" gorm:"type:varchar(100); index; not null"`
	Author               string           `json:"author" gorm:"type:varchar(100); not null"`
	PathToFile           *string          `json:"path_to_file" gorm:"type:varchar(255); default:''"`
	TextLanguage         *string          `json:"text_language" gorm:"type:varchar(20); default:''"`
	Alpha3Code           *string          `json:"alpha3_code" gorm:"type:varchar(10); default:''"`
	LanguageProbability  *float64         `json:"language_probability" gorm:"type:float"`
	LanguageRatio        *float64         `json:"language_ratio" gorm:"type:float"`
	TextTelevancy        *string          `json:"text_relevancy" gorm:"type:varchar(30); default:''"`
	RelevancyProbability *string          `json:"relevancy_probability" gorm:"type:varchar(60); default:''"`
	Keywords             *string          `json:"keywords" gorm:"type:text; default:''"`
	Morphology           *string          `json:"morphology" gorm:"type:text; default:''"`
	CntMorphology        *string          `json:"cnt_morphology" gorm:"type:text; default:''"`
	PathToSyntaxResult   *string          `json:"path_to_syntax_result" gorm:"type:varchar(255); default:''"`
	Annotation           *string          `json:"annotation" gorm:"type:text; default:''"`
	TargetTask           *string          `json:"target_task" gorm:"type:varchar(100); default:''"`
	Route                *string          `json:"route" gorm:"type:varchar(50); default:''"`
	Translate            *string          `json:"translate" gorm:"type:text; default:''"`
	ProjectName          string           `json:"project_name" gorm:"type:varchar(150);index;not null"`
	StatusClass          *ListListStrings `json:"status_class" gorm:"type:text[]"`
	KnowledgeElement     *ListJsonb       `json:"knowledge_element" gorm:"type:JSONB[]"`
	Images               *pq.StringArray  `json:"images" gorm:"type:bytea;serializer:json"`
	SendingDataStatus    *string          `json:"sending_data_status" gorm:"type:text; default:''"`
	ChainedData          *string          `json:"chained_data" gorm:"type:bytea;serializer:json"`
	NerResults           []NerResults     `json:"ner_results" gorm:"foreignKey:ReportId; constraint:OnDelete:CASCADE;"`
}

type NerResults struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Event        *string `json:"event" gorm:"type:text; default:''"`
	Organization *string `json:"organization" gorm:"type:text; default:''"`
	Person       *string `json:"person" gorm:"type:text; default:''"`
	Location     *string `json:"location" gorm:"type:text; default:''"`
	ReportId     *int    `json:"report_id" gorm:"foreignKey:ID; uniqueIndex:compositeindex"`
}
