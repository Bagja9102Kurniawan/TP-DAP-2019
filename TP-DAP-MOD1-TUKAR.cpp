/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini bertujuan untuk memindahkan input pertama yang awalnya di depan kalimat berubah menjadi di akhir kalimat
	sedangkan input ke 2 berubah menjadi di depan kalimat 
	dan input yang ke 3 menjadi di tengah kalimat
*/
#include <iostream>
using namespace std;

string a,b,c,temp;
int main (){
	cout << "Masukkan input string: ";
	cin >> a;
	cout << "Masukkan input string: ";
	cin >> b;
	cout << "Masukkan input string: ";
	cin >> c;
	cout << "output awal\n" << a << " " << b << " " << c;
	temp = a;
	a = b;
	b = c;
	c = temp;
	cout << "\noutput akhir\n" << a << " " << b << " " << c;
	return 0;
}
