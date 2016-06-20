# go-media-server
A simple media server by golang

publish:
./ffmpeg -re -i test.flv -c copy -f flv http://localhost:8888/live/aa.flv

play:
./ffplay http://localhost:8888/live/aa.flv

reference:
go-oryx: https://github.com/ossrs/go-oryx

nginx-rtmp-module: https://github.com/arut/nginx-rtmp-module

simple-rtmp-server: https://github.com/ossrs/srs / https://github.com/wenjiegit/srs

ffmpeg: https://github.com/FFmpeg/FFmpeg