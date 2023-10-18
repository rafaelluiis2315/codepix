package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/rafael/codepix/application/grpc"
	"github.com/rafael/codepix/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
