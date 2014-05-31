#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include<fcntl.h>

#include<sys/syscall.h>
#include <sys/ptrace.h>
#include <sys/user.h>
#include <sys/reg.h>
#include <sys/types.h>
#include <sys/wait.h>

#include <iostream>
#include <fstream>

