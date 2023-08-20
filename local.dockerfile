FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go get github.com/cosmtrek/air
RUN go install github.com/cosmtrek/air
CMD ["air"]