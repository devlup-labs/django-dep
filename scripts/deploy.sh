#!/bin/bash

PROJECT_DIR="gymkhana"

CI_COMMIT_SHA=$1
FEATURE_ARGS=$2
RESTART_ARGS=$3
cd "/var/www/"${PROJECT_DIR}"/dist"
git pull origin master
git checkout $CI_COMMIT_SHA
pipenv install
cd src
case $FEATURE_ARGS$RESTART_ARGS in

"c")
  pipenv run python manage.py collectstatic --noinput
  ;;&
"m")
  pipenv run python manage.py migrate
  ;;&
"u")
  sudo service uwsgi restart
  ;;&
"n")
  sudo systemctl restart nginx
  ;;&
esac
