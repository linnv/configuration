[
CREATE USER 'qnzs'@'127.0.0.1' IDENTIFIED BY 'Qnzs@2018';
GRANT ALL PRIVILEGES ON *.* TO 'qnzs'@'127.0.0.1' IDENTIFIED BY 'Qnzs@2018';
FLUSH PRIVILEGES;
]
	sqlQuery := `select coalesce(standerdAsk,'') as one ,coalesce(levelFour,'')as two from  t_knowledge_tree`


[

delete from t_business_info where enterpriseID='qy20000000022';
delete from t_business_info where enterpriseID='20180802';
delete from t_business_info where enterpriseID='30000000022';
delete from t_business_info where enterpriseID='20190101'    ;
delete from t_business_info where enterpriseID='40000000022' ;
delete from t_business_info where enterpriseID='50000000022' ;
delete from t_business_info where enterpriseID='60000000022' ;
delete from t_business_info where enterpriseID='2018071701'  ;
delete from t_business_info where enterpriseID='2018071702'  ;
delete from t_business_info where enterpriseID='9000000006'  ;
delete from t_business_info where enterpriseID='19000000006' ;
delete from t_business_info where enterpriseID='89000000016' ;
delete from t_business_info where enterpriseID='90000000016' ;
delete from t_business_info where enterpriseID='8890000000016';
delete from t_business_info where enterpriseID='9890000000016' ;
delete from t_business_info where enterpriseID='qy10000000051';
delete from t_business_info where enterpriseID='qy100000011'   ;
delete from t_business_info where enterpriseID=' '              ;
delete from t_business_info where enterpriseID='1'             ;
]

[
drop database  qy100000010    ;
drop database qy100000011    ;
drop databse 10000000051;
]
use qy100000011;
select * from t_ai_water;
select * from t_similar_problems;
select * from t_unsolved_questions;



update-alternatives: using /etc/mysql/mysql.cnf to provide /etc/mysql/my.cnf (my.cnf) in auto mode
