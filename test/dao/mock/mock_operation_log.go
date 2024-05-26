package mock

import (
	"context"
	"math/rand"

	"data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OperationLogDaoMock struct {
	OperationLogMap     map[primitive.ObjectID]*entity.OperationLogModel
	OperationLogIDs     []primitive.ObjectID
	OperationLogDao     mods.OperationLogDao
	UserMock            UserDaoMock
	InstructionDataMock InstructionDataDaoMock
	NoticeMock          NoticeDaoMock
	DocumentationMock   DocumentationDaoMock
}

func NewOperationLogDaoMock(
	operationLogDao mods.OperationLogDao, userMock UserDaoMock, instructionDataMock InstructionDataDaoMock,
	noticeMock NoticeDaoMock, documentationMock DocumentationDaoMock,
) *OperationLogDaoMock {
	return &OperationLogDaoMock{
		OperationLogMap:     make(map[primitive.ObjectID]*entity.OperationLogModel),
		OperationLogDao:     operationLogDao,
		UserMock:            userMock,
		InstructionDataMock: instructionDataMock,
		NoticeMock:          noticeMock,
		DocumentationMock:   documentationMock,
	}
}

func NewOperationLogDaoMockWithRandomData(
	n int, operationLogDao mods.OperationLogDao, userMock UserDaoMock, instructionDataMock InstructionDataDaoMock,
	noticeMock NoticeDaoMock, documentationMock DocumentationDaoMock,
) *OperationLogDaoMock {
	operationLogDaoMock := NewOperationLogDaoMock(
		operationLogDao, userMock, instructionDataMock, noticeMock, documentationMock,
	)
	for i := 0; i < n; i++ {
		operationLog := operationLogDaoMock.GenerateOperationLogModel()
		operationLogDaoMock.OperationLogMap[operationLog.OperationLogID] = operationLog
		operationLogDaoMock.OperationLogIDs = append(operationLogDaoMock.OperationLogIDs, operationLog.OperationLogID)
	}
	return operationLogDaoMock
}

func (m *OperationLogDaoMock) GenerateOperationLogWithUserID(userID primitive.ObjectID) *entity.OperationLogModel {
	ipAddress, userAgent, operation, description, status := GenerateOperationLog()
	entityID, entityType := RandomEntity(m.InstructionDataMock, m.NoticeMock, m.DocumentationMock, m.UserMock)
	operationLogID, err := m.OperationLogDao.InsertOperationLog(
		context.Background(), userID, entityID, ipAddress, userAgent, operation, entityType, description, status,
	)
	if err != nil {
		panic(err)
	}

	operationLog, err := m.OperationLogDao.GetOperationLogByID(context.Background(), operationLogID)
	if err != nil {
		panic(err)
	}
	return operationLog
}

func (m *OperationLogDaoMock) GenerateOperationLogWithIpAddress(ipAddress string) *entity.OperationLogModel {
	userID := m.UserMock.RandomUserID()
	_, userAgent, operation, description, status := GenerateOperationLog()
	entityID, entityType := RandomEntity(m.InstructionDataMock, m.NoticeMock, m.DocumentationMock, m.UserMock)
	operationLogID, err := m.OperationLogDao.InsertOperationLog(
		context.Background(), userID, entityID, ipAddress, userAgent, operation, entityType, description, status,
	)
	if err != nil {
		panic(err)
	}

	operationLog, err := m.OperationLogDao.GetOperationLogByID(context.Background(), operationLogID)
	if err != nil {
		panic(err)
	}
	return operationLog
}

func (m *OperationLogDaoMock) GenerateOperationLogWithEntityID(entityID primitive.ObjectID) *entity.OperationLogModel {
	userID := m.UserMock.RandomUserID()
	ipAddress, userAgent, operation, description, status := GenerateOperationLog()
	entityType := randomEnum([]string{"INSTRUCTION", "NOTICE", "DOCUMENTATION", "USER"})
	operationLogID, err := m.OperationLogDao.InsertOperationLog(
		context.Background(), userID, entityID, ipAddress, userAgent, operation, entityType, description, status,
	)
	if err != nil {
		panic(err)
	}

	operationLog, err := m.OperationLogDao.GetOperationLogByID(context.Background(), operationLogID)
	if err != nil {
		panic(err)
	}
	return operationLog
}

func (m *OperationLogDaoMock) Create(operationLog *entity.OperationLogModel) {
	m.OperationLogMap[operationLog.OperationLogID] = operationLog
}

func (m *OperationLogDaoMock) Get(operationLogID primitive.ObjectID) (*entity.OperationLogModel, error) {
	operationLog, ok := m.OperationLogMap[operationLogID]
	if !ok {
		return nil, nil
	}
	return operationLog, nil
}

func (m *OperationLogDaoMock) RandomOperationLogID() primitive.ObjectID {
	return m.OperationLogIDs[rand.Intn(len(m.OperationLogIDs))]
}

func (m *OperationLogDaoMock) GenerateOperationLogModel() *entity.OperationLogModel {
	userID := m.UserMock.RandomUserID()
	ipAddress, userAgent, operation, description, status := GenerateOperationLog()
	entityID, entityType := RandomEntity(m.InstructionDataMock, m.NoticeMock, m.DocumentationMock, m.UserMock)
	operationLogID, err := m.OperationLogDao.InsertOperationLog(
		context.Background(), userID, entityID, ipAddress, userAgent, operation, entityType, description, status,
	)
	if err != nil {
		panic(err)
	}

	operationLog, err := m.OperationLogDao.GetOperationLogByID(context.Background(), operationLogID)
	if err != nil {
		panic(err)
	}
	return operationLog
}

func (m *OperationLogDaoMock) Delete() {
	for _, operationLogID := range m.OperationLogIDs {
		_ = m.OperationLogDao.DeleteOperationLog(context.Background(), operationLogID)
	}
}

// GenerateOperationLog generate an operation log and return ip address, user agent, operation, description and status
func GenerateOperationLog() (string, string, string, string, string) {
	return randomIp(), randomEnum([]string{"Chrome", "Firefox", "Safari", "Edge"}), randomEnum(
		[]string{
			"CREATE", "UPDATE", "DELETE",
		},
	), randomString(20), randomEnum([]string{"SUCCESS", "FAILURE"})
}

// RandomEntity return an entity id and entity type randomly
func RandomEntity(
	instructionDataMock InstructionDataDaoMock, noticeMock NoticeDaoMock, documentationMock DocumentationDaoMock,
	userMock UserDaoMock,
) (primitive.ObjectID, string) {
	instructionDataID := instructionDataMock.RandomInstructionDataID()
	noticeID := noticeMock.RandomNoticeID()
	documentID := documentationMock.RandomDocumentationID()
	userID := userMock.RandomUserID()
	entityType := randomEnum([]string{"INSTRUCTION", "NOTICE", "DOCUMENTATION", "USER"})
	switch entityType {
	case "INSTRUCTION":
		return instructionDataID, entityType
	case "NOTICE":
		return noticeID, entityType
	case "DOCUMENTATION":
		return documentID, entityType
	case "USER":
		return userID, entityType
	default:
		return primitive.NilObjectID, ""
	}
}
