# -*- coding: utf-8 -*-
import os
from worker import run
import logging
from multiprocessing import Process, set_start_method
from handlers import Base, engine

# WEB_CONCURRENCY defines the number of worker processes, independent of how many CPU cores are available in the server.
# https://github.com/tiangolo/uvicorn-gunicorn-fastapi-docker#web_concurrency
WEB_CONCURRENCY = int(os.environ.get("WEB_CONCURRENCY", 4))
WORKDERS = WEB_CONCURRENCY * 2

def initialize():
    # create DBs after def ORM
    Base.metadata.create_all(engine)
    logging.info("Created DBs")


if __name__ == "__main__":
    set_start_method('fork')
    logging.info("Start controller")
    initialize()
    run("single")
    # for i in range(WORKDERS):
    #     p = Process(target=run, args=(i,))
    #     p.start()
