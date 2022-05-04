# -*- coding: utf-8 -*-
import os
import logging

# Custom formatter returns a structure, than a string
class CustomFormatter(logging.Formatter):
    def format(self, record):
        return {'msg': super(CustomFormatter, self).format(record), 'args':record.args}

logger = None

if eval(os.getenv("CLOUD_LOGGER", default="True")):
    logger = logging.getLogger("cloudLogger")

    import google.cloud.logging
    # Instantiates a client
    client = google.cloud.logging.Client()
    # client.setup_logging()
    handler = client.get_default_handler()
    handler.setFormatter(CustomFormatter())

    logger.setLevel(logging.getLevelName(os.getenv("LOGLEVEL", default="INFO")))
    logger.addHandler(handler)
else:
    logging.basicConfig(level=logging.getLevelName(os.getenv("LOGLEVEL", default="DEBUG")))
    logger = logging
