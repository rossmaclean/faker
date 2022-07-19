#!/bin/bash

echo "Starting Faker"
ls -la /app/pkg/api/

exec /usr/local/bin/gosu cloudron:cloudron /app/pkg/api/faker