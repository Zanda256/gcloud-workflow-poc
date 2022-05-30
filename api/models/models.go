package models

const (
	QcApproved = "QC_STATUS_APPROVED"
	QcRejected = "QC_STATUS_REJECTED"
	QcPreApproved = "QC_STATUS_PRE_APPROVED"
	DocsApproved = "DOCS_STATUS_APPROVED"
	DocsRejected = "DOCS_STATUS_REJECTED"
)

type Status struct {
	QcStatus string `json:"qc_status"`
	DocStatus string `json:"doc_status"`
}

type Order struct {
	State Status `json:"state"`
	A , B int
}

type Item struct {
	State Status `json:"state"`
	C, D int
}

type QcInput struct {
	Action      string          `json:"action"`
	AnswerSheet map[string]bool `json:"answer_sheet"`
}

type QcWfCallback struct {
	Url string `json:"url"`
}

type MainInput struct {
	CustomerId string `json:"customer_id"`
	DBUrl    string `json:"db_url"`
	Order Order `json:"order"`
	Item Item `json:"item"`
	State Status `json:"state"`
}
