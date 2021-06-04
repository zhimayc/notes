# ES6新特性

### 变量的解构赋值

> ```javascript
> // 数组的解构
> const F4 = ['黑暗之女','九尾妖狐','皎月女神','曙光女神'];
> let [a,b,c,d] = F4;
> 
> // 对象的解构
> const hero = {
>     name : '安妮',
>     nickname : '黑暗之女',
>     age : 10,
>     skill : function(){
>         console.log('提伯斯');
>     }
> }
> 
> let {name,nickname,age,skill} = hero;
> ```
>
> 

### 模板字符串

> ```javascript
> // 内容中可以直接出现换行符
> let str = `<ul><li>暗影之拳</li>
> <li>影流之主</li><li>疾风剑豪</li><li>冰晶凤凰</li></ul>`;
> 
> 
> // 变量拼接
> let title = '英雄联盟英雄名：';
> let Annie = '黑暗之女';
> let Zed = '影流之主';
> let Ashe = '寒冰射手';
> let output = `${title} ${Annie} , ${Zed} , ${Ashe}`;
> ```
>
> 

### 对象的简化写法

> ```javascript
> let name = '拉克丝';
> let change = function(){
>     console.log('我们可以改变你！');
> }
> 
> const school = {
>     name,
>     change,
>     improve(){
>         console.log('我们可以提高你的技能！');
>     }
> }
> ```

### 箭头函数以及声明特点

> ```javascript
> // 箭头函数 this是静态的，this始终指向函数声明时所在作用域下的this的值
> let fn = (a,b) => {
>     return a + b;
> }
> 
> function getName(){
>     console.log(this.name);
> }
> 
> let getName2 = () => {
>     console.log(this.name);
> }
> 
> window.name = '九尾妖狐';
> const school = {
>     name : '阿狸',
> }
> 
> getName();
> getName2();
> 
> // call 方法调用
> getName.call(school);
> getName2.call(school);
> 
> 
> // 不能作为构造函数实例化对象
> let Person = (name,age) => {
>     this.name = name;
>     this.age = age;
> }
> 
> // 这里会报错，不能实例化对象
> //let me = new Person('tom',20);
> 
> 
> // 不能使用arguments变量，这里会报错
> //    let fn1 = () => {
> //      console.log(arguments);
> //    }
> 
> 
> 
> // 箭头函数的简写
> // (1) 省略小括号，当形参只有一个的时候
> let add = n => {
>     return n + n;
> }
> 
> // (2) 省略花括号，当代码体只有一条语句的时候，此时return必须省略，而且语句的执行结果就是函数的返回值
> let pow = n =>  n * n
> 
> // 箭头函数适合与this无关的回调，定时器，数组的方法回调
> // 箭头函数不适合与this有关的回调，事件回调，对象的方法
> ```

### 函数参数的默认值

> ```javascript
> // 形参初始值 具有默认值的参数，一般位置要靠后（潜规则）
> function add(a,b,c=10) {
>     return a + b + c;
> }
> 
> let result = add(1,2);
> 
> // 与解构赋值结合
> function connect({host="127.0.0.1",username,password,port}){
>     console.log(host);
>     console.log(username);
>     console.log(password);
>     console.log(port);
> }
> 
> connect({
>     username : 'root',
>     password : 'root',
>     port : 3306
> })
> ```

### rest参数

> ```javascript
> // ES6引入rest参数，用于获取函数的实参，用来代替arguments
> function date(...args) {
>     console.log(args);
> }
> 
> // 如果函数有多个参数，rest参数必须放到参数的最后
> function fn (a,b,...args){
>     console.log(a);
>     console.log(b);
>     console.log(args);
> }
> 
> date('Ahri','Annie','Zed');
> fn();
> ```

### 扩展运算符

> ```javascript
> const heroes = ['九尾妖狐','赏金猎人','刀锋意志'];
> 
> function pick(){
>     console.log(arguments);
> }
> 
> pick(...heroes);     // 等同于  pick('九尾妖狐','赏金猎人','刀锋意志');
> 
> // 数组的合并
> const hero1 = ['寒冰射手','众星之子','仙灵女巫'];
> const hero2 = ['唤潮鲛姬','琴瑟仙女','风暴之怒'];
> const combine = [...hero1,...hero2];
> 
> // 数组的克隆
> const arr = ['E','G','M'];
> 
> // 如果数组arr中有引用类型的话，则执行的是浅拷贝，否则执行的是深拷贝
> const arrCopy = [...arr];
> 
> // 将伪数组转为真正的数组
> const divs = document.querySelectorAll('div');
> const divArr = [...divs];
> ```

### Symbol

