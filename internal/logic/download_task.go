package logic

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/doug-martin/goqu/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goload/internal/configs"
	"goload/internal/dataaccess/database"
	goload "goload/internal/generated/goload"
	
)


type CreateDownloadTaskParams struct {
	Token        string
	DownloadType goload.DownloadType
	URL          string
}

type CreateDownloadTaskOutput struct {
	DownloadTask *goload.DownloadTask
}

type GetDownloadTaskListParams struct {
	Token  string
	Offset uint64
	Limit  uint64
}

type GetDownloadTaskListOutput struct {
	DownloadTaskList       []*goload.DownloadTask
	TotalDownloadTaskCount uint64
}

type UpdateDownloadTaskParams struct {
	Token          string
	DownloadTaskID uint64
	URL            string
}

type UpdateDownloadTaskOutput struct {
	DownloadTask *goload.DownloadTask
}

type DeleteDownloadTaskParams struct {
	Token          string
	DownloadTaskID uint64
}

type GetDownloadTaskFileParams struct {
	Token          string
	DownloadTaskID uint64
}

type DownloadTask interface {
	CreateDownloadTask(context.Context, CreateDownloadTaskParams) (CreateDownloadTaskOutput, error)
	GetDownloadTaskList(context.Context, GetDownloadTaskListParams) (GetDownloadTaskListOutput, error)
	UpdateDownloadTask(context.Context, UpdateDownloadTaskParams) (UpdateDownloadTaskOutput, error)
	DeleteDownloadTask(context.Context, DeleteDownloadTaskParams) error
	
}

type downloadTask struct {
	tokenLogic                  Token
	downloadTaskDataAccessor    database.DownloadTaskDataAccessor
	goquDatabase                *goqu.Database
	logger                      *zap.Logger

}

func NewDownloadTask(
	tokenLogic Token,
	downloadTaskDataAccessor database.DownloadTaskDataAccessor,
	goquDatabase *goqu.Database,
	logger *zap.Logger,
) DownloadTask {
	return &downloadTask{
		tokenLogic:                  tokenLogic,
		downloadTaskDataAccessor:    downloadTaskDataAccessor,
		goquDatabase:                goquDatabase,
		logger:                      logger,
	}
}

func (c downloadTask) CreateDownloadTask(ctx context.Context, params CreateDownloadTaskParams) (CreateDownloadTaskOutput, error) {

}

func (c downloadTask) GetDownloadTaskList(ctx context.Context, params GetDownloadTaskListParams) (GetDownloadTaskListOutput, error) {

}

func (c downloadTask) UpdateDownloadTask(ctx context.Context, params UpdateDownloadTaskParams) (UpdateDownloadTaskOutput, error) {

}

func (c downloadTask) DeleteDownloadTask(ctx context.Context, params DeleteDownloadTaskParams) error {	

}	
