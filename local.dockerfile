FROM golang:1.21.1 as go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go get github.com/cosmtrek/air
RUN go install github.com/cosmtrek/air
CMD ["air"]