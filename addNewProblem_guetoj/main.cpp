#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
using namespace std;
int writeFromString( string &fileName, const string& buffer, size_t count);
int readToString(string &fileName, string* str);
int main(int argc, char *argv[])
{
	/*
	string str=" jialin wu";
	char a[str.length()];
	sprintf(a,"%s",str.c_str());
	printf("%s\n", a);
	cout<<sizeof(a)<<"str's length:"<<str.length()<<endl;
//	cout<<strlen(a)<<"str's length:"<<str.length()<<endl;
//	*/
	SQL* sqlconn =  new SQL();
	sqlconn->setHost("tcp://127.0.0.1:3306");
	sqlconn->setUser("root");
	sqlconn->setPasswd("a");
	sqlconn->connectSQL();
	sqlconn->useDatabase("goj");
	
	/*
	sqlconn->querySQL("select * from id");
	sql::ResultSet *res = sqlconn->getResultSet();
	cout<<"colunm num: "<<sqlconn->getColunmsCount()<<endl;
	cout<<"row count:"<<sqlconn->getRowsCount()<<endl;

	*/
	string file = "./code.txt";
	string nstr;
	//string sql = "update tbl_testcase_problem  set input = '"+nstr+"' where problem_id =1042";
	readToString(file,&nstr);
	//java
//	string sql = "INSERT INTO `tbl_run` (`run_id`, `problem_id`, `user_id`, `language_id`, `time_used`, `memory_used`, `source_code`, `submit_date`, `status`, `auto_judge`, `compile_error`) VALUES (NULL, '1000', '1000000', '4', '0', '0', '"+nstr+"', '2014-04-18 00:02:01', '100000', '1', NULL)";
//	//cpp
	string sql = "INSERT INTO `tbl_run` (`run_id`, `problem_id`, `user_id`, `language_id`, `time_used`, `memory_used`, `source_code`, `submit_date`, `status`, `auto_judge`, `compile_error`) VALUES (NULL, '1000', '1000000', '2', '0', '0', '"+nstr+"', '2014-04-18 00:02:01', '100000', '1', NULL)";
	//wrong
	//string sql = "INSERT INTO `tbl_run` (`run_id`, `problem_id`, `user_id`, `language_id`, `time_used`, `memory_used`, `source_code`, `submit_date`, `status`, `auto_judge`, `compile_error`) VALUES (NULL, '1000', '1000000', '1', '0', '0', '"+nstr+"', '2014-04-18 00:02:01', '100000', '1', NULL)";
	//
	int giveRecord= atoi(argv[1]);
	for (int i = 0; i < giveRecord; ++i)
	//sqlconn->updateSQL("insert into id set id = 27,name = 'jialin'");
//	int i =10;
//	while(i)
	{
	//	i--;
	
		cout<<"push record"<<endl;
	sqlconn->updateSQL(sql);
	//sleep(10);
	}	
	/*
	//insert
	//update
       	sqlconn->updateSQL("update id set name = 'new name' where id =8");
	*/


	/*
	string str="";
	string file= "./in";
	
	while (res->next()){				
		for (int i = 1; i <= sqlconn->getRowsCount(); ++i)
			{
				cout<<res->getString("name")<<" ";
			//	str+=res->getString("name");
				//cout<<res->getString(i)<<" ";
			}
		cout<<endl;
	}
	
	writeFromString(file,str,str.length());

	string nstr;
	readToString(file,&nstr);
	*/	
	sqlconn->closeSQL();
	return 0;
}
int readToString(string &fileName, string* str){

	FILE *fd = fopen(fileName.c_str(),"rb");
	char c;
	while((c = fgetc(fd)) != EOF){
	*str +=c;	
	}	

	return (*str).length();

}
int writeFromString( string &fileName, const string& buffer, size_t count){
	FILE *fd = fopen(fileName.c_str(),"wb+");
    const char*p = buffer.c_str();
//cout<<"*c_str: "<<buffer.c_str();
    while (count > 0 ) {      
   int num = fwrite(buffer.c_str(),sizeof(char),count,fd);
            if (num == -1) {    
   printf("Fail to write from file");
               return -1;
           }
            p += num;
            count -= num;
        }
    fclose(fd);
    return 0;
}
