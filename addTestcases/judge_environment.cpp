#include <iostream>
using namespace std;
class JudgeEnvironment{
	
	private:
		int _timeLimit;
		int _memoryLimit;
		int  _fileSizeLimit;
		int _problemId;
		int _languageId;
	public:
		void setTimeLimit(const int &id){
			_timeLimit= id;
		}
		int getTimeLimit(){
			return _timeLimit;
		}

		void setMemoryLimit(const int &id){
			_memoryLimit= id;
		}
		int getMemoryLimit(){
			return _memoryLimit;
		}

		void setFileSizeLimit(const int &id){
			_fileSizeLimit= id;
		}
		int getFileSizeLimit(){
			return _fileSizeLimit;
		}
		void setLanguageId(const int &id){
			_languageId =id;
		}
		int getLanguageId(){
			return _languageId;
		}
		void setProblemId(const int &id){
			_problemId= id;
		}
		int getProblemId(){
			return _problemId;
		}

};

