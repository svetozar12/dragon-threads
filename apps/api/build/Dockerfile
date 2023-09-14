FROM golang:1.19 as deps

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./apps/api ./apps/api
COPY ./libs ./libs

RUN CGO_ENABLED=0 GOOS=linux go build -o /api apps/api/main.go

FROM scratch

COPY --from=deps api .

EXPOSE 3333

# Command to run when starting the container.
CMD ["/api"]