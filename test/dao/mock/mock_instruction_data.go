package mock

import (
	"context"
	"math/rand"

	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InstructionDataDaoMock struct {
	InstructionDataMap map[primitive.ObjectID]*entity.InstructionDataModel
	InstructionDataIDs []primitive.ObjectID
	UserMock           *UserDaoMock
	InstructionDataDao mods.InstructionDataDao
}

func NewInstructionDataDaoMock(
	userMock *UserDaoMock, instructionDataDao mods.InstructionDataDao,
) *InstructionDataDaoMock {
	return &InstructionDataDaoMock{
		InstructionDataMap: make(map[primitive.ObjectID]*entity.InstructionDataModel),
		UserMock:           userMock,
		InstructionDataDao: instructionDataDao,
	}
}

func NewInstructionDataDaoMockWithRandomData(
	n int, userMock *UserDaoMock, instructionDataDao mods.InstructionDataDao,
) *InstructionDataDaoMock {
	instructionDataDaoMock := NewInstructionDataDaoMock(userMock, instructionDataDao)
	for i := 0; i < n; i++ {
		instructionData := instructionDataDaoMock.GenerateInstructionDataModel()
		instructionDataDaoMock.InstructionDataMap[instructionData.InstructionDataID] = instructionData
		instructionDataDaoMock.InstructionDataIDs = append(
			instructionDataDaoMock.InstructionDataIDs, instructionData.InstructionDataID,
		)
	}
	return instructionDataDaoMock
}

func (m *InstructionDataDaoMock) Create(instructionData *entity.InstructionDataModel) error {
	m.InstructionDataMap[instructionData.InstructionDataID] = instructionData
	return nil
}

func (m *InstructionDataDaoMock) Get(instructionDataID primitive.ObjectID) (*entity.InstructionDataModel, error) {
	instructionData, ok := m.InstructionDataMap[instructionDataID]
	if !ok {
		return nil, nil
	}
	return instructionData, nil
}

func (m *InstructionDataDaoMock) RandomInstructionDataID() primitive.ObjectID {
	return m.InstructionDataIDs[rand.Intn(len(m.InstructionDataIDs))]
}

func (m *InstructionDataDaoMock) GenerateInstructionDataModel() *entity.InstructionDataModel {
	userID := m.UserMock.RandomUserID()
	username := m.UserMock.UserMap[userID].Username
	instruction, input, output, theme, source, note, statusCode, statusMessage := randomInstructionData()
	instructionDataID, err := m.InstructionDataDao.InsertInstructionData(
		context.Background(), userID, username, instruction, input, output, theme, source, note, statusCode,
		statusMessage,
	)
	if err != nil {
		panic(err)
	}
	instructionData, err := m.InstructionDataDao.GetInstructionDataByID(context.Background(), instructionDataID)
	if err != nil {
		panic(err)
	}
	return instructionData
}

func (m *InstructionDataDaoMock) Delete() {
	for _, instructionDataID := range m.InstructionDataIDs {
		_ = m.InstructionDataDao.DeleteInstructionData(context.Background(), instructionDataID)
	}
}

// GenerateInstructionData generates a new instruction data, returns instruction, input, output, theme, source, note, status code, status message
func randomInstructionData() (string, string, string, string, string, string, string, string) {
	return randomString(10), randomString(10), randomString(10), randomEnum(
			[]string{
				"THEME1", "THEME2", "THEME3",
			},
		), "https://" + randomString(10) + ".com", randomString(10), randomEnum(
			[]string{
				"PENDING", "APPROVED", "REJECTED",
			},
		), randomString(10)
}
