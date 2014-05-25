
#include <iostream>
#include <vector>
using namespace std;
class Collection{
	private:
		long int _timeLimit;
		long int _memoryLimit;
		int  _fileSizeLimit;
		int _state;
		int _judgeState; //state after diff judge
		long int _runId;
		int _languageId;
		long int _timeConsumpted;
		long int _memoryConsumpted;
		int _userId;
		int _problemId;
		int _testCaseId;

		string _compilerName;
		string _compilerOption;
		string _compileError;
		string _sCodeSuffix;
		string _executeSuffix;
		string _sCode;
		string _stdInput;
		string _stdOutput;
		string _userOutput;
		vector<string> _testCase;
	public:
		Collection(){
			_judgeState = 100000;
			_timeConsumpted = 0;
			_memoryConsumpted=0;
			_compileError="";
		}
		void setTimeLimit(const int &id){
			_timeLimit= id;
		}
		long int getTimeLimit(){
			return _timeLimit;
		}

		void setMemoryLimit(const int &id){
			_memoryLimit= id;
		}
		long int getMemoryLimit(){
			return _memoryLimit;
		}

		void setFileSizeLimit(const int &id){
			_fileSizeLimit= id;
		}
		int getFileSizeLimit(){
			return _fileSizeLimit;
		}


		void setJudgeState(const int &id){
			_judgeState= id;
		}
		int getJudgeState(){
			return _judgeState;
		}

		void setRunId(const int &id){
			_runId= id;
		}
		long int getRunId(){
			return _runId;
		}
		void setLanguageId(const int &id){
			_languageId= id;
		}
		int getLanguageId(){
			return _languageId;
		}

		void setTimeConsumption(const int &id){

			_timeConsumpted = (id>_timeConsumpted?id:_timeConsumpted);
		}
		long int getTimeConsumption(){
			return _timeConsumpted;
		}

		void setMemoryConsumption(const int &id){
			_memoryConsumpted = (id>_memoryConsumpted?id:_memoryConsumpted);
		}
		long int getMemoryConsumption(){
			return _memoryConsumpted;
		}

		void setUserId(const int &id){
			_userId = id;
		}
		int getUserId(){
			return _userId;
		}

		void setExecuteSuffix(const string &id){
			_executeSuffix = id;
		}
		string getExecuteSuffix(){
			return _executeSuffix;
		}

		void setSCodeSuffix(const string &id){
			_sCodeSuffix = id;
		}
		string getSCodeSuffix(){
			return _sCodeSuffix;
		}


		void setCompilerName(const string &id){
			_compilerName = id;
		}
		string getCompilerName(){
			return _compilerName;
		}

		void setCompilerOption(const string &id){
			_compilerOption = id;
		}
		string getCompilerOption(){
			return _compilerOption;
		}

		void setCompilerError(const string &id){
			_compileError = id;
		}
		string getCompilerError(){
			return _compileError;
		}




		void setProblemId(const int &id){
			_problemId = id;
		}
		int getProblemId(){
			return _problemId;
		}

		void setLastState(const int&s){
		_state = s;	
		}
		int getLastState(){
		return _state;	
		}
		
		void setSourceCode(string & str){
		 _sCode = str;	
		}
		string getSourceCode(){
			return _sCode;	
		}
		
		void setSTDIput(const string &str){
		_stdInput = str;	
		}
		string getSTDIput(){
			return _stdInput;	
		}

		void setSTDOutput(string &str){
		_stdOutput = str;	
		}
		string getSTDOutput(){
			return _stdOutput;	
		}

		void setUserOutput(string &str){
		_userOutput= str;	
		}
		string getUserOutput(){
			return _userOutput;	
		}

		void setTestcaseID(const int &id){
			_testCaseId = id;	
		}
		int getTestcaseID(){
			return _testCaseId;	
		}

};
