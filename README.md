# COMAHO: Comic and Manga Hosting Manager

## Context

I have a [Kobo Libra 2](https://www.amazon.com/Kobo-Touchscreen-Waterproof-Adjustable-Temperature/dp/B09HSRGZRL), and it has a browser that has this old outdate webkit version for it. So I created a fileserver (?) that converts manga files into epub.

## Setup

First, create a folder similar to `/home/pc/Documents/server/media/library` beforehand

A `compose` working example of a self-hosted instance would look like:

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

Then, adding totally legally obtained manga content as `.cbz` files into your media directory:

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

Recently added: Checkmarks that let you see which files have already a converted version for you to click `Download`. Clicking `Convert` again will attempt a new re-trigger of the conversion of the `.cbz`

![Screenshot from 2025-05-27 22-30-25](https://github.com/user-attachments/assets/ad89762f-dfba-482e-8dcd-54a5ea2ecaa7)

Since it is really tricky to give feedback on the frontend using the Kobo's browser, please be aware that conversions are slow. When self hosting this, you may check the logs of the running container to give you an idea:
```
2025/05/28 03:35:46 Expected output file: /app/media/(MAL-745) Pluto/Pluto_v03.kepub.epub
2025/05/28 03:35:46 Attempting conversion: kcc-c2e.py -p KoL -m '/app/media/(MAL-745) Pluto/Pluto_v03.cbz' -o '/app/media/(MAL-745) Pluto/Pluto_v03.kepub.epub'
2025/05/28 03:36:30 Entered: /MAL-745-Pluto
2025/05/28 03:36:30 Metadata file for media/(MAL-745) Pluto/metadata.json already exist. Skipping API fetching.
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: MAL-745-Pluto
2025/05/28 03:36:30 Entered: /favicon.ico
Type is not MAL or ID is missing. Skipping metadata refresh.
2025/05/28 03:36:30 URL ID type and value are nil
2025/05/28 03:36:46 Conversion completed: comic2ebook v7.0.0 - Written by Ciro Mattia Gonano and Pawel Jastrzebski.
Working on /app/media/(MAL-745) Pluto/Pluto_v03.cbz...
Preparing source images...
Checking images...
Processing images...
Creating EPUB file...
```


The space after the closing parenthesis is also mandatory.
From your media folder, you will notice a new file `metadata.json` that is used to render this information.

## FAQ

> How are the updates doing? Can we have multi-threaded safe conversions parallelized on quantum realms?

I have to pay bills, please be patient on that

> Why does the UI look like poppy booppy

The Kobo Libra 2 has a really old outdated engine for its browser. I need to work with what it has, plus the usual refresh rate of a e-reader. So no fancy CSS or Javascript for us

## Credits

* [KCC (Kindle Comic Converter)](https://github.com/ciromattia/kcc)
Currently there is no mechanism to re-trigger the metadata pulling, so in case you messed up the MAL's id, you just need to delete the `metadata.json`, rename your folder and navigate again into the content page of your manga, to re-trigger the pulling.
