FROM golang:latest
ENV TZ=Asia/Tokyo

WORKDIR /workspace
COPY . /workspace/
RUN go build .
RUN chmod +x main
CMD ./main
