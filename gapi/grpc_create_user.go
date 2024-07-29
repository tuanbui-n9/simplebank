package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	db "github.com/tuanbui-n9/simplebank/db/sqlc"
	"github.com/tuanbui-n9/simplebank/pb"
	"github.com/tuanbui-n9/simplebank/util"
	"github.com/tuanbui-n9/simplebank/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if len(violations) > 0 {
		invalidArgumentError(violations)
	}
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	// args := db.CreateUserTxParams{
	// 	CreateUserParams: db.CreateUserParams{
	// 		Username:       req.GetUsername(),
	// 		HashedPassword: hashedPassword,
	// 		FullName:       req.GetFullName(),
	// 		Email:          req.GetEmail(),
	// 	},
	// 	AfterCreate: func(user db.User) error {
	// 		// send verification email
	// 		taskPayload := &worker.PayloadSendVerifyEmail{
	// 			Username: user.Username,
	// 		}
	// 		opts := []asynq.Option{
	// 			asynq.MaxRetry(10),
	// 			asynq.ProcessIn(10 * time.Second),
	// 			asynq.Queue(worker.QueueCritical),
	// 		}
	// 		return server.taskDistributor.DisTributeTasSendVerifyEmail(ctx, taskPayload, opts...)
	// 	},
	// }
	args := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}
	startTime := time.Now()

	user, err := server.store.CreateUser(ctx, args)
	elapsedTime := time.Since(startTime)
	log.Info().Msgf("CreateUser took %s", elapsedTime)

	// txResult, err := server.store.CreateUserTx(ctx, args)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := validator.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}
