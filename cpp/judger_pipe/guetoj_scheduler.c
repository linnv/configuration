#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <mysql.h>
//#include </usr/include/mysql/mysql.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <dirent.h>

#define ROOT "."
#define TMP "/tmp"
#define USER_ROOT "/user"
#define PROBLEM_ROOT "/problem"
#define PROBLEM_INPUT "/input_"
#define PROBLEM_OUTPUT "/output_"
#define COMPILE_ERROR_OUTPUT "compiler_error_"
#define RESULT_FILE_NAME "/result_"
#define DIFF_RESULT_FILE_NAME "/diff_result_"
#define HOST "localhost"
#define GUETOJ "goj"
#define USER "guetoj"
#define PASSWORD "guetoj"
#define JUDGER_NAME "./guetoj_judger"
#define TIMEMEMORYRESULT "/RESULT"

#define CONCURRENT_RUN_NUM 10
#define SLEEP_TIME (CONCURRENT_RUN_NUM)
#define BUFFER_LENGTH 65535
#define STATUS_LENGTH 50
#define LANGUAGE_LENGTH 50
#define COMPILER_OPTION 80
#define COMPILER_COMMAND_LENGTH 150
#define FILE_PATH_LENGTH 200
#define INPUT_LENGTH 65535
#define OUTPUT_LENGTH INPUT_LENGTH
#define COMPILE_BUFFER_LENGTH INPUT_LENGTH

#define UPDATE_RUN_STATUS "UPDATE tbl_run SET status = ?, compile_error = ? WHERE Run_ID = ?"

struct stStatusMeta
{
	int statusID;
	char statusName[STATUS_LENGTH];
} * statusList;
int numStatus;

struct stLanguageMeta
{
	int languageID;
	char languageName[LANGUAGE_LENGTH];
	char compilerName[LANGUAGE_LENGTH];
	char compilerOption[COMPILER_OPTION];
	char executableSuffix[LANGUAGE_LENGTH];
	char sourceSuffix[LANGUAGE_LENGTH];
} * languageList;
int numLanguage;

struct stRunMeta
{
	unsigned long runID;
	int languageID;
	int userID;
	int problemID;
	char sourceCode[BUFFER_LENGTH];

} * runList;
int numRun;

struct stTestcaseMeta
{
	int testcaseID;
	char input[INPUT_LENGTH];
	char output[OUTPUT_LENGTH];
} * testcaseList;
int numTestcase;

struct stTimeMemoryLimit
{
	int timeLimit;
	int memoryLimit;
	int problemID;
};

struct stTimeMemoryLimit timeMemory;

/*
   struct stRunStatus
   {
   long long timeUsed;
   long memoryUsed;
   long  runExitStatus;
   };
   */
struct stRunStatus
{
	int timeUsed;
	int memoryUsed;
	int runExitStatus;
};

int piped[2]; //IPC:pipe
void pipeReadToStruct(int file,struct stRunStatus *result);
void pipeWriteFromStruct(int file,const struct stRunStatus result);



struct stRunStatus currentRunStatus;
struct stRunStatus totalRunStatus;

enum exitStatus {COMPILING = 100000, ACCEPTED, PRESENTATION_ERROR, TIME_LIMIT_ERROR, MEMORY_LIMIT_ERROR, WRONG_ANSWER, RUNTIME_ERROR, OUTPUT_LIMIT_ERROR, COMPILE_ERROR, SYSTEM_ERROR, VALIATOR_ERROR, EXIT_NORMALLY,SYSCALL_RESTRICTION,SEGMENTATION_FAULT};

char buffer[BUFFER_LENGTH];
char stmt[BUFFER_LENGTH];
char compileError[BUFFER_LENGTH];

char compileErrorFileName[BUFFER_LENGTH];

char filePath[FILE_PATH_LENGTH];
char caseID[FILE_PATH_LENGTH];
char compilerCommand[COMPILER_COMMAND_LENGTH];

char tempPath[FILE_PATH_LENGTH];
char userBasePath[FILE_PATH_LENGTH];
char problemBasePath[FILE_PATH_LENGTH];

char userPath[FILE_PATH_LENGTH];
char problemPath[FILE_PATH_LENGTH];

MYSQL mysql;
MYSQL_RES *mysql_res;
MYSQL_ROW mysql_row;

//Following variables are used for prepared statement of MYSQL
MYSQL_STMT *mysql_stmt;
MYSQL_BIND bind[3];
int Readn(int fd, void* buffer, size_t count);
int Writen(int fd, const void* buffer, size_t count) ;
void fetchStatusList();
void fetchLanguageList();
int fetchRunInfo(int);
int connect_mysql();
int initialize();

int finalize_mysql();
int finalizeStatusList();
int finalizeLanguageList();
int finalizeRunInfo();
int finalizeTestcaseList();

int finalize();

char * generateProblemPath();
char * generateUserPath();
char * generateTempPath();
void generateBasePath();
struct stTestcaseMeta * generateTestcase(char * problemID);
char * getSourceCodeSuffix(int languageID);

