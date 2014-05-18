#include "include.h"
#define MAXLINE	1024
using namespace std;
//void infoExchangeWrite(int,char*);
void infoExchangeRead(int);

void pipeWriteCstyle(int file,char* cstring);
void pipeReadCstyle(int file);

void pipeReadCstyle(int file,struct rSet *result);
void pipeReadCstyle(int fd,char** str);
void pipeWriteCstyle(int fd,char* str);

struct rSet{
	int statu;
	int timeConsumption;
	int memoryConsumption;
};
int main(int argc, char *argv[])
{
	int piped[2];
	pid_t childpid;
	pipe(piped);
	if ((childpid = fork()) == 0)
	{
		/*
		 * childe can only write so close the reading port
		 */
		close(piped[0]);
		struct rSet wRet;
		wRet.statu = 81000;
		wRet.memoryConsumption=1111;
		wRet.timeConsumption=222;
		/*
		char arr[]="demoString";
		*/
		char stringChild[MAXLINE];
		sprintf(stringChild,"%d %d %d",wRet.statu,wRet.timeConsumption,wRet.memoryConsumption);
		pipeWriteCstyle(piped[1],stringChild);

		//pipeWriteCstyle(piped[1],stringChild);
		/*
		string stringChild;
		//sprintf(stringChild,"chile pid is %d\nsecond line\n",getpid());
		//stringChild +=" "+getpid();
		stringChild =std::to_string(getpid())+"\n";  //int to string
		cout<<"string:"<<stringChild<<endl;
		infoExchangeWrite(piped[1],stringChild);
		infoExchangeWrite(piped[1],stringChild);
		*/
		exit(0);
	}
	/*
	 * parent can only read so close the writing port
	 */
	close(piped[1]);
	/*
	string readstr;
	readToString(piped[0],&readstr);
	cout<<"from pipe:"<<stoi(readstr)<<endl;  //string to int
	*/
//	char buff[MAXLINE];
	struct rSet result;
	pipeReadCstyle(piped[0],&result);
	cout<<"statu:"<<result.statu<<" time:"<<result.timeConsumption<<" memory:"<<result.memoryConsumption<<endl;
	waitpid(childpid,NULL,0);
	
	return 0;
}
/*=================cString with struct======================*/
void pipeReadCstyle(int file,struct rSet *result){
	FILE* pstream = fdopen(file,"r");
	//if (fscanf(pstream,"%s%d",getArr,&getInt) == EOF)
	if (fscanf(pstream,"%d%d%d",&result->statu,&result->timeConsumption,&result->memoryConsumption) == EOF)
	{

	}
	//cout<<"in read function statu:"<<(*result).statu<<" time:"<<result->timeConsumption<<" memory:"<<result->memoryConsumption<<endl;
	//printf("demo getInt:%d getArr:%s\n",getInt,getArr );
}

void pipeWriteCstyle(int file,struct rSet &result){
	char cstring[MAXLINE];
	sprintf(cstring,"%d %d %d",result.statu,result.timeConsumption,result.memoryConsumption);
	FILE* pWStream = fdopen(file,"w");
	fprintf(pWStream, "%s",cstring);

}
/*=================end cString with struct======================*/

/*=================cString========================*/
void pipeWriteCstyle(int file,char* cstring){
	FILE* pWStream = fdopen(file,"w");
	fprintf(pWStream, "%s",cstring);

}
/*=================end cString======================*/

/*=================cppString======================*/
void pipeWriteString(int fd,const string  &str){
	/*
	char buff[MAXLINE];
	sprintf(buff,"write to pipe: %s\n",str);
	write(fd,buff,sizeof(buff));
	*/
//	cout<<"writing string:"<<str<<endl;
	write(fd,str.c_str(),str.length());
}
int pipeReadString(int pipeName, string* str){

	*str = "";
	FILE *fd = fdopen(pipeName,"rb");
	if (fd == NULL)
	{
		cout<<"can't open file: "<<*str<<endl;
	}
	char c;
	while((c = fgetc(fd)) != EOF){
	*str +=c;
	}
	fclose(fd);
	return (*str).length();

}

/*=================end cppString======================*/
