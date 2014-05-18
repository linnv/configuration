#include <iostream>
using namespace std;
void swap(int *soure, int *dest){
	//You can treat the soure and the dest as each array's first address
	while(*soure != 0){
	*dest = *soure;
	soure++;
	dest ++;
	}
	
}
int main(int argc, char *argv[])
{
	int s[10]={1,2,3,4,5,6,7,8,9,0};
	int d[10];
	swap(s,d);
	for (int i = 0; i < 9; ++i)
	{
		cout<<s[i]<<" ";
		cout<<d[i]<<" ";
		/* code */
		cout<<endl;
	}
	return 0;
}
