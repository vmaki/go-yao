BINARY="go-yao"

build:
	sh app.sh ${BINARY} build

start:
	sh app.sh ${BINARY} start

stop:
	sh app.sh ${BINARY} stop

restart:
	sh app.sh ${BINARY} restart

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make start - 启动项目"
	@echo "make stop - 停止项目"
	@echo "make restart - 重启项目"
