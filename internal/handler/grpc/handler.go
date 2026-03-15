package grpc

import (
	"context"

	"goload/internal/generated/grpc/goload"
	"goload/internal/logic"
	"google.golang.org/grpc"
)

type Handler struct {
	goload.UnimplementedGoLoadServiceServer
	accountLogic logic.Account 
	downloadTaskLogic logic.DownloadTask
}

func NewHandler(
	accountLogic logic.Account,
) goload.GoLoadServiceServer {
	return &Handler{
		accountLogic: accountLogic,
	}
}

func (a Handler) CreateAccount(ctx context.Context, request *go_load.CreateAccountRequest) (*go_load.CreateAccountResponse, error) {
	// TODO: implement
	output, err := a.accountLogic.CreateAccount(ctx, logic.CreateAccountParams{
		AccountName: request.GetAccountName(),
		Password:    request.GetPassword(),
	}) 
	if err != nil{
		return nil, err
	}
	return &goload.CreateAccountResponse{
		AccountId: output.ID,
	}, nil
}

func (h *Handler) CreateSession(ctx context.Context, req *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	// TODO: implement
	return &goload.CreateSessionResponse{}, nil
}

func (h *Handler) CreateDownloadTask(ctx context.Context, req *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	// TODO: implement
	return &goload.CreateDownloadTaskResponse{}, nil
}

func (h *Handler) GetDownloadTaskList(ctx context.Context, req *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	// TODO: implement
	return &goload.GetDownloadTaskListResponse{}, nil
}

func (h *Handler) UpdateDownloadTask(ctx context.Context, req *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	// TODO: implement
	return &goload.UpdateDownloadTaskResponse{}, nil
}

func (h *Handler) DeleteDownloadTask(ctx context.Context, req *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	// TODO: implement
	return &goload.DeleteDownloadTaskResponse{}, nil
}

func (h *Handler) GetDownloadTaskFile(req *go_load.GetDownloadTaskFileRequest, stream grpc.ServerStreamingServer[go_load.GetDownloadTaskFileResponse]) error {
	// TODO: implement streaming
	return nil
}