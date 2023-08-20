# Design

Quick design sketch for reference.

## What it needs to do

Given a filename, serve the file on an HTML page at 100% zoom (configurable
later).

When open, the page should open a websocket with the server and wait for any
updates.

When the file changes, the server should notify any listening websocket
connections which will trigger a refresh on the page.

## Components to build

### HTML page with image and script

Some templated HTML page which has an image element.

It needs to include a script that opens a websocket connection with the server
on the handler described below. When it receives a notification, it refreshes
the page automatically.

### Server that serves both template and actual image

File system server with just the desired image with the file name, and `/`
sends the template.

### File watcher

A long-running process that watches the target file for changes.

### Websocket handler

A websocket handler that waits for messages from the file watcher at `/watch`

## Checkpoints

### Static serve

Create a simple HTML page that shows the image statically. No hot reloading yet.

This is done when:

- [x] `quickview <img>` starts a server
- [x] The page can be viewed from browser
- [x] The image is visible on the page

### Hot reload

Add a websocket endpoint that emits a message on file modifications to reload the image.

This is done when:

- [x] The page opens a websocket connection successfully
- [x] The server emits an updated unix nano timestamp on the websocket when the file is updated
- [x] The page updates the image source using the nano timestamp to reload the image
