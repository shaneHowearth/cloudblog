package main

import (
	"context"
	"log"
	"net"
	"os"

	article "github.com/shanehowearth/cloudblog/readarticle/internal/pkg/articleservice"
	repo "github.com/shanehowearth/cloudblog/readarticle/internal/repository/redis"

	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var ss = article.NewArticleService(new(repo.Redis))

func main() {

	portNum := os.Getenv("PORT_NUM")
	lis, err := net.Listen("tcp", "0.0.0.0:"+portNum)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterArticleServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// GetArticle -
func (s *server) GetArticle(ctx context.Context, req *v1.ArticleRequest) (*v1.Article, error) {
	st, err := ss.GetArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// GetArticles -
func (s *server) GetArticles(ctx context.Context, req *v1.ArticleOffset) (*v1.ArticleList, error) {
	st, err := ss.GetArticles(ctx, req)
	if err != nil {
		return nil, err
	}
	return st, nil
}
