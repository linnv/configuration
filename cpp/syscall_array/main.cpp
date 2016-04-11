#include <iostream>
#ifdef __i386
#define SYSCALL_ orig_eax
//#define  MACHINE_TYPE 32
#include "disabled_syscall_x32.h"
#else
#ifdef __x86_64
#define SYSCALL_ orig_rax
//#define  MACHINE_TYPE 64
#include "disabled_syscall_x64.h"
#endif
#endif

//#include <string.h>
using namespace std;
int main(int argc, char *argv[])
{
	/*
	const char *a[2]=
	//const string a[2]=
	{" fje"," jjj"};

	cout<<a[0]<<endl;
	cout<<a[1];
	*/	

	for (int i = 0; i < 315; ++i)
	{
		cout<<i<<"-->"<<syscallName[i]<<"\n";
	}
	//new code
	return 0;
}
