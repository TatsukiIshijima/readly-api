package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"readly/controller"
	sqlc "readly/db/sqlc"
	"readly/env"
	"readly/middleware"
	"readly/pb"
	"readly/repository"
	"readly/router"
	"readly/server"
	"readly/service/auth"
	"readly/usecase"
)

func main() {
	config, err := env.Load(filepath.Join(env.ProjectRoot(), "/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	a := sqlc.Adapter{}
	db, q := a.Connect(config.DBDriver, config.DBSource)
	t := repository.New(db)

	bookRepo := repository.NewBookRepository(q)
	userRepo := repository.NewUserRepository(q)
	readingHistoryRepo := repository.NewReadingHistoryRepository(q)
	sessionRepo := repository.NewSessionRepository(q)

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	registerBookUseCase := usecase.NewRegisterBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	deleteBookUseCase := usecase.NewDeleteBookUseCase(t, bookRepo, readingHistoryRepo, userRepo)
	signUpUseCase := usecase.NewSignUpUseCase(config, maker, t, sessionRepo, userRepo)
	signInUseCase := usecase.NewSignInUseCase(config, maker, t, sessionRepo, userRepo)
	refreshTokenUseCase := usecase.NewRefreshAccessTokenUseCase(config, maker, sessionRepo)

	//runGinServer(
	//	config,
	//	maker,
	//	registerBookUseCase,
	//	deleteBookUseCase,
	//	signUpUseCase,
	//	signInUseCase,
	//	refreshTokenUseCase,
	//)

	runGRPCServer(
		config,
		maker,
		registerBookUseCase,
		deleteBookUseCase,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)

	//go runGatewayServer()
}

func runGinServer(
	config env.Config,
	maker auth.TokenMaker,
	registerBookUseCase usecase.RegisterBookUseCase,
	deleteBookUseCase usecase.DeleteBookUseCase,
	signUpUseCase usecase.SignUpUseCase,
	signInUseCase usecase.SignInUseCase,
	refreshTokenUseCase usecase.RefreshAccessTokenUseCase,
) {
	bookController := controller.NewBookController(registerBookUseCase, deleteBookUseCase)
	userController := controller.NewUserController(config, maker, signUpUseCase, signInUseCase, refreshTokenUseCase)

	r := router.Setup(middleware.Authorize(maker), bookController, userController)
	err := r.Run(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runGRPCServer(
	config env.Config,
	maker auth.TokenMaker,
	registerBookUseCase usecase.RegisterBookUseCase,
	deleteBookUseCase usecase.DeleteBookUseCase,
	signUpUseCase usecase.SignUpUseCase,
	signInUseCase usecase.SignInUseCase,
	refreshTokenUseCase usecase.RefreshAccessTokenUseCase,
) {
	grpcServer := grpc.NewServer()

	userServer := server.NewUserServer(
		config,
		maker,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)
	bookServer := server.NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
	)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	pb.RegisterBookServiceServer(grpcServer, bookServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}

// 43~46, 49, 53まではみてみる
// 46はgrpcに移行したためIPアドレスやUserAgentの取得方法っぽい
// 49のバリデーションはginからgrpcへ移行したので独自実装になるっぽい
// grpcの移行は徐々にできなそうなのでGinでのコードは残しつつやるようだ
// 53updateUserAPIのための認可。metadataのセクションが関連している
// これも独自実装になるっぽい。HTTPとGRPCどちらも対応するため。

// コース45を参考
func runGatewayServer(
	config env.Config,
	maker auth.TokenMaker,
	registerBookUseCase usecase.RegisterBookUseCase,
	deleteBookUseCase usecase.DeleteBookUseCase,
	signUpUseCase usecase.SignUpUseCase,
	signInUseCase usecase.SignInUseCase,
	refreshTokenUseCase usecase.RefreshAccessTokenUseCase,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userServer := server.NewUserServer(
		config,
		maker,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)
	bookServer := server.NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
	)

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	// HTTPリクエストをgRPCのリクエストに変換する
	grpcMux := runtime.NewServeMux(jsonOption)

	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userServer)
	if err != nil {
		log.Fatalf("cannot register handle server: %v", err)
	}
	err = pb.RegisterBookServiceHandlerServer(ctx, grpcMux, bookServer)
	if err != nil {
		log.Fatalf("cannot register handle server: %v", err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, httpMux)
	if err != nil {
		log.Fatalf("cannot start HTTP gateway server: %v", err)
	}
}
