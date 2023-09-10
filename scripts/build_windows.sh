#!/usr/bin/env bash
SCRIPTS_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

source $SCRIPTS_ROOT/cfg/services.cfg
source $SCRIPTS_ROOT/cfg/docker-composes.cfg

LARK_APP=$(dirname "${SCRIPTS_ROOT}")
BUILD_BIN=${LARK_APP}"/build/bin"

length=${#service_names[@]}
echo $length
for (( i=0; i <= $length; i++ )); do
  service=${service_names[$i]}
  if [ -z ${service} ];then
      continue
  fi
  app_cmd=${LARK_APP}/apps/${service}/cmd
  echo "build "${app_cmd}
  cd ${app_cmd}
  go build -o ${LARK_APP}/lark_${service}.exe
done
echo "build success..."
