# Prefix Words Finder in Go
A small, somewhat interactive app with no Javascript and 100% backend HTML generation

## Stack
* Go
* GoFiber
* Pug/Jade templates
* htmx (for AJAX)
* Bulma (CSS)

## Build and run
```bash
docker build . -t prefixfinder
docker run --rm -p 3000:3000 --name prefix prefixfinder
```
Browse to http://localhost:3000
