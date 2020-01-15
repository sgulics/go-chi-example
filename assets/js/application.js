import "../../node_modules/normalize.css/normalize.css";
import "../css/application.sass";
import "../css/another.sass";
import Mountain from '../images/mountain.jpg';


document.addEventListener("DOMContentLoaded", () => {
    console.log("Congratulations, go-webpack is working!");
    document.getElementById("output").innerHTML = "Congratulations, go-webpack is working!";
    const element = document.createElement('div');
    const myIcon = new Image();
    myIcon.src = Mountain;
    element.appendChild(myIcon);
});