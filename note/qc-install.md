3mf79d

## SmartQC
参考 $SmartQC/config/config.ini.sample

## FilecutProxy
参考 $FilecutProxy/config/config.ini.sample


## FilecutProxy/cutapp/AudioSpliter
config configure_8k_offline:  model to absolute path

----
## php-web: install Phalcon
```
curl -s "https://packagecloud.io/install/repositories/phalcon/stable/script.deb.sh" | sudo bash

sudo apt-get install php5-phalcon
```

install php-mysql php-redis php-curl


config mysql redis host in config-production(php file)
	check app/config/constant: ENVIRONMENT

php:
	app/modules/qcadmin/controllers/BaseController.php
		$this->url->setBaseUri("http://qc.test:8085/");
frontend:
	app\modules\qcadmin\views\frame\index.phtml

issue:
	frontend shoudl use hostname only , don't add port when do request

### todo unicomm
//skip handling task in Audito2TextConvertor since, taskgen/asr maintain convertion job; notify judger there are tasks to work only



{

curl -X PUT "http://47.95.6.82:9200/qc_index" -H 'Content-Type: application/json' -d'
{
  "mappings": {
      "qc_search": {
        "properties": {
          "AudioSource": {
            "type": "string"
          },
          "FilePathAbs": {
            "type": "string"
          },
          "FileSplitedPathHttpAbs": {
            "type": "string"
          },
          "JudgeCount ": {
            "type": "long"
          },
          "TaskFormat": {
            "type": "long"
          },
          "agentID": {
            "type": "string"
          },
          "ai_status": {
            "type": "long"
          },
          "all_check_status": {
            "type": "long"
          },
          "appeal_status": {
            "type": "long"
          },
          "archive_path": {
            "type": "string"
          },
          "archive_time": {
            "type": "long"
          },
          "asrUploadCount": {
            "type": "long"
          },
          "calling_number": {
            "type": "string"
          },
          "center_id": {
            "type": "string"
          },
          "check_reault": {
            "type": "string"
          },
          "check_status": {
            "type": "long"
          },
          "check_type": {
            "type": "string"
          },
          "content": {
            "properties": {
              "bps": {
                "type": "string"
              },
              "eps": {
                "type": "string"
              },
              "hit_content": {
                "type": "string"
              },
              "hit_tactic": {
                "type": "string"
              },
              "index": {
                "type": "string"
              },
              "tactic_first_dir": {
                "type": "string"
              },
              "tactic_second_dir": {
                "type": "string"
              },
              "tactic_second_dir_name": {
                "type": "string"
              },
              "text": {
                "type": "string"
              },
              "type": {
                "type": "string"
              }
            }
          },
          "create_time": {
            "type": "long"
          },
          "create_user": {
            "type": "string"
          },
          "defendant_user": {
            "type": "string"
          },
          "error_type": {
            "type": "string"
          },
          "filePathHttpAbs": {
            "type": "string"
          },
          "filePathLocalAbs": {
            "type": "string"
          },
          "first_dir_code": {
            "type": "string"
          },
          "groupArea": {
            "type": "string",
             "index": "not_analyzed"
          },
          "groupAreaName": {
            "type": "string",
             "index": "not_analyzed"
          },
          "groupCity": {
            "type": "string",
             "index": "not_analyzed"
          },
          "groupCityName": {
            "type": "string",
             "index": "not_analyzed"
          },
          "groupID": {
            "type": "string"
          },
          "groupName": {
            "type": "string",
             "index": "not_analyzed"
          },
          "groupProvince": {
            "type": "string"
          },
          "groupProvinceName": {
            "type": "string",
             "index": "not_analyzed"
          },
          "group_id1": {
            "type": "string"
          },
          "group_id2": {
            "type": "string"
          },
          "group_names": {
            "type": "string"
          },
          "intro": {
            "type": "string"
          },
          "operator": {
            "type": "string",
             "index": "not_analyzed"
          },
          "operator_status": {
            "type": "long"
          },
          "operator_time": {
            "type": "long"
          },
          "owner": {
            "type": "string"
          },
          "path_name": {
            "type": "string"
          },
          "qnEndTime": {
            "type": "long"
          },
          "qnFilePathHttpAbs": {
            "type": "string"
          },
          "qnStartTime": {
            "type": "long"
          },
          "quality_label": {
            "type": "string"
          },
          "record_id": {
            "type": "string"
          },
          "record_url": {
            "type": "string"
          },
          "related_content": {
            "type": "string"
          },
          "reply_content": {
            "type": "string"
          },
          "score": {
            "type": "string"
          },
          "service_channel": {
            "type": "string"
          },
          "service_id": {
            "type": "string"
          },
          "tactics": {
            "properties": {
              "first_dir": {
                "type": "string"
              },
              "second_dir": {
                "type": "string"
              },
              "second_dir_name": {
                "type": "string"
              }
            }
          },
          "taskFormat": {
            "type": "long"
          },
          "task_id": {
            "type": "string",
             "index": "not_analyzed"
          },
          "task_type": {
            "type": "string"
          },
          "type": {
            "type": "string"
          },
          "userid": {
            "type": "string",
             "index": "not_analyzed"
          }
        }
      }
      }
      }
'
}

ffmpeg -i sound.mp3 -af volumedetect -f null -y nul &> original.txt