char * getCompilerCommand(int languageID);
char * getProblemPath(char * problemID);
char * getUserPath(char * userDir);
char * getSourcePath(char * sourcePath);
void generateSourceCode(char * sourcePath, int runIndex);
int compileSourceCode(int runIndex);
struct stTimeMemoryLimit getTimeMemoryLimit(int problemID);

size_t get_file_size(const char *path);

void updateRunStatus(int runIndex, int status, struct stRunStatus runStatus);
char * getCompileError(int runIndex);
void executeRun(int runIndex);
int getRunStatus(int runIndex, struct stRunStatus * runStatus);
enum exitStatus diffResult(int runIndex, int testcaseIndex);

void fetchStatusList()
{
	int i;

	MYSQL_RES *mysql_res;
	MYSQL_ROW mysql_row;

	strcpy(stmt, "SELECT feedback_id, feedback_name FROM tbl_feedback");

	if (!mysql_query(&mysql, stmt))
	{
		printf("Starting fetch Status!\n");
	}

	mysql_res = mysql_store_result(&mysql);

	numStatus = mysql_num_rows(mysql_res);

	printf("There are %d records in the Feedbacks Table!\n", numStatus);

	statusList = (struct stStatusMeta *) malloc(sizeof(struct stStatusMeta) * numStatus);
	i = 0;
	while (mysql_row = mysql_fetch_row(mysql_res))
	{
		statusList[i].statusID = atoi(mysql_row[0]);
		strcpy(statusList[i].statusName, mysql_row[1]);
		printf("%d\t%s\n", statusList[i].statusID, statusList[i].statusName);
		i++;
	}

	printf("Finished fetch Status!\n");
}

void fetchLanguageList()
{
	int i;

	MYSQL_RES *mysql_res;
	MYSQL_ROW mysql_row;

	strcpy(stmt, "SELECT language_id, language_name, Compiler_Name, Compiler_Option, Executable_Suffix, Source_Suffix FROM tbl_language WHERE status = 1");

	if (!mysql_query(&mysql, stmt))
	{
		printf("Starting fetch Language!\n");
	}

	mysql_res = mysql_store_result(&mysql);

	numLanguage = mysql_num_rows(mysql_res);

	printf("There are %d records in the Language Table!\n", numLanguage);

	languageList = (struct  stLanguageMeta *) malloc(sizeof(struct stLanguageMeta) * numLanguage);

	i = 0;
	while (mysql_row = mysql_fetch_row(mysql_res))
	{
		languageList[i].languageID = atoi(mysql_row[0]);
		strcpy(languageList[i].languageName, mysql_row[1]);
		strcpy(languageList[i].compilerName, mysql_row[2]);
		strcpy(languageList[i].compilerOption, mysql_row[3]);
		strcpy(languageList[i].executableSuffix, mysql_row[4]);
		strcpy(languageList[i].sourceSuffix, mysql_row[5]);

		printf("%d\t%s\t%s\t%s\t%s\n", languageList[i].languageID, languageList[i].languageName,
				languageList[i].compilerName, languageList[i].executableSuffix, languageList[i].sourceSuffix);
		i++;
	}

	printf("Finished fetch Language!\n");
}

int fetchRunInfo(int runNum)
{
	int i;
	int codeLength;
	FILE * source_file;
	char * temp;
	char sourcePath[FILE_PATH_LENGTH];

	//    struct stTimeMemoryLimit timeMemory;

	sprintf(stmt, "SELECT Run_ID, Problem_ID, User_ID, Language_ID, Source_Code FROM tbl_run WHERE Status = 100000 AND Auto_Judge = 1 ORDER BY Run_ID ASC Limit 0, %d", runNum);

	if (!mysql_query(&mysql, stmt))
	{
		//        printf("mysql_query successed!\n");
	}

	mysql_res = mysql_store_result(&mysql);

	numRun = mysql_num_rows(mysql_res);

	printf("There are %d records found!\n", numRun);

	if (runList != NULL)
	{
		free(runList);
	}

	runList = (struct stRunMeta *) malloc(sizeof(struct stRunMeta) * numRun);

	i = 0;
	while (mysql_row = mysql_fetch_row(mysql_res))
	{
		runList[i].runID = atoi(mysql_row[0]);
		runList[i].problemID = atoi(mysql_row[1]);
		runList[i].userID = atoi(mysql_row[2]);
		runList[i].languageID = atoi(mysql_row[3]);
		strcpy(runList[i].sourceCode, mysql_row[4]);

		sprintf(sourcePath, "%d", runList[i].userID);

		temp = getUserPath(sourcePath);


		sprintf(sourcePath, "%s/%ld", temp, runList[i].runID);


		generateSourceCode(sourcePath, i);
		if (compileSourceCode(i) != 0)
		{
			sprintf(sourcePath, "%d", runList[i].problemID);
			generateTestcase(sourcePath);

			timeMemory = getTimeMemoryLimit(runList[i].problemID);

			executeRun(i);
		}
		else // syntax error encounted
		{
			updateRunStatus(i, 100008, currentRunStatus);
		}
		i++;
	}

	mysql_free_result(mysql_res);
	return numRun;
}

struct stTestcaseMeta * generateTestcase(char * problemID)
{
	int i;
	char * temp;

