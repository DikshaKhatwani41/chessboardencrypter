

function display(text){
    curIndex=0; // current index of my text

    for (i=0;i<8;i++){
        for(j=0;j<8;j++){
             c = text.charAt(curIndex);
             if(text.charCodeAt(curIndex)>=49 && text.charCodeAt(curIndex)<=57){
                 j+= (text.charCodeAt(curIndex)-48); // 49-48 = 1, 50-48=2
                 curIndex++;
                 
             }
             switch (c) {
                case 'r':
                    document.getElementById(""+i+j).innerHTML="&#9820;";
                    
                    curIndex++;
                    break;
                case 'n':
                    document.getElementById(""+i+j).innerHTML="&#9822;";
                    curIndex++;
                    break;
                case 'b':
                    document.getElementById(""+i+j).innerHTML="&#9821;";
                    curIndex++;
                    break;
                case 'q':
                    document.getElementById(""+i+j).innerHTML="&#9819;";
                    curIndex++;
                    break;
                case 'k':
                    document.getElementById(""+i+j).innerHTML="&#9818;";
                    curIndex++;
                    break;
                case 'p':
                    document.getElementById(""+i+j).innerHTML="&#9823;";
                    curIndex++;
                    break;
                case 'R':
                    document.getElementById(""+i+j).innerHTML="&#9814;";
                    curIndex++;
                    break;
                case 'N':
                    document.getElementById(""+i+j).innerHTML="&#9816;";
                    curIndex++;
                    break;
                case 'B':
                    document.getElementById(""+i+j).innerHTML="&#9815;";
                    curIndex++;
                    break;
                case 'Q':
                    document.getElementById(""+i+j).innerHTML="&#9813;";
                    curIndex++;
                    break;
                case 'K':
                    document.getElementById(""+i+j).innerHTML="&#9812;";
                    curIndex++;
                    break;
                case 'P':
                    document.getElementById(""+i+j).innerHTML="&#9817;";
                    curIndex++;
                    break;
                
             }
             
             
        }
        curIndex++; //ignoring /
        
    }
}

function encodestr(str){
    fetch("/encode?str="+str)
    .then(response => response.json())
    .then(data => {
        console.log(data.fen);
        document.getElementById("fen").innerHTML=data.fen;
        lbl=document.createElement('label').innerHTML=""+data.player_choice;
    
        document.appendChild(lbl);
        
    })
    
}