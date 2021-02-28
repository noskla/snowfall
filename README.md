# ❄️ Snowfall

### Project is abandoned
Deadlines were way too short and had very little time to spend on this project so it will stay archived indefinitely
and most likely won't be developed anymore in this form on this repository.

### About
A project made for "Snowfall Ponymeet" event performed online in 2021 by Bronies Silesia.

Building on Windows:
```
git clone git://github.com/noskla/snowfall.git
cd .\snowfall\
.\build.ps1
```
Building on Linux/FreeBSD:
```
git clone git://github.com/noskla/snowfall.git
cd snowfall
chmod +x build.sh
./build.sh
```
Building and deploying a Docker image:
```
git clone git://github.com/noskla/snowfall.git
cd snowfall
chmod +x build.sh
./build.sh
docker build -t snowfall .
docker run --name snowfall -p 20000:20000 \
    -v /path/to/config.json:/data/config.json \
    -d snowfall
```
