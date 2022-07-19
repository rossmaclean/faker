#!/bin/bash

echo "Starting Faker"
ls -la /app/code/api/

exec /usr/local/bin/gosu cloudron:cloudron /app/code/api/faker