> ```javascript
> // ES6引入了一种新的原始数据类型Symbol，表示独一无二的值，是一种类似于字符串的数据类型
> // Symbol特点
> // 1) Symbol的值是唯一的，用来解决命名冲突的问题
> // 2) Symbol值不能与其他数据进行运算
> // 3) Symbol定义的对象属性不能使用for...in循环遍历，但是可以使用Reflect.ownKeys来获取对象的所有键名
> 
> 
> // 创建Symbol
> let s = Symbol();
> // s2和s3不相等
> let s2 = Symbol('伊泽瑞尔');
> let s3 = Symbol('伊泽瑞尔');
> console.log(s2 === s3);
> 
> // s4和s5相等
> let s4 = Symbol.for('伊泽瑞尔');
> let s5 = Symbol.for('伊泽瑞尔');
> console.log(s4 === s5);
> 
> // 不能和其他数据进行运算，这里会报错
> // let result = s + 100;
> 
> // 向对象中添加方法 up down(当向对象中添加方法，有可能造成命名冲突，利用Symbol的唯一性可以解决冲突的问题)
> let game = {
>     up : function(){
>         console.log('origin up method is called!');
>     },
>     down : function(){
>         console.log('orgin down method is called!');
>     }
> }
> 
> let methods = {
>     up : Symbol(),
>     down : Symbol()
> }
> 
> game[methods.up] = function(){
>     console.log("up method is called!");
> }
> 
> game[methods.down] = function(){
>     console.log("down method is called!");
> }
> 
> game.up();
> game[methods.up]();
> ```

### 迭代器

> ```javascript
> const heroes = ['疾风剑豪','封魔剑魂','刀锋之影','影流之主'];
> for (let v of heroes){
>     console.log(v);
> }
> 
> let iterator = heroes[Symbol.iterator]();
> // console.log(iterator.next());
> 
> while(true){
>     let res = iterator.next();
>     if(res.done){
>         break;
>     } 
>     console.log(res.value);
> }
> 
> const fruits = {
>     name : '水果',
>     lists : [
>         'apple',
>         'banana',
>         'orange',
>         'strawberry'
>     ],
>     [Symbol.iterator](){
>         let index = 0;
>         let _this = this;
>         return {
>             next : function() {
>                 if(index < _this.lists.length){
>                     const result = {value : _this.lists[index],done : false}
>                     index++;
>                     return result;
>                 }else{
>                     return {value : undefined,done : true}
>                 }
>             }
>         }
>     }
> }
> 
> for (let v of fruits){
>     console.log(v);
> }
> ```

### 生成器

> ```javascript
> function * gen(arg){
> 
>     console.log('111111',arg);
>     let one = yield 'part1';
>     console.log('222222',one);
>     yield 'part2';
> 
>     yield 'part3';
> 
> 
> }
> 
> let iterator = gen('AAA');
> iterator.next();
> 
> 
> function getUsers() {
> 
>     setTimeout( ()=>{
>         let data = 'user data';
>         iterator2.next(data);
>     },1000);
> 
> }
> 
> function getOrders() {
>     setTimeout( ()=> {
>         let data = 'order data';
>         iterator2.next(data);
>     },1000);
> }
> 
> function getGoods() {
>     setTimeout( ()=>{
>         let data = 'goods data';
>         iterator2.next(data);
>     },1000);
> }
> 
> function * gen2(){
> 
>     let users = yield getUsers();
>     console.log(users);
>     let orders = yield getOrders();
>     console.log(orders);
>     let gooods = yield getGoods();
>     console.log(gooods);
> }
> 
> let iterator2 = gen2();
> iterator2.next();
> ```
>
> 

### Promise

> ```javascript
> const p = new Promise(function(resolve,reject){
>     setTimeout(function(){
>         let data = 'user data';
>         // 调用resolve表示成功返回数据，reject表示获取数据失败
>         // resolve(data);
>         let err = 'get data error';
>         reject(err);
>     },1000);
> });
> 
> // 调用promise对象的then方法
> p.then(function(value){     // 调用resolve()方法，这个函数会被执行
>     console.log(value);
> },function(value){      // 调用reject()方法，这个函数会被执行
>     console.log(value);
> });
> 
> // promise封装ajax请求
> const pse = new Promise((resolve,reject) => {
>        const xhr = new XMLHttpRequest();
>        const url = 'https://api.apiopen.top/getJoke';
>        xhr.open("GET",url);
>        xhr.send();
>        xhr.onreadystatechange = function(){
>            if(xhr.readyState === 4){
>                if(xhr.status === 200){
>                    resolve(JSON.parse(xhr.responseText));
>                }else{
>                    reject(xhr.status);
>                }
>            } 
>        };
> });
> 
> pse.then(data => {
>        console.log(data);
> },err => {
>        console.error(err);
> });
> ```

