//
//  main.cpp
//  traceClass
//
//  Created by wujialin on 11/3/14.
//  Copyright (c) 2014 wujialin. All rights reserved.
//

//#include <iostream>
#include "trace.h"
#include "include.h"
#include "trace.cpp"
using namespace std;
int main(int argc, const char * argv[])
{
    trace t;
    t.SetLimitInfo();
    t.CreateProcess();
    t.waitForChild();
    

    cout<<"time usage: "<<t.getTime()<<"ms memory usage: "<<t.GetMemory()<<"kb"<<endl;
    // insert code here...
//    std::cout << "Hello, World!\n";
    return 0;
}

