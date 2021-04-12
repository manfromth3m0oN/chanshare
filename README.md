# Chanshare

A Go program levearaging MPV and OpenGL to share the viewing of a 4chan thread

## Usage

To use the program as a client simply launch it.

For hosting use the -h flag when starting the program

There have been issues in windows where the version of the MPV library used has not implemented GL rendering, 
therefore it will launch as a seperate window on the loading of correct media.
The program can also take a while to load on windows, 
things like streaming on discord seem to also impact this.

The input takes this form:
* Play - to play and pause the content
* Prev & Next - Move to the next or previous video in the board
* Input 1 - Input the shortname (e.g. gif, g, b, etc) of the desired board
* Input 2 - Input the ID of the thread on the board obove
* Load thread - loads the media stream of the thread denoted by the above fields
* Volume - Adjusts the volume of the media in the player

## Sharing protocol design

Uppon launching the program in host mode, a TCP server will be started. 

The protocol will follow a request response architecture.
Each request (at the moment) can be any of:
* `skip` - request to skip the current media
* `connect` - join the hosts network, without connecting you cannot submit to skip
* `leave` - be removed from the hosts list of clients

Data will be transfered using Go's `gob` encoding format