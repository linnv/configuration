#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
#include "collection.cpp"
#include "judge_environment.cpp"
using namespace std;
int writeFromString( string &fileName, const string& buffer, size_t count);
int readToString(string &fileName, string* str);
int startExecution(Collection *col);
int diffCasesJudge(Collection* col);
void find_next_nonspace(int * c1, int * c2, FILE * stdf, FILE * usrf,Collection * col);
enum exitStatus {COMPILING = 100000, ACCEPTED=100001, PRESENTATION_ERROR, TIME_LIMIT_ERROR, MEMORY_LIMIT_ERROR, WRONG_ANSWER, RUNTIME_ERROR, OUTPUT_LIMIT_ERROR, COMPILE_ERROR, SYSTEM_ERROR, VALIATOR_ERROR, EXIT_NORMALLY};
int main(int argc, char *argv[])
{
	SQL* sqlconn =  new SQL();
	sqlconn->setHost("tcp://127.0.0.1:3306");
	sqlconn->setUser("root");
	sqlconn->setPasswd("a");
	sqlconn->connectSQL();
	sqlconn->useDatabase("goj");
	
	char tmpsql[200];
	string sql;

//	sqlconn->querySQL("select * from id where id =2");
	//sqlconn->querySQL("SELECT Run_ID, Problem_ID, User_ID, Language_ID, Source_Code FROM tbl_run WHERE Status = 100000 AND Auto_Judge = 1 ");
	sqlconn->querySQL("SELECT * FROM tbl_run WHERE Status = 100000 AND Auto_Judge = 1 ");
	sql::ResultSet *res = sqlconn->getResultSet();
	cout<<"colunm num: "<<sqlconn->getColunmsCount()<<endl;
	//num of new collection according to the rows
	cout<<"row count:"<<sqlconn->getRowsCount()<<endl<<endl;

	int waitintCount = sqlconn->getRowsCount();

	Collection *col[waitintCount];
	for (int i = 0; i < waitintCount; ++i)
	{
		col[i] = new Collection();

	}
	int count = waitintCount;
	string tmp;
	int tmpInt;
	while (count){				
			{

				res->next();
				tmpInt = res->getInt("Run_ID");
				col[waitintCount-count]->setRunId(tmpInt);
				tmpInt = res->getInt("User_ID");
				col[waitintCount-count]->setUserId(tmpInt);
	        		tmpInt = res->getInt("Problem_ID");
				col[waitintCount-count]->setProblemId(tmpInt);



				/*
				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);

				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);
				*/
				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);


			//cout<<res->getString("name")<<" ";
			//	str+=res->getString("name");
				//cout<<res->getString(i)<<" ";
				count--;
			}
	}

	string content="";
	string file= "";
    string compileCommand;
	

	for (int i = 0; i < waitintCount; ++i)
	{
	cout<<i<<"user id:"<<col[i]->getUserId()<<" run id:"<<col[i]->getRunId()<<" problem id:"<<col[i]->getProblemId()<<endl;//" source code: \n"<<col[i]->getSourceCode()<<endl;
	content = col[i]->getSourceCode();	
	file = std::to_string(col[i]->getUserId()) + "_"+ std::to_string(col[i]->getRunId())+".cpp";
	writeFromString(file,content,content.length());
    compileCommand = "g++ "+ file + " -o  Main 2>1";
    system(compileCommand.c_str());
    //remove the source file when it's needless
    file = "rm "+file;
    system(file.c_str());
    FILE* executeFile =    fopen("./Main","r") ;
    if(executeFile == NULL){
    cout<<"compile error" <<endl;
    }
    else{
	sqlconn->querySQL("SELECT Time_Limit, Memory_Limit FROM tbl_problem WHERE status = 1 and Problem_ID ="+std::to_string(col[i]->getProblemId()));
	sql::ResultSet *res = sqlconn->getResultSet();
	while(res->next()){
		tmpInt = res->getInt("Time_Limit");
		col[i]->setTimeLimit(tmpInt);
		tmpInt = res->getInt("Memory_Limit");
		col[i]->setMemoryLimit(tmpInt);

	}	

	sqlconn->querySQL("SELECT * FROM tbl_testcase_problem  WHERE Problem_ID ="+std::to_string(col[i]->getProblemId()));
	res = sqlconn->getResultSet();
	while(res->next()){
		tmp = res->getString("input");
		col[i]->setSTDIput(tmp);
		tmp = res->getString("output");
		col[i]->setSTDOutput(tmp);

	string stdFile = "./in";
	writeFromString(stdFile,col[i]->getSTDIput(),col[i]->getSTDIput().length());
	stdFile = "./stdOut";
	writeFromString(stdFile,col[i]->getSTDOutput(),col[i]->getSTDOutput().length());


	    //compile done start to  judge
	startExecution(col[i]);

	string nstr;
	stdFile = "./userOut";
	readToString(stdFile,&nstr);
//	cout<<"input:"<<col[i]->getSTDIput()<<" output:"<<nstr<<endl;
	
	if ( col[i]->getState() != EXIT_NORMALLY)
	{
		//write to database;
	}
	else{
		//cout<<"start to diff judge"<<endl;
		diffCasesJudge(col[i]);
		//col[i]->setState( diffCasesJudge(col[i]));	

		}
//	cout<<"size"<<col[i]->getSTDIput().length()<<endl;
	cout<<"time used:"<<col[i]->getTimeComsupted()<<"ns   memory used:"<<col[i]->getMemoryComsupted()<<"b state: "<<col[i]->getState()<<endl;
		}	
	

	   	
	sprintf(tmpsql,"UPDATE tbl_run SET Status = %d, Time_Used = %ld, Memory_Used = %ld WHERE Run_ID = %ld",col[i]->getState(),col[i]->getTimeComsupted(),col[i]->getMemoryComsupted(),col[i]->getRunId());
	//sql ="UPDATE tbl_run SET status ="+ std::to_string(col[i]->getState()) + "time_used =" +std::to_string(col[i]->getTimeComsupted())+ " memory_used="+std::to_string(col[i]->getMemoryComsupted())+" WHERE Run_ID = " + std::to_string(col[i]->getRunId());
	sql = tmpsql;
	sqlconn->updateSQL(sql);
	    system("rm ./Main ./1 ") ;
	    //system("rm ./Main 1 in userOut stdOut") ;
    	}
	}

	/*
	string tmp;
	while (res->next()){				
		for (int i = 1; i <= sqlconn->getRowsCount(); ++i)
			{
				tmp = res->getString("name");
				col->setSourceCode(tmp);
				//cout<<res->getString("name")<<" ";
			//	str+=res->getString("name");
				//cout<<res->getString(i)<<" ";
			}
		cout<<endl;
	}
	cout<<col->getSourceCode()<<endl;

	string str="";
	string file= "./in";
	writeFromString(file,str,str.length());
	
	//insert
	sqlconn->updateSQL("insert into id set id = 27,name = 'jialin'");
	//update
       	sqlconn->updateSQL("update id set name = 'new name' where id =8");
	*/


	/*
	string nstr;
	readToString(file,&nstr);
	*/	

	sqlconn->closeSQL();
	return 0;
}
int readToString(string &fileName, string* str){

	FILE *fd = fopen(fileName.c_str(),"rb");
	char c;
	while((c = fgetc(fd)) != EOF){
	*str +=c;	
	}	

	return (*str).length();

}
int writeFromString( string &fileName, const string& buffer, size_t count){
	FILE *fd = fopen(fileName.c_str(),"wb+");
    const char*p = buffer.c_str();
//cout<<"*c_str: "<<buffer.c_str();
    while (count > 0 ) {      
   int num = fwrite(buffer.c_str(),sizeof(char),count,fd);
            if (num == -1) {    
   printf("Fail to write from file");
               return -1;
           }
            p += num;
            count -= num;
        }
    fclose(fd);
    return 0;
}