	MYSQL_RES * mysql_res;
	MYSQL_ROW mysql_row;

	FILE * input_file, * output_file;

	char testcaseInputPath[FILE_PATH_LENGTH];
	char testcaseOutputPath[FILE_PATH_LENGTH];

	sprintf(stmt, "SELECT Testcase_ID, Input, Output FROM tbl_testcase_problem WHERE Status = 1 AND Problem_ID = %s ORDER BY Testcase_ID", problemID);

	if (!mysql_query(&mysql, stmt))
	{
		//        printf("mysql_query successed!\n");
	}

	mysql_res = mysql_store_result(&mysql);

	numTestcase = mysql_num_rows(mysql_res);

	temp = getProblemPath(problemID);

	if (testcaseList != NULL)
	{
		free(testcaseList);
	}

	testcaseList = (struct stTestcaseMeta *) malloc (sizeof(struct stTestcaseMeta) * numTestcase);

	i = 0;
	while (mysql_row = mysql_fetch_row(mysql_res))
	{
		testcaseList[i].testcaseID = atoi(mysql_row[0]);
		strcpy(testcaseList[i].input, mysql_row[1]);
		//        strcpy(testcaseList[i].output, mysql_row[2]);
		//        printf("Testcase ID:%d\n", testcaseList[i].testcaseID);
		//        printf("Input: %s\n", testcaseList[i].input);
		//        printf("Output: %s\n", testcaseList[i].output);
		sprintf(testcaseInputPath, "%s%s%d", temp, PROBLEM_INPUT, i);
		sprintf(testcaseOutputPath, "%s%s%d", temp, PROBLEM_OUTPUT, i);

		input_file = fopen(testcaseInputPath, "w");

		if (input_file == NULL)
		{
			printf("Error creating file %s\n", testcaseInputPath);
		}
		else
		{
			fprintf(input_file, "%s\n", testcaseList[i].input);
			fclose(input_file);
			sprintf(stmt, "dos2unix %s\n", testcaseInputPath);
			system(stmt);
		}

		output_file = fopen(testcaseOutputPath, "w");

		if (output_file == NULL)
		{
			printf("Error creating file %s\n", testcaseOutputPath);
		}
		else
		{
			//@author lliu
			//@date 02/05/2014
			//@describtion  testcaseList not used in other places
			//
			//fprintf(output_file, "%s", testcaseList[i].output);
			fprintf(output_file, "%s", mysql_row[2]);
			fclose(output_file);
		}

		//        printf("%s\n", testcaseInputPath);
		//        printf("%s\n", testcaseOutputPath);
		i++;
	}


	mysql_free_result(mysql_res);
	printf("Testcases generated successfully!!!!\n");
	return testcaseList;
}

struct stTimeMemoryLimit getTimeMemoryLimit(int problemID)
{
	MYSQL_RES * mysql_res;
	MYSQL_ROW mysql_row;
	struct stTimeMemoryLimit timeMemory;

	sprintf(stmt, "SELECT Time_Limit, Memory_Limit FROM tbl_problem WHERE Problem_ID = %d", problemID);

	if (!mysql_query(&mysql, stmt))
	{
		//        printf("mysql_query successed!\n");
	}

	mysql_res = mysql_store_result(&mysql);

	mysql_row = mysql_fetch_row(mysql_res);

	if (mysql_row != NULL)
	{
		timeMemory.timeLimit = atoi(mysql_row[0]);
		timeMemory.memoryLimit = atoi(mysql_row[1]);
		timeMemory.problemID = problemID;

		/*
		   printf("Time and Memory Limit for Problem : %d TL = %d, ML = %d\n",
		   timeMemory.problemID, timeMemory.timeLimit, timeMemory.memoryLimit);
		   */
	}
	else
	{
		printf("Error encountted while fetching time and memory limit for problem %d\n",
				problemID);
	}

	mysql_free_result(mysql_res);
	return timeMemory;
}

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

	fetchStatusList();
	fetchLanguageList();
	generateBasePath();

	return 0;
}

int finalize_mysql()
{
	mysql_close(&mysql);
	return 0;
}

int finalizeStatusList()
{
	if (statusList != NULL)
	{
		free(statusList);
		statusList = NULL;
	}

	return 0;
}

int finalizeLanguageList()
{
	if (languageList != NULL)
	{
		free(languageList);
		languageList = NULL;
	}

	return 0;
}

int finalizeRunInfo()
{
	if (runList != NULL)
	{
		free(runList);
		runList = NULL;
	}

	return 0;
}

int finalizeTestcaseList()
{
	if (testcaseList != NULL)
	{
		free(testcaseList);
		testcaseList = NULL;
	}

	return 0;
}

int finalize()
{
	finalizeStatusList();
	finalizeLanguageList();
	// finalizeRunInfo();
	// finalizeTestcaseList();
	finalize_mysql();
}

char * getCompilerCommand(int languageID)
{
	int i;

	for (i = 0; i < numLanguage; i++)
	{
		if (languageList[i].languageID == languageID)
		{
			sprintf(compilerCommand, "%s %s", languageList[i].compilerName, languageList[i].compilerOption);
			break;
		}
	}

	return compilerCommand;
}

