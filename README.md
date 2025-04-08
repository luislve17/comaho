Need to create the folder `/home/pc/Documents/server/media/library` beforehand

Working example of self hosted instance:

```yaml
comaho:
    image: ghcr.io/luislve17/comaho:0.1-alpha
    container_name: comaho
    environment:
      - COMAHO_DOCKER_VOLUME_PATH=/home/pc/Documents/server/media/library
    volumes:
      - type: bind
        source: /home/pc/Documents/server/media/library
        target: /app/media
    ports:
      - 9090:8080
```

Finally, going to `http://<server-ip>:9090/dashboard` should prompt you the initial page
