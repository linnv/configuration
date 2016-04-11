#include<iostream>
using namespace std;
int fac(int n)
{
int h=0,i;
if(n==1)
return 1;
if(n%2!=0)
h=fac(n-1);
else
{
for(i=1;i<=n/2;i++)
h+=fac(i);
h++;}
return h;
}
int main()
{
int n,h,i;
cin>>n;
if(n==0)
cout<<0<<endl;
else
{
h=fac(n);
cout<<h<<endl;
}
system("pause");
return 0;
}