char * getSourceCodeSuffix(int languageID)
{
	int i;

	for (i = 0; i < numLanguage; i++)
	{
		if (languageList[i].languageID == languageID)
		{
			return languageList[i].sourceSuffix;
			break;
		}
	}

	return NULL;
}

char * generateTempPath()
{
	DIR * tempFolder;

	sprintf(tempPath, "%s%s", ROOT, TMP);

	tempFolder = opendir(tempPath);
	if (tempFolder == NULL)
	{
		mkdir(tempPath, 0777);
		printf("Creating folder: %s\n", tempPath);
	}
	else
	{
		closedir(tempFolder);
		// printf("%s exist!\n", tempPath);
	}
	return tempPath;
}

char * generateProblemPath()
{
	DIR * problemFolder;

	sprintf(problemBasePath, "%s%s%s", ROOT, TMP, PROBLEM_ROOT);

	problemFolder = opendir(problemBasePath);
	if (problemFolder == NULL)
	{
		mkdir(problemBasePath, 0777);
		//    printf("Creating folder: %s\n", problemPath);
	}
	else
	{
		closedir(problemFolder);
		// printf("%s exist!\n", problemPath);
	}
	return problemBasePath;
}

char * generateUserPath()
{
	DIR * userFolder;
	sprintf(userBasePath, "%s%s%s", ROOT, TMP, USER_ROOT);

	userFolder = opendir(userBasePath);

	if (userFolder == NULL)
	{
		mkdir(userBasePath, 0777);
		//    printf("Creating folder: %s \n", userPath);
	}
	else
	{
		closedir(userFolder);
		//printf("%s exist!\n", userPath);
	}
	return userBasePath;
}

void generateBasePath()
{
	generateTempPath();
	generateProblemPath();
	generateUserPath();
}

char * getUserPath(char * userDir)
{
	DIR * userFolder;

	sprintf(userPath, "%s/%s", userBasePath, userDir);

	userFolder = opendir(userPath);

	if (userFolder == NULL)
	{
		mkdir(userPath, 0777);
		//    printf("Creating folder: %s \n", userPath);
	}
	else
	{
		closedir(userFolder);
		// printf("%s exist!\n", userPath);
	}
	return userPath;
}

char * getProblemPath(char * problemID)
{
	DIR * problemFolder;

	sprintf(problemPath, "%s/%s", problemBasePath, problemID);

	problemFolder = opendir(problemPath);

	if (problemFolder == NULL)
	{
		mkdir(problemPath, 0777);
		//    printf("Creating folder: %s\n", problemPath);
	}
	else
	{
		closedir(problemFolder);
		//printf("%s exist!\n", problemPath);
	}
	return problemPath;
}

char * getSourcePath(char * sourcePath)
{
	DIR * sourceFolder;

	sourceFolder = opendir(sourcePath);

	if (sourceFolder == NULL)
	{
		mkdir(sourcePath, 0777);
		//    printf("Creating folder: %s \n", sourcePath);
	}
	else
	{
		closedir(sourceFolder);
		//printf("%s exist!\n", sourcePath);
	}
	return sourcePath;
}

void generateSourceCode(char * sourcePath, int runIndex)
{
	char temp[FILE_PATH_LENGTH];
	FILE * source_file;

	getSourcePath(sourcePath);

	sprintf(temp, "%s/%ld%s", 
			sourcePath, runList[runIndex].runID, 
			getSourceCodeSuffix(runList[runIndex].languageID));

	source_file = fopen(temp, "w");
	if (source_file == NULL)
	{
		printf("Error creating file %s\n", temp);
	}
	else
	{
		fprintf(source_file, "%s", runList[runIndex].sourceCode);
		//        printf("Creating file: %s\n", temp);
		fclose(source_file);
	}
}

int compileSourceCode(int runIndex)
{
	char sourceFilePath[FILE_PATH_LENGTH];
	char executableFileName[FILE_PATH_LENGTH];
	char sourceFileName[FILE_PATH_LENGTH];
	char compileError[FILE_PATH_LENGTH];

	FILE *execFile;
	sprintf(sourceFilePath, "%s/%d/%ld", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID);

	sprintf(executableFileName, "%s/%d/%ld/%ld", userBasePath,
			runList[runIndex].userID, runList[runIndex].runID,
			runList[runIndex].runID);

	sprintf(sourceFileName, "%s/%d/%ld/%ld%s", userBasePath,
			runList[runIndex].userID, runList[runIndex].runID,
			runList[runIndex].runID, getSourceCodeSuffix(runList[runIndex].languageID));

	//    printf("%s\n", executableFileName);
	//    printf("%s\n", sourceFileName);

	getCompilerCommand(runList[runIndex].languageID);
	sprintf(compileError, "%s/%s%ld", sourceFilePath, COMPILE_ERROR_OUTPUT, runList[runIndex].runID);

	//    printf("%s\n", compileError);
	//addjava
	if(runList[runIndex].languageID < 4){
		//compile command for c/gcc/g++
		printf("%s\n",sourceFileName);
		sprintf(stmt, "%s -o %s %s 2>%s\n", compilerCommand, executableFileName, sourceFileName, compileError);
		printf("\n%s\n", stmt);
	}
	if(runList[runIndex].languageID == 4){


		char tmpName[200];
		//once the class name must be same with the java file's name,so here will rename the sourcefileName to java.java
		sprintf(tmpName,"%s/Main.java",sourceFilePath);
		if(!rename(sourceFileName,tmpName)){
			printf("rename done!\n");	
		}
		//compile command for java
		sprintf(stmt, "%s %s  2>%s\n", compilerCommand, tmpName, compileError);
		/// printf("compile command for java:%s\n java file:%s",compilerCommand,tmpName);
		// sprintf(executableFileName,"%s.class",executableFileName);
		//    printf("%s\n",executableFileName);
	}
	//end
	//   printf("%s", stmt);

	system(stmt);

	if(runList[runIndex].languageID == 4){
		sprintf(executableFileName,"%s/Main.class",sourceFilePath);
	}
	execFile = fopen(executableFileName, "r");
	if (execFile == NULL)
	{
		printf("%ld Compiling error!\n", runList[runIndex].runID);
		return 0;
	}
	else
	{
		//    printf("%ld Executable file generated!\n", runList[runIndex].runID);
		fclose(execFile);
	}

	return 1;
}

