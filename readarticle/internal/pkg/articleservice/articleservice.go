package articleservice

import (
	"context"
	"log"

	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
	repo "github.com/shanehowearth/cloudblog/readarticle/integration/repository/cache/v1"
)

// articleServiceServer is implementation of v1.ArticleServiceServer proto interface
type articleServiceServer struct {
	Cache repo.Cache
}

// NewArticleService creates Article service
func NewArticleService(c repo.Cache) v1.ArticleServiceServer {
	if c == nil {
		log.Panic("NewArticleService has no cache to get articles from")
	}
	a := articleServiceServer{Cache: c}
	return &a
}

// GetArticle(context.Context, *ArticleRequest) (*Article, error)
func (a *articleServiceServer) GetArticle(ctx context.Context, req *v1.ArticleRequest) (*v1.Article, error) {
	article := a.Cache.GetByTitle(req.GetTitle())

	return article, nil
}

// For pagination
const MaxRetrieve = 50

// GetArticles(context.Context, *ArticleOffset) (*ArticleList, error)
func (a *articleServiceServer) GetArticles(ctx context.Context, req *v1.ArticleOffset) (*v1.ArticleList, error) {
	offset := req.GetOffset()
	// negative offsets not supported
	if offset < 0 {
		return &(v1.ArticleList{}), nil
	}
	articleList := a.Cache.GetFifty(offset, MaxRetrieve)

	return articleList, nil
}
