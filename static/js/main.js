let imgArray = [
    "static/img/rock_icon_v3.png",
    "static/img/paper_icon_v2.png",
    "static/img/scissors_icon_v2.png"
];


function choose(x){
    fetch("/play?c=" + x)
    .then(response => response.json())
    .then(data => {
        if (x == 0){
            document.getElementById("player_choice").innerHTML = "El jugador eligió PIEDRA.";
        }else if (x == 1){
            document.getElementById("player_choice").innerHTML = "El jugador eligió PAPEL.";
        }else{
            document.getElementById("player_choice").innerHTML = "El jugador eligió TIJERA.";
        }

        document.getElementById("player_score").innerHTML = data.player_score;
        document.getElementById("computer_score").innerHTML = data.computer_score;
       
        document.getElementById("computer_choice").innerHTML = data.computer_choice;
        
        document.getElementById("round_result").innerHTML = data.round_result;
        document.getElementById("message").innerHTML = data.message;

        var imgElement = document.getElementById("img_computer");
        imgElement.src = imgArray[data.computer_choice_int];
    })
}