var maxImageCount;      // maximum number of images on this page
var allPhotos = [];     // array of all the photos
var currentIndex = 0;   // index into allPhotos that is being displayed

// Photo class with path of thumbnail and full resolution
class Photo {
    constructor(thumb, full) {
        this.thumbPath = thumb;
        this.fullPath = full;
    }
}

// helped to shuffle an array randomly
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
    index = Number(this.id.replace("img_",""));
    currentIndex = index
    setCurrentImage()
}

function setCurrentImage() {
    targetImg = document.getElementById("topImage"), 
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

function imageLoaded() {
    l = this.getBoundingClientRect().left;
    l = (Math.round(l)) + "px";
    document.getElementById("prevDiv").style.left = l;

    r = this.getBoundingClientRect().right;
    
    nextDiv = document.getElementById("nextDiv");
    
    l = r - nextDiv.getBoundingClientRect().width;
    l = Math.round(l) - 2;
    l = l + "px";
    nextDiv.style.left = l;
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
        photoThumbPath = imgFolder + indices[j].toString() + '_thumb.jpg';
        photoFullPath = imgFolder + indices[j].toString() + '.jpg';
        photo = new Photo(photoThumbPath, photoFullPath)
        allPhotos.push(photo)

        // Create and show the thumbnail
        var img = document.createElement("img");
        img.src = photoThumbPath;
        // NOTE! the id is parsed later for finding image index for navigation. Don't change without changing that
        img.id = "img_" + j;
        img.className = "imgThumb";
                
        img.onclick = imageClick;
       
        targetDiv.appendChild(img);
    }

    maxImageCount = count;
    setCurrentImage();

    // setup handlers
    document.body.onkeydown = keyHandler;
    document.getElementById("topImage").onload = imageLoaded;
    document.getElementById("previous").onclick = setPrevPhoto;
    document.getElementById("next").onclick = setNextPhoto;
}