int startExecution(Collection * col){

    int pid;
    int status;
    float passTime;
    long tmp;
	struct rusage executableResourceUsage;    
    struct rlimit executableLimit;

    struct timeval before, after;
    struct timeval before1, after1;
     pid = fork();
    if (pid)
    {

         gettimeofday(&before1, NULL);
	 wait(&status);

	 //cout<<"status:"<<status<<endl;
	 // col->setState(RUNTIME_ERROR);
	  if (WIFEXITED(status))
	  {
		
	  	col->setState(EXIT_NORMALLY);
		
	  }
	  else if (WIFSIGNALED(status))
	  {
		if (SIGXCPU == WTERMSIG(status))
		{
			col->setState(TIME_LIMIT_ERROR);
		}
		else if ( SIGSEGV == WTERMSIG(status))
		{
			col->setState(MEMORY_LIMIT_ERROR);
		}
		else if ( SIGKILL == WTERMSIG(status))
		{
			col->setState(SYSTEM_ERROR);
		}

	  }

	 gettimeofday(&after1, NULL);
	tmp = ((long long)after1.tv_sec)*1000*1000 + 
              ((long long)after1.tv_usec) -
              ((long long)before1.tv_sec)*1000*1000 - 
              ((long long)before1.tv_usec);
	col->setTimeComsupted(tmp);
	 getrusage(RUSAGE_CHILDREN, &executableResourceUsage);
	col->setMemoryComsupted(executableResourceUsage.ru_maxrss);
	}
	 else
    	{
	if ( getrlimit(RLIMIT_AS,&executableLimit) == 0)
	{
		executableLimit.rlim_cur = 2 * col->getMemoryLimit() * 1024;
		if (setrlimit(RLIMIT_AS, &executableLimit) == 0)
		{
		//	cout<<"set memory limit done!"<<endl;
		}
	}	
	if ( getrlimit(RLIMIT_CPU,&executableLimit) == 0)
	{
	//	cout<<"time limit:"<<col->getTimeLimit()<<endl;
		executableLimit.rlim_cur = 2 * col->getTimeLimit()/1000 ;
		if (setrlimit(RLIMIT_CPU, &executableLimit) == 0)
		{
	//		cout<<"set time limit done!"<<endl;
		}
	}	
        
        freopen("./in", "r", stdin);
        freopen("./userOut", "w+", stdout);
        execl("./Main", " ", (char *) NULL);
    	}
}