char * getCompileErrorFileName(int runIndex)
{
	sprintf(compileErrorFileName, "%s/%d/%ld/%s%ld", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID,
			COMPILE_ERROR_OUTPUT, runList[runIndex].runID);

	return compileErrorFileName;
}

char * getCompileError(int runIndex)
{
	char compileErrorFileName[FILE_PATH_LENGTH];
	char line[4096];
	FILE * compileErrorFile;

	sprintf(compileErrorFileName, "%s/%d/%ld/%s%ld", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID,
			COMPILE_ERROR_OUTPUT, runList[runIndex].runID);

	compileErrorFile = fopen(compileErrorFileName, "r");

	printf("Compile Error File Name \n%s\n", compileErrorFileName);
	if (compileErrorFile == NULL)
	{
		return NULL;
	}
	else
	{
		memset(compileError, 0, BUFFER_LENGTH);
		printf("Start reading Error messages\n");
		while (fgets(line, 4096, compileErrorFile) != NULL)
		{
			sprintf(compileError, "%s%s", compileError, line);
		}
		//   printf("Reading Error messages Done!\n");
		fclose(compileErrorFile);
	}

	//printf("%s\n", compileError);
	return compileError;
}

size_t get_file_size(const char *path)
{
	unsigned long filesize = -1;
	struct stat statbuff;

	if (stat(path, &statbuff) < 0)
	{
		return filesize;
	}
	else
	{
		filesize = statbuff.st_size;
	}

	return filesize;
}

void updateRunStatus(int runIndex, int status, struct stRunStatus currentStatus)
{
	FILE *dest_file;

	MYSQL_RES *mysql_res;
	MYSQL_ROW mysql_row;

	size_t filesize, buf_length;

	char * ptrCompileError;
	char * buffer;

	if (status == COMPILE_ERROR)
	{
		//ptrCompileError = getCompileError(runIndex);

		filesize = get_file_size(getCompileErrorFileName(runIndex));
		if (filesize != -1)
		{
			printf("Compile Error Message fetched!\n");
		}

		buffer = (char *) malloc(sizeof(char)*(filesize+1));

		memset(buffer, 0, filesize+1);
		printf("File size :%d\n", filesize);

		if (buffer != NULL)
		{
			dest_file = fopen(getCompileErrorFileName(runIndex), "rb");
			if (dest_file != NULL)
			{
				fread(buffer, 1, filesize, dest_file);

				printf("Buffer length : %d\n", strlen(buffer));

				printf("File Content \n%s\n", buffer);
				mysql_stmt = mysql_stmt_init(&mysql);

				if (mysql_stmt != NULL)
				{
					if (!mysql_stmt_prepare(mysql_stmt, UPDATE_RUN_STATUS, strlen(UPDATE_RUN_STATUS)))
					{
						printf("Total %d parameters for mysql_stmt\n", mysql_stmt_param_count(mysql_stmt));

						memset(bind, 0, sizeof(bind));

						//binding the parameter for tbl_run.status
						bind[0].buffer_type = MYSQL_TYPE_LONG;
						bind[0].buffer = (char *) &status;
						bind[0].is_null = 0;
						bind[0].length = 0;

						//binding the parameter for tbl_run.compile_error
						bind[1].buffer_type = MYSQL_TYPE_BLOB;
						bind[1].buffer = (char *) buffer;
						bind[1].buffer_length = filesize+1;
						bind[1].is_null = 0;
						buf_length = filesize+1;
						bind[1].length = &buf_length;

						//binding the parameter for tbl_run.run_id
						bind[2].buffer_type = MYSQL_TYPE_LONG;
						bind[2].buffer = (char *) &runList[runIndex].runID;
						bind[2].is_null = 0;
						bind[2].length = 0;

						if (mysql_stmt_bind_param(mysql_stmt, bind))
						{
							printf("Bind param error!!\n");
						}
						else
						{
							if (mysql_stmt_execute(mysql_stmt))
							{
								printf("Execute prepared statement failed!!\n");
							}
							else
							{
								printf("Execute prepared statement finished!!\n");
								mysql_stmt_close(mysql_stmt);

								mysql_stmt = NULL;

								fclose(dest_file);
								free(buffer);
							}
						}
					}
					else
					{
						printf("mysql_stmt_prepare() failed!\n");
					}
				}
				else
				{
					printf("mysql_stmt_init() failed!\n");
				}

			}
		}
		/*
		 * commented by QIN Xingguo
		 sprintf(stmt, "UPDATE tbl_run SET Status = %d WHERE Run_ID = %ld",
		 status, runList[runIndex].runID);
		 */
		//sprintf(stmt, "UPDATE Runs SET Status = %d, Compile_Error = '%s' WHERE Run_ID = %ld",
		//   status, ptrCompileError, runList[runIndex].runID);
	}
	else
	{
		sprintf(stmt, "UPDATE tbl_run SET Status = %d, Time_Used = %lld, Memory_Used = %ld WHERE Run_ID = %ld", 
				status, currentStatus.timeUsed, currentStatus.memoryUsed,
				runList[runIndex].runID);
	}

	//    printf("%s\n", stmt);

	if (!mysql_query(&mysql, stmt))
	{
		//	printf("writen result %d to database\n", status);
		printf(" Updated Runs while Run_ID = %ld with status=%d!\n", runList[runIndex].runID,status);
	}
}

