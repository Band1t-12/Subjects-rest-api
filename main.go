package main

import (
	"StudyApi/myhttp"
	objects "StudyApi/objects/subject"
	"StudyApi/sql"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	conn, err := sql.CheckConnection(ctx)
	if err != nil {
		fmt.Println("Failed to connect!")
		return
	}
	db := objects.NewDatabase(ctx, conn)
	httpHandlers := myhttp.NewHTTPHandlers(db)
	server := myhttp.NewServer(httpHandlers)

	if err := server.StartServer(); err != nil {
		panic(err)
	}

}
