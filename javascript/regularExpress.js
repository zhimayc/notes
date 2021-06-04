/*

var pattern = /g..gle/;  //  .表示匹配除换行符以外的任意字符
var str = 'gbcgle';
alert(pattern.test(str));


var pattern = /go*gle/;  //  o*,表示0个，1个，或者多个o
var str = 'gooogle';
alert(pattern.test(str));


var pattern = /go+gle/;  //  o+,表示匹配至少一个o
var str = 'google';
alert(pattern.test(str));


var pattern = /go?gle/;  //  o?,表示1个，或者0个
var str = 'ggle';
alert(pattern.test(str));


var pattern = /g.?gle/;  //  .?,表示1个，或者0个的任意字符
var str = 'gagle';
alert(pattern.test(str));


var pattern = /go{2,4}gle/;  //  o{2,4},表示匹配o 2到4次，包含2和4
var str = 'goooogle';
alert(pattern.test(str));


var pattern = /go{3}gle/;  //  o{3},表示只能限定为3个o
var str = 'gooogle';
alert(pattern.test(str));


var pattern = /go{3,}gle/;  //  o{3,},表示为3或者3个以上的o
var str = 'goooooooogle';
alert(pattern.test(str));


var pattern = /[a-z]oogle/;  //[a-z]表示26个小写字母都可以匹配  
var str = 'coogle';
alert(pattern.test(str));


var pattern = /[A-Z]oogle/;  //[A-Z]表示26个大写字母都可以匹配  
var str = 'Boogle';
alert(pattern.test(str));


var pattern = /[0-9]oogle/;  //[0-9]表示0-9中任意一个数字都可以匹配  
var str = '5oogle';
alert(pattern.test(str));


var pattern = /[0-9]*oogle/;  //[0-9]*表示0次，1次或者多次  
var str = '55oogle';
alert(pattern.test(str));



var pattern = /[a-zA-Z0-9]oogle/;  //[a-zA-Z0-9]表示匹配大小写的a-zA-Z0-9 
var str = 'xoogle';
alert(pattern.test(str));


var pattern = /[^0-9]oogle/;  //[^0-9]表示匹配非0-9的任意字符
var str = 'aoogle';
alert(pattern.test(str));


var pattern = /[a-z]+/;  //[^0-9]表示匹配1个或者多个a-z之间的任意字符
var str = 'gga';
alert(pattern.test(str));


var pattern = /^[0-9]oogle/;  // /^[0-9]/ 表示0-9之间的任意字符在行首匹配
var str = '6oogle';
alert(pattern.test(str));


var pattern = /\woogle/;  // \w表示匹配字母，数字及_
var str = '_oogle';
alert(pattern.test(str));


var pattern = /\Woogle/;  // \W表示匹配非字母，数字及_
var str = '!oogle';
alert(pattern.test(str));

var pattern = /\doogle/;  // \d表示匹配数字
var str = '1oogle';
alert(pattern.test(str));


var pattern = /\Doogle/;  // \D表示匹配非数字
var str = 'aoogle';
alert(pattern.test(str));


var pattern = /google$/;  // /google$/表示行尾匹配google
var str = 'afsafagoogle';
alert(pattern.test(str));


var pattern = /goo\sgle/;  // \s表示匹配空白字符、空格、制表符和换行符
var str = 'goo gle';
alert(pattern.test(str));



var pattern = /google\b/;  // \b表示到达边界
var str = 'google';
alert(pattern.test(str));



var pattern = /google|baidu|bing/;  // |表示或选择模式
var str = 'this is bing';  //匹配概念不是相等,包含关系
alert(pattern.test(str));



var pattern = /(google){2,5}$/;  // ()表示分组的概念 括号里面的字符可以看做一个字符
var str = 'googlegooglegoogle';  
alert(pattern.test(str));


var pattern = /8(.*)8/;  // RegExp.$1表示获取模式中第一个分组对应的匹配字符串
var str = 'this is 8sfsfgg8';  //此字符串必须要先运行RegExp.$1才能获取到
pattern.test(str);
alert(RegExp.$1);


var pattern = /8(.*)8/;  
var str = 'this is 8yahoo8';  
document.write(str.replace(pattern,'<strong>$1</strong>')); //$1表示分组获取到的字符串匹配的内容



var pattern = /(.*)\s(.*)/;  
var str = 'google yahoo';  
alert(str.replace(pattern,'$2 $1'));  //位置交换



var pattern = /[a-z]+/;  //这里使用了贪婪模式
var str = 'abcdefg';  
alert(str.replace(pattern,'1'));  //所有的字符串变成了1



var pattern = /[a-z]+?/;  //这里使用了惰性模式
var str = 'abcdefg';  
alert(str.replace(pattern,'1'));  //只有第一个字符变成了1，后面没有匹配


var pattern = /[a-z]+?/g;  //开启全局，并且使用惰性模式
var str = 'abcdefg';  
alert(str.replace(pattern,'1'));  //每一个字符都被替换成了1


var pattern = /8(.*)8/;  //使用了贪婪
var str = '8google8 8google8 8google8';  //匹配到的是google8 8google8 8google
document.write(str.replace(pattern,'<strong>$1</strong>')); 


var pattern = /8(.*?)8/g;  //开启全局，使用惰性
var str = '8google8 8google8 8google8';  
document.write(str.replace(pattern,'<strong>$1</strong>')); 


var pattern = /8([^8]*)8/g;  //另一种惰性，屏蔽了8的匹配，也就是两边的包含字符
var str = '8google8 8google8 8google8';  
document.write(str.replace(pattern,'<strong>$1</strong>')); 


var pattern = /^[a-z]+\s[0-9]{4}$/;  
var str = 'google 2012aaaaa';  
alert(pattern.exec(str));   //返回一个包含字符串的数组


var pattern = /^([a-z]+)\s([0-9]{4})$/;    //使用了分组
var str = 'google 2012';  
alert(pattern.exec(str));   //返回google 2012，google，2012


var pattern = /(\d+)([a-z])/;    //捕获性分组，所有的分组都返回
var str = '123abc';  
alert(pattern.exec(str));   //返回123a,123,a


var pattern = /(\d+)(?:[a-z])/;    //非捕获性分组，只需要在不需要捕获返回的分组加上?:
var str = '123abc';  
alert(pattern.exec(str));   //返回123a,123


var pattern = /(a?(b?(c?)))/;    //嵌套分组，从外往内获取
var str = 'abc';  
alert(pattern.exec(str));   
//第一步：a[0],整个匹配到的字符串abc
//第二步：a[1],匹配第一个分组(a?(b?(c?))),abc
//第三步：a[2],匹配第二个分组(b?(c?)),bc
//第四步：a[3],匹配第三个分组(c?),c


var pattern = /goo(?=gle)/;  //goo后面必须是gle才能返回goo
var str = 'google';
alert(pattern.exec(str));  //返回的是goo，而不是google，这是前瞻性捕获


var pattern = /\[/;  //用\来转义正则里的特殊字符，才能匹配
var str = '[';
alert(pattern.exec(str));  


var pattern = /^\d+/gm;  //限定首匹配，并且开启换行模式
var str = '1.google\n2.yahoo\n3.bing';
alert(str.replace(pattern,'#'));  



*/

var pattern = /^\d+/gm;  //限定首匹配，并且开启换行模式
var str = '1.google\n2.yahoo\n3.bing';
alert(str.replace(pattern,'#'));  