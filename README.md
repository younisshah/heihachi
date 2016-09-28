# heihachi
Site Checker with Time Scheduling in Golang

## Steps
 - Configure __config.json__ file with the sites to monitor and your email settings.
 - Run the program. BOOM!

[![3592161-animepaper.net_picture_standard_video_ga.jpg](https://s9.postimg.org/xu0h64gv3/3592161_animepaper_net_picture_standard_video_ga.jpg)](https://postimg.org/image/i8j5m64wr/)

## **This is a work in progress** :construction_worker:


# TODOs
- [X] Make the ticker timeout configurable.
     ```JSON
        "every" : {
            "duration" : 1,
            "unit" : "S" // S -> Second, M -> Minutes, H -> Hours
          }
     ```
- [ ] Find a more sophisticated method to check the status of a website.
- [ ] Check whether go channels are being garbage collected.
