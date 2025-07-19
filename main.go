package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"readly/configs"
	sqlc "readly/db/sqlc"
	"readly/db/transaction"
	bookDomain "readly/feature/book/domain"
	bookRepo "readly/feature/book/repository"
	bookServer "readly/feature/book/server"
	bookUseCase "readly/feature/book/usecase"
	imageRepo "readly/feature/image/repository"
	imageServer "readly/feature/image/server"
	imageUseCase "readly/feature/image/usecase"
	sessionRepo "readly/feature/user/repository"
	userRepo "readly/feature/user/repository"
	userServer "readly/feature/user/server"
	userUseCase "readly/feature/user/usecase"
	"readly/middleware/auth"
	"readly/middleware/image"
	"readly/pb/readly/v1"
	router "readly/router"
)

func main() {
	config, err := configs.Load(filepath.Join(configs.ProjectRoot(), "/configs/env"))
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	a := sqlc.Adapter{}
	db, q := a.Connect(config.DBDriver, config.DBSource)
	t := transaction.New(db)

	bookRepository := bookRepo.NewBookRepository(q)
	userRepository := userRepo.NewUserRepository(q)
	readingHistoryRepository := bookRepo.NewReadingHistoryRepository(q)
	sessionRepository := sessionRepo.NewSessionRepository(q)
	imgRepo := imageRepo.NewImageRepository()

	maker, err := auth.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	createGenresUseCase := bookUseCase.NewCreateGenresUseCase(t, bookRepository)
	registerBookUseCase := bookUseCase.NewRegisterBookUseCase(t, bookRepository, readingHistoryRepository)
	deleteBookUseCase := bookUseCase.NewDeleteBookUseCase(t, bookRepository, readingHistoryRepository)
	getBookUseCase := bookUseCase.NewGetBookUseCase(bookRepository, readingHistoryRepository)
	getBookListUseCase := bookUseCase.NewGetBookListUseCase(readingHistoryRepository)
	signUpUseCase := userUseCase.NewSignUpUseCase(config, maker, t, sessionRepository, userRepository)
	signInUseCase := userUseCase.NewSignInUseCase(config, maker, t, sessionRepository, userRepository)
	refreshTokenUseCase := userUseCase.NewRefreshAccessTokenUseCase(config, maker, sessionRepository)
	uploadImgUseCase := imageUseCase.NewUploadImgUseCase(config, imgRepo)

	// Register genres at application startup
	registerGenres(createGenresUseCase)

	// メインルーチンでgRPC Serverの起動しているとそこでブロックしてしまい、
	//HTTP Gatewayの起動ができないため、別のルーチンで起動する
	go runGatewayServer(
		config,
		maker,
		registerBookUseCase,
		deleteBookUseCase,
		getBookUseCase,
		getBookListUseCase,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
		uploadImgUseCase,
	)

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
		getBookUseCase,
		getBookListUseCase,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)
}

//func runGinServer(
//	config configs.Config,
//	maker auth.TokenMaker,
//	registerBookUseCase usecase.RegisterBookUseCase,
//	deleteBookUseCase usecase.DeleteBookUseCase,
//	signUpUseCase usecase.SignUpUseCase,
//	signInUseCase usecase.SignInUseCase,
//	refreshTokenUseCase usecase.RefreshAccessTokenUseCase,
//) {
//	bookController := controller.NewBookController(registerBookUseCase, deleteBookUseCase)
//	userController := controller.NewUserController(config, maker, signUpUseCase, signInUseCase, refreshTokenUseCase)
//
//	r := router.Setup(middleware.Authorize(maker), bookController, userController)
//	err := r.Run(config.HTTPServerAddress)
//
//	if err != nil {
//		log.Fatal("cannot start server:", err)
//	}
//}

func runGRPCServer(
	config configs.Config,
	maker auth.TokenMaker,
	registerBookUseCase bookUseCase.RegisterBookUseCase,
	deleteBookUseCase bookUseCase.DeleteBookUseCase,
	getBookUseCase bookUseCase.GetBookUseCase,
	getBookListUseCase bookUseCase.GetBookListUseCase,
	signUpUseCase userUseCase.SignUpUseCase,
	signInUseCase userUseCase.SignInUseCase,
	refreshTokenUseCase userUseCase.RefreshAccessTokenUseCase,
) {
	grpcServer := grpc.NewServer()

	userSrv := userServer.NewUserServer(
		config,
		maker,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)
	bookSrv := bookServer.NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
		getBookUseCase,
		getBookListUseCase,
	)
	pb.RegisterUserServiceServer(grpcServer, userSrv)
	pb.RegisterBookServiceServer(grpcServer, bookSrv)
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

// registerGenres registers predefined genres in the database
func registerGenres(createGenresUseCase bookUseCase.CreateGenresUseCase) {
	// Get the list of genres from entity
	genreNames := bookDomain.GetGenres()

	// Create a request with the genre names
	request := bookUseCase.NewCreateGenresRequest(genreNames)

	// Call the use case to register the genres
	err := createGenresUseCase.CreateGenres(context.Background(), request)
	if err != nil {
		log.Fatal("failed to register genres:", err)
	}

	log.Println("Successfully registered genres")
}

func runGatewayServer(
	config configs.Config,
	maker auth.TokenMaker,
	registerBookUseCase bookUseCase.RegisterBookUseCase,
	deleteBookUseCase bookUseCase.DeleteBookUseCase,
	getBookUseCase bookUseCase.GetBookUseCase,
	getBookListUseCase bookUseCase.GetBookListUseCase,
	signUpUseCase userUseCase.SignUpUseCase,
	signInUseCase userUseCase.SignInUseCase,
	refreshTokenUseCase userUseCase.RefreshAccessTokenUseCase,
	uploadImgUseCase imageUseCase.UploadImgUseCase,
) {
	userSrv := userServer.NewUserServer(
		config,
		maker,
		signUpUseCase,
		signInUseCase,
		refreshTokenUseCase,
	)
	bookSrv := bookServer.NewBookServer(
		maker,
		registerBookUseCase,
		deleteBookUseCase,
		getBookUseCase,
		getBookListUseCase,
	)
	imgServer := imageServer.NewImageServer(uploadImgUseCase)

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, userSrv)
	if err != nil {
		log.Fatalf("cannot register handle server: %v", err)
	}
	err = pb.RegisterBookServiceHandlerServer(ctx, grpcMux, bookSrv)
	if err != nil {
		log.Fatalf("cannot register handle server: %v", err)
	}

	// クライアントから実際のHTTPリクエストを受け取る
	httpMux := http.NewServeMux()
	// HTTPリクエストをgRPCのリクエストに変換するするためにgrpcMuxにルーティングする
	httpMux.Handle("/", grpcMux)

	// REST APIのルーティング（画像アップロードAPIはgRPC未対応のため）
	r := router.Setup(auth.AuthenticateHTTP(maker), image.ValidateImageUpload(), imgServer)
	httpMux.Handle("/v1/image/upload", r)

	// CORSの設定
	c := cors.Default()
	handler := c.Handler(httpMux)

	httpServer := &http.Server{
		Handler: handler,
		Addr:    config.HTTPServerAddress,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start HTTP server: %v", err)
	}

	// log.Printf("start HTTP gateway server at %s", httpServer.Addr)
}
