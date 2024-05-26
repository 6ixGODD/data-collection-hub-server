package dao_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertNotice(t *testing.T) {
	// t.Skip("Skip TestInsertNotice")
	title := "Title"
	content := "Content"
	noticeType := "NORMAL"

	noticeID, err = noticeDao.InsertNotice(noticeDaoCtx, title, content, noticeType)
	assert.NoError(t, err)
	assert.NotEmpty(t, noticeID)

	notice, err := noticeDao.GetNoticeByID(noticeDaoCtx, noticeID)
	assert.NoError(t, err)
	assert.NotNil(t, notice)
	assert.Equal(t, title, notice.Title)
	assert.Equal(t, content, notice.Content)
	assert.Equal(t, noticeType, notice.NoticeType)
}

func TestGetNotice(t *testing.T) {
	// t.Skip("Skip TestGetNotice")
	notice, err := noticeDao.GetNoticeByID(noticeDaoCtx, noticeID)
	assert.NoError(t, err)
	assert.NotNil(t, notice)
	assert.NotEmpty(t, notice.NoticeID)
	assert.NotEmpty(t, notice.Title)
	assert.NotEmpty(t, notice.Content)
	assert.NotEmpty(t, notice.NoticeType)
	assert.NotEmpty(t, notice.CreatedAt)
	assert.NotEmpty(t, notice.UpdatedAt)
}

func TestGetNoticeList(t *testing.T) {
	// t.Skip("Skip TestGetNoticeList")
	var (
		createStartTime = time.Now().Add(-time.Hour)
		createEndTime   = time.Now().Add(time.Hour)
		updateStartTime = time.Now().Add(-time.Hour)
		updateEndTime   = time.Now().Add(time.Hour)
		noticeType      = "NORMAL"
	)

	noticeList, count, err := noticeDao.GetNoticeList(
		noticeDaoCtx, 0, 10, false, nil, nil, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	assert.NotEmpty(t, noticeList)
	t.Logf("Notice Count: %d", *count)
	t.Logf("Notice List: %v", noticeList)
	t.Logf("=====================================")

	noticeList, count, err = noticeDao.GetNoticeList(
		noticeDaoCtx, 0, 10, false, &createStartTime, &createEndTime, nil, nil, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Create Start Time: %s", createStartTime)
	t.Logf("Create End Time: %s", createEndTime)
	t.Logf("Notice Count: %d", *count)
	t.Logf("Notice List: %v", noticeList)
	t.Logf("=====================================")

	noticeList, count, err = noticeDao.GetNoticeList(
		noticeDaoCtx, 0, 10, false, nil, nil, &updateStartTime, &updateEndTime, nil,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Update Start Time: %s", updateStartTime)
	t.Logf("Update End Time: %s", updateEndTime)
	t.Logf("Notice Count: %d", *count)
	t.Logf("Notice List: %v", noticeList)
	t.Logf("=====================================")

	noticeList, count, err = noticeDao.GetNoticeList(
		noticeDaoCtx, 0, 10, false, nil, nil, nil, nil, &noticeType,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Notice Type: %s", noticeType)
	t.Logf("Notice Count: %d", *count)
	t.Logf("Notice List: %v", noticeList)
	t.Logf("=====================================")

	noticeList, count, err = noticeDao.GetNoticeList(
		noticeDaoCtx, 0, 10, false, &createStartTime, &createEndTime, &updateStartTime, &updateEndTime, &noticeType,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, count)
	t.Logf("Create Start Time: %s", createStartTime)
	t.Logf("Create End Time: %s", createEndTime)
	t.Logf("Update Start Time: %s", updateStartTime)
	t.Logf("Update End Time: %s", updateEndTime)
	t.Logf("Notice Type: %s", noticeType)
	t.Logf("Notice Count: %d", *count)
	t.Logf("Notice List: %v", noticeList)
	t.Logf("=====================================")
}

func TestUpdateNotice(t *testing.T) {
	// t.Skip("Skip TestUpdateNotice")
	title := "New Title"
	content := "New Content"
	noticeType := "NORMAL"

	err := noticeDao.UpdateNotice(noticeDaoCtx, noticeID, &title, &content, &noticeType)
	assert.NoError(t, err)

	notice, err := noticeDao.GetNoticeByID(noticeDaoCtx, noticeID)
	assert.NoError(t, err)
	assert.NotNil(t, notice)
	assert.Equal(t, title, notice.Title)
	assert.Equal(t, content, notice.Content)
	assert.Equal(t, noticeType, notice.NoticeType)
}

func TestDeleteNotice(t *testing.T) {
	// t.Skip("Skip TestDeleteNotice")
	err := noticeDao.DeleteNotice(noticeDaoCtx, noticeID)
	assert.NoError(t, err)

	notice, err := noticeDao.GetNoticeByID(noticeDaoCtx, noticeID)
	assert.Error(t, err)
	assert.Nil(t, notice)
}

func TestDeleteNoticeList(t *testing.T) {
	// t.Skip("Skip TestDeleteNoticeList")
	noticeType := "NORMAL"
	count, err := noticeDao.DeleteNoticeList(noticeDaoCtx, nil, nil, nil, nil, &noticeType)
	assert.NoError(t, err)
	assert.NotNil(t, count)
	t.Logf("Notice Type: %s", noticeType)
	t.Logf("Delete Count: %d", *count)
	t.Logf("=====================================")

	noticeList, count, err := noticeDao.GetNoticeList(
		noticeDaoCtx, 0, 10, false, nil, nil, nil, nil, &noticeType,
	)
	assert.Empty(t, noticeList)
}
