
<?php
namespace App\Modules\Qcadmin\Services;

use App\Services\Services;
use App\Libraries\Encrypt;
use App\Libraries\Dataapi;
use App\Services\Service\Es;
use App\Libraries\BackendApi;
use App\Libraries\StringHelper;
use App\Models\Entities\UserModel;
use App\Modules\Qcadmin\Models\QcconditionModel;
use App\Modules\Qcadmin\Services\Intention;
use App\Modules\Qcadmin\Models\OperatorQcLogModel;

require APP_PATH.'../vendor/autoload.php';

use Elasticsearch\Client as ElasticsearchClient;
use Elasticsearch\ClientBuilder;

/**
 *
 * @author linklin
 *
 */
class CommonService
{
	public $es_task_index = 'taskindex';
	public $es_task_index_type = 'task';
	public $es_qctask_index = 'qc_index';
	public $es_qctask_index_type = 'qc_search';

	/**
	 * 工单表：标签 ID 和字段对应关系
	 * @var arary
	 */
	public $task_int = [
			'1852'=>'city',
			'1854'=>'province',
			'1880'=>'task_id',
			'1882'=>'service_channel',
			'1884'=>'calling_number',
			'1886'=>'userid',
			'1958'=>'carrier_id',
			'2120'=>'archive_path',
			'2130'=>'create_time',
			'2132'=>'archive_time',
			'2298'=>'owner',
			'2856'=>'type',
			'2866'=>'first_dir_code',
			'3450'=>'service_id',
			'3766'=>'is_handled',
			'4274'=>'service_vip',
			'5074'=>'call_id',
			'5075'=>'mobile',
			'5076'=>'contact_phone',
			'5078'=>'problemqq',
			'5079'=>'weixin_id',
			'5080'=>'openid',
			'5081_noindex'=>'title',
			'5082_noindex'=>'intro',
			'5083_noindex'=>'content',
			'5084_noindex'=>'reply_content',
			'5085'=>'called_id',
			'5089'=>'create_user',
			'5091'=>'reply_id',
			'5092'=>'owner_group',
			'5093'=>'owner_group_name',
			'5094'=>'dest_group',
			'5095'=>'vip',
			'5096'=>'record_time',
			'5098'=>'call_time',
			'5099'=>'access_times',
			'5100'=>'path',
			'5101'=>'pathname',
			'5104'=>'last_reply',
			'5105'=>'reply_type',
			'5106'=>'related_task',
			'5110'=>'begin_time',
			'5112'=>'resp_time',
			'5113'=>'end_time',
			'5114'=>'upgrade_time',
			'5117'=>'access_time',
			'5122'=>'sync_type',
			'5123'=>'second_dir_code',
			'5134'=>'descrip',
			'5135'=>'satis_level',
			'5138'=>'archivepath_name',
			'5141'=>'archive_type',
			'5142'=>'record_id',
			'5148'=>'archive_operator',
			'5149'=>'group_reply',
			'5195'=>'urge_task',
			'5196'=>'urge_time',
			'5374'=>'exp_time',
			'5420'=>'task_label',
	];

	/**
	 * 产品 code 和name对应数组
	 * @var array
	 */
	public $product_names = [];

	/**
	 * 员工所在的中心、大组、小组
	 * @var array
	 */
	public $staff_groups = [];

	/**
	 * 工单字段对应标签信息
	 * @var array
	 */
	public $task_field_tags = [];

	/**
	 * 工单标签信息
	 * @var array
	 */
	public $task_tags = [];

	public function __construct(){
		$this->task_int = array_flip($this->task_int);
	}

	/**
	 * 分页查询数据
	 * @param array $es_cfg    es 配置：host、port
	 * @param array $where     查询条件
	 * @param number $page
	 * @param number $limit
	 * @param string $index   es 库索引名
	 * @param string $type    es 库索引对应的 type
	 * @return array
	 */
	public function paginate($es_cfg,$where,$page=0,$limit=15,$index='qc_index',$type='qc_search'){
		$ret = Services::getService('Es',$es_cfg)->paginate($index,$type,$where,'',$page,$limit);
		return $ret;
	}

	/**
	 * 抓取工单
	 * 1、随机抓取
	 * 2、每次抓取量
	 * 3、平均被质检人数
	 * 4、抓取数据临时的
     * @param array $es_cfg   es 配置：host、端口
     * @param array $where    查询条件
     *              例：$where['check_status'] = ['operator'=>'term','value'=>'0'];
     * @param int $limit   每次抓取量
     * @param string $loginName   当前登录用户
     * @return array
	 */
	public function fetchTask($es_cfg,$where = [],$limit = 15,$loginName){

		$es = Services::getService('Es',$es_cfg);

		// 查询是否还有未质检的单据


			if(empty($where)){
				$where['check_status'] = ['operator'=>'term','value'=>'0'];
			}
			//$where['service_channel'] = ['operator'=>'=','value'=>'IVR'];
			//$where['content'] = ['operator'=>'!=','value'=>''];
			//$where['ai_status'] = ['operator'=>'=','value'=>'1'];
			//$where['archive_time'] = ['operator'=>'range','value'=>['sdate'=>'2018-03-01','edate'=>'2018-03-05']];
			//$where['task_id'] = ['operator'=>'=','value'=>'18032109064787846074'];

			$where = [];
			$body = [];
			if($limit > 50){
				$limit = 50;
			}
			$body['size'] = $limit*50;// 这里主要是随机多点数据，当部分数据不合格时，可以用剩下的数据
			$body['from'] = 0;
			$body['query'] = [];
			$query = $es->setBobyByWhere($where);
			// 随机查询数据
	 		$body['query']['function_score'] = [
					'query'=>$query,
					"functions"=>[
					[
						"random_score"=>[
							"seed"=>rand(1,999999)
						]
					]
					],
					"score_mode"=> "sum"
			];


			$ret = $es->search($this->es_qctask_index,$this->es_qctask_index_type,$body);
			print_r($ret);die();

		$result = [];
		if(!empty($ret['hits']['hits'])){
			$i = 0;
			foreach ($ret['hits']['hits'] as $row){
				if($i > ($limit-1)){
					break;
				}
				$i ++;

				// 查询被质检人是否已经达到设定的阀值
				$result[] = $row['_source'];
			}
		}
		if(!empty($result)){
			$result = $this->packageQcTask($result,'',$es_cfg);
			foreach ($result as $key=>$val){
				// 更新质检人
				$this->updateQcIndexField($es_cfg, ['operator'=>$loginName,'task_id'=>$val['task_id']]);
			}
		}
		return $result;
	}

