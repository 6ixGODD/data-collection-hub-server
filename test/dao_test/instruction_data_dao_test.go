package dao_test

import (
	"testing"
	"time"

	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var instructionDataID primitive.ObjectID

func TestInsertInstructionData(t *testing.T) {
	// t.Skip("Skip TestInsertInstructionData")
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		userID             = injector.UserDaoMock.RandomUserID()
		instruction        = "Instruction"
		input              = "Input"
		output             = "Output"
		theme              = "Theme"
		source             = "Source"
		note               = "Note"
		statusCode         = "PENDING"
		statusMsg          = "Pending for review"
		err                error
	)

	instructionDataID, err = instructionDataDao.InsertInstructionData(
		ctx, userID,
		instruction, input, output, theme, source, note,
		statusCode, statusMsg,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, instructionDataID)

	instructionData, err := instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, instruction, instructionData.Row.Instruction)
	assert.Equal(t, input, instructionData.Row.Input)
	assert.Equal(t, output, instructionData.Row.Output)
	assert.Equal(t, theme, instructionData.Theme)
	assert.Equal(t, source, instructionData.Source)
	assert.Equal(t, note, instructionData.Note)
	assert.Equal(t, statusCode, instructionData.Status.Code)
	assert.Equal(t, statusMsg, instructionData.Status.Message)

}

func TestGetInstructionData(t *testing.T) {
	// t.Skip("Skip TestGetInstructionData")
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		err                error
	)
	instructionData, err := instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.NotEmpty(t, instructionData.InstructionDataID)
	assert.NotEmpty(t, instructionData.UserID)
	assert.NotEmpty(t, instructionData.Username)
	assert.NotEmpty(t, instructionData.Row.Instruction)
	assert.NotEmpty(t, instructionData.Row.Input)
	assert.NotEmpty(t, instructionData.Row.Output)
	assert.NotEmpty(t, instructionData.Theme)
	assert.NotEmpty(t, instructionData.Source)
	assert.NotEmpty(t, instructionData.Note)
	assert.NotEmpty(t, instructionData.Status.Code)
	assert.NotEmpty(t, instructionData.Status.Message)
	assert.NotEmpty(t, instructionData.CreatedAt)
	assert.NotEmpty(t, instructionData.UpdatedAt)
	assert.False(t, instructionData.Deleted)
}

func TestGetInstructionDataList(t *testing.T) {
	// t.Skip("Skip TestGetInstructionDataList")
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		userID             = injector.UserDaoMock.RandomUserID()
		theme              = "THEME1"
		statusCode         = "PENDING"
		createStartTime    = time.Now().AddDate(0, 0, -1)
		createEndTime      = time.Now().AddDate(0, 0, 1)
		updateStartTime    = time.Now().AddDate(0, 0, -1)
		updateEndTime      = time.Now().AddDate(0, 0, 1)
		query              = "a"
		err                error
	)
	instructionDataList, count, err := instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, nil, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotEmpty(t, *count)
	assert.NotEmpty(t, instructionDataList)
	assert.Equal(t, 10, len(instructionDataList))
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, &userID, nil, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("User ID: %s", userID)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, &theme, nil,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Theme: %s", theme)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, nil, &statusCode,
		nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Status code: %s", statusCode)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, nil, nil,
		&createStartTime, &createEndTime, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Create time start: %v", createStartTime)
	t.Logf("Create time end: %v", createEndTime)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, nil, nil,
		nil, nil, &updateStartTime, &updateEndTime, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Update time start: %v", updateStartTime)
	t.Logf("Update time end: %v", updateEndTime)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, nil, nil,
		nil, nil, nil, nil, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Query: %s", query)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, &userID, &theme, &statusCode,
		&createStartTime, &createEndTime, &updateStartTime, &updateEndTime, &query,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("User ID: %s", userID)
	t.Logf("Theme: %s", theme)
	t.Logf("Status code: %s", statusCode)
	t.Logf("Create time start: %v", createStartTime)
	t.Logf("Create time end: %v", createEndTime)
	t.Logf("Update time start: %v", updateStartTime)
	t.Logf("Update time end: %v", updateEndTime)
	t.Logf("Query: %s", query)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("Instruction data list: %v", instructionDataList)
	t.Logf("=====================================")
}

