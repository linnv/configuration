#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
#include "collection.cpp"
#include <iostream>
//#include "judge_environment.cpp"
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
	int caseCount;
	long int totolTimeconsumption=0;
	long int totolMemoryConsumption=0;
	string sql;
	string strReadFromFile;
	string stdFile;
	string content="";
	string file= "";
    	string compileCommand;
	string delteComand;
	string mainName;
	string errorfile;
    FILE* generateFile;
    
    /*
     * loop for headp problems
     */
    while(1){
	    
	/*
	 * query problem fromdatabase
	 */
	sqlconn->querySQL("SELECT * FROM tbl_run WHERE Status = 100000 AND Auto_Judge = 1 ");
	sql::ResultSet *res = sqlconn->getResultSet();
	//cout<<"colunm num: "<<sqlconn->getColunmsCount()<<endl;
	//cout<<"row count:"<<sqlconn->getRowsCount()<<endl<<endl;
	/*
	 * get a heap problems once time
	 */
	cout<<"dealing with "<<sqlconn->getRowsCount()<<" source code(s)"<<endl<<endl;

	int waitintCount = sqlconn->getRowsCount();
	int count = waitintCount;
	string tmp;
	int tmpInt;

	Collection *col[waitintCount];
	/*
	 * new collection object occording to the count of problem head
	 */
	for (int i = 0; i < waitintCount; ++i)
	{
		col[i] = new Collection();

	}

	while (count){	
		/*
		 * get run_id,user_id,problem_id,source_code,language_id to each collection object
		 */
			{
				res->next();
				tmpInt = res->getInt("Run_ID");
				col[waitintCount-count]->setRunId(tmpInt);
				tmpInt = res->getInt("User_ID");
				col[waitintCount-count]->setUserId(tmpInt);
	        		tmpInt = res->getInt("Problem_ID");
				col[waitintCount-count]->setProblemId(tmpInt);
				tmp = res->getString("Source_Code");
				col[waitintCount-count]->setSourceCode(tmp);
				tmpInt = res->getInt("language_id");
				col[waitintCount-count]->setLanguageId(tmpInt);
				count--;
			}
	}
	for (int i = 0; i < waitintCount; ++i){
        /*
         *now dealing with one by one
         */
		cout<<i<<"user id:"<<col[i]->getUserId()<<" run id:"<<col[i]->getRunId()<<" problem id:"<<col[i]->getProblemId()<<endl;//" source code: \n"<<col[i]->getSourceCode()<<endl;
		sqlconn->querySQL("SELECT * FROM tbl_language  WHERE language_id = "+std::to_string(col[i]->getLanguageId()));
		res = sqlconn->getResultSet();
		while(res->next()){
			/*
			 * get compiler_name,language fix,compiler_option, only one set
			 */
			tmp = res->getString("compiler_name");
		//	cout<<"compiler name:"<<tmp<<endl;
			col[i]->setCompilerName(tmp);
			tmp = res->getString("source_suffix");
			col[i]->setSCodeSuffix(tmp);
			tmp = res->getString("compiler_option");
			col[i]->setCompilerOption(tmp);
			}

		//cout<<i<<"user id:"<<col[i]->getUserId()<<" run id:"<<col[i]->getRunId()<<" problem id:"<<col[i]->getProblemId()<<" source code: \n"<<col[i]->getSourceCode()<<endl;
		/*
		 * every souce code file will be renamed to Main.suffix
		 * and the compile error file to compile_error
		 */
		mainName = "Main";
		errorfile="compile_error";
		content = col[i]->getSourceCode();
		file = mainName+col[i]->getSCodeSuffix();
		/*
		 * write souce code to file Main.suffix (mainName)
		 */
		writeFromString(file,content,content.length());
		if (col[i]->getLanguageId() == 4)// java
		{
			compileCommand = col[i]->getCompilerName() +" "+ file +" "+  col[i]->getCompilerOption() + " 2>" + errorfile;
		}
		else{
			compileCommand = col[i]->getCompilerName() +" "+ file +" "+  col[i]->getCompilerOption()+" -o "+  mainName +"  2>" +errorfile;
			cout<<"compile command is: "<<endl<<compileCommand<<endl;
		}
		
		/*
         //write error to database
	    if(strReadFromFile.length() >10){
		    cout<<"write error to sql,conten:"<<strReadFromFile<<endl;
		col[i]->setCompilerError(strReadFromFile);
	    }
	    else{
		    strReadFromFile = " ";
			col[i]->setCompilerError(strReadFromFile);
	    }
	    */
        
	    system(compileCommand.c_str());

         /*
          *rename mainName to execute app's name
          */
		if (col[i]->getLanguageId() ==4)
		{
		    mainName = "./Main.class";
		}
		else{
		    mainName ="./Main";
		}

        generateFile =fopen(mainName.c_str(),"r");
        
	    if(generateFile == NULL)
        {
            /*
             *compile failed
             *
             */
		readToString(errorfile,&strReadFromFile);
		strReadFromFile+=" from null";
		col[i]->setCompilerError(strReadFromFile);
		    col[i]->setState(COMPILE_ERROR);
            /*
             *delete the source file and error file when it's needless
             */
            errorfile = "./"+errorfile;
		    delteComand="rm "+file + " " +errorfile;
		    system(delteComand.c_str());
		//cout<<"time used:"<<col[i]->getTimeComsupted()<<"ns   memory used:"<<col[i]->getMemoryComsupted()<<"kb state: "<<col[i]->getState()<<endl;

		sql::PreparedStatement *pstm=sqlconn->con->prepareStatement("UPDATE tbl_run SET Status = ?,compile_error = ?  WHERE Run_ID = ?");
		pstm->setInt(1,col[i]->getState());
		pstm->setString(2,col[i]->getCompilerError());
		//pstm->setInt(2,totolTimeconsumption);
		//pstm->setInt(2,col[i]->getTimeComsupted());
		//pstm->setInt(3,col[i]->getMemoryComsupted());
		//pstm->setInt(3,totolMemoryConsumption);
		//pstm->setString(4,col[i]->getCompilerError());
		pstm->setInt(3,col[i]->getRunId());
		pstm->executeUpdate();
		delete pstm;

		    /*
		//sprintf(tmpsql,"UPDATE tbl_run SET Status = %d, Time_Used = %ld, Memory_Used = %ld ,compile_error = '%s'  WHERE Run_ID = %ld",col[i]->getState(),col[i]->getTimeComsupted(),col[i]->getMemoryComsupted(),col[i]->getCompilerError().c_str(),col[i]->getRunId());
        sprintf(tmpsql,"UPDATE tbl_run SET Status = %d,compile_error = '%s'  WHERE Run_ID = %ld",col[i]->getState(),col[i]->getCompilerError().c_str(),col[i]->getRunId());
		//cout<<"status: "<<col[i]->getState()<<"compile error: "<<col[i]->getCompilerError()<<endl;
		//cout<<tmpsql<<endl;
			sql = tmpsql;
			//cout<<sql<<endl;
		sqlconn->updateSQL(tmpsql);
		*/
	    }
	    else{
		/*
		 * compile source code successfully
		 * now get time limit and memory limit for  test case preperation,only one set
		 */
		fclose(generateFile);
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
			caseCount = sqlconn->getRowsCount();
			cout<<"there are "<<caseCount<<"cases to test"<<endl;
			while(res->next()){
					/*
					 * use every test data to test the executive app
					 * default:
					 * stdIn:file stored data for exectuive app's stdin
					 * stdOut:file stored std answer
					 * userOut:file stored the answer redirect from executive app's stdout
					 */
					tmp = res->getString("input");
					col[i]->setSTDIput(tmp);
					tmp = res->getString("output");
					col[i]->setSTDOutput(tmp);
					stdFile = "./stdIn";
					writeFromString(stdFile,col[i]->getSTDIput(),col[i]->getSTDIput().length());
					stdFile = "./stdOut";
					writeFromString(stdFile,col[i]->getSTDOutput(),col[i]->getSTDOutput().length());
					cout<<"stdOut: "<<col[i]->getSTDOutput()<<endl;
				    /*
                     *compile done start to  run user app and redirect the stdout to the file userOut file
                     *here will get the time consumption and the memory consumption etc.
                     */
					startExecution(col[i]);
					/*
					 * executation done, now get the executation result
					 */
					stdFile = "./userOut";
					strReadFromFile.clear();
					readToString(stdFile,&strReadFromFile);
					cout<<"userOut: "<<strReadFromFile<<endl;
					if ( col[i]->getState() != EXIT_NORMALLY)
					{
						//write to database;
					}
					else
					{
						/*
						 * ok the app can run normally now check the app's answer with the std answer: comparing the conten of  userOut and the conten of stdOut
						 *
						 */
						diffCasesJudge(col[i]);
					}
				   /*
				    * get the exection condicton data:time comsuption memory cumsuption etc.
				    * delete the file preparing for next test case
                    *don't have to delete but set them to empty?
				    */
				    delteComand="rm ./stdIn userOut stdOut";
				    system(delteComand.c_str());
                    /*
                     *if choose the max consumption, the following will not need anymore,just
modiry the colleciont's setTimeConsumption() and setMemoryConsumption()
                     */
				    totolTimeconsumption +=col[i]->getTimeComsupted();
				    totolMemoryConsumption +=col[i]->getMemoryComsupted();
				    cout<<"test case time used:"<<col[i]->getTimeComsupted()<<"ns   memory used:"<<col[i]->getMemoryComsupted()<<"kb state: "<<col[i]->getState()<<endl;
				}
			/*
			 * delete the source code file error file and execute file when they are needless
             * for example: rm ./Main main.cpp error
			 */
		    delteComand="rm "+mainName+" "+file + " " +errorfile;
		    system(delteComand.c_str());
            
			totolTimeconsumption /=(caseCount*1000);   //milis but not micros?
			totolMemoryConsumption /=(caseCount*1024); //kb or M?
			/*
			 *write result to database using 占位符;
			 */
			sql::PreparedStatement *pstm=sqlconn->con->prepareStatement("UPDATE tbl_run SET Status = ?, Time_Used = ?, Memory_Used = ? ,compile_error =?  WHERE Run_ID = ?");
		//	cout<<"error coneten:   "<<col[i]->getCompilerError()<<endl;
			pstm->setInt(1,col[i]->getState());
			pstm->setInt(2,totolTimeconsumption);
			pstm->setInt(3,totolMemoryConsumption);
			pstm->setString(4,col[i]->getCompilerError());
			pstm->setInt(5,col[i]->getRunId());
			pstm->executeUpdate();
			delete pstm;
			cout<<"tTime: "<<totolTimeconsumption<<"ms tMemory: "<<totolMemoryConsumption<<"M"<<endl;
			 totolTimeconsumption = 0;
			 totolMemoryConsumption=0;
			}
	    delete col[i];
	    col[i] = NULL;
        /*
         *ok,deal with one problem done!
         */
		}

	sleep(3);
    }//end loop
		cout<<"cloing sql"<<endl;
		sqlconn->closeSQL();
		return 0;
}
int readToString(string &fileName, string* str){

	*str = "";
	FILE *fd = fopen(fileName.c_str(),"rb");
	if (fd == NULL)
	{
		cout<<"can't open file: "<<*str<<endl;
	}
	char c;
	while((c = fgetc(fd)) != EOF){
	*str +=c;
	}
	fclose(fd);

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
	 //microseconds:us
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
	    	freopen("./stdIn", "r", stdin);
       		 freopen("./userOut", "w+", stdout);

		 if (col->getLanguageId() ==4)
		 {
            		execl("/usr/java/bin/java", "/usr/java/bin/java","Main", (char *) NULL);
		 }
		 else{
	if ( getrlimit(RLIMIT_AS,&executableLimit) == 0)
	{
		executableLimit.rlim_cur = 2 * col->getMemoryLimit() * 1024;
	//	if (setrlimit(RLIMIT_AS, &executableLimit) == 0)
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
            execl("./Main", "./Main", (char *) NULL);
		 }
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
            {
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
                            }
                        }
                        if (c2 == EOF)
                        {
                            c1 = fgetc(stdf);
                            if (!isspace(c1))
                            {
                                col->setState(WRONG_ANSWER);
                            }
                        }
                        break;
                    }
                    c1 = toupper(c1);
                    c2 = toupper(c2);
                    if (c1 != c2)
                    {
                        col->setState(WRONG_ANSWER);
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
                        goto end;
                    }
                    if ((c1 == '\n' || !c1) && (c2 == '\n' || !c2))
                    {
                        break;
                    }
                }
    //for end
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
