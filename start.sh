#!/bin/bash

echo "Starting Faker"
ls -la /app/code/api/

echo "2"
/app/code/api/faker
echo "3"
exec /usr/local/bin/gosu cloudron:cloudron /app/code/api/faker