# --- Stage 1: Get official image of KCC
FROM ghcr.io/ciromattia/kcc:7.0.0 AS kcc-container
# ------------------------------------------------

# --- Stage 2: Build the Go application (Comaho)
FROM golang:1.22.8-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN go mod tidy
RUN go build -o comaho .
# ------------------------------------------------

# --- Stage 3: Final image to run the application with KCC
FROM alpine:latest

WORKDIR /app

# Install needed python dependencies 
COPY py-requirements.txt /app/py-requirements.txt

# Install dependencies
RUN apk --no-cache add git wget build-base gcc python3-dev musl-dev linux-headers ca-certificates py3-pip libpng libjpeg-turbo p7zip && \
    pip3 install --break-system-packages -r /app/py-requirements.txt && \
    wget -qO- https://archive.org/download/kindlegen_linux_2_6_i386_v2_9/kindlegen_linux_2.6_i386_v2_9.tar.gz | tar -xz && \
    mv kindlegen /usr/local/bin/kindlegen && \
    chmod +x /usr/local/bin/kindlegen && \
    rm -r *.txt docs manual.html

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/src/comaho .

# Copy the frontend templates (index.html)
COPY --from=builder /app/src/templates /app/templates

# Get scripts from KCC project (https://github.com/ciromattia/kcc)
COPY --from=kcc-container /opt/kcc/kcc-c2e.py /usr/local/bin/kcc-c2e.py
COPY --from=kcc-container /opt/kcc/kcc-c2p.py /usr/local/bin/kcc-c2p.py
COPY --from=kcc-container /opt/kcc/kindlecomicconverter /usr/local/bin/kindlecomicconverter

EXPOSE 8080
CMD ["./comaho"]
# ------------------------------------------------

