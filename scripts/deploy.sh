#!/bin/bash

PROJECT_DIR="gymkhana"

CI_COMMIT_SHA=$1
echo $CI_COMMIT_SHA

#cd "/var/www/"${PROJECT_DIR}"/dist"
#git pull origin master
#git checkout $CI_COMMIT_SHA
#pipenv install
#cd src
#pipenv run python manage.py migrate
#pipenv run python manage.py collectstatic --noinput
#sudo service uwsgi restart
