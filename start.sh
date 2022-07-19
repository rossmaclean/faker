#!/bin/bash

echo "Starting Faker"
exec /usr/local/bin/gosu cloudron:cloudron /app/code/api/faker /app/code/api/configs/