	/**
	 * 二次处理工单字段内容
	 * @param array $data
	 */
	public function packageQcTask($data,$user='',$es_cfg=''){
		$feature_cfg = [2=>"升级单",4=>"重复单",8=>"不满意单",16=>"超时单"];
		foreach ($data as $key=>$val){

			if(!empty($val['quality_label'])){
				$val['quality_label'] = json_decode($val['quality_label'],true);
			}
			if(!empty($val['tactics'])){
				$tactics = json_decode($val['tactics'],true);
				if(!empty($tactics) && ($val['service_channel'] == 'IVR' || $val['service_channel'] == 'IMC')){
					$tmp = [];
					foreach ($tactics as $t){
						if(isset($t['second_dir'])){///////////////////////////////////////
							$tmp[$t['second_dir']] = $t;
						}
					}
					$tactics = array_values($tmp);
				}
				$val['tactics'] = $tactics;
			}
			if($val['service_channel'] == 'IVR' || $val['service_channel'] == '外呼渠道' || $val['service_channel'] == 'IPCC' || $val['service_channel'] == '电话'){

				$val['decode_content'] = json_decode($val['content']);
				if($val['decode_content'] != ''){
					$val['content'] = $val['decode_content'];
				}
				//$url = 'http://recdl.cm.com/voc_record_listen.php?from=newivr&recordId=' . $record_id;
				$url = 'http://recdl.cm.com/kf_record_listen.php?from=newivr&recordId=' . str_replace('\\','*',$val['record_id']);
				$val['record_url'] = $url;
			}else{
				$val['content'] = $this->replace_hongbao_imgurl($val['content']);
				$val['reply_content'] = $this->replace_hongbao_imgurl($val['reply_content']);
			}
			if(!empty($val['intro'])){
				$val['intro'] = html_entity_decode($val['intro']);
			}
			if($val['service_channel'] == 'IMC'){
				//$val['content'] = $this->imcContentFormat($result[$key]['content']);
				$val['content'] = $this->imcContentFormat($val['content']);
			}
			// 特征转
			if(isset($val['feature']) && !empty($val['feature'])){
				$val['feature_zh'] = '';
				foreach ($val['feature'] as $feature){
					$val['feature_zh']  .= isset($feature_cfg[$feature])?$feature_cfg[$feature]."  ":'';
				}
			}
			$val['create_time'] = "20".substr($val['task_id'], 0,2)."-".substr($val['task_id'], 2,2)."-".substr($val['task_id'], 4,2)." ".substr($val['task_id'], 6,2).":".substr($val['task_id'], 8,2).":".substr($val['task_id'], 10,2);
			// 看看是不是转单
			// 统计转单的
			if(isset($val['operator_status']) && $val['operator_status'] == '10'){
				$OperatorQcLogModel = new OperatorQcLogModel();
				$where = [];
				$where['conditions']['dest_user']['operator'] = "=";
				$where['conditions']['dest_user']['value'] = $user;
				$where['conditions']['task_id']['operator'] = "=";
				$where['conditions']['task_id']['value'] = $val['task_id'];
				$where['conditions']['operator_type']['operator'] = "=";
				$where['conditions']['operator_type']['value'] = '10';
				$ret_move = $OperatorQcLogModel->fetchOne($where);
				if(!empty($ret_move)){
					$val['move'] = $ret_move->operator;
				}
			}

			$data[$key] = $val;
		}
		return $data;
	}


	/**
	 * 查询工单的流转记录
	 */
	public function getTaskLog($es_cfg,$task_id,$method){
		//$where['method'] = ['operator'=>'=','value'=>$method];     method 没加索引
		$body['query']['bool']['must'][]['term']['task_id'] = $task_id;
		$ret = Services::getService('Es',$es_cfg)->search("tasklogindex", "log",$body);
		$encrypt = new Encrypt();
		$content = '';
		if(!empty($ret)){// 查不到这个工单认定风险
			if(isset($ret['hits']['hits'])){
				$hits = $ret['hits']['hits'];
				if(!empty($hits)){
					foreach ($hits as $val){
						if(isset($val['_source'])){
							$_source = $val['_source'];
							if(isset($_source['method']) && $_source['method'] == 'batch_archive_task'){
								$content = $encrypt->decryptBase64($_source['content'], _AES_ECB_KEY_);
							}
						}

					}
				}
			}
		}
		return $content;
	}

	/**
	 * 根据指定人获取配置的条件拉取工单
	 * @param string $username
	 */
	public function fetchOrderByCondition($es_cfg,$username,$limit = 15,$isGet = FALSE){
		// 添加是否拉取的逻辑
		date_default_timezone_set('PRC');
		$es = Services::getService('Es',$es_cfg);

			// 超过限制单量的人排除
			/* $full_user = $this->statsDefendantUserCount($es_cfg);
			if(!empty($full_user)){
				$query['query']['bool']['must_not']['terms']['owner'] = $full_user;
			} */
			//var_dump(json_encode($query));
			//exit;
			// $body = [];
			// $$body['body']['query']['match']['task_id'] = '18011510135186083430';
			// $$body['query']['match']['task_id'] = '18011510135186083430';
		$hosts = [
    '47.95.6.82:9200'
];
$client = ClientBuilder::create()           // Instantiate a new ClientBuilder
                    ->setHosts($hosts)      // Set the hosts
                    ->build();              // Build the client object
$params = [
    'index' => $this->es_qctask_index,
    'type' => $this->es_qctask_index_type,
    'body' => [
        'query' => [
            'match' => [
                // '_id' => '18011510135186083430'
                '_id' => '19031509173581056654'
            ]
        ]
    ]
];

$ret = $client->search($params);
		// $body = [
		// 	'task_id'=>'18011510135186083430'
		// ];
// return $results['hits']['hits'][0];

			if(!empty($ret['hits']['hits'])){
				$i = 0;
				foreach ($ret['hits']['hits'] as $row){
					if($i > ($limit-1)){
						break;
					}
					$i ++;

					// 查询被质检人是否已经达到设定的阀值
//$row['_source']['content']=json_decode($row['_source']['content']);
					$result[] = $row['_source'];

					//var_dump(json_decode($row['_source']['content'],true));
					//var_dump($i);
				}
				return $result;
			}
			// if($limit > 50){
			// 	$limit = 50;
			// }
			// $body['size'] = $limit*50;// 这里主要是随机多点数据，当部分数据不合格时，可以用剩下的数据
			// $body['from'] = 0;
			// $body['query'] = [];
			// $query = $es->setBobyByWhere([]);
			// // 随机查询数据
			// //var_dump(json_encode($query));
			// //exit;
			// $body['query']['function_score'] = [
			// 		'query'=>$query,
			// 		"functions"=>[
			// 				[
			// 						"random_score"=>[
			// 								"seed"=>rand(1,999999)
			// 						]
			// 				]
			// 		],
			// 		"score_mode"=> "sum"
			// ];
			var_dump($body);
			$ret = $es->search($this->es_qctask_index,$this->es_qctask_index_type,$body);
			if(!empty($ret['hits']['hits'])){
				$i = 0;
				foreach ($ret['hits']['hits'] as $row){
					if($i > ($limit-1)){
						break;
					}
					$i ++;

					// 查询被质检人是否已经达到设定的阀值
//$row['_source']['content']=json_decode($row['_source']['content']);
					$result[] = $row['_source'];

					//var_dump(json_decode($row['_source']['content'],true));
					//var_dump($i);
				}
				return $result;
			}
			if(!empty($result)){
				$result = $this->packageQcTask($result,$username,$es_cfg);
				foreach ($result as $key=>$val){
					// 更新质检人
					$this->updateQcIndexField($es_cfg, ['operator'=>$username,'task_id'=>$val['task_id']]);
				}
					//var_dump($result);
					//var_dump("stat new result");
					//var_dump($result);
					//var_dump("done new result");
			}
					//var_dump($result);
		return $result;
	}

