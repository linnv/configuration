#include<iostream>

using namespace std; 

int shu(int n)
{
	if (n == 1)
	{
		return 1;
	}
	else if ( n == 2||n==3)
	{
		return 1 + shu(1);
	} 
	else if (n==4||n==5)
	{
		return 1 + shu(1) + shu(2);
	} 
	else if (n == 6 || n == 7)
	{
		return 1 + shu(1) + shu(2) + shu(3);
	}
	else if (n == 8 || n == 9)
	{
		return 1 + shu(1) + shu(2) + shu(3) + shu(4);
	}
	else
	{
		return 0;
	}
}
int main()
{
	int n = 0;
	cin >> n;

	if (n / 10 == 0)
	{
		cout<<shu(n);
	}
	else
	{
		cout<<shu(n / 10);
	}

}