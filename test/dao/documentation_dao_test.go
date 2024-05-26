package dao_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertDocumentation(t *testing.T) {
	// t.Skip("Skip TestInsertDocumentation")
	title := "Title"
	content := "Content"

	documentID, err = documentationDao.InsertDocumentation(documentationDaoCtx, title, content)
	assert.NoError(t, err)
	assert.NotEmpty(t, documentID)

	documentation, err := documentationDao.GetDocumentationByID(documentationDaoCtx, documentID)
	assert.NoError(t, err)
	assert.NotNil(t, documentation)
	assert.Equal(t, title, documentation.Title)
	assert.Equal(t, content, documentation.Content)
}

func TestGetDocumentation(t *testing.T) {
	// t.Skip("Skip TestGetDocumentation")
	documentation, err := documentationDao.GetDocumentationByID(documentationDaoCtx, documentID)
	assert.NoError(t, err)
	assert.NotNil(t, documentation)
	assert.NotEmpty(t, documentation.DocumentID)
	assert.NotEmpty(t, documentation.Title)
	assert.NotEmpty(t, documentation.Content)
	assert.NotEmpty(t, documentation.CreatedAt)
	assert.NotEmpty(t, documentation.UpdatedAt)
}

func TestGetDocumentationList(t *testing.T) {
	// t.Skip("Skip TestGetDocumentationList")
	var (
		createStartTime = time.Now().Add(-time.Hour)
		createEndTime   = time.Now().Add(time.Hour)
		updateStartTime = time.Now().Add(-time.Hour)
		updateEndTime   = time.Now().Add(time.Hour)
	)
	documentationList, count, err := documentationDao.GetDocumentationList(
		documentationDaoCtx, 0, 10, false, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, documentationList)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("Documentation List: %v", documentationList)
	t.Logf("=====================================")

	documentationList, count, err = documentationDao.GetDocumentationList(
		documentationDaoCtx, 0, 10, false, &createStartTime, &createEndTime, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Create Start Time: %v", createStartTime)
	t.Logf("Create End Time: %v", createEndTime)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("Documentation List: %v", documentationList)
	t.Logf("=====================================")

	documentationList, count, err = documentationDao.GetDocumentationList(
		documentationDaoCtx, 0, 10, false, nil, nil, &updateStartTime, &updateEndTime,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Update Start Time: %v", updateStartTime)
	t.Logf("Update End Time: %v", updateEndTime)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("Documentation List: %v", documentationList)
	t.Logf("=====================================")

	documentationList, count, err = documentationDao.GetDocumentationList(
		documentationDaoCtx, 0, 10, false, &createStartTime, &createEndTime, &updateStartTime, &updateEndTime,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Create Start Time: %v", createStartTime)
	t.Logf("Create End Time: %v", createEndTime)
	t.Logf("Update Start Time: %v", updateStartTime)
	t.Logf("Update End Time: %v", updateEndTime)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("Documentation List: %v", documentationList)
	t.Logf("=====================================")
}

func TestUpdateDocumentation(t *testing.T) {
	// t.Skip("Skip TestUpdateDocumentation")
	title := "New Title"
	content := "New Content"

	err := documentationDao.UpdateDocumentation(documentationDaoCtx, documentID, &title, &content)
	assert.NoError(t, err)

	documentation, err := documentationDao.GetDocumentationByID(documentationDaoCtx, documentID)
	assert.NoError(t, err)
	assert.NotNil(t, documentation)
	assert.Equal(t, title, documentation.Title)
	assert.Equal(t, content, documentation.Content)
}

func TestDeleteDocumentation(t *testing.T) {
	// t.Skip("Skip TestDeleteDocumentation")
	err := documentationDao.DeleteDocumentation(documentationDaoCtx, documentID)
	assert.NoError(t, err)

	documentation, err := documentationDao.GetDocumentationByID(documentationDaoCtx, documentID)
	assert.Error(t, err)
	assert.Nil(t, documentation)
}

func TestDeleteDocumentationList(t *testing.T) {
	// t.Skip("Skip TestDeleteDocumentationList")
	var (
		createStartTime = time.Now().Add(-time.Hour)
		createEndTime   = time.Now().Add(time.Hour)
		updateStartTime = time.Now().Add(-time.Hour)
		updateEndTime   = time.Now().Add(time.Hour)
	)

	documentationList, count, err := documentationDao.GetDocumentationList(
		documentationDaoCtx, 0, 10, false, &createStartTime, &createEndTime, &updateStartTime, &updateEndTime,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, documentationList)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("Documentation List: %v", documentationList)
	t.Logf("=====================================")

	count, err = documentationDao.DeleteDocumentationList(
		documentationDaoCtx, &createStartTime, &createEndTime, &updateStartTime, &updateEndTime,
	)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("=====================================")

	documentationList, count, err = documentationDao.GetDocumentationList(
		documentationDaoCtx, 0, 10, false, &createStartTime, &createEndTime, &updateStartTime, &updateEndTime,
	)
	assert.NoError(t, err)
	assert.Empty(t, documentationList)
	t.Logf("Documentation Count: %d", *count)
	t.Logf("Documentation List: %v", documentationList)
	t.Logf("=====================================")
}
