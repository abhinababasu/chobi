var maxImageIndex
function imageClick() {
    newSrc= this.src.replace("_thumb", "");
    setMainImageByPath(newSrc)
}

// TODO: Actually fix this image sequencing!!!

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

function setMainImage(index) {
    var imgFolder = './Gallery/';
    targetImg = document.getElementById("topImage");
    targetImg.src = imgFolder + index.toString() + '.jpg';
}

function setMainImageByPath(path) {
    targetImg = document.getElementById("topImage"), 
    targetImg.src = path;
}

function setRandomPhoto(count) {
    randomIndex = Math.floor(Math.random() * count);
    setMainImage(randomIndex)
}

function setPrevPhoto() {
    randomIndex-- 
    if (randomIndex < 0) {
        randomIndex = maxImageIndex
    }
    
    setMainImage(randomIndex)
}

function setNextPhoto() {
    randomIndex++
    if (randomIndex > maxImageIndex) {
        randomIndex = 0
    }
    
    setMainImage(randomIndex)
}

function addImages(targetDiv, count){
    var imgFolder = './Gallery/';
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
        var img = document.createElement("img");
        img.src = imgFolder + indices[j].toString() + '_thumb.jpg';
        img.id = "imgThumbId";
        img.class = "imgthumb";
        
        img.onclick = imageClick;
       
        targetDiv.appendChild(img);
        var image = [];
        image.push(imgFolder + indices[j].toString() + '.jpg');
        images.push(image);
    }

    maxImageIndex = count - 1
}