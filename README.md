# Quickview

A simple quick hot-reloading image viewer in the browser.

## Why this exists

I like to create diagrams using things like PlantUML and Mermaid. Sometimes I
do this remotely on another machine, and I'd like to see my changes when the
generated image changes without having to switch to the browser and hit refresh
every time.  That annoys me enough to make this.

## How to use it

Host a specific image at `http://localhost:8083`:

```bash
# quickview <img-name> -a <address>
quickview diagrams/seq.svg -a localhost:8083
```
