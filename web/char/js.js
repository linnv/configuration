// 基于准备好的dom，初始化echarts实例
var myChart = echarts.init(document.getElementById('echar'));

// 指定图表的配置项和数据
var option = {
	title: {
		text: 'ECharts dynamic demo'
	},
	tooltip: {},
	legend: {
		data:['sale']
	},
	xAxis: {
		data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
	},
	yAxis: {},
	series: [{
		name: '销量',
		type: 'bar',
		data: [5, 20, 36, 10, 10, 20]
	}]
};

myChart.setOption(option);

var counter
var interval=500
function executor(){
	counter= new Date();
	var a=Math.floor((Math.random() * 100) + 1);
	// 使用刚指定的配置项和数据显示图表。
	myChart.setOption(
			{
				series:[{
					name: '销量',
					type: 'bar',
					data: [5, a*1.5, 36, 10, 10, a]
					// data: [a, a, a, a, a+10, a]
				}]}
			);
	document.getElementById("interval").innerHTML=counter.toTimeString();
}

var intervalFunction=function(){
	executor();

	setTimeout(intervalFunction,interval);
}

setTimeout(intervalFunction,interval);