void executeRun(int runIndex)
{
	char executableFileName[FILE_PATH_LENGTH];
	char resultFileName[FILE_PATH_LENGTH];
	char inputFileName[FILE_PATH_LENGTH];
	char timeMemoryFileName[FILE_PATH_LENGTH];
	int timeLimitation, memoryLimitation;
	int i;
	enum exitStatus currentStatus;

	char javaPath[200];
	sprintf(javaPath, "%s/%d/%ld", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID
	       );



	sprintf(executableFileName, "%s/%d/%ld/%ld", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID,
			runList[runIndex].runID);

	/*    sprintf(resultFileName, "%s/%d/%ld%s%ld", userBasePath, 
	      runList[runIndex].userID, runList[runIndex].runID,
	      RESULT_FILE_NAME, runList[runIndex].runID);
	      */

	timeLimitation = timeMemory.timeLimit;
	memoryLimitation = timeMemory.memoryLimit;

	totalRunStatus.timeUsed = 0;
	totalRunStatus.memoryUsed = 0;
	totalRunStatus.runExitStatus = COMPILING;

	char input[150];//addByJialin 
	for (i = 0; i < numTestcase; i++)
	{

		sprintf(inputFileName, "%s/%d%s%d", problemBasePath, 
				runList[runIndex].problemID, PROBLEM_INPUT, i);
		//addByJialin:  get the data of inputfile
		sprintf(caseID, "%d:%d:%ld:%d",runList[runIndex].userID,runList[runIndex].problemID,runList[runIndex].runID,i );  
		int ip = open(inputFileName,O_RDONLY);
		if(Readn(ip,input,sizeof(input)))
		{
			//put the data to tbl_testCases' do with following	
			//	close(ip);
		}	 

		sprintf(resultFileName, "%s/%d/%ld%s%ld_%d", userBasePath, 
				runList[runIndex].userID, runList[runIndex].runID,
				RESULT_FILE_NAME, runList[runIndex].runID, i);

		sprintf(timeMemoryFileName, "%s/%d/%ld%s", userBasePath, 
				runList[runIndex].userID, runList[runIndex].runID,
				TIMEMEMORYRESULT);

		/*
		   sprintf(buffer, "%s < %s > %s\n", 
		   executableFileName, inputFileName, 
		   resultFileName);
		   */

		//system(buffer);
		/*
		 * using pipe to transfer run status
		 */
		pipe(piped);	

		//addjava
		if (4 == runList[runIndex].languageID)
		{
			sprintf(buffer, "%s %s %d %d %s %s %s %d\n", 
					JUDGER_NAME, javaPath,
					timeLimitation, memoryLimitation, timeMemoryFileName, 
					inputFileName, resultFileName,runList[runIndex].languageID);
		}
		else{


			sprintf(buffer, "%s %s %d %d %s %s %s %d %d\n", 
					JUDGER_NAME, executableFileName,
					timeLimitation, memoryLimitation, timeMemoryFileName, 
					inputFileName, resultFileName,runList[runIndex].languageID,piped[1]);
		}
		//end
		/*
		 * judger will write status:time consumption,memory consumption and run statu to from pipe[1]
		 */

		system(buffer);
		/*
		 * must close the writing port or the schedule will sleep
		 */
		close(piped[1]);
		pipeReadToStruct(piped[0],&currentRunStatus);
		printf("status read from  pipe %d %d %d\n", currentRunStatus.timeUsed,currentRunStatus.memoryUsed,currentRunStatus.runExitStatus);
		close(piped[0]);

		/*
		 * using pipe to get run staus instead of writing and reading file
		 *getRunStatus(runIndex, &currentRunStatus);
		 */


		totalRunStatus.timeUsed += currentRunStatus.timeUsed;
		totalRunStatus.memoryUsed += currentRunStatus.memoryUsed;
		totalRunStatus.runExitStatus = currentRunStatus.runExitStatus;

		/*
		 * addByJialin   insert the userID+problem+runID+userID+casei into tbl_testCases preparing for the following operation
		 * store every test case's resulte
		 *
		 sprintf(stmt, "INSERT INTO tbl_testCases SET id = '%s',userInput='%s',timeUsed=%lld,memoryUsed=%ld,status='%ld'",caseID,input,currentRunStatus.timeUsed/1000,currentRunStatus.memoryUsed,currentRunStatus.runExitStatus);
		 if (!mysql_query(&mysql, stmt))
		 {
		//printf("insert caseID done!\n");
		}
		*/
		/*
		 * if the app false to pass one of the test cases
		 * stop test and write the  test resutl as final result
		 */
		if (totalRunStatus.runExitStatus != EXIT_NORMALLY)
		{
			updateRunStatus(runIndex, totalRunStatus.runExitStatus, totalRunStatus);
			printf("test case%d break\n",i );
			currentStatus=totalRunStatus.runExitStatus;
			break;
		}
		else{
			currentStatus = diffResult(runIndex, i);
			printf("after diffJusge result is:%d\n",currentStatus );

			//test
			if (currentStatus != ACCEPTED)
			{
				updateRunStatus(runIndex, currentStatus, totalRunStatus);

				/*
				 * check whether runid exit,if exit skip inserting or not.
				 *
				 * insert to testCase wronge log
				 */
				sprintf(stmt,"select * from tbl_run_testcase where run_id=%ld",runList[runIndex].runID);
				MYSQL_RES * mysql_res;
				if (!mysql_query(&mysql, stmt))
				{
					mysql_res = mysql_store_result(&mysql);
					int caseRecordCount = mysql_num_rows(mysql_res);
					if (caseRecordCount)
					{
						printf("  %d recorde already exist which run_id is:%d\n",caseRecordCount,runList[runIndex].runID);

					}
					else{
						sprintf(stmt, "INSERT INTO tbl_run_testcase  SET run_id = %ld,testcase_id=%d",runList[runIndex].runID,testcaseList[i].testcaseID);
						if (!mysql_query(&mysql, stmt))
						{
							printf("wrong output:insert caseID done!\n");
						}
					}
					mysql_free_result(mysql_res);


				}
				/*
				   if ( !mysql_query(&mysql, stmt))
				   {
				   printf("wrong answer test case ID:%d\n",testcaseList[i].testcaseID );	
				   mysql_res= (&mysql);
				   if (mysql_res)
				   {
				   printf(" testing: recorde exist\n");
				   }

				   }

*/


				break;
			}
		}
	}

	printf("Judge Status for Run_ID %ld for Problem_ID %d is %ld\n", runList[runIndex].runID,
			runList[runIndex].problemID, (long) currentStatus);

	if (currentStatus == ACCEPTED)
	{
		printf("Run_ID %ld ACCEPTED\n", runList[runIndex].runID);
		totalRunStatus.timeUsed /= (i+1)*1000;
		totalRunStatus.timeUsed++;

		totalRunStatus.memoryUsed /= (i+1);
		updateRunStatus(runIndex, currentStatus, totalRunStatus);
	}
	else{
		printf("final result unnormal break at case %d\n", i); 
	}

	/*
	   printf("Time Used for RunID = %ld is %lld MS\n",
	   runList[runIndex].runID, totalRunStatus.timeUsed);
	   printf("Memory Used for RunID = %ld is %ld KB\n", 
	   runList[runIndex].runID, totalRunStatus.memoryUsed);
	   printf("Exit Status for RunID = %ld is %ld\n",
	   runList[runIndex].runID, totalRunStatus.runExitStatus);
	   */
}

