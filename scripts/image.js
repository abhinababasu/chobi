var maxImageCount;      // maximum number of images on this page
var allPhotos = [];     // array of all the photos
var currentIndex = 0;   // index into allPhotos that is being displayed
var transitionTimer;
const tansitionInterval = 5000;

// html elements manipulated by this script
const imageIdPrefix = 'img_';
const topImageId = 'topImage';
const prevDiv = 'previous';
const nextDiv = 'next';
const imgThumbClass = 'imgThumb';
const imageThumbFileNameSuffix = '_thumb.jpg';
const imageFileExtension = '.jpg';
const imageThumbClass = 'imgThumb';

// Photo class with path of thumbnail and full resolution
class Photo {
    constructor(thumb, full) {
        this.thumbPath = thumb;
        this.fullPath = full;
    }
}

// shuffle an array randomly
function shuffle(array) {
    var currentIndex = array.length, temporaryValue, randomIndex;
    // While there remain elements to shuffle...
    while (0 !== currentIndex) {
        // Pick a remaining element...
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex -= 1;
        // And swap it with the current element.
        temporaryValue = array[currentIndex];
        array[currentIndex] = array[randomIndex];
        array[randomIndex] = temporaryValue;
    }
    
    return array;
}

// Image navigation on the page
function imageClick() {
    index = Number(this.id.replace(imageIdPrefix,""));
    currentIndex = index
    setCurrentImage()
}

function imageonLoad() {
    document.getElementById("loadingDiv").style.display = "none";
}

function setCurrentImage() {
    document.getElementById("loadingDiv").style.display = "block";
    targetImg = document.getElementById(topImageId);
    targetImg.onload = imageonLoad;
    targetImg.src = allPhotos[currentIndex].fullPath;

    // reset counting down for next change. This ensures a periodic tick doesn't change a 
    // photo immediately after user selects a new one directly
    clearInterval(transitionTimer);
    transitionTimer = setTimeout(setNextPhoto, tansitionInterval);
}

function setPrevPhoto() {
    currentIndex-- 
    if (currentIndex < 0) {
        currentIndex = maxImageCount - 1
    }
    
    setCurrentImage()
}

function setNextPhoto() {
    currentIndex = (currentIndex+1) % maxImageCount

    setCurrentImage()
}

function keyHandler(e) {
    // left arrow or "p"
    if (e.keyCode == 37 || e.keyCode == 80) {
        setPrevPhoto();
    } 
    // right arrow or "n"
    else if (e.keyCode == 39 || e.keyCode == 78) {
        setNextPhoto();
    }
}

// Add all images into the page
function addImages(name, targetDiv, count){
    var imgFolder =  './' + name + '/';
    // fill an array with 0..count-1
    var indices = [];
    for (var i = 0; i < count; i++) {
        indices.push(i);
    }
    

    // shuffle the array so that the numbers 0..count-1 is random in the array
    indices = shuffle(indices);

    // loop through the random array and crate image objects
    var images = [];
    for(var j = 0; j < indices.length; j++){

        // Create paths and store them
        photoThumbPath = imgFolder + indices[j].toString() + imageThumbFileNameSuffix;
        photoFullPath = imgFolder + indices[j].toString() + imageFileExtension;
        photo = new Photo(photoThumbPath, photoFullPath)
        allPhotos.push(photo)

        // Create and show the thumbnail
        var img = document.createElement("img");
        img.src = photoThumbPath;
        // NOTE! the id is parsed later for finding image index for navigation. Don't change without changing that
        img.id = imageIdPrefix + j;
        img.className = imageThumbClass;
                
        img.onclick = imageClick;
       
        targetDiv.appendChild(img);
    }

    maxImageCount = count;
    setCurrentImage();

    // setup handlers
    document.body.onkeydown = keyHandler;
    document.getElementById(prevDiv).onclick = setPrevPhoto;
    document.getElementById(nextDiv).onclick = setNextPhoto;

    document.getElementById("thumbnailScrollRight").onmouseover = scrollRight;
    document.getElementById("thumbnailScrollRight").onmouseout = scrollStop;
    document.getElementById("thumbnailScrollRight").onclick = scrollStepRight;
    
    document.getElementById("thumbnailScrollLeft").onmouseover = scrollLeft;
    document.getElementById("thumbnailScrollLeft").onmouseout = scrollStop;
    document.getElementById("thumbnailScrollLeft").onclick = scrollStepLeft;
}

var scrolling = false;

function scrollRight() {
    scrolling = true;
    scrollContent("right");
}

function scrollStepLeft() {
    $("#thumbnailDiv").animate({
        scrollLeft: "-=650px"
    }, 1, function(){});
}


function scrollStepRight() {
    $("#thumbnailDiv").animate({
        scrollLeft: "+=650px"
    }, 1, function(){});
}

function scrollLeft() {
    scrolling = true;
    scrollContent("left");
}

function scrollStop() {
    scrolling = false;
}


function scrollContent(direction) {
    var amount = (direction === "left" ? "-=5px" : "+=5px");
    $("#thumbnailDiv").animate({
        scrollLeft: amount
    }, 1, function() {
        if (scrolling) {
            // If we want to keep scrolling, call the scrollContent function again:
            scrollContent(direction);
        }
    });
}