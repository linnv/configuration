#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
#include "collection.cpp"
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
	

//	sqlconn->querySQL("select * from id where id =2");
	sqlconn->querySQL("SELECT * FROM tbl_run WHERE Status = 100000 AND Auto_Judge = 1 ");
	//sqlconn->querySQL("SELECT Run_ID, Problem_ID, User_ID, Language_ID, Source_Code FROM tbl_run WHERE Status = 100000 AND Auto_Judge = 1 ");
	sql::ResultSet *res = sqlconn->getResultSet();
	cout<<"colunm num: "<<sqlconn->getColunmsCount()<<endl;
	//num of new collection according to the rows
	cout<<"row count:"<<sqlconn->getRowsCount()<<endl<<endl;

	int waitintCount = sqlconn->getRowsCount();

	Collection *col[waitintCount];
	for (int i = 0; i < waitintCount; ++i)
	{
		col[i] = new Collection();

	}
	int count = waitintCount;
	string tmp;
	int tmpInt;
	while (count){				
			{

				res->next();
				tmpInt = res->getInt("Run_ID");
				col[waitintCount-count]->setRunId(tmpInt);

				/*
				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);

				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);
				*/
				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);


			//cout<<res->getString("name")<<" ";
			//	str+=res->getString("name");
				//cout<<res->getString(i)<<" ";
				count--;
			}
	}
	for (int i = 0; i < waitintCount; ++i)
	{
		cout<<i<<"run id:"<<col[i]->getRunId()<<" source code: \n"<<col[i]->getSourceCode()<<endl;

	}

	/*
	string tmp;
	while (res->next()){				
		for (int i = 1; i <= sqlconn->getRowsCount(); ++i)
			{
				tmp = res->getString("name");
				col->setSourceCode(tmp);
				//cout<<res->getString("name")<<" ";
			//	str+=res->getString("name");
				//cout<<res->getString(i)<<" ";
			}
		cout<<endl;
	}
	cout<<col->getSourceCode()<<endl;
	
	//insert
	sqlconn->updateSQL("insert into id set id = 27,name = 'jialin'");
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
