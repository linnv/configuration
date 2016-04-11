#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include </usr/include/mysql/mysql.h>

#include <sys/types.h>
#include <sys/stat.h>
#include <dirent.h>

#define ROOT "."
#define TMP "/tmp"
#define HOST "localhost"
#define GUETOJ "test"
#define USER "guetoj"
#define PASSWORD "guetoj"
#define BUFFER_LENGTH 5000

struct stRunMeta
{
    unsigned long runID;
    int languageID;
    int userID;
    int problemID;
    char sourceCode[BUFFER_LENGTH];

} * runList;
int numRun;

char buffer[BUFFER_LENGTH];
char stmt[BUFFER_LENGTH];

MYSQL mysql;
MYSQL_RES *mysql_res;
MYSQL_ROW mysql_row;
MYSQL_STMT *mysql_stmt;
MYSQL_BIND bind[1];

unsigned long get_file_size(const char *path);

int connect_mysql();
int initialize();

int finalize_mysql();
int finalize();

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

    printf("Disconnecting from mysql successed!\n");
    return 0;
}

int finalize()
{
    finalize_mysql();
}

int main (int argc, char * argv[])
{
    FILE * dest_file;
    size_t filesize, buf_length;
    char *buffer;

    if (argc < 2)
    {
	    printf("usage:\n");
	    printf("%s infilename [outfilename]\n", argv[0]);
	    return 1;
    }

    if (initialize() == 1)
    {
        return 1;
    }
    
    printf("Connecting to mysql successed!\n");

    filesize = get_file_size(argv[1]);

    if (filesize != -1)
    {
	    buffer = (char *) malloc(sizeof(char)*(filesize+1));

	    memset(buffer, 0, filesize+1);
	    printf("File size :%d\n", filesize);

	    if (buffer != NULL)
	    {
		    dest_file = fopen(argv[1], "rb");
		    if (dest_file != NULL)
		    {
			    fread(buffer, 1, filesize, dest_file);

	    		    printf("Buffer length : %d\n", strlen(buffer));

			    printf("File Content \n%s\n", buffer);
			    mysql_stmt = mysql_stmt_init(&mysql);
			    if (mysql_stmt != NULL)
			    {
				    sprintf(stmt, "insert into tbl_run(user_id, source_code, compile_error, compile_error2) values(1000000, 'hello world', 'compile_error', ?)");
				    if (!mysql_stmt_prepare(mysql_stmt, stmt, strlen(stmt)))
				    {
					    printf("Total %d parameters for mysql_stmt\n", mysql_stmt_param_count(mysql_stmt));

					    memset(bind, 0, sizeof(bind));
					    bind[0].buffer_type = MYSQL_TYPE_BLOB;
					    bind[0].buffer = (char *) buffer;

					    //memset(buffer, 1, strlen(buffer));

					    bind[0].buffer_length = filesize+1;
					    bind[0].is_null = 0;
					    buf_length = filesize+1;
					    bind[0].length = &buf_length;

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
    }

    if (argc > 2)
    {
	    if (mysql_stmt != NULL)
	    {
		    mysql_stmt_close(mysql_stmt);

		    mysql_stmt = mysql_stmt_init(&mysql);
		    sprintf(stmt, "select compile_error2, compile_error from tbl_run");
		    if (!mysql_stmt_prepare(mysql_stmt, stmt, strlen(stmt)))
		    {
		    }
	    }
    }

    finalize();
    return 0;
}
