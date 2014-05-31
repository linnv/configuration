#include <iostream>
#include <unistd.h>
using namespace std;
char a[1024*1024*100];
int main(int argc, char *argv[])
{
	/*
	for (int i = 0; i < 100000; ++i)
	{
		cout<<i<<endl;
	//	cout<<"\t";
	}
	*/
//	while(1);
	int a,b;
	cin>>a>>b;
	cout<<a+b<<endl;
	chdir("~/");
	return 0;
}
