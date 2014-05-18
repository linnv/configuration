#include<iostream>

using namespace std;

int fun(int n)
{
	int result = 0;
	if (n == 0 || n == 1)
	{
		result = 0;
	}
	else
	{
		for (int i = n / 2; i > 0; i--)
		{
			result = fun(i) + result;
		}
	}
	return result + 1;
}

int main()
{
	int N = 0;
	int result = 0;
	cin >> N;

	if (N / 10 == 0)
	{
		result = fun(N);
	}
	else
	{
		result = fun(N / 10);
	}

	cout << result;

	return 0;
}