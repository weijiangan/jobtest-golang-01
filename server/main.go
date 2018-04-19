package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	pb "github.com/weijiangan/bruno-test/brunotest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

var (
	db    *pg.DB
	sugar *zap.SugaredLogger
)

type server struct{}

type AuditEvent struct {
	Id       int32
	ClientIp string
	ServerIp string
	Tag      map[string]string
	Message  string
}

type Log struct {
	Id  int32
	Log zapcore.Entry
}

func (s *server) Send(ctx context.Context, in *pb.AuditEvent) (*pb.Response, error) {
	event := &AuditEvent{
		ClientIp: in.ClientIp,
		ServerIp: in.ServerIp,
		Tag:      in.Tag,
		Message:  in.Message,
	}
	err := db.Insert(event)
	if err != nil {
		return &pb.Response{StatusCode: 400, Message: "Bad Request"}, err
	}
	return &pb.Response{StatusCode: 200, Message: "OK"}, nil
}

func (s *server) Query(in *pb.QueryParam, stream pb.App_QueryServer) error {
	var events []AuditEvent
	query := db.Model(&events).
		Where("?0 IS NULL OR client_ip = ?0", NullStringify(in.ClientIp)).
		Where("?0 IS NULL OR server_ip = ?0", NullStringify(in.ServerIp))
	for k, v := range in.Tag {
		query = query.Where("tag->>? = ?", k, v)
	}
	if err := query.Select(); err != nil {
		sugar.Fatalf("%v", err)
		return err
	}
	fmt.Printf("%+v\n", events)

	for _, event := range events {
		res := &pb.AuditEvent{
			ClientIp: event.ClientIp,
			ServerIp: event.ServerIp,
			Tag:      event.Tag,
			Message:  event.Message,
		}
		if err := stream.Send(res); err != nil {
			sugar.Fatalf("%v", err)
			return err
		}
	}

	return nil
}

func NullStringify(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  len(s) > 0,
	}
}

func TestDB_Model() {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "brunotest",
	})
	err := createSchema(db, []interface{}{&AuditEvent{}})
	if err != nil {
		panic(err)
	}
	// db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
	// 	query, err := event.FormattedQuery()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Printf("%s %s", time.Since(event.StartTime), query)
	// })
}

func createSchema(db *pg.DB, models []interface{}) error {
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar = logger.Sugar()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		sugar.Fatalf("Failed to listen port %s: %v", port, err)
	}
	TestDB_Model()
	defer db.Close()
	s := grpc.NewServer()
	pb.RegisterAppServer(s, &server{})
	reflection.Register(s)
	sugar.Infof("Server listening: http://localhost%s", port)
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("Failed to serve: %v", err)
	}
}