	/**
	 * IMC 内容格式化
	 * @param string $content
	 * @return array
	 */
	public function imcContentFormat($content){
		$content_arr = explode("<br><br>", $content);

		$info = [];
		foreach ($content_arr as $row){
			if(trim($row)){
				$tmp = explode("<br>", $row);
				$str = $tmp[0]?$tmp[0]:$tmp[1];
				if(strstr($str,"客服") || substr($str, 0,3) == 400){
					$info[] = [
							'type'=>1,
							'text'=>$row,
							'bps'=>'',
							'eps'=>'',
							'hit_tactic'=>'',
							'index'=>'',
							'tactic_second_dir'=>'',
							'tactic_second_dir_name'=>'',
					];
				}else{
					$info[] = [
							'type'=>2,
							'text'=>$row,
							'bps'=>'',
							'eps'=>'',
							'hit_tactic'=>'',
							'index'=>'',
							'tactic_second_dir'=>'',
							'tactic_second_dir_name'=>'',
					];
				}
			}
		}
		return $info;
	}

	/**
	 * 红包需求替换对应的图片地址
	 * @param string $content
	 */
	public function replace_hongbao_imgurl($content){
		if(empty($content)){
			return '';
		}
		//$url = str_replace('["PICURL:', '<img src="', $url);
		//$url = str_replace('.jpg]','.jpg" />', $url);
		//$url = str_replace('.jpeg]','.jpeg" />', $url);
		//$url = str_replace('.gif]','.gif" />', $url);
		//$url = str_replace('.png]','.png" />', $url);
		//$url = str_replace('.bmp]','.bmp" />', $url);

		// 对于多图情况处理，如：["PICURL:http://imc.cm.com/pic/minipic/20171008/201710081833381526011584.jpg","PICURL:http://imc.cm.com/pic/minipic/20171008/2017100818335514710112204.jpg"]
		$content = str_replace(array('["PICURL:','"PICURL:'), '<img src="', $content);
		$content = str_replace(array('.jpg"]','.jpg",'),'.jpg" />', $content);
		$content = str_replace(array('.jpeg"]','.jpeg",'),'.jpeg" />', $content);
		$content = str_replace(array('.gif"]','.gif",'),'.gif" />', $content);
		$content = str_replace(array('.png"]','.png",'),'.png" />', $content);
		$content = str_replace(array('.bmp"]','.bmp",'),'.bmp" />', $content);
		return $content;
	}

	/**
	 * 同步归档工单
	 * 1、按归档时间段拉取工单
	 * 2、工单数据处理
	 * 3、数据入库
	 * 4、调用后台接口（1、ivr 语音转文本接口，2、异步智能质检接口）
	 * 注：系统同时要提供接口给后台回写质检结果
	 * 兼容处理：
	 * 1、先判断单据是否已同步，已同步则不再同步
	 * result 格式：
	 * Array
(
    [_scroll_id] => cXVlcnlUaGVuRmV0Y2g7MzszNjY6WnFnVG9rd3JRMDJWLVV1dkI5SWs4ZzszNjE6aDIxNktOZDZUSy1BNzV2QWNycDNZZzszNjI6aDIxNktOZDZUSy1BNzV2QWNycDNZZzswOw==
    [took] => 8
    [timed_out] =>
    [_shards] => Array
        (
            [total] => 3
            [successful] => 3
            [failed] => 0
        )

    [hits] => Array
        (
            [total] => 217
            [max_score] => 1
            [hits] => Array
                (
                    [0] => Array
                        (
                            [_index] => taskindex
                            [_type] => task
                            [_id] => 17120516074282433850
                            [_score] => 1
                            [_source] => Array
                                (
                                    [5088] => 400201
                                    [5100] => 2;958;A2069;1038095;1038125;1038125
                                 )
                         )
                 )
          )
     )
     * @param array $es_cfg   es 配置：host、端口
     * @param array $time     归档时间条件
	 * @return
	 */
	public function syncTask($es_cfg,$time = [],$where=[]){
		//$es_cfg['host'] = '10.198.144.103';
		$es = Services::getService('Es',$es_cfg);
		$es->createClient($es_cfg);

		if(empty($time) OR count($time) != 2){
			$time = [];
			$time[0] = date("Y-m-d 00:00:00",strtotime('-1 days'));
			$time[1] = date("Y-m-d 23:59:59",strtotime('-1 days'));
		}
		if(empty($where)){
			// 归档时间
			//$where['1882'] = ['operator'=>'=','value'=>'IVR'];
			$where['2132'] = ['operator'=>'range','value'=>['sdate'=>$time[0],'edate'=>$time[1]]];
			//$where['5089'] = ['operator'=>'!=','value'=>'kfhb'];
			$where['1882'] = ['operator'=>'!=','value'=>'心悦平台'];
			$where['2298'] = ['operator'=>'!=','value'=>'admin'];
			//$where['3450'] = ['operator'=>'=','value'=>'A1959'];//-------------- 只是微信支付的
			//$where['1882'] = ['operator'=>'in','value'=>['微信异步','手 Q 异步','微信渠道','手Q']];
			//$where['1880'] = ['operator'=>'term','value'=>'18031723233081686699'];
			//$where = [];
		}
		$total = 0;
		try {
			$this->getStaffGroups();
			$response = $es->searchScroll($this->es_task_index,$this->es_task_index_type,$where,'30s',10);
			$count = count($response['hits']['hits']);
			if(!isset($response['hits']['hits']) OR  $count <= 0){
				return false;
			}
			$ret = $this->insertTask($es_cfg, $response['hits']['hits']);
			$total += $count;
			//print_r($ret);
			//print_r($response['hits']['hits']);die();
			while (isset($response['hits']['hits']) && count($response['hits']['hits']) > 0) {
				$scroll_id = $response['_scroll_id'];
				$response =$es->scrolling($scroll_id,'30s');

				$ret = $this->insertTask($es_cfg, $response['hits']['hits']);
				$total += count($response['hits']['hits']);
			}
		} catch (Exception $e) {
			print_r($e);
		}
		return $total;
	}


