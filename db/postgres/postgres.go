package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)


var DB *pgxpool.Pool

func DBConnect(){
	fmt.Println(os.Getenv("DB_URL"))
	var err error
	DB,err = pgxpool.New(context.Background(),os.Getenv("DB_URL"))
	if err!=nil{
		log.Panicln("db connection error",err)
	}
}  