Need to create the folder `/home/pc/Documents/server/media/library` beforehand

Working example of self hosted instance:

```yaml
comaho:
    image: ghcr.io/luislve17/comaho:0.1-alpha
    container_name: comaho
    environment:
      - COMAHO_DOCKER_VOLUME_PATH=/home/pc/Documents/server/media/library  # Any proper existent directory should work
    volumes:
      - type: bind
        source: /home/pc/Documents/server/media/library  # Map from the directory, to the required app->media lookup
        target: /app/media
    ports:
      - 9090:8080
```

Finally, going to `http://<server-ip>:9090/dashboard` should prompt you the initial page
