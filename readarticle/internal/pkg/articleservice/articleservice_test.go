package articleservice

import (
	"context"
	"testing"

	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
	repo "github.com/shanehowearth/cloudblog/readarticle/integration/repository/cache/v1"

	// as "github.com/shanehowearth/cloudblog/readarticle/internal/pkg/articleservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyTestCache struct {
	mock.Mock
}

func (c *MyTestCache) GetByTitle(title string) *v1.Article {
	args := c.Called(title)
	return args.Get(0).(*v1.Article)
}

func (c *MyTestCache) GetFifty(offset, maxretrieve int32) *v1.ArticleList {
	args := c.Called(offset, maxretrieve)
	return args.Get(0).(*v1.ArticleList)
}

func TestNewArticleService(t *testing.T) {
	testcases := map[string]struct {
		cache repo.Cache
		err   string
	}{
		"Happy path": {
			cache: new(MyTestCache),
		},
		"No Cache": {
			cache: nil,
			err:   "NewArticleService has no cache to get articles from",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if tc.err != "" {
				assert.PanicsWithValue(t, tc.err, func() { NewArticleService(tc.cache) }, "New Article did not panic")
			} else {
				output := NewArticleService(tc.cache)
				_, isa := output.(v1.ArticleServiceServer)
				assert.True(t, isa, "output is not an ArticleServiceServer")
			}
		})
	}
}

func TestGetArticle(t *testing.T) {
	testcases := map[string]struct {
		title  string
		exists bool
	}{
		"Happy Path": {
			title:  "Test title",
			exists: true,
		},
		"Nothing found": {
			title: "Does not exist",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			testService := NewArticleService(new(MyTestCache))
			if tc.exists {
				testService.(*articleServiceServer).Cache.(*MyTestCache).On("GetByTitle", tc.title).Return(&(v1.Article{Title: tc.title}))
			} else {
				testService.(*articleServiceServer).Cache.(*MyTestCache).On("GetByTitle", tc.title).Return(&(v1.Article{}))

			}
			req := v1.ArticleRequest{Title: tc.title}
			a, e := testService.GetArticle(context.Background(), &req)
			assert.Nil(t, e, "Error should be nil")
			assert.NotNil(t, a, "Article should not be nil")
		})
	}
}

func TestGetArticles(t *testing.T) {
	testcases := map[string]struct {
		offset      int32
		maxRetrieve int32
	}{
		"Happy Path":      {offset: 0, maxRetrieve: 50},
		"Negative offset": {offset: -17, maxRetrieve: 50},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			testService := NewArticleService(new(MyTestCache))
			testService.(*articleServiceServer).Cache.(*MyTestCache).On("GetFifty", tc.offset, tc.maxRetrieve).Return(&(v1.ArticleList{}))

			req := v1.ArticleOffset{Offset: tc.offset}
			a, e := testService.GetArticles(context.Background(), &req)
			assert.Nil(t, e, "Error should be nil")
			assert.NotNil(t, a, "Article should not be nil")
		})
	}
}
