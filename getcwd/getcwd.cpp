#include <iostream>
#include <stdio.h>
using namespace std;
int main(int argc, char *argv[])

{
	char path [150];

	    if (getcwd(path, sizeof(path)) == NULL) {
	    	    cout<<"Fail to get the current working dir";
		    return 1;
	        }
	    printf("\n%s\n",path );
	return 0;
}
