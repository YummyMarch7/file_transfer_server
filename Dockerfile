FROM golang:alpine
LABEL authors="aaa"

ENTRYPOINT ["top", "-b"]