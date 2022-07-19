#!/bin/bash

chown -R cloudron:cloudron /app/code

echo "Starting Faker"
exec /usr/local/bin/gosu cloudron:cloudron /app/code/api/faker