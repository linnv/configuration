
#include <iostream>
#include <vector>
using namespace std;
class Collection{
	private:
		long int _timeLimit;
		long int _memoryLimit;
		int  _fileSizeLimit;
		int _state;
		int _judgeState;
		long int _runId;
		int _languageId;
		long int _timeComsupted;
		long int _memoryComsupted;
		int _userId;
		int _problemId;

		string _compilerName;
		string _sCodeSuffix;
		string _executeSuffix;
		string _sCode;
		string _stdInput;
		string _stdOutput;
		string _userOutput;
		vector<string> _testCase;
	public:
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

		void setTimeComsupted(const int &id){
			_timeComsupted = id;
		}
		long int getTimeComsupted(){
			return _timeComsupted;
		}

		void setMemoryComsupted(const int &id){
			_memoryComsupted = id;
		}
		long int getMemoryComsupted(){
			return _memoryComsupted;
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


		void setProblemId(const int &id){
			_problemId = id;
		}
		int getProblemId(){
			return _problemId;
		}

		void setState(const int&s){
		_state = s;	
		}
		int getState(){
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

};
