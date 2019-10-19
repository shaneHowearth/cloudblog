#!/bin/bash
docker stop ra 
docker rm ra 
docker rmi readarticle 
./builder.sh

docker run -d -p 5100:5100 --name ra  readarticle

# Insert some test artifacts
docker exec -d ra /bin/bash -c "echo \"hset articles:001 created '2017-12-19 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:001'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:002 created '2018-01-07 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:002'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:003 created '2018-01-08 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:003'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:004 created '2018-01-09 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:004'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:005 created '2018-01-10 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:005'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:006 created '2018-01-11 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:006'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:007 created '2018-01-12 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:007'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:008 created '2018-01-13 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:008'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:009 created '2018-01-14 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:009'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:010 created '2018-01-15 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:010'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:011 created '2018-01-16 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:011'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:012 created '2018-01-17 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:012'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:013 created '2018-01-18 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:013'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:014 created '2018-01-19 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:014'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:015 created '2018-01-20 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:015'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:016 created '2018-01-21 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:016'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:017 created '2018-01-22 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:017'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:018 created '2018-01-23 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:018'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:019 created '2018-01-24 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:019'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:020 created '2018-01-25 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:020'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:021 created '2018-01-26 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:021'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:022 created '2018-01-27 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:022'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:023 created '2018-01-28 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:023'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:024 created '2018-01-29 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:024'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:025 created '2018-01-30 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:025'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:026 created '2018-01-31 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:026'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:027 created '2018-02-01 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:027'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:028 created '2018-02-02 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:028'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:029 created '2018-02-03 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:029'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:030 created '2018-02-04 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:030'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:031 created '2018-02-05 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:031'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:032 created '2018-02-06 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:032'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:033 created '2018-02-07 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:033'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:034 created '2018-02-08 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:034'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:035 created '2018-02-09 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:035'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:036 created '2018-02-10 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:036'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:037 created '2018-02-11 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:037'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:038 created '2018-02-12 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:038'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:039 created '2018-02-13 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:039'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:040 created '2018-02-14 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:040'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:041 created '2018-02-15 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:041'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:042 created '2018-02-16 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:042'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:043 created '2018-02-17 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:043'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:044 created '2018-02-18 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:044'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:045 created '2018-02-19 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:045'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:046 created '2018-02-20 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:046'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:047 created '2018-02-21 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:047'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:048 created '2018-02-22 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:048'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:049 created '2018-02-23 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:049'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:050 created '2018-02-24 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:050'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:051 created '2018-02-25 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:051'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:052 created '2018-02-26 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:052'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:053 created '2018-02-27 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:053'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:054 created '2018-02-28 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:054'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:055 created '2018-03-01 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:055'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:056 created '2018-03-02 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:056'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:057 created '2018-03-03 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:057'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:058 created '2018-03-04 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:058'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:059 created '2018-03-05 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:059'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:060 created '2018-03-06 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:060'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:061 created '2018-03-07 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:061'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:062 created '2018-03-08 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:062'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:063 created '2018-03-09 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:063'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:064 created '2018-03-10 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:064'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:065 created '2018-03-11 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:065'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:066 created '2018-03-12 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:066'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:067 created '2018-03-13 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:067'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"
# Second
docker exec -d ra /bin/bash -c "echo \"hset articles:068 created '2018-03-14 05:22:35+11' title 'Javascript Routing'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:068'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'Javascript Routing' articles:002\" | redis-cli"
#Third
docker exec -d ra /bin/bash -c "echo \"hset articles:069 created '2018-03-15 22:01:37+11' title 'FUBAR'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:069'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'FUBAR' articles:003\" | redis-cli"

docker exec -d ra /bin/bash -c "echo \"hset articles:070 created '2018-03-16 22:01:37+11' title 'First Article'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"lpush article_keys 'articles:070'\" | redis-cli"
docker exec -d ra /bin/bash -c "echo \"hset articles:lookup:title 'First Article' articles:001\" | redis-cli"


# test grpc commands
grpcurl -v -plaintext localhost:5100 list

# Test our methods
# grpcurl -d '{"id": 1234, "tags": ["foo","bar"]}' \
    # grpc.server.com:443 my.custom.server.Service/Method
# grpcurl -v -plaintext -d '{"title": "Javascript Routing"}' \
#     localhost:5100 v1.ArticleService/GetArticle
# echo "############ Get Articles, no offset ###############"
# grpcurl -v -plaintext -d '' \
#     localhost:5100 v1.ArticleService/GetArticles
# echo "############ Get Articles, offset ###############"
# grpcurl -v -plaintext -d '{"offset": 2}' \
#     localhost:5100 v1.ArticleService/GetArticles
echo "############ Get Articles, nothing to fetch ###############"
grpcurl -v -plaintext -d '{"title": "Does not exist"}' \
    localhost:5100 v1.ArticleService/GetArticle

echo "########################### LOGS ######################################"
docker logs ra
