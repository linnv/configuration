#include "include.h"
using namespace std;
class SQL{
	private:
		bool isConnected;
		string host;
		string user;
		string passwd;
		int numColunm;
		int numRow;
	sql::Driver *driver;
 	//sql::Connection *con;
  	//sql::Statement *stmt;
  	sql::ResultSet *res;
 	sql::PreparedStatement *pstmt;
	
	public:
		
 	sql::Connection *con;

  	sql::ResultSet* getResultSet(){
		return res;
	}
	int getRowsCount(){
	
		return res->rowsCount();
	}
	int getColunmsCount(){
		return numColunm;	
	}
		void setHost(const string & h ){
			host =h;	
		}
		void setUser(const string & u){
			user = u;	
		}
		void setPasswd(const string& p){
			passwd = p;	
		}
		void connectSQL(){
			
  			/* Create a connection */
			driver = get_driver_instance();
			con = driver->connect(host,user,passwd);
			isConnected = true;
			
		}
		void useDatabase(const string& d){
			con->setSchema(d);	
		}
		void querySQL(const string & sql){
			if (isConnected)
			{
				/* code */
			try{
			/* Select in ascending order */
			  pstmt = con->prepareStatement(sql);
    			res = pstmt->executeQuery();
	      /* Fetch in reverse = descending order! */
      		//	res->afterLast();
			sql::ResultSetMetaData *res_meta = res->getMetaData();
			numColunm = res_meta->getColumnCount();
			/*
        		while (res->next()){
				for (int i = 1; i <= numColunm; ++i)
				{
					cout<<res->getString(i)<<" ";
				}
				cout<<endl;
	    		//cout << "\t" << res->getInt("id");
	    		//cout << "\t... MySQL counts: " << res->getInt("id") << endl;
			}
			*/
	 		 } 
			catch (sql::SQLException &e) {
			  cout << "# ERR: SQLException in " << __FILE__;
		    	cout << "(" << __FUNCTION__ << ") on line " << __LINE__ << endl;
		    	cout << "# ERR: " << e.what();
		    	cout << " (MySQL error code: " << e.getErrorCode();
		    	cout << ", SQLState: " << e.getSQLState() << " )" << endl;
		  					}
			}
			else
			{
				cout<<"connect to mysql firstly"<<endl;
			}
			
		}
	
		void updateSQL(const string & sql){
			if (isConnected)
			{
			  pstmt = con->prepareStatement(sql);
			  //pstmt = con->prepareStatement("INSERT INTO test(id) VALUES (?)");
    			  pstmt->executeUpdate();
		//	delete pstmt;	  
			}
			else{
				cout<<"connect to mysql firstly"<<endl;
			}
	}
		void closeSQL(){
		try{
			delete res;
			res = NULL;
	      		delete pstmt;
			pstmt = NULL;
		  	delete con;
			con = NULL;
			 } 
			catch (sql::SQLException &e) {
			  cout << "# ERR: SQLException in " << __FILE__;
		    	cout << "(" << __FUNCTION__ << ") on line " << __LINE__ << endl;
		    	cout << "# ERR: " << e.what();
		    	cout << " (MySQL error code: " << e.getErrorCode();
		    	cout << ", SQLState: " << e.getSQLState() << " )" << endl;
		  					}
		}
	
};
/*
int mainn(void)
{
	SQL* sqlDemo = new SQL();
	sqlDemo->setHost("tcp://127.0.0.1:3306");
	sqlDemo->setUser("root");
	sqlDemo->setPasswd("a");
	sqlDemo->connectSQL();
	sqlDemo->useDatabase("test");



//	sqlDemo->updateSQL("insert into test set id = 88888");
	sqlDemo->querySQL("SELECT * FROM id");
	sqlDemo->closeSQL();
	

	cout << endl;
return EXIT_SUCCESS;
}
*/
