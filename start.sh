#!/bin/bash

echo "Starting Faker"
ls -la /app/pkg
ls -la /app/code
ls -la /app/code/api
/app/code/api/faker
#exec /usr/local/bin/gosu cloudron:cloudron /app/code/api/faker