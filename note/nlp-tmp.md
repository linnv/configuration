[
@ OutCallModel.py:212 @ def textSimilarity(doc,nodeMap,nodeId):
        doc_token = jt.tokens(text)
        doc_feat = fb.compute(doc_token)
        doc_fl = DocFeatLoader(smb, doc_feat)

        logger.info("text:%s"%(text))
        logger.info("doc_token:%s"%(doc_token))
        logger.info("doc_feat:%s"%(doc_feat))
        logger.info("doc_fl:%s"%(doc_fl))

        contentFlListMap = nodeMap
        if nodeId in contentFlListMap.keys():
            nodeFlList = contentFlListMap[nodeId]
            nodeMaxDist = 0
            nodeMaxIndex = 0
            for i in range(len(nodeFlList)):
                # logger.info("computin:%s"%("xxx"))
                # logger.info("nodeFlList[i][lableDataFeatureVector].feat_vec:%s"%(nodeFlLis
t[i]["lableDataFeatureVe
]

curl http://119.23.32.100:8001/OutCall/testnode?nodeId=1&flowId=281
[
{
  "FlowId": "",
  "Result": [
    {
      "NodeId": "xxxNodeId",
      "DisMatchConditions": [
        {
          "ConditioId": "xxx",
          "ConditioTitle": "xxx",
          "Keywords": [
            "1"
          ],
          "Labels": [
            "2"
          ]
        }
      ]
    }
  ]
}
]

{

logger = logging.getLogger()

fileHandler = logging.FileHandler("./OutCallModel.log.{0}".format(logSufix))
fileHandler.setFormatter(logFormatter)
logger.addHandler(fileHandler)

consoleHandler = logging.StreamHandler()
consoleHandler.setFormatter(logFormatter)
logger.addHandler(consoleHandler)

}


[
{
  "Level": "L03",
  "Node": "HMD01-L03N3006",
  "Caption": "是的 =YY|现在就参加 =DIN|好的=YES|我试试看=ITI",
  "Text": "是这样的， 为迎接猪年，我们交行特地邀请您领取新年大礼包。通过手机银行免费领养一只萌猪，就有机会当场直接获得最高 300 元的大福袋。现在就让我来告诉您怎么参加，好吗？",
  "VoiceType": "V/T",
  "Action": "NULL/OVER/ZRG",
  "Result": "客户挂断 - 长时间静音 - 产品介绍"
}
]
