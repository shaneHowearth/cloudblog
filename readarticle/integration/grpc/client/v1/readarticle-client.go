package blogclient

import (
	"context"
	"log"
	"time"

	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
	"google.golang.org/grpc"
)

// BlogClient -
type BlogClient struct {
	Address string
}

func (s *BlogClient) newConnection() (v1.ArticleServiceClient, *grpc.ClientConn) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(s.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return v1.NewArticleServiceClient(conn), conn
}

// GetArticles -
func (s *BlogClient) GetArticles() (*v1.ArticleList, error) {
	c, conn := s.newConnection()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.GetArticles(ctx, nil)
}
