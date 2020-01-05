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

## TODO
Template
1. Fix moving location of prev next button
1. Right size the top image
1. Image transitions

Generation
1. Tumbnails for portrait should be not from the center (cmd line option)
1. Proper cmd line parsing
1. Make it work for mac as well