int diffCasesJudge(Collection* col){

 FILE *stdf, *usrf;
    col->setState(ACCEPTED); //puzzle whern comment this line state become 10011 
    stdf = fopen("./stdOut", "r");
    usrf = fopen("./userOut", "r");
    if (stdf == NULL || usrf == NULL)
    {
	    col->setState(RUNTIME_ERROR);
        //currentExitStatus = RUNTIME_ERROR;
    }
    else
    {
for (; ;)
        {
            int c1 = fgetc(stdf);
            int c2 = fgetc(usrf);
            
            find_next_nonspace(&c1, &c2, stdf, usrf, col);
            
            for (; ;)
            {       //judge the answer one character by one character
                while ((!isspace(c1) && c1) || (!isspace(c2) && c2))
                {
                    if (c1 == EOF && c2 == EOF)
                    {
                        goto end;
                    }

                    if (c1 == EOF || c2 == EOF)
                    {
                        if (c1 == EOF)
                        {
                            c2 = fgetc(usrf);
                            if (!isspace(c2))
                            {
	    			col->setState(WRONG_ANSWER);
                               // currentExitStatus = WRONG_ANSWER;
                            }
                        }
                        if (c2 == EOF)
                        {
                            c1 = fgetc(stdf);
                            if (!isspace(c1))
                            {
	    			col->setState(WRONG_ANSWER);
                               // currentExitStatus = WRONG_ANSWER;
                            }
                        }
                        break;
                    }
                    c1 = toupper(c1);
                    c2 = toupper(c2);
                    if (c1 != c2)
                    {
	    		col->setState(WRONG_ANSWER);
                       // currentExitStatus = WRONG_ANSWER;
                        goto end;
                    }
                    c1 = fgetc(stdf);
                    c2 = fgetc(usrf);
                    find_next_nonspace(&c1, &c2, stdf, usrf, col);
                    if (c1 == EOF && c2 == EOF)
                    {
                        goto end;
                    }
                    if (c1 == EOF || c2 == EOF)
                    {
	    		col->setState(WRONG_ANSWER);
                       // currentExitStatus = WRONG_ANSWER;
                        goto end;
                    }
                    if ((c1 == '\n' || !c1) && (c2 == '\n' || !c2))
                    {
                        break;
                    }
                }
            }
        }

   }
    end:
    if (stdf)
    {
        fclose(stdf);
    }

    if (usrf)
    {
        fclose(usrf);
    }

    return 1;
}

void find_next_nonspace(int * c1, int * c2, FILE * stdf, FILE * usrf,Collection * col)
{
    while ((isspace(*c1)) || (isspace(*c2)))
    {
        if (*c1 != *c2)
        {
            if (*c2 == EOF)
            {
                do
                {
                    *c1 = fgetc(stdf);
                } while (isspace(*c1));
                continue;
            }
            else if (*c1 == EOF)
            {
                do
                {
                    *c2 = fgetc(usrf);
                } while (isspace(*c2));
                continue;
            }
            else if (*c1 == '\r' && *c2 == '\n')
            {
                *c1 = fgetc(stdf);
            }
            else
            {
		    col->setState(PRESENTATION_ERROR);
                //*currentExitStatus = PRESENTATION_ERROR;
            }
        }
        if (isspace(*c1))
        {
            *c1 = fgetc(stdf);
        }
        if (isspace(*c2))
        {
            *c2 = fgetc(usrf);
        }
    }
}
