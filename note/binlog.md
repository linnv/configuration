{
items[indexID               ]
items[indexSessonID         ]
itemsindexDataTim         ]
items[indexEnterpriseID     ]
items[indexUsername         ]
items[indexSource           ]
items[indexQuery            ]
items[indexAnswer           ]
items[indexStatus           ]
items[indexLeveOne          ]
items[indexLevelTwo         ]
items[indexLevelThree       ]
items[indexLevelFour        ]
items[indexLevelOneScore    ]
items[indexLevelTwoScore    ]
items[indexLevelThreeScore  ]
items[indexLevelFourScore   ]
items[indexMatchType        ]
items[indexNeedIntoAgentFlag]
items[indexOPerationType    ]


+-----------------+--------------+------+-----+-------------------+----------------+
| Field           | Type         | Null | Key | Default           | Extra          |
+-----------------+--------------+------+-----+-------------------+----------------+
| id              | int(11)      | NO   | PRI | NULL              | auto_increment |
| enterpriseID    | varchar(255) | NO   |     | NULL              |                |
| businessNameCh  | varchar(50)  | NO   |     |                   |                |
| businessNameEn  | varchar(255) | NO   |     |                   |                |
| createTime      | datetime     | NO   |     | CURRENT_TIMESTAMP |                |
| token           | varchar(255) | NO   |     | NULL              |                |
| timestamp       | int(11)      | NO   |     | NULL              |                |
| modelServerPort | int(6)       | NO   |     | 8010              |                |
+-----------------+--------------+------+-----+-------------------+----------------+


{
Create Table: CREATE TABLE `t_similar_problems` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `question` varchar(255) NOT NULL DEFAULT '' COMMENT '问题名称',
  `standerAsk` varchar(255) NOT NULL DEFAULT '' COMMENT '匹配标准问法',
  `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态，0- 待学习，1-添加为新问题，2-合并到标准问法，3-忽略',
  `addtime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modifytime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8
}



| t_knowledge_tree | CREATE TABLE `t_knowledge_tree` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `codeId` varchar(255) DEFAULT NULL COMMENT 'L1xxxxxx 这样的 code',
  `codeName` varchar(255) DEFAULT '' COMMENT '问题标题',
  `levelMain` varchar(50) DEFAULT NULL COMMENT 'L1  L2  L3  L4  ',
  `levelOne` varchar(255) DEFAULT NULL COMMENT '一级目录   L1xxxxx',
  `levelTwo` varchar(255) DEFAULT NULL COMMENT '二级目录   L2xxxxx',
  `levelThree` varchar(255) DEFAULT NULL COMMENT '三级目录   L3xxxxx',
  `levelFour` varchar(255) DEFAULT NULL COMMENT '四级目录   L4xxxxx',
  `standerdAsk` varchar(255) DEFAULT NULL COMMENT '标准问法',
  `answer` text COMMENT '答案',
  `relationList` text COMMENT '关联问题',
  `addUser` varchar(255) DEFAULT NULL COMMENT '添加人',
  `modifiedUser` varchar(255) DEFAULT NULL COMMENT '修改人',
  `mutiAnswerSet` tinyint(1) DEFAULT '0' COMMENT '多轮问答开关 ',
  `mutiAnsWerList` text COMMENT '多轮问答列表配置',
  `interfaceId` int(11) DEFAULT NULL COMMENT '接口 id ',
  `clickTimes` int(10) DEFAULT '0' COMMENT '点击次数',
  `createTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modifiedTime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `enabled` tinyint(1) DEFAULT NULL COMMENT '是否可用 ',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除标志',
  `reask` varchar(600) DEFAULT NULL,
  `reaskRelationList` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8
}
