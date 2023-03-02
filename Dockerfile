# 사용할 이미지 태그
FROM golang:alpine as builder

WORKDIR /app

# Copy local code to the container image.
COPY . .

# Initialize a new Go module.
RUN go mod tidy

# Build the command inside the container.
RUN go build -o mwitter-backend-rest

FROM gcr.io/distroless/base-debian11

# Change the working directory.
WORKDIR /

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/mwitter-backend-rest .

ENTRYPOINT ["/mwitter-backend-rest"]

# gcloud builds submit \
#   --tag asia-northeast3-docker.pkg.dev/thematic-scene-379107/hello-repo/mwitter-backend-rest.gke .
