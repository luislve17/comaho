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

# Install runtime dependencies (Python, libraries required by KCC tools)
RUN apk --no-cache add git wget build-base gcc python3-dev musl-dev linux-headers ca-certificates python3 py3-pip libpng libjpeg-turbo p7zip

WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/src/comaho .

# Copy the frontend templates (index.html)
COPY --from=builder /app/src/templates /app/templates

# Install needed python dependencies 
RUN pip3 install --break-system-packages \
    certifi==2024.8.30 \
    cffi==1.17.1 \
    charset-normalizer==3.4.0 \
    distro==1.9.0 \
    fastnumbers==5.1.0 \
    idna==3.10 \
    mozjpeg-lossless-optimization==1.1.5 \
    natsort==8.4.0 \
    numpy==1.26.4 \
    packaging==24.0 \
    pillow==11.0.0 \
    psutil==6.1.0 \
    pycparser==2.22 \
    pyparsing==3.1.2 \
    python-slugify==8.0.4 \
    raven==6.10.0 \
    requests==2.32.3 \
    setuptools==70.3.0 \
    text-unidecode==1.3 \
    urllib3==2.2.3

# Download and install KindleGen
RUN wget -qO- https://archive.org/download/kindlegen_linux_2_6_i386_v2_9/kindlegen_linux_2.6_i386_v2_9.tar.gz | tar -xz \
    && mv kindlegen /usr/local/bin/kindlegen \
    && chmod +x /usr/local/bin/kindlegen
RUN rm -r *.txt docs manual.html

# Get scripts from KCC project (https://github.com/ciromattia/kcc)
COPY --from=kcc-container /opt/kcc/kcc-c2e.py /usr/local/bin/kcc-c2e.py
COPY --from=kcc-container /opt/kcc/kcc-c2p.py /usr/local/bin/kcc-c2p.py
COPY --from=kcc-container /opt/kcc/kindlecomicconverter /usr/local/bin/kindlecomicconverter

EXPOSE 8080
CMD ["./comaho"]
# ------------------------------------------------

