function getData(){
	alert("js works");
}

function ajaxClear(){
			document.getElementById("get-response").innerHTML="";
}

var xmlr
function ajaxGet(){
	xmlr=new XMLHttpRequest();
	xmlr.onreadystatechange=callback;
	// xmlr.open("GET","http://127.0.0.1:8099/get",false);
	// async
	xmlr.open("GET","http://127.0.0.1:8099/get",true);
	xmlr.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
	xmlr.send(null);
}

function callback(){
	if (xmlr.readyState==4) {
		if (xmlr.status==200) {

			var obj = JSON.parse(xmlr.responseText);
			var tablestr="<tr><td>id</td><td>name</td></tr>";
			for (var i in obj) {
				tablestr+="<tr>";
				tablestr=tablestr+"<td>"+String(obj[i].id)+"</td>"+"<td>"+String(obj[i].name)+"</td>";
				tablestr+="</tr>";
			}

			document.getElementById("get-response").innerHTML=tablestr;
		}
	}
}


