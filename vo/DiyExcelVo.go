package vo

type DiyExcelHeadVo struct {
	Title          string           `json:"title"`           //标题
	Position       string           `json:"position"`        //位置
	PositionLetter string           `json:"position_letter"` //位置
	IsTable        bool             `json:"is_table"`        //是否表格
	Option         []string         `json:"option"`          //下拉
	Required       bool             `json:"required"`        //是否必填
	DiyField       string           `json:"diy_field"`
	MaxLine        int              `json:"max_line"`
	Children       []DiyExcelHeadVo `json:"children"` //子集
}

type DataListStruct struct {
	Position  string `json:"position"` //位置
	Line      uint32 `json:"line"`     //行
	MaxLine   uint32 `json:"max_line"` //最大行
	Value     string `json:"value"`
	FieldType string `json:"field_type"` // 类型有input text等
}

type ImportData struct {
	Line      int              `json:"line"`
	Postion   string           `json:"postion"`    //位置
	MergeLine int              `json:"merge_line"` //合并行
	Result    bool             `json:"result"`     //结果
	ResultID  uint32           `json:"result_id"`  //插入ID
	ResultMsg string           `json:"result_msg"` //消息
	Data      []SubmitFormData `json:"data"`
}

type ErrMsg struct {
	Line     int    `json:"line"`     //
	Position string `json:"position"` //
	Msg      string `json:"msg"`
}
