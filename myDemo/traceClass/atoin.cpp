/* Conversation Standards */
#define KB_SIZE (1 << 10)
#define MB_SIZE (1 << 20)
#define GB_SIZE (1 << 30)
#define ERR_NULL_PTR 1
/* Status Definations */
#define SIZE_ATOI_STATUS_INIT 0
#define SIZE_ATOI_STATUS_NUMBER 1
#define SIZE_ATOI_STATUS_K 2
#define SIZE_ATOI_STATUS_M 3
#define SIZE_ATOI_STATUS_G 4
#define SIZE_ATOI_STATUS_B 5
#define SIZE_ATOI_STATUS_BY 6
#define SIZE_ATOI_STATUS_BYT 7
#define SIZE_ATOI_STATUS_BYTE 8
#define SIZE_ATOI_STATUS_BYTES 9
#define SIZE_ATOI_STATUS_ERROR -1
int size_atoin(long *size_ptr, char *str, size_t len)
{    long result = 0;
    int status = SIZE_ATOI_STATUS_INIT;
	char *str_p, *str_endp;

    if (size_ptr == NULL) return -ERR_NULL_PTR;
    if (str == NULL) return -ERR_NULL_PTR;

    str_p = str;
	str_endp = str_p + len;
    *size_ptr = 0;
    while (str_p != str_endp)
    {        if (*str_p == ' ')
            {            /* skip space */
	            }
            else
            {            switch (status)
	                {                case SIZE_ATOI_STATUS_INIT:
			                    if (('0' <= *str_p) && (*str_p <= '9'))
			                    {                        result = result * 10 + ((int)(*str_p) - (int)('0'));
					                            status = SIZE_ATOI_STATUS_NUMBER;
					                        }
			                    else 
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_NUMBER:
			                    if (('0' <= *str_p) && (*str_p <= '9'))
			                    {                        result = result * 10 + ((int)(*str_p) - (int)('0'));
					                            /* Status doesn't required to change */
					                        }
			                    else if ((*str_p == 'k') || (*str_p == 'K'))
			                    {                        result *= KB_SIZE;
					                            status = SIZE_ATOI_STATUS_K;
					                        }
			                    else if ((*str_p == 'm') || (*str_p == 'M'))
			                    {                        result *= MB_SIZE;
					                            status = SIZE_ATOI_STATUS_M;
					                        }
			                    else if ((*str_p == 'g') || (*str_p == 'G'))
			                    {                        result *= GB_SIZE;
					                            status = SIZE_ATOI_STATUS_G;
					                        }
			                    else if ((*str_p == 't') || (*str_p == 'T'))
			                    {                        /* TB doesn't supported currently */
					                            status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    else 
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_K:
			                case SIZE_ATOI_STATUS_M:
			                case SIZE_ATOI_STATUS_G:
			                    if ((*str_p == 'b') || (*str_p == 'B'))
			                    {                        status = SIZE_ATOI_STATUS_B;
					                        }
			                    else
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_B:
			                    if ((*str_p == 'y') || (*str_p == 'Y'))
			                    {                        status = SIZE_ATOI_STATUS_BY;
					                        }
			                    else
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_BY:
			                    if ((*str_p == 't') || (*str_p == 'T'))
			                    {                        status = SIZE_ATOI_STATUS_BYT;
					                        }
			                    else
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_BYT:
			                    if ((*str_p == 'e') || (*str_p == 'E'))
			                    {                        status = SIZE_ATOI_STATUS_BYTE;
					                        }
			                    else
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_BYTE:
			                    if ((*str_p == 's') || (*str_p == 'S'))
			                    {                        status = SIZE_ATOI_STATUS_BYTES;
					                        }
			                    else
			                    {                        status = SIZE_ATOI_STATUS_ERROR;
					                        }
			                    break;
			                case SIZE_ATOI_STATUS_BYTES:
			                    status = SIZE_ATOI_STATUS_ERROR;
			                    break;
			                case SIZE_ATOI_STATUS_ERROR:
			                default:
			                    return -1;
			                    break;
			            }
	            }
            str_p++;
        }
    switch (status)
    {        case SIZE_ATOI_STATUS_NUMBER:
            case SIZE_ATOI_STATUS_K:
            case SIZE_ATOI_STATUS_M:
            case SIZE_ATOI_STATUS_G:
            case SIZE_ATOI_STATUS_B:
            case SIZE_ATOI_STATUS_BYTE:
            case SIZE_ATOI_STATUS_BYTES:
                *size_ptr = result;
                break;
            default:
                *size_ptr = 0;
                return -1;
                break;
        }
    return 0;
}