	/**
	 * 工单入库
	 * @param array $data
	 */
	public function insertTask($es_cfg,$data){
		$es = Services::getService('Es',$es_cfg);
		$datas = [];
		if(count($data)==count($data,1)){
			$datas[] = $data;
		}else{
			$datas = $data;
		}
		$task_int = $this->task_int;
		$this->task_tags = $this->getTaskTags();

		$insertData = [];
		$ivrData = [];
		$encrypt = new Encrypt();
		foreach ($datas as $val){
			if(!isset($val['_source'])){
				continue;
			}

			$row = $val['_source'];

			// 查询质检单据是否存在，存在则不再拉取
			/*************************************/
			//$es_cfg['host'] = '10.133.8.94';
			//$es->createClient($es_cfg);
			/*************************************/

			//$query_qc = $es->searchById($this->es_qctask_index,$this->es_qctask_index_type,$row['1880']);

			if(!empty($query_qc)){
				continue;
			}
			/*************************************/
			//$es_cfg['host'] = '10.198.144.103';
			//$es->createClient($es_cfg);
			/*************************************/

			//var_dump($row);
			//exit;
			$related_task = '';
			$related_content = '';
			$task_type = 0;
			$package_data = $this->packageInsertData($row,$related_content,$task_type);
			$insertData[] = $package_data['insertData'];
			$ivrData[] = $package_data['ivrData'];
			if(isset($row['1864'])){// 投诉单
				// 拉取投诉单据出来标识，同时入库。当前单则标识 check_status=3无需人工质检
				$task_id = '';
				$task_id = $row['1864'];
				if(strlen($row['1864'])>20){
					$task_id = substr($row['1864'], 0,20);
				}
				if(!is_numeric($task_id)){
					continue;// 如果对象单不存在，则做任何处理
				}
				$query = $es->searchById($this->es_task_index,$this->es_task_index_type,$task_id);
				if(empty($query)){
					continue;// 如果对象单不存在，则做任何处理
				}
				$related_content .= "记录单号：".$row['1880'].";<br/>";
				if(isset($row['1862'])){
					$related_content .= '被投诉人：'.$row['1862'].";<br/>详细描述：<br/>";
				}
				if(!empty($row[$task_int['intro']])){
					$related_content .= $encrypt->decryptBase64($row[$task_int['intro']], _AES_ECB_KEY_);
				}
				$task_type = '1';
				$package1864_data = $this->packageInsertData($query['_source'],$related_content,$task_type);
				$insertData[] = $package1864_data['insertData'];
				$ivrData[] = $package1864_data['ivrData'];
				//$row = $query['_source'];
			}else if(isset($row['1860'])){// 表扬单
				// 拉取表扬单据出来标识，同时入库。当前单则标识 check_status=3无需人工质检
				$task_id = '';
				$task_id = $row['1860'];
				if(strlen($row['1860'])>20){
					$task_id = substr($row['1860'], 0,20);
				}
				if(!is_numeric($task_id)){
					continue;// 如果对象单不存在，则做任何处理
				}
				$query = $es->searchById($this->es_task_index,$this->es_task_index_type,$task_id);
				if(empty($query)){
					continue;// 如果对象单不存在，则做任何处理
				}
				$related_content .= "记录单号：".$row['1880'].";<br/>";
				if(isset($row['1858'])){
					$related_content .= '被表扬人：'.$row['1858'].";<br/>详细描述：<br/>";
				}
				if(!empty($row[$task_int['intro']])){
					$related_content .= $encrypt->decryptBase64($row[$task_int['intro']], _AES_ECB_KEY_);
				}
				$task_type = '2';
				$package1860_data = $this->packageInsertData($query['_source'],$related_content,$task_type);
				$insertData[] = $package1860_data['insertData'];
				$ivrData[] = $package1860_data['ivrData'];
				//$row = $query['_source'];
			}

		}
		$ret = false;
		if(!empty($insertData)){

			/*************************************/
			//$es_cfg['host'] = '10.133.8.94';
			//$es->createClient($es_cfg);
			/*************************************/

			$ret = $es->storeBatch($this->es_qctask_index,$this->es_qctask_index_type,$insertData,'task_id');
			if(!empty($ivrData)){
				//ivr 单据要调用后台语音转文本入库接口
				$backendApi = new BackendApi();
				$ret_api = $backendApi->aiqc(['data'=>json_encode($ivrData)]);
				//print_r($ret_api);
			}

			$insertData = [];
		}
		return $ret;
	}

