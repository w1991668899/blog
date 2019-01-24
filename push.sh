#!/usr/bin/env bash

git status && git add .

git commit -m "master tools auto push"

echo "commit Success"

git pull origin  master

echo "pull Success"

git push origin  master

echo "push Success"
