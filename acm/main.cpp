#ifdef __i386
#define SYSCALL_ orig_eax
//#define  MACHINE_TYPE 32
#include "disabled_syscall_x32.h"
#else
#ifdef __x86_64
#define SYSCALL_ orig_rax
//#define  MACHINE_TYPE 64
#include "disabled_syscall_x64.h"
#endif
#endif

#include <iostream>
#include <sys/syscall.h>
#include <sys/resource.h>
#include <sys/ptrace.h>
#include <sys/user.h>
#include <sys/reg.h>
#include <string.h>
#include <fcntl.h>
#include <sys/stat.h>

#include "include.h"
#include "mysql_connection.h"
#include "mysqlcpp.cpp"
#include "collection.cpp"

using namespace std;



int ReadTimeConsumption(pid_t pid);
int ReadMemoryConsumption(pid_t pid);
void updateConsumption(pid_t pid,Collection* col);
void daemon(void);
void helpInfo();

int writeFromString( string &fileName, const string& buffer, size_t count);
int readToString(string &fileName, string* str);
int startExecution(Collection *col);
int diffCasesJudge(Collection* col);
void find_next_nonspace(int * c1, int * c2, FILE * stdf, FILE * usrf,Collection * col);
enum exitStatus {COMPILING = 100000, ACCEPTED=100001, PRESENTATION_ERROR, TIME_LIMIT_ERROR, MEMORY_LIMIT_ERROR, WRONG_ANSWER, RUNTIME_ERROR, OUTPUT_LIMIT_ERROR, COMPILE_ERROR, SYSTEM_ERROR, VALIATOR_ERROR, EXIT_NORMALLY,SYSCALL_RESTRICTION,SEGMENTATION_FAULT};
int main(int argc, char *argv[])
{
	string argv1_str;
	bool outputInfo=true;

	if (argc >=2)
	{
		argv1_str= argv[1];
		if (argv1_str =="-d"||argv1_str =="--daemon")
		{
			outputInfo =false;	
		}
		else if (argv1_str =="-h"||argv1_str =="--help")
		{
			helpInfo();
			return 1;
		}
		else{
		   cout<<"invalid option -- '"<<argv1_str<<"'"<<endl;
		   cout<<"try 'guetoj_judger -h or --help' for more information"<<endl;
		   return 1;
		   }
	}
	//	else
	{
		if (!outputInfo)
		{
			daemon();
		}
		helpInfo();

		SQL* sqlconn =  new SQL();
		sqlconn->setHost("tcp://127.0.0.1:3306");
		sqlconn->setUser("root");
		sqlconn->setPasswd("a");
		sqlconn->connectSQL();
		sqlconn->useDatabase("goj");

		//char tmpsql[200];
		int caseCount;
		//long int totolTimeconsumption=0;
		//long int totolMemoryConsumption=0;
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
			if (outputInfo)
			{
				cout<<"dealing with "<<sqlconn->getRowsCount()<<" source code(s)"<<endl<<endl;

			}
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
				if (outputInfo)
				{
					cout<<"dealing with the "<<i<<"th "<<"user id:"<<col[i]->getUserId()<<" run id:"<<col[i]->getRunId()<<" problem id:"<<col[i]->getProblemId()<<endl;//" source code: \n"<<col[i]->getSourceCode()<<endl;
				}
				sqlconn->querySQL("SELECT * FROM tbl_language  WHERE language_id = "+std::to_string(col[i]->getLanguageId()));
				res = sqlconn->getResultSet();
				while(res->next()){
					/*
					 * get compiler_name,language fix,compiler_option, only one set
					 */
					tmp = res->getString("compiler_name");
					if (outputInfo)
					{
						//	cout<<"compiler name:"<<tmp<<endl;
					}
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
				//cout<<"source code: "<<endl<<content<<endl;
				writeFromString(file,content,content.length());
				//cout<<"demo"<<endl;
				if (col[i]->getLanguageId() == 4)// java
				{
					compileCommand = col[i]->getCompilerName() +" -Wall "+ file +" "+  col[i]->getCompilerOption() + " 2>" + errorfile;
				}
				else{
					//cout<<"start compiling "<<endl;
					compileCommand = col[i]->getCompilerName() +" -o "+ mainName+" "+ file +" 2>" +errorfile;
					//compileCommand = col[i]->getCompilerName() +" "+ file +" "+  col[i]->getCompilerOption()+" -o "+  mainName +"  2>" +errorfile;
				}
				if (outputInfo)
				{
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

				//cout<<"finish exectue compile command"<<endl;
				if(generateFile == NULL)
				{
					cout<<"error compiling"<<endl;
					/*
					 *compile failed
					 *
					 */
					readToString(errorfile,&strReadFromFile);
					strReadFromFile+=" from null";
					col[i]->setCompilerError(strReadFromFile);
					col[i]->setLastState(COMPILE_ERROR);
					/*
					 *delete the source file and error file when it's needless
					 */
					errorfile = "./"+errorfile;
					delteComand="rm "+file + " " +errorfile;
					system(delteComand.c_str());
					//cout<<"time used:"<<col[i]->getTimeComsupted()<<"ns   memory used:"<<col[i]->getMemoryComsupted()<<"kb state: "<<col[i]->getLastState()<<endl;
					//cout<<"last result of run ID "<<col[i]->getRunId()<<" time consumption:"<<col[i]->getTimeConsumption()<<" memory consumption: "<<
					//	col[i]->getMemoryConsumption()<<" status: "<<col[i]->getLastState()<<endl<<endl;

					sql::PreparedStatement *pstm=sqlconn->con->prepareStatement("UPDATE tbl_run SET Status = ?,compile_error = ?  WHERE Run_ID = ?");
					pstm->setInt(1,col[i]->getLastState());
					pstm->setString(2,col[i]->getCompilerError());
					//pstm->setInt(2,totolTimeconsumption);
					//pstm->setInt(2,col[i]->getTimeComsupted());
					//pstm->setInt(3,col[i]->getMemoryComsupted());
					//pstm->setInt(3,totolMemoryConsumption);
					//pstm->setString(4,col[i]->getCompilerError());
					pstm->setInt(3,col[i]->getRunId());
					pstm->executeUpdate();
					delete pstm;

					//totolTimeconsumption = 0;

					/*
					//sprintf(tmpsql,"UPDATE tbl_run SET Status = %d, Time_Used = %ld, Memory_Used = %ld ,compile_error = '%s'  WHERE Run_ID = %ld",col[i]->getLastState(),col[i]->getTimeComsupted(),col[i]->getMemoryComsupted(),col[i]->getCompilerError().c_str(),col[i]->getRunId());
					sprintf(tmpsql,"UPDATE tbl_run SET Status = %d,compile_error = '%s'  WHERE Run_ID = %ld",col[i]->getLastState(),col[i]->getCompilerError().c_str(),col[i]->getRunId());
					//cout<<"status: "<<col[i]->getLastState()<<"compile error: "<<col[i]->getCompilerError()<<endl;
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

					int tmpStatu=0;

					while(res->next()){
						/*
						 * use every test data to test the executive app
						 * default:
						 * stdIn:file stored data for exectuive app's stdin
						 * stdOut:file stored std answer
						 * userOut:file stored the answer redirect from executive app's stdout
						 */
						col[i]->setTestcaseID(res->getInt("testcase_id"));
						tmp = res->getString("input");
						col[i]->setSTDIput(tmp);
						tmp = res->getString("output");
						col[i]->setSTDOutput(tmp);
						stdFile = "./tmp/stdIn";
						writeFromString(stdFile,col[i]->getSTDIput(),col[i]->getSTDIput().length());
						stdFile = "./tmp/stdOut";
						writeFromString(stdFile,col[i]->getSTDOutput(),col[i]->getSTDOutput().length());
						if (outputInfo)
						{
							cout<<"stdIn: \n"<<col[i]->getSTDIput()<<endl;
							cout<<"stdOut: \n"<<col[i]->getSTDOutput()<<endl;
						}
						/*
						 *compile done start to  run user app and redirect the stdout to the file userOut file
						 *here will get the time consumption and the memory consumption etc.
						 */
						startExecution(col[i]);
						/*
						 * executation done, now get the executation result
						 */
						stdFile = "./tmp/userOut";
						strReadFromFile.clear();
						readToString(stdFile,&strReadFromFile);
						if (outputInfo)
						{
							cout<<"userOut: \n"<<strReadFromFile<<endl;
						}
						if ( col[i]->getJudgeState() != EXIT_NORMALLY)
							//if ( col[i]->getJudgeState() != EXIT_NORMALLY)
						{
							if (outputInfo)
							{

								cout<<"app running with failure and statu is: "<<col[i]->getJudgeState()<<endl;
							}

							col[i]->setLastState(col[i]->getJudgeState() );
							//dontSetAccptedNextTime = true;
							break;

						}
						else
						{
							/*
							 * ok the app can run normally now check the app's answer with the std answer: comparing the conten of  userOut and the conten of stdOut
							 *
							 */
							if (outputInfo)
							{
								cout<<"befor diff judge status: "<<col[i]->getJudgeState()<<endl;
							}

							diffCasesJudge(col[i]);
							if (outputInfo)
							{
								cout<<"after diff judge status: "<<col[i]->getJudgeState()<<endl;
							}

							if (col[i]->getJudgeState() !=ACCEPTED ) {
								//wtd
								//select * from tbl_run_testcase where run_id=%ld and testcase_id=%d
								tmpStatu =col[i]->getJudgeState();

								sqlconn->querySQL("select * from tbl_run_testcase where run_id="+std::to_string(col[i]->getRunId()) +" and testcase_id= "+std::to_string(col[i]->getTestcaseID()));
								//sql::ResultSet *res = sqlconn->getResultSet();
								//cout<<"dealing with "<<sqlconn->getRowsCount()<<" source code(s)"<<endl<<endl;
								if (sqlconn->getRowsCount())
								{
									if (outputInfo)
									{
										cout<<"case record already exist"<<endl;
									}

								}
								else{
									sql::PreparedStatement *pstm=sqlconn->con->prepareStatement("INSERT INTO tbl_run_testcase  SET run_id = ?,testcase_id=?");
									pstm->setInt(1,col[i]->getRunId());
									/*
									   pstm->setInt(2,totolTimeconsumption);
									   pstm->setInt(3,totolMemoryConsumption);
									   */
									pstm->setInt(2,col[i]->getTestcaseID());
									pstm->executeUpdate();
									cout<<"WA case record!"<<endl;
									delete pstm;
								}






								/*
								 * judge statu now can not be exit normally
								 * just check it's AC or not
								 *insert into tbl_run_testcase
								 *set last statu to WA
								 */
							}
							else{
								col[i]->setLastState(col[i]->getJudgeState());
							}


						}
						/*
						 * get the exection condicton data:time comsuption memory cumsuption etc.
						 * delete the file preparing for next test case
						 *don't have to delete but set them to empty?
						 */
						delteComand="rm ./tmp/stdIn ./tmp/userOut ./tmp/stdOut";
						system(delteComand.c_str());
						/*
						 * use the max consumiption

						 *if choose the max consumption, the following will not need anymore,just
						 modiry the colleciont's setTimeConsumption() and setMemoryConsumption()
						 *
						 totolTimeconsumption +=col[i]->getTimeComsupted();
						 totolMemoryConsumption +=col[i]->getMemoryComsupted();
						 */
						if (outputInfo)
						{
							cout<<"run ID:"<<col[i]->getRunId()<<" testcase ID:"<<col[i]->getTestcaseID()<<" time used:"<<col[i]->getTimeConsumption()<<"ms   memory used:"<<col[i]->getMemoryConsumption()<<"kb judge state: "<<col[i]->getJudgeState()<<endl;
						}//all testcases are tested
					}

					/*
					 * delete the source code file error file and execute file when they are needless
					 * for example: rm ./Main main.cpp error
					 */
					delteComand="rm "+mainName+" "+file + " " +errorfile;
					system(delteComand.c_str());

					//totolTimeconsumption /=(caseCount*1000);   //milis but not micros?
					//totolMemoryConsumption /=(caseCount*1024); //kb or M?
					/*
					 *write result to database using 
					 */
					if (tmpStatu)
					{
						col[i]->setLastState(tmpStatu);

					}



					col[i]->setTimeConsumption(1);
					sql::PreparedStatement *pstm=sqlconn->con->prepareStatement("UPDATE tbl_run SET Status = ?, Time_Used = ?, Memory_Used = ? ,compile_error =?  WHERE Run_ID = ?");
					//	cout<<"error coneten:   "<<col[i]->getCompilerError()<<endl;
					pstm->setInt(1,col[i]->getLastState());
					/*
					   pstm->setInt(2,totolTimeconsumption);
					   pstm->setInt(3,totolMemoryConsumption);
					   */
					pstm->setInt(2,col[i]->getTimeConsumption());
					pstm->setInt(3,col[i]->getMemoryConsumption());


					pstm->setString(4,col[i]->getCompilerError());
					pstm->setInt(5,col[i]->getRunId());
					pstm->executeUpdate();
					delete pstm;
					cout<<"time limitation: "<<col[i]->getTimeLimit()<<" memory limitation: "<<col[i]->getMemoryLimit()<<endl;
					cout<<"last result of run ID "<<col[i]->getRunId()<<" time consumption:"<<col[i]->getTimeConsumption()<<" memory consumption: "<<
						col[i]->getMemoryConsumption()<<" status: "<<col[i]->getLastState()<<endl<<endl;
					//totolTimeconsumption = 0;
					//totolMemoryConsumption=0;
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
	if (fd == NULL)
	{
		cout<<"open file to write faild"<<endl;
		//		return 1;
	}
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
	{

		int pid;
		int status;
		//long tmp;

		//		cout<<"start executing: time limitation: "<<col->getTimeLimit()<<" memory limitation: "<<col->getMemoryLimit()<<endl;

		struct rlimit executableLimit;

		pid = fork();
		if (pid)
		{
			//timeConsumption=0;
			//memoryConsumption=0;	

			struct user_regs_struct regs;
			int firstExecute = 1;
			/*
			 * comment by: Jialin Wu
			 * refer to ZOJ,but not completely same with ZOJ.[https://code.google.com/p/zoj/source/browse/trunk/judge_client/client/tracer.c]
			 * judge time exceedance by signal SIGXCPU,if time consumption exceexed,set time consumption to zero and memory consumption to zero.
			 * judge memory exceedance by comparing memory consumption and memoryLimitaion every time getting memory consumption  from proc/$pid/status,if memory consumption is greater than the memoryLimitaion,use ptrace(PTRACE_KILL,pid,NULL,NULL) to stop the user app, and set time consumption to zero and memory consumption to zero.
			 */
			while(waitpid(pid,&status,0) > 0){
				updateConsumption(pid,col);

				if (WIFSIGNALED(status))
				{
					if(WTERMSIG(status) == SIGKILL){
						cout<<"signal kill"<<endl;
						col->setJudgeState(RUNTIME_ERROR);
						//col->setLastState(RUNTIME_ERROR);

					}
					break;
				}
				if(!WIFSTOPPED(status)){
					if (WTERMSIG(status))
					{
						col->setJudgeState(RUNTIME_ERROR);
						//col->setLastState(RUNTIME_ERROR);

					}	
					break;
				}
				/*
				 * comment by: Jialin Wu
				 *copy from referrence manual: While being traced, the tracee will stop each time a signal is delivered, even if the signal is being ignored. (An exception is SIGKILL, which has its usual effect.) The tracer will be notified at its next call to waitpid(2) (or one of the related "wait" system calls); that call will return a status value containing information that indicates the cause of the stop in the tracee. While the tracee is stopped, the tracer can use various ptrace requests to inspect and modify the tracee. The tracer then causes the tracee to continue, optionally ignoring the delivered signal (or even delivering a different signal instead).
				 */
				int sig =  WSTOPSIG(status);
				if (sig != SIGTRAP)
				{
					updateConsumption(pid,col);	
					if (sig == SIGXCPU || (col->getTimeConsumption() > col->getTimeLimit()))
					{

						col->setJudgeState(TIME_LIMIT_ERROR);
						printf("time exceeded\n");
						//timeConsumption =(timeLimitation)*1000+1;
						col->setTimeConsumption(col->getTimeLimit()+1);
						break;
						//col->setTimeConsumption((col[i]->getTimeLimit()+1));
					}
					else if( sig == SIGXFSZ){
						col->setJudgeState(OUTPUT_LIMIT_ERROR);
						break;
						//col->setLastState(OUTPUT_LIMIT_ERROR);

					}
					/*
					   else if( sig == SIGKILL){
					   programStatus = RUNTIME_ERROR;

					   }	
					   */
					else if( sig == SIGILL){
						col->setJudgeState(RUNTIME_ERROR);
						break;
						//col->setLastState(RUNTIME_ERROR);

					}
					else if( sig == SIGSEGV){

						col->setJudgeState(SEGMENTATION_FAULT);
						break;
						//col->setLastState(SEGMENTATION_FAULT);

					}
					/*
					   else
					   {
					   cout<<"unknow error"<<endl;
					   col->setJudgeState(RUNTIME_ERROR);
					   break;
					//col->setLastState(RUNTIME_ERROR);

					}
					*/
					ptrace(PTRACE_SYSCALL, pid, NULL, sig);
				}
				   if (ReadTimeConsumption(pid) >= 2*col->getTimeLimit());
				   {
					   cout<<"TLC: ts:"<<ReadTimeConsumption(pid)<<endl;
				   col->setJudgeState(TIME_LIMIT_ERROR);
				   col->setTimeConsumption(col->getTimeLimit()+1);
				   ptrace(PTRACE_KILL,pid,NULL,NULL);	
				   break;

				   }	
				if (ReadMemoryConsumption(pid) >= col->getMemoryLimit())
				{
					col->setJudgeState(MEMORY_LIMIT_ERROR);
					col->setTimeConsumption(col->getMemoryLimit()+1);
					//col->setLastState(MEMORY_LIMIT_ERROR);

					//memoryConsumption = memoryLimitation+1;
					/*
					 * don't use ptrace_kill wati fix
					 */
					ptrace(PTRACE_KILL,pid,NULL,NULL);	
					break;
				}
				ptrace(PTRACE_GETREGS,pid,NULL,&regs);
				/*
				 *comment by: Jialin Wu
				 *regs.orig_rax(in ubuntu_64bit) or regs.orig_ax(int centOS_32bit) will hold the syscall num that was invoked.
				 *
				 *compare with the diabled_syscall array in "disabled_syscall.h" to judge whether the syscall invoked is permitted or not,if the values of disabled_syscall[regs.orig_rax]==1,the syscall is forbidden.
				 */
				/*
				 * for x32
				 if (disabled_syscall[regs.orig_eax]==1)
				 * for x64
				 if (disabled_syscall[regs.orig_rax]==1)
				 */
				if (disabled_syscall[regs.SYSCALL_]==1)
					//if (disabled_syscall[regs.orig_rax]==1)
				{
					if (firstExecute)
					{
						/*
						 *comment by: Jialin Wu
						 * the first execv comes from from parent process
						 * so first exec is permitted to run user app
						 */
						firstExecute =0;
						ptrace(PTRACE_SYSCALL,pid,NULL,NULL);	
					}
					else{
						//the app is do somethin evil,so kill it!
						col->setJudgeState(SYSCALL_RESTRICTION);

						ptrace(PTRACE_KILL,pid,NULL,NULL);	
						break;
					}

				}	
				if (regs.SYSCALL_ == SYS_exit ||regs.SYSCALL_ ==SYS_exit_group)
				{

					/*
					 * comment by:Jialin Wu
					 * if the user app is exit normally,it will call SYS_exit or SYS_exit_group
					 * detect these two syscall to judge the app is exit normally or not
					 */

					updateConsumption(pid,col);
					col->setJudgeState(EXIT_NORMALLY);
					//	col->setJudgeState(EXIT_NORMALLY);


				}
				ptrace(PTRACE_SYSCALL,pid,NULL,NULL);
			}


		}
		else
		{
			ptrace(PTRACE_TRACEME,0,NULL,NULL);
			freopen("./tmp/stdIn", "r", stdin);
			freopen("./tmp/userOut", "w+", stdout);

			if (col->getLanguageId() ==4)
			{
				execl("/usr/java/bin/java", "/usr/java/bin/java","Main", (char *) NULL);
			}
			else{
				if ( getrlimit(RLIMIT_AS,&executableLimit) == 0)
				{
					// executableLimit.rlim_cur = 2 * 1024;
					// byte
					executableLimit.rlim_cur =  col->getMemoryLimit() * 1024;
					//	if (setrlimit(RLIMIT_AS, &executableLimit) == 0)
					{
						//	cout<<"set memory limit done!"<<endl;
					}
				}
				if ( getrlimit(RLIMIT_CPU,&executableLimit) == 0)
				{
					//	cout<<"time limit:"<<col->getTimeLimit()<<endl;
					//executableLimit.rlim_cur = 2;
					//second
					if (col->getTimeConsumption() > 1000)
					{
						executableLimit.rlim_cur = (col->getTimeLimit()/1000);
					}
					else
					{
						executableLimit.rlim_cur =1;
					}
					if (setrlimit(RLIMIT_CPU, &executableLimit) == 0)
					{
						//		cout<<"set time limit done!"<<endl;
					}
				}

				execl("./Main", "./Main", (char *) NULL);
			}
		}
	}
}

int diffCasesJudge(Collection* col){

	FILE *stdf, *usrf;
	col->setJudgeState(ACCEPTED);
	//col->setLastState(ACCEPTED); //puzzle whern comment this line state become 10011
	stdf = fopen("./tmp/stdOut", "r");
	usrf = fopen("./tmp/userOut", "r");
	if (stdf == NULL || usrf == NULL)
	{
		col->setJudgeState(RUNTIME_ERROR);
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
								col->setJudgeState(WRONG_ANSWER);
							}
						}
						if (c2 == EOF)
						{
							c1 = fgetc(stdf);
							if (!isspace(c1))
							{
								col->setJudgeState(WRONG_ANSWER);
							}
						}
						break;
					}
					c1 = toupper(c1);
					c2 = toupper(c2);
					if (c1 != c2)
					{
						col->setJudgeState(WRONG_ANSWER);
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
						col->setJudgeState(WRONG_ANSWER);
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
				col->setJudgeState(PRESENTATION_ERROR);
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
int ReadTimeConsumption(pid_t pid){
	char buffer[64];
	sprintf(buffer,"/proc/%d/stat",pid);
	FILE* fp = fopen(buffer,"r");
	if (fp == NULL)
	{
		//		printf("no stat found in proc\n");
		return -1;
	}
	int stime,utime;
	int d1;
	char buff[24],c1;
	//	fgetc(fp);
	if (fscanf(fp,"%d %s %c %*d %*d %*d %*d %*d %*u %*u %*u %*u %*u %d %d",&d1,buff,&c1,&utime,&stime) < 2)
	{
		printf("fail read time\n");
		fclose(fp);
		return -1;
		/* code */
	}
	//	printf("%d,%s,%c,%d,%d\n",d1,buff,c1,utime,stime );
	fclose(fp);
	static int clktck = 0;
	if (clktck == 0)
	{
		//Inquire about the number of clock ticks per second
		//the number of clock ticks per second
		clktck = sysconf(_SC_CLK_TCK); 
	}
	//unit: million second
	return (int) (((double)(stime + utime + 0.0) / clktck) * 1000);
}

int ReadMemoryConsumption(pid_t pid){
	char buffer[64];
	sprintf(buffer,"/proc/%d/status",pid);
	FILE* fp = fopen(buffer, "r");
	if (fp == NULL) {   
		return -1;
	}
	int vmPeak = 0,
	    vmSize = 0,
	    vmExe = 0,
	    vmLib = 0,
	    vmStack = 0;


	while (fgets(buffer, 32, fp)) 
	{ 
		if (!strncmp(buffer, "VmPeak:", 7)) 
		{ 
			sscanf(buffer + 7, "%d", &vmPeak);

		} else if (!strncmp(buffer, "VmSize:", 7)) 
		{
			sscanf(buffer + 7, "%d", &vmSize);

		} else if (!strncmp(buffer, "VmExe:", 6)) 
		{ 
			sscanf(buffer + 6, "%d", &vmExe);

		} else if (!strncmp(buffer, "VmLib:", 6)) 
		{           
			sscanf(buffer + 6, "%d", &vmLib);

		} else if (!strncmp(buffer, "VmStk:", 6)) 
		{            
			sscanf(buffer + 6, "%d", &vmStack);

		}

	}
	//		printf("VmPeak:%d VmSize:%d\n",vmPeak,vmSize );
	fclose(fp);
	if(vmPeak){
		vmSize = vmPeak;	
	}
	//		printf("vmsize:%d,vmExe:%d,vmLib:%d,vmStack:%d\n",
	//vmSize,vmExe,vmLib,vmStack );
	//unit:kb
	return vmSize - vmExe -vmLib -vmStack;
}

void updateConsumption(pid_t pid,Collection* col){

	col->setMemoryConsumption(ReadMemoryConsumption(pid));
	col->setTimeConsumption(ReadTimeConsumption(pid));

	/*
	   if (ReadMemoryConsumption(pid) > memoryConsumption)
	   {
	   memoryConsumption = ReadMemoryConsumption(pid);
	   }

	   if (ReadTimeConsumption(pid) > timeConsumption)
	   {
	   timeConsumption = ReadTimeConsumption(pid);
	   }	
	   */
}
void daemon(void) {    
	pid_t  pid;

	/*
	 *      * Become a session leader to lose controlling TTY.
	 *           */

	if ((pid = fork()) < 0) {        perror("fork");
		exit(1);
	} else if (pid != 0) /* parent */
		exit(0);
	setsid();

	/*
	 *      * Change the current working directory to the root.
	 *           * if you are using database  just ignore this and comment them.
	 *                */

	/*if (chdir("/") < 0) {     perror("chdir");
	 *      exit(1);
	 *           }
	 *                */
	/*
	 *      * Attach file descriptors 0, 1, and 2 to /dev/null.
	 *           */

	umask(0);
	close(0);
	close(1);
	close(2);

	open("/dev/null", O_RDWR);

	dup2(0, 1);
	dup2(0, 2);
}




void helpInfo(){

	cout<<"GUETOJ Judger help information\nMandatory arguments to long options are mandatory for short options too."<<endl;
	cout<<" -v, --verbose,none              verbosely list running detail"<<endl;
	cout<<" -d, --daemon               run GUETOJ_Judger as daemon process"<<endl;


}
