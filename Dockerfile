FROM registry.access.redhat.com/ubi8/go-toolset:latest AS builder
WORKDIR /app
COPY  --chown=1001:0 go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o buildconfig-server .



# Final stage
FROM registry.access.redhat.com/ubi8/ubi-minimal

# Install wget for health checks
RUN microdnf install -y wget ca-certificates && \
    microdnf clean all

# Create non-root user
# RUN useradd -r -u 1001 -g root -m -d /app appuser

WORKDIR /app
COPY --from=builder /app/buildconfig-server .
COPY --from=builder /app/crc-linux-amd64.tar.xz* ./
RUN chown -R appuser:root /app
USER appuser
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1
CMD ["./buildconfig-server"]
