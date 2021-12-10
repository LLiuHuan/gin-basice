package system

import (
	"gin-basice/global"
	"gin-basice/model/system"
)

type OperationRecordService struct {
}

// CreateSysOperationRecord 创建接口请求信息
func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.AdpDb.Create(&sysOperationRecord).Error
	return err
}
