package vo

// 表单字段
type Fields struct {
	CmpType     string `json:"cmpType"`     //常用组件-common 定制组件-custom
	Label       string `json:"label"`       //表单label
	Tag         string `json:"tag"`         //组件名
	DiyField    string `json:"diyField"`    //表单唯一
	TagIcon     string `json:"tagIcon"`     //组件icon
	Placeholder string `json:"placeholder"` //占位提示
	Span        int    `json:"span"`        //栅格
	LabelWidth  int    `json:"labelWidth"`  //标签宽度
	Them        string `json:"them"`        //主题

	Clearable       bool     `json:"clearable"`       //是否可清除
	Disabled        bool     `json:"disabled"`        //是否禁用
	Required        bool     `json:"required"`        //是否必填
	Filterable      bool     `json:"filterable"`      //是否可搜索
	DefaultValue    string   `json:"defaultValue"`    //默认值
	DefaultArrValue []string `json:"defaultArrValue"` //默认值
	AsSummary       bool     `json:"asSummary"`       //是否为摘要

	Multiple     bool          `json:"multiple"`     //是否多选
	Options      []Options     `json:"options"`      //select选项
	ProCondition bool          `json:"proCondition"` //是否被流程图作为条件必填项
	FormId       string        `json:"formId"`       //表单ID
	RegList      []interface{} `json:"regList"`
	ChangeTag    bool          `json:"changeTag"`
	//表格下的
	ActionText string `json:"actionText"` //添加参数
	Layout     string `json:"layout"`     //显示方式
	ButtonType string `json:"buttonType"` //展示方式，组织架构
	Title      string `json:"title"`      //展示按钮名称
	Type       string `json:"type"`       //类型 用于表格展示方式
	TableType  string `json:"tableType"`  //类型 tableType 用于table 或者是list
	//子级，如果表格或者是布局的情况下
	Children []Fields `json:"children"`
	//关联业务
	BusinessType []int `json:"businessType"` //业务类型
	ShowDuration bool  `json:"showDuration"` //时间范围，显示时长
	//关联表
	DbTableName string `json:"db_table_name"` //关联表
	DbField     string `json:"db_field"`      //关联表字段
	DbType      uint8  `json:"db_type"`       //关联类型 1表示关联发起 2表示显示关联数据
	//系统变量
	Variable string `json:"variable"` //关联表字段
}

// 选项
type Options struct {
	ID       string    `json:"id"`
	Label    string    `json:"label"`
	Value    string    `json:"value"`
	Children []Options `json:"children"` //联级时使用
}

// 部门或者是成员
type SubmitFormData struct {
	DiyField   string        `json:"diy_field"`
	Value      string        `json:"value"`              // 值
	GroupIndex uint32        `json:"group_index"`        // 组ID
	IsArr      bool          `json:"is_arr"`             // 是否为数组
	ValueArr   []interface{} `gorm:"-" json:"value_arr"` //数组
}

// 自定义列表
type CustomListVo struct {
	Field        string        `json:"field"`         //字段
	Name         string        `json:"name"`          //字段名称
	Width        int           `json:"width"`         //宽度
	Style        uint8         `json:"style"`         //显示方式 1 2
	DefaultValue string        `json:"default_value"` //默认值
	IsKeyWord    uint8         `json:"is_key_word"`   //关键词
	Options      []interface{} `json:"options"`       //关键词
	Tag          string        `json:"tag"`           //关键词
	PField       string        `json:"p_field"`       //上级
}

type Keywords struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
