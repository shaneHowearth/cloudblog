package rediscache_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
	rediscache "github.com/shanehowearth/cloudblog/readarticle/internal/repository/redis"
	"github.com/stretchr/testify/assert"
)

// Create global (to these tests) Redis and connection objects
var (
	testR rediscache.Redis
	conn  *redigomock.Conn
)

func TestMain(m *testing.M) {

	conn = redigomock.NewConn()
	mockPool := &redis.Pool{
		Dial:    func() (redis.Conn, error) { return conn, nil },
		MaxIdle: 10,
	}
	testR = rediscache.Redis{Pool: mockPool}
	os.Exit(m.Run())
}

func TestGetByTitle(t *testing.T) {
	testcases := map[string]struct {
		title  string
		expect bool
	}{
		"Happy Path":        {title: "Success", expect: true},
		"Nothing to return": {title: "Nothing"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			conn.Clear()
			if tc.expect {
				conn.Command("HGET", "articles:lookup:title", tc.title).Expect("1")
				conn.Command("HGETALL", "1").ExpectMap(map[string]string{"title": tc.title})
			} else {
				conn.Command("HGET", "articles:lookup:title", tc.title)
				conn.Command("HGETALL").ExpectMap(map[string]string{})

			}
			output := testR.GetByTitle(tc.title)
			if tc.expect {
				assert.Equal(t, &v1.Article{Title: tc.title, UrlTitle: "1"}, output, "Article returned did not match expected")
			} else {

				assert.Equal(t, &v1.Article{}, output, "Article returned did not match expected")
			}
		})
	}
}

func TestGetFifty(t *testing.T) {
	testcases := map[string]struct {
		title       string
		offset      int32
		maxRetrieve int32
		expect      bool
		numArticles int
	}{
		"Happy Path":        {title: "success", offset: 0, maxRetrieve: 50, expect: true, numArticles: 2},
		"Filled list":       {title: "Nothing", offset: 0, maxRetrieve: 1, expect: true, numArticles: 1},
		"Nothing to return": {title: "Nothing", offset: 0, maxRetrieve: 50},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			conn.Clear()
			if tc.expect {
				conn.Command("LRANGE", "article_keys", tc.offset, tc.maxRetrieve).ExpectStringSlice("1", "2")
				conn.Command("HGETALL", "1").ExpectMap(map[string]string{"title": tc.title})
				conn.Command("HGETALL", "2").ExpectMap(map[string]string{"title": tc.title, "created": "2019-10-11 10:11:12"})
			} else {
				conn.Command("LRANGE", "article_keys", tc.offset, tc.maxRetrieve).ExpectError(fmt.Errorf("unexpected Error"))
				conn.Command("HGET", "articles:lookup:title", tc.title)
				conn.Command("HGETALL").ExpectMap(map[string]string{})

			}
			output := testR.GetFifty(tc.offset, tc.maxRetrieve)
			if tc.expect {
				a := []*v1.Article{&v1.Article{Title: tc.title, UrlTitle: "1"}, &v1.Article{Title: tc.title, UrlTitle: "2", Year: 2019, Month: 10, Day: 11, Hour: 10, Minute: 11}}

				assert.Equal(t, &v1.ArticleList{Articles: a[:tc.numArticles]}, output, "Article returned did not match expected")
			} else {

				assert.Equal(t, &v1.ArticleList{Articles: []*v1.Article{}}, output, "Article returned did not match expected")
			}
		})
	}

}

func TestGetByID(t *testing.T) {
	testcases := map[string]struct {
		title string
	}{
		"Error from redis": {title: "1"},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			conn.Clear()
			conn.Command("HGETALL").ExpectError(fmt.Errorf("unexpected Error"))

			output := testR.GetByID(tc.title)

			assert.Equal(t, &v1.Article{}, output, "Article returned did not match expected")

		})
	}
}
