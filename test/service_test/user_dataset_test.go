package service_test

import (
	"context"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/test/mock"
	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertInstructionData(t *testing.T) {
	var (
		injector       = wire.GetInjector()
		ctx            = injector.Ctx
		datasetService = injector.UserDatasetService
		instruction    = mock.RandomString(10)
		input          = mock.RandomString(10)
		output         = mock.RandomString(10)
		theme          = "THEME1"
		status         = "PENDING"
		note           = ""
	)

	ctx = context.WithValue(ctx, config.UserIDKey, injector.UserDaoMock.RandomUserID().Hex())
	resp, err := datasetService.InsertInstructionData(ctx, &instruction, &input, &output, &theme, &status, &note)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
	instructionDataID, err := primitive.ObjectIDFromHex(resp)
	assert.NoError(t, err)
	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, instruction, instructionData.Row.Instruction)
	assert.Equal(t, input, instructionData.Row.Input)
	assert.Equal(t, output, instructionData.Row.Output)
	assert.Equal(t, theme, instructionData.Theme)
	assert.Equal(t, status, instructionData.Status.Code)
}

func TestUserGetInstructionData(t *testing.T) {
	var (
		injector          = wire.GetInjector()
		ctx               = injector.Ctx
		datasetService    = injector.UserDatasetService
		instructionDataID = injector.InstructionDataDaoMock.RandomInstructionDataID()
	)

	resp, err := datasetService.GetInstructionData(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
}

func TestUserGetInstructionDataList(t *testing.T) {
	var (
		injector       = wire.GetInjector()
		ctx            = injector.Ctx
		datasetService = injector.UserDatasetService
		pageSize       = int64(10)
		page           = int64(1)
		theme          = "test_theme"
		status         = "PENDING"
	)

	ctx = context.WithValue(ctx, config.UserIDKey, injector.UserDaoMock.RandomUserID().Hex())
	resp, err := datasetService.GetInstructionDataList(ctx, &page, &pageSize, nil, nil, &theme, &status)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("Response Data: %+v", resp)
}

func TestUserUpdateInstructionData(t *testing.T) {
	var (
		injector          = wire.GetInjector()
		ctx               = injector.Ctx
		datasetService    = injector.UserDatasetService
		instructionDataID = injector.InstructionDataDaoMock.RandomInstructionDataID()
		instruction       = "test_instruction"
		input             = "test_input"
		output            = "test_output"
		theme             = "test_theme"
		source            = "test_source"
		note              = "test_note"
	)

	err := datasetService.UpdateInstructionData(
		ctx, &instructionDataID, &instruction, &input, &output, &theme, &source, &note,
	)
	assert.NoError(t, err)
	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, instruction, instructionData.Row.Instruction)
	assert.Equal(t, input, instructionData.Row.Input)
	assert.Equal(t, output, instructionData.Row.Output)
	assert.Equal(t, theme, instructionData.Theme)
	assert.Equal(t, source, instructionData.Source)
	assert.Equal(t, note, instructionData.Note)
	t.Logf("Instruction Data: %+v", instructionData)
}

func TestUserDeleteInstructionData(t *testing.T) {
	var (
		injector       = wire.GetInjector()
		ctx            = injector.Ctx
		datasetService = injector.UserDatasetService
		instruction    = mock.RandomString(10)
		input          = mock.RandomString(10)
		output         = mock.RandomString(10)
		theme          = "THEME1"
		status         = "PENDING"
		note           = ""
	)

	ctx = context.WithValue(ctx, config.UserIDKey, injector.UserDaoMock.RandomUserID().Hex())
	resp, err := datasetService.InsertInstructionData(ctx, &instruction, &input, &output, &theme, &status, &note)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	instructionDataID, err := primitive.ObjectIDFromHex(resp)
	assert.NoError(t, err)

	err = datasetService.DeleteInstructionData(ctx, &instructionDataID)
	assert.NoError(t, err)
	instructionData, err := injector.InstructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.Error(t, err)
	assert.Nil(t, instructionData)
}
