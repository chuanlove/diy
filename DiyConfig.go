package diy

import (
	"github.com/chuanlove/diy/vo"
	"gorm.io/gorm"
	"im_open/pkg/helper"
	"im_open/pkg/tool"
)

// 字段类型
// 文本
const FIELD_TYPE_INPUT = "input"

// 多行文本
const FIELD_TYPE_TEXTAREA = "textarea"

// 数字
const FIELD_TYPE_NUMBER = "number"

// 金额
const FIELD_TYPE_AMOUNT = "amount"

// 下拉
const FIELD_TYPE_SELECT = "select"

// 联级
const FIELD_TYPE_CASADER = "cascader"

// 单选框
const FIELD_TYPE_RADIO = "radio"

// check
const FIELD_TYPE_CHECKBOX = "checkbox"

// 文件
const FIELD_TYPE_UPLOAD = "upload"

// 时间
const FIELD_TYPE_TIME = "time"

// 日期
const FIELD_TYPE_DATE = "date"

// 评分
const FIELD_TYPE_RATE = "rate"

// 日期范围
const FIELD_TYPE_DATE_RANGE = "date-range"

// 时间范围
const FIELD_TYPE_TIME_RANGE = "time-range"

// 表格
const FIELD_TYPE_TIME_TABLE = "table"

// 布局
const FIELD_TYPE_TIME_ROW = "row"

// 组织架构
const FIELD_TYPE_DEPT = "dept"

// //组织成员
const FIELD_TYPE_USER = "user"

// 关联业务
const FIELD_LINK = "link"

// 反馈与评论
const COMMENT = "comment"

// 文字备注
const TEXTDESC = "textDesc"

// 关联表
const CORRELATION = "linkTable"

// 系统变量
const SYSVARIABLE = "sysVariable"

var FIELD_TYPE_ALL = tool.Array{FIELD_TYPE_INPUT, FIELD_TYPE_TEXTAREA, FIELD_TYPE_NUMBER, FIELD_TYPE_AMOUNT, FIELD_TYPE_RADIO,
	FIELD_TYPE_SELECT, FIELD_TYPE_CASADER, FIELD_TYPE_CHECKBOX, FIELD_TYPE_UPLOAD, FIELD_TYPE_TIME, FIELD_TYPE_DATE, FIELD_TYPE_RATE, FIELD_TYPE_DATE_RANGE,
	FIELD_TYPE_TIME_RANGE, FIELD_TYPE_TIME_TABLE, FIELD_TYPE_TIME_ROW, FIELD_TYPE_DEPT, FIELD_TYPE_USER, FIELD_LINK, COMMENT, TEXTDESC,
	CORRELATION, SYSVARIABLE}

// 默认值为数组的
var OPTION_TYPE_ALL = tool.Array{FIELD_TYPE_CASADER, FIELD_TYPE_DATE_RANGE, FIELD_TYPE_TIME_RANGE, FIELD_TYPE_CHECKBOX, FIELD_TYPE_UPLOAD, FIELD_TYPE_DEPT, FIELD_TYPE_USER}

type Extended struct {
	LabelWidth     int    `json:"label_width"`     //标签宽度
	LabelAlignment string `json:"label_alignment"` // 标签对齐
	FormSize       string `json:"form_size"`       //表单尺寸
	FormStyle      string `json:"form_style"`      //表单样式
}

// / DiyConfig 自定义设计表单
type DiyConfig struct {
	ID                   uint32                    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"id"`
	MchID                uint32                    `gorm:"index:uqKey;column:mch_id;type:int(10) unsigned;not null;default:0" json:"mch_id"` // 商家ID
	Module               string                    `gorm:"index:uqKey;column:module;type:varchar(32);not null;default:''" json:"module"`     // 模块
	Name                 string                    `gorm:"column:name;type:varchar(150);not null;default:''" json:"name"`                    // 表单名称
	FormData             string                    `gorm:"column:form_data;type:text" json:"form_data"`                                      // 表单数据
	FormDataArr          interface{}               `gorm:"-" json:"fields"`                                                                  // 表单数据
	FormExtended         string                    `gorm:"column:form_extended;type:varchar(8000);" json:"-"`                                // 表单样式扩展
	FormExtendedArr      Extended                  `gorm:"-" json:"form_extended"`                                                           // 表单样式扩展
	IsExport             uint8                     `gorm:"column:is_export;type:tinyint(3) unsigned;not null;default:0" json:"is_export"`
	ListColumns          string                    `gorm:"column:list_columns;type:varchar(2000);" json:"-"`      // 列显示
	ListColumnsArr       []vo.CustomListVo         `gorm:"-" json:"list_columns"`                                 // 表单数据
	SearchColumns        string                    `gorm:"column:search_columns;type:varchar(2000);" json:"-"`    // 搜索
	SearchColumnsArr     []vo.CustomListVo         `gorm:"-" json:"search_columns"`                               // 表单数据
	CorrelationConfig    string                    `gorm:"column:correlation_config;type:varchar(200);" json:"-"` // 关联表数据
	CorrelationConfigArr []vo.DiyCorrelationConfig `gorm:"-" json:"correlation_config"`
	IsDelete             uint8                     `gorm:"index:uqKey;column:is_delete;type:tinyint(1) unsigned;not null;default:0" json:"is_delete"`
	CreatedAt            helper.LocalTime          `gorm:"column:created_at;type:timestamp;not null;default:0000-00-00 00:00:00" json:"created_at"` // 创建时间
	UpdatedAt            helper.LocalTime          `gorm:"column:updated_at;type:timestamp;not null;default:0000-00-00 00:00:00" json:"updated_at"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *DiyConfig) TableName() string {
	return "diy_config"
}

// 创建
func (m *DiyConfig) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = helper.Timer()
	m.UpdatedAt = m.CreatedAt
	return
}

// 更新
func (m *DiyConfig) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = helper.Timer()
	return
}
