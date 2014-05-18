#include "include.h"
#define MAXLINE	1024
using namespace std;

void pipeReadToStruct(int file,struct rSet *result);
void pipeWriteFromStruct(int file,struct rSet &result);
struct rSet{
	int statu;
	int timeConsumption;
	int memoryConsumption;
};
int main(int argc, char *argv[])
{

	int piped=atoi(argv[1]);
	cout<<"pid is:"<<getpid()<<endl<<"pipe fd:"<<piped<<endl;
	struct rSet wRet;
	wRet.statu = 1000;
	wRet.memoryConsumption=81111;
	wRet.timeConsumption=8222;
	pipeWriteFromStruct(piped,wRet);
	//exit(0);

	return 0;
}
void pipeReadToStruct(int file,struct rSet *result){
	FILE* pstream;
	pstream = fdopen(file,"r");
	if (pstream==NULL)
	{
		printf("error occur opening pipe \n");
		exit(0);
	}

	//if (fscanf(pstream,"%s%d",getArr,&getInt) == EOF)
	if (fscanf(pstream,"%d%d%d",&result->statu,&result->timeConsumption,&result->memoryConsumption) == EOF)
	{

	}
	//cout<<"in read function statu:"<<(*result).statu<<" time:"<<result->timeConsumption<<" memory:"<<result->memoryConsumption<<endl;
	//printf("demo getInt:%d getArr:%s\n",getInt,getArr );
}

void pipeWriteFromStruct(int file,struct rSet &result){
	char cstring[MAXLINE];
	sprintf(cstring,"%d %d %d",result.statu,result.timeConsumption,result.memoryConsumption);
	FILE*pWStream; 
	pWStream = fdopen(file,"w");
	if (pWStream==NULL)
	{
		printf("error occur opening pipe \n");
		exit(0);
	}
	fprintf(pWStream, "%s",cstring);

}
