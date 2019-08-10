
### サンプル用のDBコンテナ立ち上げる時
docker container run --rm -d \
    -v $PWD/database/init:/docker-entrypoint-initdb.d \
    -e MYSQL_ROOT_PASSWORD=root \
    -p 43306:3306 --name mysql mysql:5.7