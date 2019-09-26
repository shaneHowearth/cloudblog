package repository

// Provide interfaces for the repository cache to use in order for the
// microservice to communicate with the cache.
// Use interfaces here so the cache depends on the microservice, not the other way round
// allowing any cache to be dropped in, as long as it satisfies the interface(s)

import (
	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
)

type Cache interface {
	GetByTitle(title string) *v1.Article
	GetFifty(offset, maxretrieve int32) *v1.ArticleList
}
