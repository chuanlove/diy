package vo

type SummaryCacheArrVo struct {
	FieldName  string ` json:"field_name"`  // 字段名称
	DiyField   string ` json:"diy_field"`   // 字段key，由系统自动生成
	FieldType  string ` json:"field_type"`  // 类型有input text等
	ValueType  uint8  ` json:"value_type"`  // 数据存储类型1表示文本 2表示json
	FieldValue string ` json:"field_value"` // 内容
}
