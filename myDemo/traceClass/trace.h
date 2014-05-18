//
//  trace.h
//  traceClass
//
//  Created by wujialin on 11/3/14.
//  Copyright (c) 2014 wujialin. All rights reserved.
//

#ifndef traceClass_trace_h
#define traceClass_trace_h

struct LimitInfo{
    int tl_;//time limit
    int vml_; //virtual memory limit
    int ml_;//memory limit
    int fsl_;//file size limit
    int sl_;//stack limit
};
class trace{
private:
    int pid_;
    struct LimitInfo limitInfo;
    int result_;
    int timeUsed_;
    int memoryUsed_;
    bool trace_;
public:
    void SetLimitInfo();
    void updateUsage();
    void waitForChild();
    bool CreateProcess();
    int getTime();
    int GetMemory();
    
    
};

#endif
