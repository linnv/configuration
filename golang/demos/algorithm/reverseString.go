package demo

//reverseString implements ...
func reverseString(s string) string {
	byteS := []rune(s)
	sLen := len(byteS)
	for i := 0; i < sLen/2; i++ {
		byteS[i], byteS[sLen-i-1] = byteS[sLen-i-1], byteS[i]
	}
	return string(byteS)
}

//in c
// char *s="abcd";
// int size=strlen(s);
// if (size <=0)
// {
// 	return NULL;
// }
// char *ret=(char*)malloc(size+1);
// for (int i = 0; i < size; ++i)
// {
// 	ret[i]=s[size-i-1];
// }
// ret[size]='\0';
// // move[0]='\0'; // compiler handle \0 as EOF of a string in s
// // move[1]='i'; //if prefix has \o, sufix will not display in cout
// std::cout << size << std::endl;
// std::cout << s<< std::endl;
// std::cout <<ret<< std::endl;
