package polaris_server

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "github.com/tko-cactus/polaris/polaris_server/pb"
)

const port = ":54511"

type UserService struct {
}

func GetRoomsStatus() {

}

func (s *UserService) RegisterUserServiceServer(ctx context.Context, in *pb.User) (*pb.Response, error) {
	// dbに接続
	db, err := sql.Open("mysql", "root@/my_database")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// ユーザー情報を取得
	var id = db.QueryRow("SELECT id FROM Users WHERE id = MAX(id)")
	var uname = in.Name

	// Prepared Statements
	stmtInsert, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		return nil, err
	}
	defer stmtInsert.Close()

	// 新しいユーザー情報を登録する
	db.Query("INSERT INTO Users VALUES (?, ?)", id, uname)
	if err != nil {
		log.Fatal(err)
	}
	return nil, errors.New("raise error")
}
