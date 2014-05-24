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
	string file = "./in.txt";
	string nstr;
	readToString(file,&nstr);
	string sql = "update tbl_testcase_problem  set input = '"+nstr+"' where problem_id =1042";
	sqlconn->updateSQL(sql);
	//sqlconn->updateSQL("insert into id set id = 27,name = 'jialin'");
		
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