	public function packageInsertData($row,$related_content,$task_type){
		$task_int = $this->task_int;
		$this->task_tags = $this->getTaskTags();
		$encrypt = new Encrypt();
		$insertData = [];
		$ivrData = [];
		if(empty($row[$task_int['archivepath_name']]) && !empty($row[$task_int['pathname']])){
			$pathname = explode("->", $row[$task_int['pathname']]);
			$pathname = array_slice($pathname,-3,3);
			$row[$task_int['archivepath_name']] = implode('->', $pathname);
		}
		$reply_content = isset($row[$task_int['reply_content']])?$row[$task_int['reply_content']]:'';
		$content = isset($row[$task_int['content']])?$row[$task_int['content']]:'';
		$intro = isset($row[$task_int['intro']])?$row[$task_int['intro']]:'';
		$userid = isset($row[$task_int['userid']])?$row[$task_int['userid']]:'';
		$calling_number = isset($row[$task_int['calling_number']])?$row[$task_int['calling_number']]:'';
		$service_channel = isset($row[$task_int['service_channel']])?$row[$task_int['service_channel']]:'';
		$archive_path = isset($row[$task_int['archive_path']])?$row[$task_int['archive_path']]:'';
		$create_user = isset($row[$task_int['create_user']])?$row[$task_int['create_user']]:'';
		$task_label = isset($row[$task_int['task_label']])?$row[$task_int['task_label']]:'';// 是否升级单
		$satis_level = isset($row[$task_int['satis_level']])?$row[$task_int['satis_level']]:'';// 不满意单
		$call_time = isset($row[$task_int['call_time']])?$row[$task_int['call_time']]:'';// 超时
		$resp_time = isset($row[$task_int['resp_time']])?$row[$task_int['resp_time']]:'';// 人工响应时间
		$begin_time = isset($row[$task_int['begin_time']])?$row[$task_int['begin_time']]:'';// 发起会话时间
		$access_time = isset($row[$task_int['access_time']])?$row[$task_int['access_time']]:'';// 接入时间
		if(!empty($reply_content)){
			$reply_content = $encrypt->decryptBase64($reply_content, _AES_ECB_KEY_);
		}
		if(!empty($content)){
			$content = $encrypt->decryptBase64($content, _AES_ECB_KEY_);
		}
		if(!empty($intro)){
			$intro = $encrypt->decryptBase64($intro, _AES_ECB_KEY_);
		}
		if(!empty($userid)){
			$userid = $encrypt->decryptBase64($userid, _AES_ECB_KEY_);
		}
		if(!empty($calling_number)){
			$calling_number = $encrypt->decryptBase64($calling_number, _AES_ECB_KEY_);
		}
		// 标签处理
		$quality_label = [];
		foreach ($row as $k=>$v){
			if(isset($this->task_tags[$k]) && $this->task_tags[$k]['meunid_level3'] == $archive_path){
				$tmp_label = [];
				$tmp_label['value'] = $v;
				$tmp_label['name'] = $this->task_tags[$k]['name'];
				$tmp_label['id'] = $this->task_tags[$k]['id'];
				$quality_label[] = $tmp_label;
			}
		}
		// 特征处理
		$feature = [];
		if($call_time>1000){
			array_push($feature, 16);
		}
		// 升级
		if($task_label == '65536'){
			array_push($feature, 2);
		}
		// 不满意
		$imc_level = ['2','21','22','1','11','12','23','24','25','26','27','28','61','62','63','64','65','66','67','68','69','6a','6b','6c','6d','6e','6f','71','72','73','74','75','76','77','78','79','7a','7b','7c','7d','7e','7f'];
		//if(($service_channel == 'IVR' && in_array($satis_level, [3,4,5])) || ($service_channel == '异步' && in_array($satis_level, [1,2])) || ($service_channel == 'IMC' && in_array($satis_level, $imc_level))){
		if((in_array($service_channel,['IVR','电话','IPCC','外呼渠道']) && in_array($satis_level, [3,4,5])) || (in_array($service_channel,['微信异步','手 Q 异步','小程序异步','异步']) && in_array($satis_level, [1,2])) || (in_array($service_channel,['IMC','微信渠道','手Q渠道']) && in_array($satis_level, $imc_level))){
			array_push($feature, 8);
		}
		//
		$tmp = [
				'task_id'=>$row[$task_int['task_id']],
				'reply_content'=>$reply_content,
				'content'=>$content,
				'intro'=>$intro,
				'record_id'=>isset($row[$task_int['record_id']])?$row[$task_int['record_id']]:'',
				'service_id'=>isset($row[$task_int['service_id']])?$row[$task_int['service_id']]:'',
				'archive_path'=>$archive_path,
				'path_name'=>isset($row[$task_int['archivepath_name']])?$row[$task_int['archivepath_name']]:'',
				'quality_label'=>'',
				'create_user'=>$create_user,
				'owner'=>isset($row[$task_int['owner']])?$row[$task_int['owner']]:'',
				'service_channel'=>$service_channel,
				'type'=>isset($row[$task_int['type']])?$row[$task_int['type']]:'',
				'userid'=>$userid,
				'feature'=>$feature,
				'first_dir_code'=>isset($row[$task_int['first_dir_code']])?$row[$task_int['first_dir_code']]:'',
				'calling_number'=>$calling_number,
				'archive_time'=>isset($row[$task_int['archive_time']])?$row[$task_int['archive_time']]:'',
				'defendant_user'=>isset($row[$task_int['create_user']])?$row[$task_int['create_user']]:'',
				'quality_label'=>empty($quality_label)?'':json_encode($quality_label)
		];
		$toAi = 1;
		if(in_array($service_channel, ['IMC']) && !empty($resp_time) && !empty($access_time) && !empty($begin_time) && strpos($create_user,'v_')===false){
			if((strtotime($resp_time) - strtotime($access_time)) > 480){// 大于 8分钟
				$toAi = 0;
			}
		}
		$h = (int)date('H');
		if($h >= 12){// 这里限制上午的数据就不发了
			$toAi = 0;
		}
		if($toAi == 1){
			if($service_channel == 'IVR' || $service_channel == '外呼渠道'){
				//if(in_array($row[$task_int['service_id']], ['A1959','A3350','A9200'])){
				$ivrData = ['taskid'=>$row[$task_int['task_id']],'type'=>'ivr'];
				//}
			}else if(in_array($service_channel, ['微信异步','手 Q 异步','微信渠道','手Q渠道']) && $row[$task_int['service_id']] == 'A1959'){
				$ivrData = ['taskid'=>$row[$task_int['task_id']],'type'=>'async'];
			}
		}

		$tmp['related_content'] = $related_content;// 保存原工单内容（表扬单、投诉单）
		$tmp['task_type'] = $task_type;
		$tmp = $this->packageQcTaskField($tmp);// 数据二次处理
		if(in_array($service_channel, ['IMC']) && !empty($resp_time) && !empty($access_time) && !empty($begin_time) && strpos($create_user,'v_')===false){// 去掉 v_开头的创建人
			if((strtotime($resp_time) - strtotime($access_time)) > 480){// 大于 8分钟
				$tmp['operator'] = "AI";
				$tmp['operator_time'] = date("Y-m-d H:i:s");
				$tmp['ai_status'] = 3;
				$tmp['error_type'] = json_encode(['type1'=>['5','34']]);
				$tmp['check_reault'] = "首次响应时长超过 8 分钟";
				$tmp['score'] = "-0.50";
				$tmp['all_check_status'] = 1;
			}
		}
		$insertData[] = $tmp;
		$data = [];
		$data['insertData'] = $tmp;
		$data['ivrData'] = $ivrData;
		return $data;
	}

