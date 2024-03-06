# WatcherRPC
The Discord RPC client that watches the running processes!

![image](https://github.com/Endg4meZer0/WatcherRPC/assets/65147948/ffcf050f-b20c-432a-8921-2f60cd4d5b50)

Right now it only works on Windows because it's using the `tasklist` command. Later it will be changed to support other systems too.

## What is this?
Just a custom Discord RPC that watches over processes listed in `processList.json`, and if one exists, then puts the data (state, details, etc.) stated in the file to the game activity in your Discord.
It requires you to create your own app [here](https://discord.com/developers) and put the client id in the `.env` file. Note that the name of the app will be displayed at the very top of game activity, as well as in your status `(Playing ...)`

![image](https://github.com/Endg4meZer0/WatcherRPC/assets/65147948/bb177125-f83c-4242-b757-f1aa60488599)

## How can I edit stuff?
Check `processList.json`! Note that for images to work you also need to upload them at Discord Developers Portal in your app in Rich Presence's Art Assets menu, then you'll be able to use the image keys to actually set images for the game activity.
```json
"processName": "StarRail.exe", // the process to watch for
"state": "while listening to cool music", // for more info on the position
"details": "building warrior Stelle", // check the screenshot at the top
"largeImageKey": "stelle",
"largeImageText": "she's pretty",
"smallImageKey": "",
"smallImageText": "",
"buttons": [{
    "text": "check out cool music",
    "url": "https://open.spotify.com/playlist/2Yt8ACA7NOwOcnaDtAD2IH?si=49ac5b8de98f453e"
}],
"useTimestamp": false // if true, will set a start timestamp on process detection
```

## TODOs
- [ ] A better interface than just an open console
- [ ] Crossplatform solutions
- [ ] More functionality maybe?
