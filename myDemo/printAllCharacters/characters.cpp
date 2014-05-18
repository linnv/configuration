#include <iostream>
using namespace std;
int main(int argc, char *argv[])
{
	int i =0;
	char c;
	// the numbers of visual characters are between 33 to 126, but you can set scope of i to get all of them
	for (int i = 33; i <= 126; ++i)
	{
		c = i;
		cout<<c<<"("<<i<<")"<<'\t';
	}
	cout<<endl;
	return 0;
}
