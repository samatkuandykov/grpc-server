FROM golang:latest
RUN mkdir /app
COPY . /app/
WORKDIR /app
EXPOSE 8081