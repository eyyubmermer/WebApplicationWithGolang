package models

//TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	StructMap map[string]interface{}
	CSRFToken string
	Flash     string
	Message   string
	Error     string
}
