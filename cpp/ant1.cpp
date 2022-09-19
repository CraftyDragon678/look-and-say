#include <iostream>
#include <string>
using namespace std;

int main(void) {
  string s = "1";
  int max;

  cin >> max;

  for (int i = 0; i < max; i++) {
    char ch = s[0];
    string next = "";
    int count = 1;
    for (int j = 1; j < s.size(); j++) {
      if (ch == s[j]) {
        count ++;
      } else {
        next += ch;
        next += to_string(count);
        ch = s[j];
        count = 1;
      }
    }
    next += ch;
    next += to_string(count);
    s = next;
  }
  cout << s << endl;
}