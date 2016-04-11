#include "include.h"
#define MAXLINE	1024
using namespace std;

void pipeWriteCstyle(int file,char* cstring);

void pipeReadCstyle(int file,struct rSet *result);
void pipeWriteCstyle(int file,struct rSet &result);

struct rSet{
	int statu;
	int timeConsumption;
	int memoryConsumption;
};
int main(int argc, char *argv[])
{
	int piped[2];
	//pid_t childpid;
	pipe(piped);
	
	/*
	 * parent can only read so close the writing port
	 */
	char executeCommand[MAXLINE];
	char app[]="./app";
	cout<<"p pid is:"<<getpid()<<endl;
	cout<<"transfer pipe:"<<piped[1]<<endl;
	sprintf(executeCommand,"%s %d\n",app,piped[1]);
	system(executeCommand);
			
	close(piped[1]);
	struct rSet result;
	pipeReadCstyle(piped[0],&result);
	cout<<"1statu:"<<result.statu<<" time:"<<result.timeConsumption<<" memory:"<<result.memoryConsumption<<endl;

//	pipeReadCstyle(piped[0],&result);
//	cout<<"2statu:"<<result.statu<<" time:"<<result.timeConsumption<<" memory:"<<result.memoryConsumption<<endl;

//	close(piped[0]);

	//waitpid(childpid,NULL,0);
	
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
