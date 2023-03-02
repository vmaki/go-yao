# 容器名
name=go-yao-docker

# 停止该镜像正在运行的Docker容器
# shellcheck disable=SC2143
if [ -n "$(docker ps | grep $name)" ]; then
	echo "存在正在运行的$name容器, 正在使其停止运行..."
	docker stop $name
	echo "$name容器, 已停止运行"
fi

# 删除该镜像的Docker容器
# line=`docker ps -a | grep $name`
# shellcheck disable=SC2143
if [ -n "$(docker ps -a | grep $name)" ]; then
	echo "存在$name容器, 对其进行删除..."
	docker rm $name
	echo "$name容器, 已被删除"
fi

docker-compose build

docker-compose up -d
