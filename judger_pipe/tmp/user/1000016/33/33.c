#include<stdio.h>
void main()
{
int i,n,m,x[50000],j,a;
a=0;
scanf("%d",&n);
scanf("%d",&m);
for(i=0;i<10;i++)
scanf("%d",&x[i]);
for(i=0;i<n;i++)
for(j=0;j<n;j++)
{
  if(x[j]==x[i])
   continue;
   if((x[i]+x[j])==m)
            a=a+1;
  

}

printf("%d",a);
}