### 集合

> ```javascript
> let s = new Set();
> let s2 = new Set(['铸星龙王','冰晶凤凰','兽灵行者','诡术妖姬','冰晶凤凰']);
> 
> // 元素个数
> console.log(s2.size);
> // 添加元素
> s2.add('皎月女神');
> // 删除元素
> s2.delete('兽灵行者');
> // 检测
> console.log(s2.has('冰晶凤凰'));
> // 清空
> s2.clear();
> // 遍历
> for(let v of s2){
>     console.log(v);
> }
> 
> // 数组去重
> let arr = [1,2,3,4,5,4,3,2,1];
> let result = [...new Set(arr)];
> console.log(result);
> 
> // 求交集
> let arr2 = [4,5,6,5,6];
> let intersectArr = [...new Set(arr)].filter(item => new Set(arr2).has(item));
> console.log(intersectArr);
> 
> // 求并集
> let union = [...arr,...arr2];
> let unionArr = [...new Set(union)];
> console.log(unionArr);
> 
> // 求差集
> let diff = [...new Set(arr)].filter(item => !(new Set(arr2).has(item)));
> console.log(diff);
> ```

### Map

> ```javascript
> // 声明map
> let m = new Map();
> 
> // 添加元素
> m.set('name','阿狸');
> m.set('change',function(){
>     console.log('I can change you!');
> });
> 
> let key = {
>     land : '艾欧尼亚'
> };
> 
> m.set(key,['班德尔城','诺克萨斯']);
> 
> // 元素个数
> console.log(m.size);
> 
> // 删除元素
> // m.delete('name');
> 
> // 获取元素
> console.log(m.get('change'));
> 
> // 清空
> // m.clear();
> 
> // 遍历
> for(let v of m) {
>     console.log(v);
> }
> 
> console.log(m,typeof m);
> ```

### 类（class）

> ```javascript
> // es6 写法
> class Phone {
>     static name = '手机';   // 静态成员
> 
>     // 构造方法，名字不能修改
>     constructor(brand,price){   // 构造方法
>         this.brand = brand;
>         this.price = price;
>     }
> 
>     test(){ // 成员方法
>         console.log('test method is called!');
>     }
> }
> 
> let onePlus = new Phone("1+",1999);
> onePlus.test();
> console.log(Phone.name);
> 
> 
> // es5 写法
> function MobilePhone(brand,price){  // 构造方法
> 
>     this.brand = brand;
>     this.price = price;    
> }
> 
> MobilePhone.testname = '移动电话';  // 静态成员
> MobilePhone.prototype.test = function(){
>     console.log('mobilePhone test method is callled!');
> }
> 
> MobilePhone.prototype.nickname = '移动电话别名';    // 实例对象成员
> 
> var mobile = new MobilePhone('iphone',5999);
> mobile.test();
> console.log(mobile.brand);
> console.log(MobilePhone.testname);
> console.log(mobile.nickname);
> ```

### 类的继承

> ```javascript
> // es5 写法
> 
> function Phone(brand,price) {
>     this.brand = brand;
>     this.price = price;
> }
> 
> Phone.prototype.call = function(){
>     console.log('我可以打电话！');
> }
> 
> function SmartPhone(brand,price,color,size) {
>     Phone.call(this,brand,price);
>     this.color = color;
>     this.size = size;
> }
> 
> // 设置子级构造函数的原型
> SmartPhone.prototype = new Phone;
> SmartPhone.prototype.constructor = SmartPhone;
> 
> // 声明子类的方法
> SmartPhone.prototype.photo = function(){
>     console.log('我可以拍照！');
> }
> 
> SmartPhone.prototype.play = function(){
>     console.log('我可以玩游戏！');
> }
> 
> var xiaomi = new SmartPhone('小米',3999,'黑色','5.5 inch');
> 
> // es6 写法
> class Phone {
>     // 构造方法
>     constructor(brand,price){
>         this.brand = brand;
>         this.price = price;
>     }
> 
>     // 父类成员属性
>     call() {
>         console.log('我可以打电话');
>     }
> }
> 
> class SmartPhone extends Phone {
> 
>     // 构造方法
>     constructor(brand,price,color,size){
>         super(brand,price);
>         this.color = color;
>         this.size = size;
>     }
> 
>     photo(){
>         console.log('拍照！');
>     }
> 
>     play(){
>         console.log('玩游戏！');
>     }
> }
> 
> const xiaomi = new SmartPhone('小米',3999,'灰色','5.5inch');
> ```

