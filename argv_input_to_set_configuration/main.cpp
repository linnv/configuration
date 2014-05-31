#include <iostream>
#include "include.h"
using namespace std;
int main(int argc, char *argv[])
{
	string argv1_str= argv[1];
	if (argv1_str!="jialin")
	{
		cout<<"second value:"<<argv1_str<<endl;
		cout<<"argc:"<<argc<<endl;
		cout<<"maybe you should input jialin as well"<<endl;
	}
	for (int i = 0; i <= argc; ++i)
	{
		cout<<"i:"<<i<<" "<<argv[i]<<'\t';
	}

	//new code
	return 0;
}
