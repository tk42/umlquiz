# -*- coding: utf-8 -*-
import os
import zmq
import time
import logging
import asyncio
import traceback
from enum import Enum
import orjson as json

from handlers import user, quiz, quiz_by_user, quizs_by_user


created = set()


def run(worker_id: str):
    asyncio.run(start_worker(worker_id))


async def start_worker(worker_id: str):
    context = zmq.Context()
    socket = context.socket(zmq.REP)
    socket.connect("tcp://localhost:5557")

    logging.info("Start worker: %s" % worker_id)

    while True:
        try:
            msg = socket.recv_json()

            _group = msg.get("group")
            _method = msg.get("method")
            _params = msg.get("params", {})

            logging.info(f"{time.time()} {_group} {_method} {_params}")

            resp = None

            if _group == "user":
                assert _method in dir(user), f"Unknown method({_method} in group{_group})"
                resp = getattr(user, _method)(**_params)
            elif _group == "quiz":
                assert _method in dir(quiz), f"Unknown method({_method} in group{_group})"
                resp = getattr(quiz, _method)(**_params)
            elif _group == "quiz_by_user":
                assert _method in dir(quiz_by_user), f"Unknown method({_method} in group{_group})"
                resp = getattr(quiz_by_user, _method)(**_params)
            elif _group == "quizs_by_user":
                assert _method in dir(quizs_by_user), f"Unknown method({_method} in group{_group})"
                resp = getattr(quizs_by_user, _method)(**_params)
            else:
                raise ValueError(f"Unknown group({_group})")
        except Exception as e:
            resp = {
                "type": str(e.__class__.__name__),
                "context": str(e.__context__),
                "error": str(e),
                "group": str(_group),
                "method": str(_method),
                "params": str(_params),
                "stacktrace": traceback.format_exc(),
            }

        if type(resp) is str:
            resp = {"timestamp": time.time(), "response": resp}
        elif resp is None:
            resp = {"timestamp": time.time(), "response": ""}
        socket.send(json.dumps(resp), copy=False)

    socket.close()
    context.destroy()