	/**
	 * 二次处理原工单信息，完善质检单据字段信息
	 * 1、被质检人所在的中心、大组、小组
	 * 2、是否重复来单
	 * 3、是否升级单
	 * 4、被投诉和被表扬标识的工单
	 * @param array $data
	 */
	public function packageQcTaskField($data){
		// 追加默认的字段保存
		$data['check_status'] = 0;
		$data['ai_status'] = 0;
		$data['all_check_status'] = 0;
		$data['appeal_status'] = 20;
		if(!isset($data['task_type'])){
			$data['task_type'] = 0;
		}
		$data['score'] = 0;
		$data['operator'] = '';
		$data['center_id'] = '';
		$data['group_id1'] = '';
		$data['group_id2'] = '';
		if($data['service_channel'] == "IMC" || $data['service_channel'] == "IVR"){
			$data['defendant_user'] = $data['create_user'];
		}else{
			$data['defendant_user'] = $data['owner'];
		}
		if(!empty($data['defendant_user']) && isset($this->staff_groups[$data['defendant_user']])){
			$data['center_id'] = $this->staff_groups[$data['defendant_user']]['center_id'];
			$data['group_id1'] = $this->staff_groups[$data['defendant_user']]['group1_id'];
			$data['group_id2'] = $this->staff_groups[$data['defendant_user']]['group2_id'];
			$data['group_names'] = $this->staff_groups[$data['defendant_user']]['path'];
		}
		if($data['owner'] == 'admin'){
			$data['check_status'] = 3;// 对于属主为 admin的不需要人工质检
		}
		return $data;
	}

	public function checkGroupStaffCount(){

	}

	/**
	 * 更新质检单据的字段内容
	 * @param array $es_cfg
	 * @param array $data
	 * @return boolean
	 */
	public function updateQcIndexField($es_cfg,$data){
		if(!isset($data['task_id'])){
			return false;
		}
		$es = Services::getService('Es',$es_cfg);
		if(!empty($data['content']) && is_array($data['content'])){
			$data['content'] = json_encode($data['content'],JSON_UNESCAPED_UNICODE);
		}
		if(!empty($data['tactics']) && is_array($data['tactics'])){
			$data['tactics'] = json_encode($data['tactics'],JSON_UNESCAPED_UNICODE);
		}
		$data['all_check_status'] = '1';// 标记为已检
		$data['operator_time'] = date("Y-m-d H:i:s");// 标记为已检
		$ret = $es->update($this->es_qctask_index,$this->es_qctask_index_type,$data,$data['task_id']);
		return $ret;
	}

	/**
	 * 添加质检单据的字段内容
	 * @param array $es_cfg
	 * @param array $data
	 * @return boolean
	 */
	public function addQcIndexField($es_cfg,$data){
		if(!isset($data['task_id'])){
			return false;
		}
		$es = Services::getService('Es',$es_cfg);
		if(!empty($data['content']) && is_array($data['content'])){
			$data['content'] = json_encode($data['content'],JSON_UNESCAPED_UNICODE);
		}
		if(!empty($data['tactics']) && is_array($data['tactics'])){
			$data['tactics'] = json_encode($data['tactics'],JSON_UNESCAPED_UNICODE);
		}
		if(!empty($data['quality_label']) && is_array($data['quality_label'])){
			$data['quality_label'] = json_encode($data['quality_label'],JSON_UNESCAPED_UNICODE);
		}
		$data['all_check_status'] = '1';// 标记为已检
		$data['operator_time'] = date("Y-m-d H:i:s");// 标记为已检
		$ret = $es->store($this->es_qctask_index,$this->es_qctask_index_type,$data,$data['task_id']);
		return $ret;
	}


	/**
	 * 获取员工中心、大组、小组信息
	 * @return array
	 */
	public function getStaffGroups(){
		if(empty($this->staff_groups)){

			$model = new \App\Models\Entities\UserModel();
			$model->setConnection('dms');
			$where['conditions']['is_valid'] = ['operator'=>'=','value'=>'1'];
			$ret = $model->fetch($where);

			$departmentModel = new \App\Models\Entities\DmsDepartmentModel();
			$query_dept = $departmentModel->fetch([
					'conditions'=>['is_valid'=>['operator'=>'=','value'=>'1']]
			]);
			$group_ids = [];
			if(!empty($query_dept)){
				foreach ($query_dept as $row){
					$group_ids[$row->id] = [
							'id'=>$row->id,
							'name'=>$row->name,
							'parent_id'=>$row->parent_id
							];
				}
			}
			$data = [];
			if(!empty($ret)){
				foreach ($ret as $row){
						$data[$row->login_name] = [
								//'department_id'=>'1',
								'center_id'=>$row->center_id,
								'group1_id'=>$group_ids[$row->group_id]['parent_id'],
								'group2_id'=>$row->group_id,
								'path'=>$group_ids[$row->center_id]['name']."->".$group_ids[$group_ids[$row->group_id]['parent_id']]['name']."->".$group_ids[$row->group_id]['name'],
						];
				}
				$this->staff_groups =  $data;
			}
		}
		return $this->staff_groups;
	}

	/**
	 * 获取指定 code 的下面子归档
	 * @param unknown $parent_code
	 */
	public function getChildrenArchive($parent_code){
		$data = array();
		if($parent_code){
			$where = array();
			$conditions['parent_code'] = array('value'=>$parent_code,'operator'=>'=');
			$conditions['deleted'] = array('value'=>0,'operator'=>'=');
			$where['conditions'] = $conditions;
			$model = new \App\Models\Entities\ArchiveModel();
			$res = $model->fetch($where);// 三级
			foreach ($res as $val){
				$tmp = array();
				$tmp['id'] = $val->code_id;
				$tmp['name'] = $val->name;
				$data[] = $tmp;
			}
		}
		return $data;
	}

	public function getStaffs(){
		$data = [];
		return $data;
	}

	/**
	 * 获取组织架构
	 */
	public function getOrganization($centers){
		$data = array();
		if(!empty($centers)){
			$center_ids = [];
			$center_names = [];
			foreach ($centers as $val){
				$center_ids[] = $val['center_id'];
				$center_names[$val['center_id']] = $val['center_name'];;
			}
			$DepartmentModel = new DepartmentModel();
			//depth =? AND is_valid='1'
			$where = array();
			$where['conditions'] = array('depth'=>array('value'=>'3','operator'=>"="),'is_valid'=>array('value'=>'1','operator'=>"="),'parent_id'=>array('value'=>$center_ids,'operator'=>"in"));
			$group1 = $DepartmentModel->fetch($where);
			$group1_arr = $group1->toArray();
			$group1_ids = array();
			$center_group1 = array();
			foreach ($group1_arr as $val){
				$group1_ids[] = $val['id'];
				$center_group1[$val['parent_id']][] = $val['id'];
			}
			$where = array();
			$where['conditions'] = array('depth'=>array('value'=>'4','operator'=>"="),'is_valid'=>array('value'=>'1','operator'=>"="),'parent_id'=>array('value'=>$group1_ids,'operator'=>"in"));
			$group2 = $DepartmentModel->fetch($where);
			$group2_arr = $group2->toArray();
			$group1_group2 = array();
			foreach ($group2_arr as $val){
				$tmp = [];
				$tmp['value'] = $val['id'];
				$tmp['label'] = $val['name'];
				$group1_group2[$val['parent_id']][] = $tmp;
			}
			foreach ($group1_arr as $val){
				$group1_ids[] = $val['id'];
			}
			$groups = array();
			$tmp = array();
			foreach ($group1_arr as $val){
				$tmp['value'] = (string)$val['id'];
				$tmp['label'] = $val['name'];
				$tmp['children'] = isset($group1_group2[$val['id']])?$group1_group2[$val['id']]:array();
				$groups[$val['id']] = $tmp;
			}
			foreach ($center_group1 as $key=>$val){
				$tmp = array();
				$tmp1 = array();
				foreach ($groups as $k=>$v){
					if(in_array($k, $val)){
						$tmp1[] = $v;
					}
				}
				$tmp['children'] = $tmp1;
				$tmp['value'] = (string)$key;
				$tmp['label'] = isset($center_names[$key])?$center_names[$key]:'';
				$data[] = $tmp;
				//$tmp['children'] = isset($groups[$val['id']])?$group1_group2[$val['id']]:array();
			}
		}
		return $data;
	}

