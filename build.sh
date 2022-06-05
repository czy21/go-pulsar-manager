#!/bin/bash

set -x

cd $(cd "$(dirname "$0")"; pwd)

web_cmd="yarn --cwd web/"

if [ ${param_npm_repo} ]; then
  web_cmd+=" --registry ${param_npm_repo}"
fi

if [ ${param_yarn_cache} ]; then
  web_cmd+=" --cache-folder ${param_yarn_cache}"
fi

${web_cmd} install --no-lockfile --update-checksums
${web_cmd} --ignore-engines build

cd api/;GOSUMDB=off go build -o build/ -x -v main.go