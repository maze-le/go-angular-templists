#!/usr/bin/bash

cd ./src/client && npm i && npm run build
cd - && docker-compose -f docker-compose.yml up --build