### get和set方法

> ```javascript
> class Phone {
> 
>     get price(){
>         console.log('pirce属性被读取了');
>         return 1999;
>     }
> 
>     set price(newVal){
>         console.log('价格属性被修改了！');
> 
>     }
> }
> 
> let s = new Phone();
> console.log(s.price);
> s.price = 5999;
> ```
>
> 

### 数值扩展

> ```javascript
> function equal(a,b) {
>     if (Math.abs(a - b) < Number.EPSILON) { 
>         return true;
>     }else {
>         return false;
>     }
> }
> 
> console.log( 0.1 + 0.2 === 0.3);
> console.log(equal(0.1 + 0.2,0.3));
> 
> // 二进制和八进制
> let b = 0b1010;
> let o = 0o17;
> console.log(o);
> 
> // Math.sign() 判断一个数为正数，负数，还是零
> console.log(Math.sign(100));
> console.log(Math.sign(-100));
> console.log(Math.sign(0));
> 
> // 将数字的小数部分抹掉
> console.log(Math.trunc(1.1912834));
> ```

### 对象方法扩展

> ```javascript
> // Object.is 判断两个值是否完全相等
> console.log(Object.is(120,120));    // true
> console.log(Object.is(NaN,NaN));    // true
> console.log(NaN === NaN);   // false
> 
> // Object.assign 对象的合并
> const config1 = {
>     host : 'localhost',
>     port : 3306,
>     name : 'root',
>     pass : 'root',
>     test : 'test'
> };
> 
> const config2 = {
>     host : 'http://example.com',
>     port : 33060,
>     name : 'charles',
>     pass : '12345',
>     test2 : 'test2'
> };
> 
> Object.assign(config1,config2);
> console.log(config1);
> 
> // Object.setPrototypeOf 设置原型对象
> const  hero = {
>     name : '影流之主'
> };
> 
> const skills = {
>     skill : ['瞬狱影杀阵','奥术跃迁','自然之力']
> };
> 
> Object.setPrototypeOf(hero,skills);
> console.log(hero);
> ```
>
> 

### ES6模块的使用

* 模块导出的使用

  > ```javascript
  > // 用法一
  > <script type="module">
  >     import * as m from "./1.js";
  > 	console.log(m.hero);
  > </script>
  > 
  > // 1.js
  > export let hero = '寒冰射手';
  > export function skill(){
  >     console.log('魔法水晶箭');
  > }
  > 
  > // 用法二
  > let hero = '寒冰射手';
  > function skill(){
  >     console.log('魔法水晶箭');
  > }
  > export {hero,skill}
  > 
  > // 用法三
  > <script type="module">
  >   import * as m from "./1.js";
  >   m.default.skill();
  > </script> 
  > 
  > export default {
  >     hero : '寒冰射手',
  >     skill : function(){
  >         console.log('魔法水晶箭');
  >     }
  > }
  > ```
  >
  > 

* 模块引入的使用

  > ```javascript
  > // 用法一
  > export let hero = '寒冰射手';
  > export let skill = function(){
  >     console.log('魔法水晶箭');
  > }
  > //==========================================
  > import { hero,skill } from './1.js';
  > skill();
  > 
  > // 用法二
  > let hero = '寒冰射手';
  > let skill = function(){
  >     console.log('魔法水晶箭');
  > }
  > export { hero,skill }
  > //===========================================
  > import { hero as hero_1,skill as skill_1} from './1.js';
  > skill_1();
  > 
  > // 用法三
  > export default {
  >     hero : '寒冰射手',
  >     skill : function(){
  >         console.log('魔法水晶箭');
  >     }
  > }
  > //===========================================
  > import { default as m } from './1.js';
  > m.skill();
  > 
  > // 简便形式，针对默认暴露
  > import m from './1.js';
  > m.skill();
  > ```
  >
  > 

### async和await

> ```javascript
> // async和await结合，可以让异步代码实现类似同步的效果
> // 创建promise对象
> const p = new Promise((resolve,reject) => {
>     setTimeout(function(){
>         resolve("成功的值111111！");
>         // reject('失败了');
>     },3000);
> 
> });
> 
> const p2 = new Promise((resolve,reject) => {
>     setTimeout(function(){
>         resolve("成功的值2222222！");
>         // reject('失败了');
>     },1000);
> 
> });
> 
> // await要放在async函数中
> async function main(){
>     try {
>         let result = await p;
>         let result2 = await p2;
>         console.log(result);
>         console.log(result2);
>     }catch(e){
>         console.log(e);
>     }
> }
> 
> // 调用函数
> main();
> ```
>
> 

