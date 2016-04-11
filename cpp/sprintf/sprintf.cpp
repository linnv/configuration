#include <iostream>
#include "include.h"
using namespace std;

#define name "name"
char com[100];
char comin[100];
int main(int argc, char *argv[])
{
	sprintf(com,"%s",name);
	printf("%s\n",com );
	sprintf(comin,"%s: from com",com);
	printf("%s\n",comin );
	return 0;
}
