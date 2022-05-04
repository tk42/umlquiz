FROM tiangolo/uvicorn-gunicorn-fastapi:latest
ENV TZ=Asia/Tokyo

RUN apt-get install -y ca-certificates

WORKDIR /workspace
COPY ./requirements.txt /workspace/

RUN pip install -r /workspace/requirements.txt
COPY . /workspace/
RUN chmod +x ./entrypoint.sh
CMD ./entrypoint.sh
