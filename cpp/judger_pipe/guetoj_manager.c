#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <mysql/mysql.h>

#include <sys/types.h>
#include <sys/stat.h>
#include <dirent.h>

#define HOST "localhost"
#define GUETOJ "guetoj"
#define USER "guetoj"
#define PASSWORD "7758521@"

#define BUFFER_LENGTH 100000


enum exitStatus {COMPILING = 100000, ACCEPTED, PRESENTATION_ERROR, TIME_LIMIT_ERROR, MEMORY_LIMIT_ERROR, WRONG_ANSWER, RUNTIME_ERROR, OUTPUT_LIMIT_ERROR, COMPILE_ERROR, SYSTEM_ERROR, VALIATOR_ERROR, EXIT_NORMALLY};

char stmt[BUFFER_LENGTH];

MYSQL mysql;
MYSQL_RES *mysql_res;
MYSQL_ROW mysql_row;

int connect_mysql();
int initialize();

int finalize_mysql();

int finalize();

int connect_mysql()
{
    printf("Initializing mysql ...\n");
    
    if (!mysql_init(&mysql))
    {
        printf("Initializing mysql failed!\n");
        return 1;
    }

    printf("Connecting to mysql ...\n");
    if (!mysql_real_connect(&mysql, HOST, USER, PASSWORD, GUETOJ, 0, NULL, 0))
    {
        printf("Connecting mysql failed!\n");
        fprintf(stderr, "Error: %s\r\n", mysql_error(&mysql));
        return 1;
    }
    
    printf("Connecting mysql successed ...\n");

    mysql_query(&mysql, "SET NAMES 'utf-8'");

    return 0;
}

int initialize()
{
    if(connect_mysql() == 1)
    {
        return 1;
    }
    
    return 0;
}

int finalize_mysql()
{
    mysql_close(&mysql);
    return 0;
}

int finalize()
{
    finalize_mysql();
}
/*
**  argv[1]: options
**  avialable options are:
**  a :auto_judge
**  p :problem
**  r :run
**  h :help
**  s :status
*/

int main (int argc, char * argv[])
{
    int runID;
    int status;
    int auto_judge, problem, run, help;
    int i;

    int numAffected;

    auto_judge = 0;
    problem = 0;
    run = 0;
    help = 0;
    status = 0;

    if (argc < 3)
    {
        printf("Argument Error!\n");
        return 1;
    }

    if (initialize() == 1)
    {
        return 1;
    }

    for (i = 0; i < strlen(argv[1]); i++)
    {
        if (argv[1][i] == 'a' || argv[1][i] == 'A')
        {
            auto_judge = 1;
        }
        else if (argv[1][i] == 'p' || argv[1][i] == 'P')
        {
            problem = 1;
        }
        else if (argv[1][i] == 'r' || argv[1][i] == 'R')
        {
            run = 1;
        }
        else if (argv[1][i] == 's' || argv[1][i] == 'S')
        {
            status = 1;
        }
        else if (argv[1][i] == 'h' || argv[1][i] == 'H')
        {
            help = 1;
        }
    }
    
    printf("auto_judge = %d\nproblem = %d\nrun = %d\nstatus = %d\nhelp = %d\n", auto_judge, problem, run, status, help);


    if (problem)
    {
        if (auto_judge)
        {
            sprintf(stmt, "UPDATE Runs SET Status = %d WHERE Problem_ID = %d", COMPILING, problemID);
        }
    }
    else if (run)
    {
    }

    if (runID == 0)
    {
        sprintf(stmt, "UPDATE Runs SET Status = %d", COMPILING);
    }
    else
    {
        sprintf(stmt, "UPDATE Runs SET Status = %d WHERE Run_ID = %d", COMPILING, runID);
    }

    if (!mysql_query(&mysql, stmt))
    {
        printf("Operation Done!!!!\n");
    }
    else
    {
        printf("Operation Faile!!!!\n");
    }

    
    finalize();
    return 0;
}