int getRunStatus(int runIndex, struct stRunStatus * runStatus)
{
	FILE *rf;
	char timeMemoryFileName[FILE_PATH_LENGTH];

	sprintf(timeMemoryFileName, "%s/%d/%ld%s", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID,
			TIMEMEMORYRESULT);
	rf = fopen(timeMemoryFileName, "r");
	if (rf == NULL)
	{
		return -1;
	}

	if (fscanf(rf, "%lld%ld%ld", &runStatus->timeUsed,
				&runStatus->memoryUsed, &runStatus->runExitStatus) == EOF)
	{
		fclose(rf);
		return -1;
	}

	/*
	   printf("Time Used for RunID = %ld is %lld\n", runList[runIndex].runID, runStatus->timeUsed);
	   printf("Memory Used for RunID = %ld is %ld\n", runList[runIndex].runID, runStatus->memoryUsed);
	   printf("Exit Status for RunID = %ld is %ld\n", runList[runIndex].runID, runStatus->runExitStatus);
	   */
	fclose(rf);

	return 0;
}

void find_next_nonspace(int * c1, int * c2, FILE * stdf, FILE * usrf, enum exitStatus * currentExitStatus)
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
				*currentExitStatus = PRESENTATION_ERROR;
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

enum exitStatus diffResult(int runIndex, int testcaseIndex)
{
	char resultFileName[FILE_PATH_LENGTH];
	char stdOutputFileName[FILE_PATH_LENGTH];
	char diffResultFileName[FILE_PATH_LENGTH];
	FILE *stdf, *usrf;


