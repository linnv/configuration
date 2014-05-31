#include <stdio.h>
#include <glog/logging.h>
using namespace std;
int main(int argc, char *argv[])
{
	google::InitGoogleLogging((const char*)argv[0]);
	LOG(INFO)<<"glog test by fjewjfwjfjweafknkdn"<<endl;
	return 0;
}
