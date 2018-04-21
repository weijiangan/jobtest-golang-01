package main

import (
	"database/sql"
	"testing"

	"github.com/go-pg/pg"
	"github.com/golang/mock/gomock"
	pb "github.com/weijiangan/bruno-test/brunotest"
	querymock "github.com/weijiangan/bruno-test/mock_bruno-test"
	"golang.org/x/net/context"
)

func TestTestDB_Model(t *testing.T) {
	TestDB_Model()
	defer db.Close()

	var got int
	_, err := db.QueryOne(pg.Scan(&got), "SELECT 1")
	if err != nil {
		t.Errorf("TestDB_Model() failed with error: %v", err)
	}
	if got != 1 {
		t.Errorf("TestDB_Model() queried 'SELECT 1', got %d", got)
	}
}

func TestSend(t *testing.T) {
	TestDB_Model()
	defer db.Close()
	s := server{}

	cases := []struct {
		in   pb.AuditEvent
		want pb.Response
	}{
		{
			pb.AuditEvent{
				Tag:     map[string]string{"key1": "value1"},
				Message: "Test1",
			},
			pb.Response{StatusCode: 200, Message: "OK"},
		},
		{
			pb.AuditEvent{
				ClientIp: "201.16.204.114",
				Tag:      map[string]string{"key1": "value1", "key2": "value2"},
				Message:  "Test2",
			},
			pb.Response{StatusCode: 200, Message: "OK"},
		},
		{
			pb.AuditEvent{
				ServerIp: "100.92.64.121",
				Tag:      map[string]string{"key3": "value3"},
				Message:  "Test3",
			},
			pb.Response{StatusCode: 200, Message: "OK"},
		},
		{
			pb.AuditEvent{
				ClientIp: "249.208.100.209",
				ServerIp: "6.103.104.214",
				Tag:      map[string]string{"key3": "value3", "key4": "value4"},
				Message:  "Test4",
			},
			pb.Response{StatusCode: 200, Message: "OK"},
		},
	}

	tx, err := db.Begin()
	if err != nil {
		t.Errorf("Unable to begin transaction mode: %v", err)
	}
	defer tx.Rollback()
	dbInsert = tx.Insert

	for _, c := range cases {
		got, err := s.Send(context.Background(), &c.in)
		if err != nil {
			t.Errorf("Send(%+v) got unexpected error: %v", c.in, err)
		}
		if *got != c.want {
			t.Errorf("Send(%+v) == %+v, want %+v", c.in, got, c.want)
		}
	}

}

func TestQueryClient(t *testing.T) {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "brunotest",
	})
	if err := createSchema(db, []interface{}{&AuditEvent{}, &Log{}}); err != nil {
		panic(err)
	}
	defer db.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	msg := &pb.AuditEvent{Tag: map[string]string{"boruto": "stream"}}

	// Create mock for the stream returned by Query
	mockStream := querymock.NewMockApp_QueryClient(ctrl)
	// set expectation on sending.
	mockStream.EXPECT().Recv().Return(msg, nil)

	queryClient := querymock.NewMockAppClient(ctrl)
	queryClient.EXPECT().Query(
		gomock.Any(),
		gomock.Any(),
	).Return(mockStream, nil)

	stream, err := queryClient.Query(context.Background(), &pb.QueryParam{
		Tag: map[string]string{"foo": "bar"},
	})
	if err != nil {
		t.Fatalf("TestQuery() got unexpected error: %v", err)
	}

	got, err := stream.Recv()
	if err != nil {
		t.Fatalf("TestQuery() got unexpected error receiving stream: %v", err)
	}
	t.Logf("%v\n", got)
}

func TestNullStringify(t *testing.T) {
	cases := []struct {
		in   string
		want sql.NullString
	}{
		{"", sql.NullString{}},
		{"test string 123", sql.NullString{String: "test string 123", Valid: true}},
	}

	for _, c := range cases {
		got := NullStringify(c.in)
		if got != c.want {
			t.Errorf("NullStringify(%q) == %+v, want %+v", c.in, got, c.want)
		}
	}
}

func benchmarkSend(input pb.AuditEvent, b *testing.B) {
	TestDB_Model()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		b.Errorf("Unable to begin transaction mode: %v", err)
	}
	defer tx.Rollback()
	dbInsert = tx.Insert
	s := server{}

	for n := 0; n < b.N; n++ {
		s.Send(context.Background(), &input)
	}
}

func BenchmarkSend(b *testing.B) {
	benchmarkSend(pb.AuditEvent{
		ClientIp: "249.208.100.209",
		ServerIp: "6.103.104.214",
		Tag:      map[string]string{"key3": "value3", "key4": "value4"},
		Message:  "Test4",
	}, b)
}
