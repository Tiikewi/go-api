FROM golang:1.16.5 as development

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest

EXPOSE 8080

# Reflex start server again when changes done to .go files.
CMD reflex -r '\.go$' go run . --start-service