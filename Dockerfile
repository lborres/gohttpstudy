ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine

# TODO Implement this
# Create new user
# Prevent running the container as root user
# Start with creating group for the app user: 
# -S -> create system user
# -G -> add user to group
# RUN addgroup app && \
#     adduser --system --group appuser

# cd into the app dir
# This will house the source code
WORKDIR /app

COPY go.mod go.sum ./

# After copying the mod files
# Recursively update app dir privs
# RUN chown -R appuser:app .

# Run further commands as appuser
# USER appuser

RUN go mod download

# ? Why copy?
COPY . .

# ? Do I need this?
# ENV GO111MODULE=on

# Ensure go OS builds to linux
# CGO_ENABLED=0 does not require C libs
ENV GOOS=linux CGO_ENABLED=0

# Now Build
RUN go build -o /gohttpstudy ./cmd/main.go

EXPOSE 8081

CMD ["/gohttpstudy"]
