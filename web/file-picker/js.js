function getData(){
	alert("js works");
}

var xmlr
var currentPath

function prePath(){
	xmlr=new XMLHttpRequest();
	xmlr.onreadystatechange=callback;
	var url;
	if (currentPath!=undefined) {
		if (currentPath.slice(-1)=='/') {
			var p1=currentPath.substr(0,currentPath.lastIndexOf("/"));
			var p2=p1.substr(0,p1.lastIndexOf("/"));
			url="http://127.0.0.1:8099/file-staus"+"?path="+p2;
		}else{
			var p1=currentPath.substr(0,currentPath.lastIndexOf("/"));
			url="http://127.0.0.1:8099/file-staus"+"?path="+p1;
		}
	}else{
		url="http://127.0.0.1:8099/file-staus";
	}
	xmlr.open("GET",url,true);
	xmlr.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
	xmlr.send(null);
}

function ajaxClear(){
	currentPath=undefined;
	ajaxGet();
}

function ajaxGet(attributes){
	xmlr=new XMLHttpRequest();
	xmlr.onreadystatechange=callback;
	var url;
	if (currentPath!=undefined) {
		if (currentPath.slice(-1)=='/') {
			url="http://127.0.0.1:8099/file-staus"+"?path="+currentPath+attributes.innerHTML;
		}else{
			url="http://127.0.0.1:8099/file-staus"+"?path="+currentPath+"/"+attributes.innerHTML;
		}
	}else{
		url="http://127.0.0.1:8099/file-staus";
	}
	xmlr.open("GET",url,true);
	xmlr.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
	xmlr.send(null);
}

function callback(){
	if (xmlr.readyState==4) {
		if (xmlr.status==200) {
			var obj = JSON.parse(xmlr.responseText);
			currentPath=obj.AbsPath;
			var tablestr="<tr><td><a onclick=prePath()>..</a></td></tr>";
			tablestr+="<tr><td>name</td><td>Mb</td><td>Bytes</td></tr>";
			tablestr+="<tr><td>"+String(obj.AbsPath)+"</td><td>"+String((parseInt(obj.TotalSize)/1024/1024).toFixed(2))+"</td><td>"+String(obj.TotalSize)+"</td></tr>";	
			for (var i = 0; i <= obj.Items.length - 1; i++) {
				tablestr+="<tr>";
				if (obj.Items[i].Name.slice(-1)=='/') {
					tablestr=tablestr+"<td>"+"<a onclick=ajaxGet(this)>"+String(obj.Items[i].Name)+"</a>"+"</td>"+"<td>"+String((parseInt(obj.Items[i].Size)/1024/1024).toFixed(2))+"</td>"+"<td>"+String(obj.Items[i].Size)+"</td>";
				}else{
					tablestr=tablestr+"<td>"+String(obj.Items[i].Name)+"</td>"+"<td>"+String((parseInt(obj.Items[i].Size)/1024/1024).toFixed(2))+"</td>"+"<td>"+String(obj.Items[i].Size)+"</td>";
				}
				tablestr+="</tr>";
			}

			document.getElementById("get-response").innerHTML=tablestr;
		}
	}
}
