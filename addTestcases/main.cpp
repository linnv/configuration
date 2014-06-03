#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
using namespace std;
int writeFromString( string &fileName, const string& buffer, size_t count);
int readToString(string &fileName, string* str);
int main(int argc, char *argv[])
{
	
	SQL* sqlconn =  new SQL();
	sqlconn->setHost("tcp://127.0.0.1:3306");
	sqlconn->setUser("root");
	sqlconn->setPasswd("a");
	sqlconn->connectSQL();
	sqlconn->useDatabase("goj");
	string inputStr;
	string outputStr;
	string inputFile = "./testCasesFolder/";
	string outputFile = "./testCasesFolder/";
	string sql;
	int i;
	for (i = 0; i < 5; ++i)
	{
		inputStr.clear();
		outputStr.clear();

		inputFile = "./testCasesFolder/";
		outputFile = "./testCasesFolder/";
		inputFile+="input" +std::to_string(i);
		outputFile+="output" +std::to_string(i);
		readToString(inputFile,&inputStr);
	readToString(outputFile,&outputStr);
	cout<<"inputStr: "<<inputStr<<endl;
	cout<<"outputStr: "<<outputStr<<endl;
	sql.clear();
	sql="INSERT INTO `tbl_testcase_problem` (`testcase_id` ,`problem_id` ,`user_id` ,`input` ,`output` ,`status`)VALUES (NULL,'1071','10000','"+inputStr+"','"+outputStr+"','1')";

	//cout<<"sql string: "<<sql<<endl;

	sqlconn->updateSQL(sql);
	}
	cout<<i<<" testcases have been added"<<endl;
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
