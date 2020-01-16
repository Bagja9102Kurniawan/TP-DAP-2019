/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini adalah program pengecekan tahun kabisat yaitu tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100
	akan mengoutput 0 jika salah dan 1 jika benar
*/
#include <iostream>
using namespace std;
int i;
int main (){
	cout << "Selamat Datang di Program Tahun Kabisat\n" << "Tahun         : ";
	cin >> i;
	cout << "Kabisat       : " << (i%400 == 0 || i%4 == 0 && i%100 != 0);
	return 0;
}
