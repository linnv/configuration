#include <iostream>
using namespace std;
class A{
	int a;
	public:
	A(int b){
		a =b;	
	}
};
int main(int argc, char *argv[])
{

	for ( int i = 0; i < 10000; ++i)
	{
	cout<<"\nnow app's here"<<endl;
	}

	/*
		int a,b;
	cin>>a>>b;
	cout<<a+b<<endl;
cout<<"\n\nin app 2"<<endl;
//	char o[10*1000*1000];
	for (int i = 0; i < 10000; ++i)
	{
//		A *a = new A(1);
		int* d = new int(1);
		cout<<endl;
	}
//	while(1);
//	execl("/bin/ls","ls",NULL);
*/
	return 0;
}