	enum exitStatus currentExitStatus = ACCEPTED;

	sprintf(resultFileName, "%s/%d/%ld%s%ld_%d", userBasePath, 
			runList[runIndex].userID, runList[runIndex].runID,
			RESULT_FILE_NAME, runList[runIndex].runID, testcaseIndex);

	//addByJialin: get the output data of user and put to tbl_testCases identified by runID

	sprintf(caseID, "%d:%d:%ld:%d",runList[runIndex].userID,runList[runIndex].problemID,runList[runIndex].runID,testcaseIndex );  
	char op[100];
	int ow=open(resultFileName,O_RDONLY);
	memset(op,0,sizeof(op));
	if(Readn(ow,op,sizeof(op))){
		//addd	
		//	printf("path name of outputuser:%s\n", resultFileName);
		//	printf("outputByUser:%s\n", op);
		sprintf(stmt, "UPDATE tbl_testCases SET userOutput='%s' where id='%s'",op,caseID);
		if (!mysql_query(&mysql, stmt))
		{
			//	        printf("insert userOutput done!\n");

		}

	}
	//addEnd
	sprintf(stdOutputFileName, "%s/%d%s%d", problemBasePath, 
			runList[runIndex].problemID, PROBLEM_OUTPUT, testcaseIndex);
	/*
	   sprintf(diffResultFileName, "%s/%d/%ld%s%ld_%d", userBasePath, 
	   runList[runIndex].userID, runList[runIndex].runID,
	   DIFF_RESULT_FILE_NAME, runList[runIndex].runID, testcaseIndex);
	   */


	//    sprintf(buffer, "diff -wB %s %s > %s\n", stdOutputFileName, resultFileName, diffResultFileName);

	//    system(buffer);
	stdf = fopen(stdOutputFileName, "r");
	usrf = fopen(resultFileName, "r");

	if (stdf == NULL || usrf == NULL)
	{
		currentExitStatus = RUNTIME_ERROR;
	}
	else
	{
		for (; ;)
		{
			int c1 = fgetc(stdf);
			int c2 = fgetc(usrf);

			find_next_nonspace(&c1, &c2, stdf, usrf, &currentExitStatus);

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
								currentExitStatus = WRONG_ANSWER;
							}
						}
						if (c2 == EOF)
						{
							c1 = fgetc(stdf);
							if (!isspace(c1))
							{
								currentExitStatus = WRONG_ANSWER;
							}
						}
						break;
					}
					c1 = toupper(c1);
					c2 = toupper(c2);
					if (c1 != c2)
					{
						currentExitStatus = WRONG_ANSWER;
						goto end;
					}
					c1 = fgetc(stdf);
					c2 = fgetc(usrf);
					find_next_nonspace(&c1, &c2, stdf, usrf, &currentExitStatus);
					if (c1 == EOF && c2 == EOF)
					{
						goto end;
					}
					if (c1 == EOF || c2 == EOF)
					{
						currentExitStatus = WRONG_ANSWER;
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

	return currentExitStatus;
}
void daemon(void) {    
	pid_t  pid;

	/*
	 *      * Become a session leader to lose controlling TTY.
	 *           
	 */

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

int main ()
{
	FILE * dest_file;

	char tempBuffer[FILE_PATH_LENGTH];
	int num;
	//	daemon();
	if (initialize() == 1)
	{
		return 1;
	}

	printf("Connecting to mysql successed!\n");

	while (1)
	{
		num = fetchRunInfo(CONCURRENT_RUN_NUM);
		sleep(SLEEP_TIME - num);
	}

	finalize();
	return 0;
}
//return: done with positive num  or not faild
int Readn(int fd, void* buffer, size_t count) { 
	char* p = (char*)buffer;
	while (count > 0 ) {   
		ssize_t num = read(fd, p, count);
		if (num == -1) { 
			printf("Fail to read from file");
			return -1;
		}
		if (num == 0) {            // EOF
			break;
		}
		p += num;
		count -= num;
	}

	return p - (char*)buffer;
}
/*=================cString with struct======================*/


void pipeReadToStruct(int file,struct stRunStatus* result){
	FILE* pstream = fdopen(file,"r");
	//if (fscanf(pstream,"%s%d",getArr,&getInt) == EOF)
	if (fscanf(pstream,"%d%d%d",&result->timeUsed,&result->memoryUsed,&result->runExitStatus) == EOF)
	{

	}
	//cout<<"in read function statu:"<<(*result).statu<<" time:"<<result->timeConsumption<<" memory:"<<result->memoryConsumption<<endl;
	//printf("demo getInt:%d getArr:%s\n",getInt,getArr );
}

void pipeWriteFromStruct(int file,const struct stRunStatus result){
	char cstring[1024];
	sprintf(cstring,"%d %d %d",result.timeUsed,result.memoryUsed,result.runExitStatus);
	FILE* pWStream = fdopen(file,"w");
	fprintf(pWStream, "%s",cstring);

}
/*=================end cString with struct======================*/