	/**
	 * 获取人员直属 leader
	 */
	public function getLeader($rtx){
		$UserModel = new UserModel();
		$where = array();
		$where['conditions'] = array('login_name'=>array('value'=>$rtx,'operator'=>"="),'is_valid'=>array('value'=>'1','operator'=>"="));
		$ret = $UserModel->fetchOne($where);
		if(!empty($ret)){
			$group_id = $ret->group_id;
			$where = array();
			$where['conditions'] = array('group_id'=>array('value'=>$group_id,'operator'=>"="),'is_valid'=>array('value'=>'1','operator'=>"="),'level_id'=>array('value'=>'2','operator'=>"="));
			$data = $UserModel->fetchOne($where);
		}
		return isset($data->login_name)?$data->login_name:'';
	}

	/**
	 * 获取主管下所有的人员
	 */
	public function getUsersByLeader($rtx){
		$UserModel = new UserModel();
		$where = array();
		$where['conditions'] = array('login_name'=>array('value'=>$rtx,'operator'=>"="),'is_valid'=>array('value'=>'1','operator'=>"="));
		$ret = $UserModel->fetchOne($where);
		$datas = [];
		if(!empty($ret)){
			$group_id = $ret->group_id;
			$where = array();
			$where['conditions'] = array('group_id'=>array('value'=>$group_id,'operator'=>"="),'is_valid'=>array('value'=>'1','operator'=>"="),'level_id'=>array('value'=>'1','operator'=>"="));
			$data = $UserModel->fetch($where);
			if(!empty($data)){
				foreach ($data as $val){
					$datas[] = $val->login_name;
				}
			}
		}
		return $datas;
	}

	/**
	 * 导出数据
	 */
	public function exportEsData1($ret_es,$heads){
		$head = $heads['head'];
		set_time_limit(0);
		ini_set ( 'memory_limit', '-1' );// 设置导出最大内存
		$ranking = [];
		if(isset($ret_es['hits']['hits'])){
			$hits = $ret_es['hits']['hits'];
			foreach ( $hits as $key => $value ) {
				$_source = $value['_source'];
				$tmp = [];
				//$_source['taskid'] = (string)$_source['taskid'];
				//$_source['is_risk'] = (string)$_source['is_risk'];
				foreach ($head as $h){
					$tmp[] = isset($_source[$h])?$_source[$h]:"";
				}
				$ranking [$key + 1] = array_values ( $tmp );
			}
		}
		$date = date("YmdHis");// 日期作为输出文件后缀
		$content = $this->getXLSFromList($heads,$ranking);// 获得输出的表格内容
		header("Content-type:application/vnd.ms-execl;charset=gb2312");// 设置导出格式
		header("Content-Disposition:attactment;filename=".$date."数据.xls");// 设置导出文件名
		header("Pragma: no-cache");
		header("Expires: 0");
		echo $content;
	}

	// 此方法建议写入公共方法 通过数组遍历得出导出报表类型结构
	function getXLSFromList($heads,$lists){
	//　　内容太大建议搜索少量再导出
	//    if(count($lists)>=20000)
	//    {
	//        header("Content-Type:text/html;charset=utf-8");
	//        echo "<br/><h1 style='color:red'>Export data is too large, please narrow your search!</h1><br/>";
	//        exit;
	//    }
		$pres = $heads['head'];
		$head_name = $heads['head_name'];
	    $keys=array_keys($pres);// 获取表头的键名
	    $content="";
	    $content.="<table border='0'><tr>";
	    foreach($head_name as $_pre){// 输出表头键值
	        $content.="<td>$_pre</td>";
	    }
	    $content.="</tr>";
	    foreach($lists as $_list){
	        $content.= "<tr>";
	            foreach($keys as $key){
	            	if($key == 4 || $key == 2 || $key == 0){
	            		$content.= "<td style='vnd.ms-excel.numberformat:@'>".$_list[$key]."</td>"; //style 样式将导出的内容都设置为文本格式 输出对应键名的键值 即内容
	            	}else{
	            		if(isset($_list[$key])){
	            			$content.= "<td>".$_list[$key]."</td>"; //style 样式将导出的内容都设置为文本格式 输出对应键名的键值 即内容
	            		}
	            	}

	            }
	        $content.="</tr>";
	    }
	    $content.="</table>";
	    return $content;
	}

	/**
	 * 返回质检周期日期
	 * @param string $date
	 * @return array
	 */
	public static function qcCycle($date = ''){
		if(empty($date)){
			$date = date("Y-m-d");
		}
		$month = date("m",strtotime($date));
		$sdate = date("Y-m-21",strtotime("-1 month $date"));
		$edate = date("Y-m-21",strtotime($date));
		if(strtotime($date) > strtotime($edate)){
			$sdate = date("Y-m-21",strtotime($date));
			$edate = date("Y-m-21",strtotime("+1 month $date"));
		}
		$data = [$sdate,$edate];
		return $data;
	}

	/**
	 * 根据日期返回质检周期月份
	 * @param string $date
	 * @return string
	 */
	public static function getQcCycleMonth($date = ''){
		if(empty($date)){
			$date = date("Y-m-d");
		}
		$d = date("d",strtotime($date));
		if((int)$d >= 21){
			return date("Y-m-01",strtotime("+1 month $date"));
		}
		return date("Y-m-01",strtotime($date));
	}

	/**
	 * 获取不同级别人员
	 */
	public function getLevelUser($level_id){
		$model = new \App\Models\Entities\UserModel();
		$where['conditions']['StatusId'] = ['operator'=>'=','value'=>'1'];
		$where['conditions']['level_id'] = ['operator'=>'in','value'=>$level_id];
		$ret = $model->fetch($where);
		$data = [];
		if(!empty($ret)){
			foreach ($ret as $val){
				$data[$val->login_name] = $val->chinese_name;
			}
		}
		return $data;
	}

