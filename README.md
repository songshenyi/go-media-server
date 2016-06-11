# go-media-server
A simple media server by golang

publish:
./ffmpeg -re -i test.flv -c copy -f flv http://localhost:8888/live/aa.flv

play:
./ffplay http://localhost:8888/live/aa.flv