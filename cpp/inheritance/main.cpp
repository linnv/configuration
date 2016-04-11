#include <iostream>
#include "include.h"
using namespace std;
class Base
{private:
        int b_number;
	public:
        Base( ){}
        Base(int i) : b_number (i) { }
        int get_number( ) {return b_number;}
     virtual   void printf( ) {cout << b_number << endl;}        
};
class Derivedx:public Base{
	private:
		int x_num;
	public:
		Derivedx(int i,int j):Base(i),x_num(j){
			
		}
		void printf(){
		
		
			cout<<get_number()<<" "<<x_num<<endl;
		}

};
class  Derived: public Base
{
	private:
		int d_num;
	public:
		Derived(int i,int j):Base(i),d_num(j){}
		
		void printf(){
		cout<<endl<<get_number()<<" "<<d_num<<endl;	
		
		}

};

int main( )
{        Base a(2);
        Derived b(6, 8);
        Derivedx c(3, 4);
        cout << "a is ";
        a.printf( );                // print( ) in Base
        cout << "b is ";
        b.printf( );                // print( ) in Derived
 //       cout << "base part of b is "; 
//        b.Base::printf( );                // print( ) in Base
	Base* bp= &b;
	bp->printf();  //if base class not using virtual prefix for printf(), this will do the base'printf() or Derived'printf()
	/*
         cout << "c is ";
        c.printf( );                // print( ) in Derived
        cout << "base part of c is "; 
        c.Base::printf( );   return 0;
	*/
}
