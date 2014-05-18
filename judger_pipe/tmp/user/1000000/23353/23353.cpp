#include <iostream>
 
using namespace std;
int fun (int n);
int check (int a);
int m = 0;
int main()
{
    int n;
    cin >> n;
    if (n >= 100)
    {
        return 0;
    }
    else
    {
        if (n >= 10)
        {
            n = n/ 10;
            n = n / 2;
            cout  << fun (n)+1;
        }
        else
        {
            n = n / 2;
            cout << fun (n)+ 1;
        }
    }
    return 0;
}
int fun (int n)
{
    int sum = 0;
    if (n == 1)
    {
        return 1;
    }
    else if (n == 2) return 3;
    else
    {
       sum++;
       sum+=fun (n/2);
       sum += fun (n-1);
    }
    return sum;
}