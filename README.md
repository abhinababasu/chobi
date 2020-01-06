# chobi 
## Introduction
A static photo gallery generator. chobi is lansliteration of "ছবি" which means picture in Bengali/Bangla.

## Build
Get the sources and then build using
```
go build .
```

## Run
The usage is
```
chobi <name> <image folder full path> <out path> <thumb-size>
```

Example
```
chobi Landscape D:\SourceImages d:\Upload 150
```

## Sample
A sample gallery
[Photography](http://bonggeek.com/Photography/)

## TODO
Template
1. Main image should better expand to fill screen
1. Main image should be centered vertically
1. Show a count down animation before changing to next pic automatically
1. Image transitions

Generation
1. Tumbnails for portrait should be not from the center (cmd line option)
1. Proper cmd line parsing
1. Make it work for mac as well
