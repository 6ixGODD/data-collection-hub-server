package service_test

import (
	"testing"
	"time"

	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
)

func TestAdminGetInstructionData(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		dataAuditService = injector.AdminDataAuditService
	)
	resp, err := dataAuditService.GetInstructionData(ctx, injector.InstructionDataDaoMock.RandomInstructionDataID())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
}

func TestGetInstructionDataList(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		dataAuditService = injector.AdminDataAuditService
		page             = int64(1)
		pageSize         = int64(10)
		desc             = false
		userID           = injector.UserDaoMock.RandomUserID()
		createTimeStart  = time.Now().AddDate(0, 0, -1)
		createTimeEnd    = time.Now()
		updateTimeStart  = time.Now().AddDate(0, 0, -1)
		updateTimeEnd    = time.Now()
		theme            = "THEME1"
		status           = "PENDING"
		query            = "a"
	)
	resp, err := dataAuditService.GetInstructionDataList(
		ctx, &page, &pageSize, &desc, &userID, &createTimeStart, &createTimeEnd, &updateTimeStart, &updateTimeEnd,
		&theme, &status, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = dataAuditService.GetInstructionDataList(
		ctx, &page, &pageSize, &desc, nil, nil, nil, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.InstructionDataList)
	assert.Equal(t, pageSize, int64(len(resp.InstructionDataList)))

	t.Logf("Response Data: %+v", resp)
}

func TestApproveInstructionData(t *testing.T) {
	var (
		injector          = wire.GetInjector()
		ctx               = injector.Ctx
		dataAuditService  = injector.AdminDataAuditService
		instructionDataID = injector.InstructionDataDaoMock.RandomInstructionDataID()
	)
	err := dataAuditService.ApproveInstructionData(ctx, &instructionDataID)
	assert.NoError(t, err)

	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, "APPROVED", instructionData.Status.Code)

	t.Logf("Instruction Data: %+v", instructionData)
}

func TestRejectInstructionData(t *testing.T) {
	var (
		injector          = wire.GetInjector()
		ctx               = injector.Ctx
		dataAuditService  = injector.AdminDataAuditService
		instructionDataID = injector.InstructionDataDaoMock.RandomInstructionDataID()
		message           = "Message"
	)
	err := dataAuditService.RejectInstructionData(ctx, &instructionDataID, &message)
	assert.NoError(t, err)

	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, "REJECTED", instructionData.Status.Code)
	assert.Equal(t, message, instructionData.Status.Message)

	t.Logf("Instruction Data: %+v", instructionData)
}

func TestUpdateInstructionData(t *testing.T) {
	var (
		injector          = wire.GetInjector()
		ctx               = injector.Ctx
		dataAuditService  = injector.AdminDataAuditService
		instructionDataID = injector.InstructionDataDaoMock.RandomInstructionDataID()
		userID            = injector.UserDaoMock.RandomUserID()
		instruction       = "Instruction"
		input             = "Input"
		output            = "Output"
		theme             = "THEME1"
		source            = "https://source.com"
		note              = "Note"
	)
	err := dataAuditService.UpdateInstructionData(
		ctx, &instructionDataID, &userID, &instruction, &input, &output, &theme, &source, &note,
	)
	assert.NoError(t, err)

	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, userID, instructionData.UserID)
	assert.Equal(t, instruction, instructionData.Row.Instruction)
	assert.Equal(t, input, instructionData.Row.Input)
	assert.Equal(t, output, instructionData.Row.Output)
	assert.Equal(t, theme, instructionData.Theme)
	assert.Equal(t, source, instructionData.Source)
	assert.Equal(t, note, instructionData.Note)

	t.Logf("Instruction Data: %+v", instructionData)
}

func TestExportInstructionData(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		dataAuditService = injector.AdminDataAuditService
		desc             = false
		userID           = injector.UserDaoMock.RandomUserID()
		createTimeStart  = time.Now().AddDate(0, 0, -1)
		createTimeEnd    = time.Now()
		updateTimeStart  = time.Now().AddDate(0, 0, -1)
		updateTimeEnd    = time.Now()
		theme            = "THEME1"
		status           = "PENDING"
	)
	resp, err := dataAuditService.ExportInstructionData(
		ctx, &desc, &userID, &createTimeStart, &createTimeEnd, &updateTimeStart, &updateTimeEnd, &theme, &status,
	)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = dataAuditService.ExportInstructionData(
		ctx, &desc, nil, nil, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.InstructionDataList)

	t.Logf("Response Data: %+v", resp)
}

func TestExportInstructionDataAsAlpaca(t *testing.T) {
	var (
		injector         = wire.GetInjector()
		ctx              = injector.Ctx
		dataAuditService = injector.AdminDataAuditService
		desc             = false
		userID           = injector.UserDaoMock.RandomUserID()
		createTimeStart  = time.Now().AddDate(0, 0, -1)
		createTimeEnd    = time.Now()
		updateTimeStart  = time.Now().AddDate(0, 0, -1)
		updateTimeEnd    = time.Now()
		theme            = "THEME1"
		status           = "PENDING"
	)
	resp, err := dataAuditService.ExportInstructionDataAsAlpaca(
		ctx, &desc, &userID, &createTimeStart, &createTimeEnd, &updateTimeStart, &updateTimeEnd, &theme, &status,
	)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	resp, err = dataAuditService.ExportInstructionDataAsAlpaca(
		ctx, &desc, nil, nil, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	t.Logf("Response Data: %+v", resp)
}

func TestDeleteInstructionData(t *testing.T) {
	var (
		injector          = wire.GetInjector()
		ctx               = injector.Ctx
		dataAuditService  = injector.AdminDataAuditService
		instructionDataID = injector.InstructionDataDaoMock.RandomInstructionDataID()
	)
	err := dataAuditService.DeleteInstructionData(ctx, &instructionDataID)
	assert.NoError(t, err)

	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.Error(t, err)
	assert.Nil(t, instructionData)

	t.Logf("Instruction Data ID: %s", instructionDataID)
}
