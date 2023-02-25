#!/bin/sh

Name=$1

startTask(){
  # shellcheck disable=SC2164
  nohup ./"${Name}" --env=prod >/dev/null 2>&1 &
  echo "starting..."
}

goBuild(){
  # shellcheck disable=SC2164
  echo "go build begin!"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "${Name}"
  echo "go build Success!"
}

# shellcheck disable=SC2039
if [[ $1 == "" ]]
then
  echo "参数错误!如 app.sh go-yao build"
  exit
fi


#查找进程id
# shellcheck disable=SC2006
# shellcheck disable=SC2009
pid=`ps -ef | grep "$1" | grep -v 'grep' | grep -v "$0" | awk '{print $2}'`
# shellcheck disable=SC2039
if [[ $2 == "start" ]]
then
  if [[ $pid != "" ]]
  then
    echo "$1" "is running in" "$pid"
  else
    startTask
    echo "$1" "Start Success..."
  fi
elif [[ $2 == "stop" ]]
then
  if [[ $pid != "" ]]
  then
    kill "${pid}"
    echo "$1" "Stop Success..."
  else
    echo "$1" "is Stop"
  fi
elif [[ $2 == "restart" ]]
then
  if [[ $pid != "" ]]
  then
    kill -1 "${pid}"
    echo "$1" "in" "${pid}" "Restart Success..."
  else
    echo "$1" "is Stop"
  fi
elif [[ $2 == "build" ]]
then
  goBuild
else
  echo "参数错误!如 app.sh go-yao build"
fi
