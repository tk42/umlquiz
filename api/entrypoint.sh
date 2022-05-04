#!/bin/bash

# turn on bash's job control
set -m

if [ -z "$REST_PORT" ]; then
  echo REST_PORT is empty. This field is mandatory
  exit 1
fi

python3 broker/broker.py &
python3 backend/controller.py &
uvicorn frontend.server:app --host 0.0.0.0 --port ${REST_PORT} &

# now we bring the primary process back into the foreground
# and leave it there
fg %1
