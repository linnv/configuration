function getData(){
	alert("js works");
}

var xmlr
var ga
var currentPath
function ajaxClear(){
			document.getElementById("get-response").innerHTML="";
			currentPath=undefined;
}

function ajaxGet(attributes){
	ga =attributes
	xmlr=new XMLHttpRequest();
	xmlr.onreadystatechange=callback;
	// xmlr.open("GET","http://127.0.0.1:8099/get",false);
	// async
	if (currentPath!=undefined) {
	var url="http://127.0.0.1:8099/file-staus"+"?path="+currentPath+"/"+ga.innerHTML;
	// var url="http://127.0.0.1:8099/file-staus";
	}else{
	var url="http://127.0.0.1:8099/file-staus";
	}
	console.log(url);
	xmlr.open("GET",url,true);
	xmlr.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
	xmlr.send(null);
}

function callback(){
	if (xmlr.readyState==4) {
		if (xmlr.status==200) {
			// console.log(ga.innerHTML);
			var obj = JSON.parse(xmlr.responseText);
			currentPath=obj.AbsPath;
			var tablestr="<tr><td>name</td><td>Mb</td><td>Bytes</td></tr>";
			// tablestr+="<tr><td>"+String(obj.AbsPath)+"</td><td>"+String((parseInt(obj.TotalSize)/1024/1024).toFixed(2))+"</td><td>"+String(obj.TotalSize)+"</td></tr>";	
			tablestr+="<tr><td>"+String(obj.AbsPath)+"</td><td>"+String((parseInt(obj.TotalSize)/1024/1024).toFixed(2))+"</td><td>"+String(obj.TotalSize)+"</td></tr>";	
			for (var i = 0; i <= obj.Items.length - 1; i++) {
				tablestr+="<tr>";
				tablestr=tablestr+"<td>"+"<a onclick=ajaxGet(this)>"+String(obj.Items[i].Name)+"</a>"+"</td>"+"<td>"+String((parseInt(obj.Items[i].Size)/1024/1024).toFixed(2))+"</td>"+"<td>"+String(obj.Items[i].Size)+"</td>";
				// tablestr=tablestr+"<td>"+String(obj.Items[i].Name)+"</td>"+"<td>"+String((parseInt(obj.Items[i].Size)/1024/1024).toFixed(2))+"</td>"+"<td>"+String(obj.Items[i].Size)+"</td>";
				tablestr+="</tr>";
			
			// }
			// for (var i in obj.Items) {
				// console.log(obj[i]);
				// tablestr+="<tr>";
				// tablestr=tablestr+"<td>"+String(obj[i].Name)+"</td>"+"<td>"+String(obj[i].Size)+"</td>";
				// tablestr+="</tr>";
			}

			document.getElementById("get-response").innerHTML=tablestr;
		}
	}
}


