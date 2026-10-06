package main

import (
	"StudyApi/myhttp"
	objects "StudyApi/objects/subject"
	"StudyApi/sql"
	"context"
)

func main() {
	ctx := context.Background()
	conn, err := sql.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	db := objects.NewDatabase(ctx, conn)
	httpHandlers := myhttp.NewHTTPHandlers(db)
	server := myhttp.NewServer(httpHandlers)

	if err := server.StartServer(); err != nil {
		panic(err)
	}

}
