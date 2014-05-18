#include <iostream>
#include "include.h"
using namespace std;
static int n= 10;
void get(int* i){
	
	*i =n;
}
int main(int argc, char *argv[])
{
	int a;
	get(&a);
	cout<<a<<endl;
	return 0;
}
