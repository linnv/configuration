#include<stdio.h>
main()
{
	int a[20];
	int b,i=0,m=0,n,z=0;
	while(1)
	{
		while(b=getchar()!='\n')
		{
			if(b>='0'&&b<='9')
			{
				a[i]=b-'0';
				i++;
			}
			else
			{
				if(i!=0) z++;
				for(n=0;n<i;n++)
				{
					m=m+a[n]*10^i;
					i--;
				}
				
				break;
			}
		}
		if(b=getchar()=='\n') break;
	}
	printf("%d %d\n",z,m);
}
