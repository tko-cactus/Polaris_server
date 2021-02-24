package polaris_server

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/tko-cactus/polaris/polaris_server/pb/github.com/tko-cactus/polaris/polaris_server/pb"
)

const port = ":54511"

type PolarisService struct {
}

func (s *PolarisService) GetRoomsStatus(ctx context.Context, in *pb.UserAndBeacon) *pb.RoomStatus {

}

func (s *PolarisService) RegisterUserServiceServer(ctx context.Context, in *pb.UserName) (*pb.User, error) {
	// dbに接続
	db, err := sql.Open("mysql", "root@/my_database")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// ユーザー情報を取得
	res := pb.User{
		Id:   0,
		Name: in.Uname,
	}
	// var usrid, err = db.QueryRow("SELECT id FROM Users WHERE id = MAX(id)")
	// var uname = in.Uname

	// Prepared Statements
	stmtInsert, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		return nil, errors.New("Prepared Statement Error")
	}
	defer stmtInsert.Close()

	// 新しいユーザー情報を登録する
	db.Query("INSERT INTO Users VALUES (?, ?)", res.Id, res.Name)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Register Error")
	}
	return &res, nil
}
