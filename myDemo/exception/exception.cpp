#include <iostream>
#include <string>
using namespace std;
void test(string &name,int &age){
	try{
	if (age <=0||age > 200)
	{
		throw  age;
	}	
	if (name.length() < 2|| name.length() > 10)
	{
		throw "name's length is illegal!";
	}
	}catch(int i){
		cout<<"illegal age!"<<endl;	
	}catch(const char * message){
		cout<<message<<endl;	
	}
}
int main(int argc, char *argv[])
{
       string Name;
       int Age;
	cin>>Name;       
	cin>>Age;
	test(Name,Age);
	return 0;
}
