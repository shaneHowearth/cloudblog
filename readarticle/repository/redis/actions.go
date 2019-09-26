package rediscache

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	v1 "github.com/shanehowearth/cloudblog/readarticle/integration/grpc/gen/v1"
)

type Redis struct{}

// GetByTitle -
func (r *Redis) GetByTitle(title string) *v1.Article {
	// get conn and put back when exit from method
	var conn redis.Conn
	if pool == nil {
		initPool()
		conn = pool.Get()
		ping(conn)
	}
	if conn == nil {
		conn = pool.Get()
	}
	defer conn.Close()
	id, err := redis.String(conn.Do("HGET", "articles:lookup:title", title))
	if err != nil {
		log.Printf("ERROR: fail get id for title %s, error %s", title, err.Error())
		return &v1.Article{}
	}
	return r.GetByID(id)
}

// GetByID -
func (r *Redis) GetByID(id string) *v1.Article {
	// get conn and put back when exit from method
	var conn redis.Conn
	if pool == nil {
		initPool()
		conn = pool.Get()
		ping(conn)
	}
	if conn == nil {
		conn = pool.Get()
	}
	defer conn.Close()

	dataset, err := redis.Values(conn.Do("HGETALL", id))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", id, err.Error())
		return &v1.Article{}
	}
	// Put dataset into an Article
	a := v1.Article{}
	var created, title bool
	for i, data := range dataset {
		fmt.Println(i, " ", string(data.([]byte)))
		v := string(data.([]byte))
		if i%2 == 0 {
			created = v == "created"
			title = v == "title"
		} else {
			if created {
				// v1.Article is a business layer defined structure, so it makes sense to
				// coerce a repository layer structure into something that it will
				// understand.
				// This decouples the business layer from the repository layer because the
				// business layer does not need to know the implementation or structures
				// used in the repository.
				sd := strings.Split(v, " ")
				ddd := strings.Split(sd[0], "-")
				hhh := strings.Split(sd[1], ":")
				year, _ := strconv.Atoi(ddd[0])
				month, _ := strconv.Atoi(ddd[1])
				day, _ := strconv.Atoi(ddd[2])
				hour, _ := strconv.Atoi(hhh[0])
				minute, _ := strconv.Atoi(hhh[1])
				a.Year = int32(year)
				a.Month = int32(month)
				a.Day = int32(day)
				a.Hour = int32(hour)
				a.Minute = int32(minute)
			}
			if title {
				a.Title = v
			}
		}
	}
	a.UrlTitle = id
	return &a
}

// GetFifty -
func (r *Redis) GetFifty(offset, maxRetrieve int32) *v1.ArticleList {
	// get conn and put back when exit from method
	var conn redis.Conn
	if pool == nil {
		initPool()
		conn = pool.Get()
		ping(conn)
	}
	if conn == nil {
		conn = pool.Get()
	}
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("LRANGE", "article_keys", offset, offset+maxRetrieve))
	if err != nil {
		log.Printf("ERROR: fail get smembers with error: %s", err.Error())
		return &v1.ArticleList{Articles: []*v1.Article{}}
	}

	al := []*v1.Article{}

	// Stop when list is full, or there are no more articles
	for _, key := range keys {
		al = append(al, r.GetByID(key))
		if int32(len(al)) >= maxRetrieve {
			break
		}
	}

	return &v1.ArticleList{Articles: al}
}
