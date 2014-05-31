#include <iostream>
#include <algorithm>
using namespace std;
void reverst(string & s){
	int len = s.length();
	s.reserve();
	for (int i = 0; i < len; ++i)
	{
		
		/* code */
	}
}
int main(int argc, char *argv[])
{
	wstring str;
	cin>>str;
	wcout<<"before"<<endl<<str<<endl;
	reverse(str.begin(),str.end() );
	wcout<<str<<endl;
	return 0;
}
