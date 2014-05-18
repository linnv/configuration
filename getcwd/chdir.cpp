#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>


using namespace std;
class A{
	public:
		A(int a,int b):a_(a),b_(b){}
	void 	print(){
			cout<<a_<< " | "<<b_<<endl;
		}
	private:
		int a_;
		int b_;
};

int main() { 
	char cdir[100];
	char pwd[100];
	strcpy(cdir,"/home/jialin/judge");
	printf("%s\n",cdir );
	if(!chdir(cdir))
			cout<<"change done!"<<endl;

	system("touch cdir");
		cout<<getcwd(pwd,sizeof(pwd))<<endl;
	/*
     	int a, b;
    while (cin>>a>>b) { 
     	    cout<<a + b<<endl<<endl;
        }
		
	   */ 
	
    return 0;
}

