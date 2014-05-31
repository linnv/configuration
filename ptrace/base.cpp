#include "include.h"
#include "functions.cpp"
#include "createp.c"
#include "wait.c"
using namespace std;
int main(int argc, char *argv[])
{
	struct SettingForProcess sfp;
	/*
	sfp.timeLimit_ = 1;    //in seconds
	sfp.memoryLimit_ = 1000;  //in kbs
	sfp.vmLimit_ = 1000;      // in kbs
	sfp.stackLimit_ = 8*1024;   //in kbs
	sfp.outputLimit_ = 0;   //file size in kbs
	*/

	sfp.timeLimit_ = -1;    //in seconds
	sfp.memoryLimit_ = -1000;  //in kbs
	sfp.vmLimit_ = -1000;      // in kbs
	sfp.stackLimit_ = -8*1024;   //in kbs
	sfp.outputLimit_ =-10;   //file size in kbs
	
	sfp.beTraced_ = 1;     // true to be trace or not false
	
	int pid = CreateProcessAndDoExecute(sfp);
	cout<<"main pid"<<getpid()<<endl;
	cout<<"app pid "<<pid<<endl;

	RunProgram(pid);
	
	return 0;
}
