package grpc

import (
	"context"

	"goload/internal/generated/grpc/go_load"
	"goload/internal/logic"
	"google.golang.org/grpc"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
	accountLogic logic.Account 
}

func NewHandler(
	accountLogic logic.Account,
) go_load.GoLoadServiceServer {
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
	return &go_load.CreateAccountResponse{
		AccountId: output.ID,
	}, nil
}

func (h *Handler) CreateSession(ctx context.Context, req *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	// TODO: implement
	return &go_load.CreateSessionResponse{}, nil
}

func (h *Handler) CreateDownloadTask(ctx context.Context, req *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	// TODO: implement
	return &go_load.CreateDownloadTaskResponse{}, nil
}

func (h *Handler) GetDownloadTaskList(ctx context.Context, req *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	// TODO: implement
	return &go_load.GetDownloadTaskListResponse{}, nil
}

func (h *Handler) UpdateDownloadTask(ctx context.Context, req *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	// TODO: implement
	return &go_load.UpdateDownloadTaskResponse{}, nil
}

func (h *Handler) DeleteDownloadTask(ctx context.Context, req *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	// TODO: implement
	return &go_load.DeleteDownloadTaskResponse{}, nil
}

func (h *Handler) GetDownloadTaskFile(req *go_load.GetDownloadTaskFileRequest, stream grpc.ServerStreamingServer[go_load.GetDownloadTaskFileResponse]) error {
	// TODO: implement streaming
	return nil
}