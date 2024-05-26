package vo

type DiymoduleConfig struct {
	Module     string `json:"module"`
	ConfigID   uint32 `json:"config_id"`
	Status     int    `json:"status"`
	StatusText string `json:"status_text"`
	Name       string `json:"name"`
}

type DiyCorrelationConfig struct {
	Name        string `json:"name"`
	DbTableName string `json:"db_table_name"`
	DbField     string `json:"db_field"`
	CreateField string `json:"create_field"` //映射字段
}

type DiyCorrelationTableColumns struct {
	Field string `json:"field"`
	Title string `json:"title"`
}
type DiyCorrelationTableSearchColumns struct {
	Field      string `json:"field"`
	Title      string `json:"title"`
	SearchType string `json:"search_type"`
}

type DiyCorrelationTable struct {
	TableName    string                       `json:"table_name"`
	TableNameaZh string                       `json:"table_name_zh"` //表中文名称
	Columns      []DiyCorrelationTableColumns `json:"columns"`       //表中文名称
	Models       interface{}
	ShowField    []string
	Search       []DiyCorrelationTableSearchColumns
}
