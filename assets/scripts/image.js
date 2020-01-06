var maxImageCount;      // maximum number of images on this page
var allPhotos = [];     // array of all the photos
var currentIndex = 0;   // index into allPhotos that is being displayed

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

function setCurrentImage() {
    targetImg = document.getElementById(topImageId), 
    targetImg.src = allPhotos[currentIndex].fullPath;
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
}