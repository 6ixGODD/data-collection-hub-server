package mods

import (
	"context"
	"time"

	dao "data-collection-hub-server/internal/pkg/dao/mods"
	"data-collection-hub-server/internal/pkg/domain/vo/common"
	"data-collection-hub-server/internal/pkg/service"
	"data-collection-hub-server/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoticeService interface {
	GetNotice(ctx context.Context, noticeID *primitive.ObjectID) (*common.GetNoticeResponse, error)
	GetNoticeList(
		ctx context.Context, page, pageSize *int64, noticeType *string, updateBefore, updateAfter *time.Time,
	) (*common.GetNoticeListResponse, error)
}

type noticeServiceImpl struct {
	core      *service.Core
	noticeDao dao.NoticeDao
}

func NewNoticeService(core *service.Core, noticeDao dao.NoticeDao) NoticeService {
	return &noticeServiceImpl{
		core:      core,
		noticeDao: noticeDao,
	}
}

func (n noticeServiceImpl) GetNotice(ctx context.Context, noticeID *primitive.ObjectID) (
	*common.GetNoticeResponse, error,
) {
	notice, err := n.noticeDao.GetNoticeByID(ctx, *noticeID)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	return &common.GetNoticeResponse{
		NoticeID:   notice.NoticeID.Hex(),
		Title:      notice.Title,
		Content:    notice.Content,
		NoticeType: notice.NoticeType,
		CreatedAt:  notice.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  notice.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (n noticeServiceImpl) GetNoticeList(
	ctx context.Context, page, pageSize *int64, noticeType *string, updateBefore, updateAfter *time.Time,
) (*common.GetNoticeListResponse, error) {
	offset := (*page - 1) * *pageSize
	notices, count, err := n.noticeDao.GetNoticeList(
		ctx, offset, *pageSize, false, nil, nil, updateBefore, updateAfter, noticeType,
	)
	if err != nil {
		return nil, errors.DBError(errors.ReadError(err))
	}
	resp := make([]*common.NoticeSummary, 0, len(notices))
	for _, notice := range notices {
		resp = append(
			resp, &common.NoticeSummary{
				NoticeID:   notice.NoticeID.Hex(),
				Title:      notice.Title,
				NoticeType: notice.NoticeType,
				CreatedAt:  notice.CreatedAt.Format(time.RFC3339),
			},
		)
	}
	return &common.GetNoticeListResponse{
		Total:             *count,
		NoticeSummaryList: resp,
	}, nil
}