func TestCountInstructionData(t *testing.T) {
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		userID             = injector.UserDaoMock.RandomUserID()
		theme              = "THEME1"
		statusCode         = "PENDING"
		createStartTime    = time.Now().AddDate(0, 0, -1)
		createEndTime      = time.Now().AddDate(0, 0, 1)
		updateStartTime    = time.Now().AddDate(0, 0, -1)
		updateEndTime      = time.Now().AddDate(0, 0, 1)
		err                error
	)
	count, err := instructionDataDao.CountInstructionData(
		ctx, nil, nil, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("=====================================")

	groupBy := "theme"
	aggregateCount, err := instructionDataDao.AggregateCountInstructionData(ctx, &groupBy, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, aggregateCount)
	t.Logf("Group by: %s", groupBy)
	t.Logf("Aggregate count: %v", aggregateCount)
	t.Logf("=====================================")

	groupBy = "status.code"
	aggregateCount, err = instructionDataDao.AggregateCountInstructionData(ctx, &groupBy, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, aggregateCount)
	t.Logf("Group by: %s", groupBy)
	t.Logf("Aggregate count: %v", aggregateCount)
	t.Logf("=====================================")

	count, err = instructionDataDao.CountInstructionData(
		ctx, &userID, &theme, &statusCode, &createStartTime, &createEndTime, &updateStartTime, &updateEndTime,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("User ID: %s", userID)
	t.Logf("Theme: %s", theme)
	t.Logf("Status code: %s", statusCode)
	t.Logf("Create time start: %v", createStartTime)
	t.Logf("Create time end: %v", createEndTime)
	t.Logf("Update time start: %v", updateStartTime)
	t.Logf("Update time end: %v", updateEndTime)
	t.Logf("Instruction data count: %d", *count)
	t.Logf("=====================================")
}

func TestUpdateInstructionData(t *testing.T) {
	// t.Skip("Skip TestUpdateInstructionData")
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		userID             = injector.UserDaoMock.RandomUserID()
		instruction        = "InstructionUpdated"
		input              = "InputUpdated"
		output             = "OutputUpdated"
		theme              = "THEME2"
		source             = "SourceUpdated"
		note               = "NoteUpdated"
		statusCode         = "APPROVED"
		statusMsg          = "Approved"
		err                error
	)

	err = instructionDataDao.UpdateInstructionData(
		ctx, instructionDataID, &userID, &instruction, &input, &output, &theme, &source, &note,
		&statusCode, &statusMsg,
	)
	assert.NoError(t, err)

	instructionData, err := instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.NoError(t, err)
	assert.NotNil(t, instructionData)
	assert.Equal(t, instruction, instructionData.Row.Instruction)
	assert.Equal(t, input, instructionData.Row.Input)
	assert.Equal(t, output, instructionData.Row.Output)
	assert.Equal(t, theme, instructionData.Theme)
	assert.Equal(t, source, instructionData.Source)
	assert.Equal(t, note, instructionData.Note)
	assert.Equal(t, statusCode, instructionData.Status.Code)
	assert.Equal(t, statusMsg, instructionData.Status.Message)
}

func TestDeleteInstructionData(t *testing.T) {
	// t.Skip("Skip TestDeleteInstructionData")
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		err                error
	)
	err = instructionDataDao.SoftDeleteInstructionData(ctx, instructionDataID)
	assert.NoError(t, err)

	instructionData, err := instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.Error(t, err)
	assert.Nil(t, instructionData)

	err = instructionDataDao.DeleteInstructionData(ctx, instructionDataID)
	assert.NoError(t, err)

	instructionData, err = instructionDataDao.GetInstructionDataByID(ctx, instructionDataID)
	assert.Error(t, err)
	assert.Nil(t, instructionData)
}

func TestDeleteInstructionDataList(t *testing.T) {
	// t.Skip("Skip TestDeleteInstructionDataList")
	var (
		injector           = wire.GetInjector()
		instructionDataDao = injector.InstructionDataDao
		ctx                = injector.Ctx
		theme              = "THEME1"
		statusCode         = "PENDING"
		err                error
	)

	count, err := instructionDataDao.SoftDeleteInstructionDataList(
		ctx, nil, &theme, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.NotEmpty(t, *count)
	t.Logf("Theme: %s", theme)
	t.Logf("Deleted instruction data count: %d", *count)
	t.Logf("=====================================")

	instructionDataList, count, err := instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, &theme, nil,
		nil, nil, nil, nil, nil,
	)
	assert.Empty(t, instructionDataList)

	count, err = instructionDataDao.SoftDeleteInstructionDataList(
		ctx, nil, nil, &statusCode, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Status code: %s", statusCode)
	t.Logf("Deleted instruction data count: %d", *count)
	t.Logf("=====================================")

	instructionDataList, count, err = instructionDataDao.GetInstructionDataList(
		ctx, 0, 10, false, nil, nil, &statusCode,
		nil, nil, nil, nil, nil,
	)
	assert.Empty(t, instructionDataList)
}
