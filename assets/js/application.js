import "../../node_modules/normalize.css/normalize.css";
import "../css/application.sass";
import "../css/another.sass";
import Mountain from '../images/mountain.jpg';


document.addEventListener("DOMContentLoaded", () => {
    let output = document.getElementById("output")
    output.innerHTML = "Congratulations, go-webpack is working! Edit assets/js/application.js for hot reload. The following image is added by javascript";
    const element = document.createElement('div');
    const myIcon = new Image();
    myIcon.src = Mountain;
    output.appendChild(myIcon);
});