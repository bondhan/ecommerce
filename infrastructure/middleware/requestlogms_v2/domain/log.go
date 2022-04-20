package requestlogmsv2_domain

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// LogRequest ...
type LogRequest struct {
	gorm.Model
	IPAddress             null.String `gorm:"column:ip_address" json:"ip_address"`
	Header                null.String `gorm:"column:header" json:"header"`
	Host                  null.String `gorm:"column:host" json:"host"`
	URL                   null.String `gorm:"column:url" json:"url"`
	URI                   null.String `gorm:"column:uri" json:"uri"`
	HTTPMethod            null.String `gorm:"column:http_method" json:"http_method"`
	HTTPRespCode          null.Int    `gorm:"column:http_resp_code" json:"http_resp_code"`
	Message               null.String `gorm:"column:message" json:"message"`
	UserAgent             null.String `gorm:"column:user_agent;type:text" json:"user_agent"`
	RequestID             null.String `gorm:"column:request_id" json:"request_id"`
	RequestTime           null.Time   `gorm:"column:request_time;" json:"request_time"`
	ResponseTime          null.Time   `gorm:"column:response_time;" json:"response_time"`
	LapsedTimeMs          null.Int    `gorm:"column:lapsed_time_ms;" json:"lapsed_time_ms"`
	RequestBody           null.String `gorm:"column:request_body" json:"request_body"`
	ResponseBody          null.String `gorm:"column:response_body" json:"response_body"`
}

// TableName sets the insert table name for this struct type
func (m *LogRequest) TableName() string {
	return "log_request"
}