	/**
	 * 获取产品 code 和name对应数组
	 * @return array
	 */
	public function getProductNames(){
		if(empty($this->product_names)){
			$model = new \App\Models\Entities\ArchiveModel();
			$where['conditions'] = "level=1";
			$query = $model->fetch($where);
			foreach ($query as $row){
				$this->product_names[$row->code] = $row->name;
			}
		}
		return $this->product_names;
	}

	/**
	 * 统计被质检人当前周期已质检的量大于指定量的人
	 * @param object $es_cfg
	 * @param array $defendant_users
	 * @return array
	 */
	public function statsDefendantUserCount($es_cfg,$defendant_users=[],$limit=10){
		$dates = $this->qcCycle();
		$es = new Es($es_cfg);
		$where['archive_time'] = ['operator'=>'range','value'=>['sdate'=>$dates[0],'edate'=>$dates[1]]];
		if(!empty($defendant_users)){
			$where['defendant_user'] = ['operator'=>'in','value'=>$defendant_users];
		}
		$where['check_status'] = ['operator'=>'in','value'=>["1","2"]];
		$body = $es->setBobyByWhere($where);
		$body['size'] = 0;
		$body['aggs'] = $es->setGroupWhere(['defendant_user']);
		$query = $es->search($this->es_qctask_index, $this->es_qctask_index_type,$body);
		$data = [];
		if(!empty($query['aggregations']['@defendant_user']['buckets'])){
			foreach ($query['aggregations']['@defendant_user']['buckets'] as $row){
				//$data[$row['key']] = $row['doc_count'];
				if($row['doc_count'] >= $limit){
					$data[] = $row['key'];
				}
			}
		}
		return $data;
	}

	/**
	 * 返回页面代码
	 * @param array $data   参数（页面使用数据）
	 * @param string $tpl   （模板路径）
	 */
	public static function getTpl($data,$tpl){
		$view = new \Phalcon\Mvc\View\Simple();
		$view->setViewsDir(APP_MODULE.'qcadmin/views/'); // 定义视图层目录
		return $view->render("tpl/emailTpl",$data);
	}

	/**
	 * 发送质检相关邮件、rtx
	 * @param array $datas
	 *              $datas['task_id']
	 *              $datas['appeal_status']  //reject（驳回）、error(错误)、pass(通过)、complain(投诉)、approval(审核)、
	 * @param string $to_user
	 * @param array $sendType
	 * @return boolean
	 */
	public function sendQcInfo($datas,$to_user,$sendType=[]){
		//return false;
		if(empty($sendType) || empty($datas)){
			return false;
		}
		$sendType = array_flip($sendType);
		$appeal_status = $datas['appeal_status'];//reject（驳回）、error(错误)、pass(通过)、complain(投诉)、approval(审核)、
		$tof_oa = new \App\Libraries\Tof_oa();

		$url = $datas['base_url']."qcadmin/frame/index?hash=/admin/orderDetail/".$datas['task_id'];
		$content = '有 1 单质检申诉需要您审批，详情请点击 『这里|'.$url.'] 查看。';
		$email_title = '质检审批提醒';
		switch ($appeal_status) {
			case 'error':
				$content = '有 1 单质检错误判定成立，详情请点击 『这里|'.$url.'] 查看。如有疑问，请在申诉期内提交申诉信息';
				$email_title = '质检错误提醒';
				break;
			case 'reject':
				$content = '有 1 单质检申诉被驳回，详情请点击 『这里|'.$url.'] 查看。';
				$email_title = '质检申诉结果提醒';
				break;
			case 'pass':
				$content = '有 1 单质检申诉已通过，详情请点击 『这里|'.$url.'] 查看。';
				$email_title = '质检申诉结果提醒';
				break;
			case 'complain':
				$content = '有 1 单投诉判定成立，详情请点击 『这里|'.$url.'] 查看。如有疑问，请在申诉期内提交申诉信息';
				$email_title = '质检投诉成立提醒';
				break;
			default:
				;
				break;
		}

		if(isset($sendType['rtx'])){
			// 发送 rtx
			$ret = $tof_oa->send_rtx('【智能质检平台】',$content,'QC-ADMIN',$to_user);
		}
		if(isset($sendType['email'])){
			$data ['task_id'] = $datas['task_id'];
			$data ['appeal_status'] = $appeal_status;
			$data ['to_user'] = $to_user;
			if(isset($datas['score'])){
				$data ['score'] = $datas['score'];
			}
			if(isset($datas['owner'])){
				$data ['owner'] = $datas['owner'];
			}
			if(isset($datas['error_type'])){
				$Intention = new Intention();
				$data ['error_type'] = $Intention->escapingErrorType($datas['error_type']);
			}

			$content = $this->getTpl($data, "tpl/emailTpl");
			// 发送 email
			$datalist['From'] = _EMAIL_FROM_USER_;
			$datalist['To'] = $to_user;
			//$datalist['CC'] = '';
			$datalist['Content'] = $content;
			$datalist['Title'] = $email_title;
			$ret = $tof_oa->send_mail($datalist);
		}
		return true;
	}



	/**
	 * 获取工单表字段对应的标签信息
	 */
	public function getTaskFieldTags(){
		if(empty($this->task_field_tags)){
			$model = new \App\Models\Entities\ElementModel();
			$where['conditions'] = [
					'type'=>['operator'=>'=','value'=>'task']
			];
			$query = $model->fetch($where,'name,id');
			if(!empty($query)){
				$query = $query->toArray();
				foreach ($query as $row){
					$this->task_field_tags[$row['id']] = $row;
				}
			}
		}
		return $this->task_field_tags;
	}

	/**
	 * 获取工单标签信息（非基础字段标签）
	 */
	public function getTaskTags(){
		if(empty($this->task_tags)){
			$model = new \App\Models\Entities\ElementModel();
			$where['conditions'] = [
					'type'=>['operator'=>'!=','value'=>'task']
			];
			//$query = $model->fetch($where,'name,id');
			$sql = "select a.name,a.id,c.meunid_level3,c.meunid_level2,c.product_code
					from t_element a
					join t_top_tag_element b on a.id=b.element_id
					join t_top_tag c on b.top_id=c.id";
			$query = $model->querySql($sql);

			if(!empty($query)){
				//$query = $query->toArray();
				foreach ($query as $row){
					$this->task_tags[$row['id']] = $row;
				}
			}
		}
		return $this->task_tags;
	}
}
