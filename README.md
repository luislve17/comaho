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
        target: /app/media  # This is always '/app/media'
    ports:
      - 9090:8080  # Map any port you may use
```

Finally, going to `http://<server-ip>:9090/dashboard` should prompt you the initial page

![Screenshot_2025-04-07-20-50-35_1920x1200](https://github.com/user-attachments/assets/11e64e1b-7ed8-4d9b-ad87-94b0c510f79a)

Then, adding totally legally manga content as cbz files into your media directory:

```
pc@pcS:~/Documents/server/media/library$ tree .
.
└── Pluto
    ├── Pluto_v01.cbz
    ├── Pluto_v02.cbz
    ├── Pluto_v03.cbz
    ├── Pluto_v04.cbz
    ├── Pluto_v05.cbz
    ├── Pluto_v06.cbz
    ├── Pluto_v07.cbz
    └── Pluto_v08.cbz

2 directories, 10 files
```

Is important to be careful about the naming. Both the inner folders and the .cbz files should not have any spaces or odd non-ascii characters
You will now see your content load in the dashboard

![Screenshot_2025-04-07-22-29-51_1920x1200](https://github.com/user-attachments/assets/139381b2-5d33-4fac-b67d-ee662b440827)

And clicking inside the entry, you will get the content for that folder:
![Screenshot_2025-04-07-22-30-13_1920x1200](https://github.com/user-attachments/assets/b190c764-e1a0-4231-87dc-a3f9e014317b)

Now, if your folder's name includes the MAL id between parenthesis, like:
```
pc@pc-MINI-S:~/Documents/server/media/library$ mv Pluto/ '(MAL-745) Pluto/'
pc@pc-MINI-S:~/Documents/server/media/library$ ls
'(MAL-745) Pluto'
```
You can trigger metadata scrapping to add some more details to your content:

![Screenshot_2025-04-07-22-37-07_1920x1200](https://github.com/user-attachments/assets/5d5ae386-5fa8-4dc3-a74c-307920e9a745)

The space after the closing parenthesis is also mandatory.
From your media folder, you will notice a new file `metadata.json` that is used to render this information.
Currently there is no mechanism to re-trigger the metadata pulling, so in case you messed up the MAL's id, you just need to delete the `metadata.json`, rename your folder and navigate again into the content page of your manga, to re-trigger the pulling.
