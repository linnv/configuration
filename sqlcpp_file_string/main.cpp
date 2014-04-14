#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
using namespace std;
int writenString( string &fileName, const string& buffer, size_t count);
int readString(string &fileName, string* str);
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
	sqlconn->initSQL();
	sqlconn->useDatabase("test");
	
	sqlconn->querySQL("select * from id");
	sql::ResultSet *res = sqlconn->getResultSet();

	cout<<"colunm num: "<<sqlconn->getColunmsCount()<<endl;
	cout<<"row count:"<<sqlconn->getRowsCount()<<endl;
	string str="";
	string file= "./in";
	
	while (res->next()){				
		for (int i = 1; i <= sqlconn->getRowsCount(); ++i)
			{
			//	cout<<res->getString("name")<<" ";
				str+=res->getString("name");

				//cout<<res->getString(i)<<" ";
			}
					cout<<endl;
			    }
	
	writenString(file,str,str.length());

	sqlconn->closeSQL();
/*
	ifstream in ("./in");
	stringstream buffer;
	buffer<<in.rdbuf();
	buffer<<"jialin\ndemo\n";
	string str(buffer.str());
	cout<<str<<endl;
		string str = "jialin' demo this's good\n";
	
	string file= "./in";
	writenString(file,str,str.length());
	string nstr;
	readString(file,&nstr);
	cout<<"length of nstr: "<<nstr.length()<<endl;
	cout<<nstr<<endl;
*/
	return 0;
}
int readString(string &fileName, string* str){

	FILE *fd = fopen(fileName.c_str(),"rb");
	char c;
	while((c = fgetc(fd)) != EOF){
	*str +=c;	
	}	

	return (*str).length();

}
int writenString( string &fileName, const string& buffer, size_t count){
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
