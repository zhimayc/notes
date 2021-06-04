@Resource、@Autowired、@Qualifier

> @Resource：有两个重要属性name和type
>
> （1）如果同时指定了name和type，则从Spring上下文中找到唯一匹配的bean进行装配，找不到则会抛出异常。
>
> （2）如果指定了name，则从上下文中查找名称（id）匹配的bean进行装配，找不到则抛出异常。
>
> （3）如果指定了type，则从上下文中找到类型匹配的唯一bean进行装配，找不到或者找到多个都会抛出异常。
>
> （4）如果既没有指定name，也没有指定type，则自动按照byName方式进行装配；如果没有找到匹配，则会回退为一个原始类型进行匹配，如果匹配则自动装配



> @Autowired：按照类型装配注入，如果同类型的bean有多个则会报错，需要配合下面这个注解装配



> @Qualifier：按照名称装配注入，需要配合Autowired使用

