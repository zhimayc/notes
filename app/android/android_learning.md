安卓开发学习笔记
========================================================


布局
--------------------------------------------------------
1. LinearLayout

```xml
    <?xml version="1.0" encoding="utf-8"?>
    <LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
        android:orientation="vertical"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:padding="10dp"
        android:background="#F5F5F5">

        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:id="@+id/et_number"
            />
        <EditText
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:layout_marginTop="10dp"
            android:layout_marginBottom="10dp"
            android:hint="请输入电话号码！"
            />
        <Button
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:id="@+id/bt_callnumber"
            android:text="拨打此号码"
            />
    </LinearLayout>

```

2. RelativeLayout

```xml
    <?xml version="1.0" encoding="utf-8"?>
    <RelativeLayout xmlns:android="http://schemas.android.com/apk/res/android"
        android:layout_width="match_parent" android:layout_height="match_parent">


        <TextView
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:id="@+id/tv_inputnum"
            />
        <EditText
            android:id="@+id/et_number"
            android:layout_below="@id/tv_inputnum"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:layout_marginTop="10dp"
            android:layout_marginBottom="10dp"
            android:hint="请输入电话号码！"
            />
        <Button
            android:layout_below="@id/et_number"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:id="@+id/bt_callnumber"
            android:text="拨打此号码"
            />

    </RelativeLayout>

```

3. FrameLayout

```xml
    <?xml version="1.0" encoding="utf-8"?>
    <FrameLayout xmlns:android="http://schemas.android.com/apk/res/android"
        android:layout_width="match_parent" android:layout_height="match_parent">

        <LinearLayout
            android:layout_width="match_parent"
            android:layout_height="match_parent"
            android:background="#ffff00"
            >
        </LinearLayout>

        <LinearLayout
            android:layout_marginTop="20dp"
            android:layout_width="match_parent"
            android:layout_height="match_parent"
            android:background="#ff00ff"
            >
        </LinearLayout>

        <ImageView
            android:layout_gravity="center"
            android:layout_width="wrap_content"
            android:layout_height="wrap_content"
            android:src="@drawable/ic_launcher_background"
            />
    </FrameLayout>

```

4. TableLayout

```xml
    <?xml version="1.0" encoding="utf-8"?>
    <TableLayout xmlns:android="http://schemas.android.com/apk/res/android"
        android:layout_width="match_parent" android:layout_height="match_parent">

        <TableRow>
            <TextView
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:text="第一行第一列"
            />
            <TextView
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:text="第一行第二列"
                />
        </TableRow>
        <TableRow>
            <TextView
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:text="第二行第一列"
                />
            <TextView
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                android:text="第二行第二列"
                />
        </TableRow>
    </TableLayout>

```

5. AbsoluteLayout（已弃用）



动态权限申请
-------------------------------------------------------

1. AndroidManifest文件修改

```xml
    <?xml version="1.0" encoding="utf-8"?>
    <manifest xmlns:android="http://schemas.android.com/apk/res/android"
        package="com.example.myapp">

        <application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="@string/app_name"
            android:roundIcon="@mipmap/ic_launcher_round"
            android:supportsRtl="true"
            android:theme="@style/AppTheme">
            <activity android:name=".MainActivity">
                <intent-filter>
                    <action android:name="android.intent.action.MAIN" />

                    <category android:name="android.intent.category.LAUNCHER" />
                </intent-filter>
            </activity>
        </application>
        <uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE"/>
        <uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE"/>
    </manifest>

```

2. 动态申请权限类

```java
    package com.example.myapp;

    import android.Manifest;
    import android.app.Activity;
    import android.content.Context;
    import android.content.pm.PackageManager;
    import android.support.v4.app.ActivityCompat;
    import android.support.v4.content.ContextCompat;

    import java.util.ArrayList;
    import java.util.List;

    /**
    * Created by zhouchaozhong on 2018/1/21.
    */

    public class GrantPermission {
        //要申请的权限清单
        private String[] permissions = {Manifest.permission.WRITE_EXTERNAL_STORAGE,Manifest.permission.READ_EXTERNAL_STORAGE};
        private static final int MY_PERMISSION_REQUEST_CODE = 10000;
        //存储被用户拒绝的权限列表
        List<String> denyList = new ArrayList<String>();


        //检查权限是否授予
        public boolean checkPermission(Context context){
            denyList.clear();
            for(String permission : permissions){
                if(ContextCompat.checkSelfPermission(context,permission) != PackageManager.PERMISSION_GRANTED){
                    denyList.add(permission);
                    return false;
                }
            }

            return true;
        }

        public void requestPermissions(Activity activity){
            ActivityCompat.requestPermissions(activity,permissions,MY_PERMISSION_REQUEST_CODE);
        }
    }


```


文件读写
--------------------------------------------------------------

1. activity_main.xml

```xml
    <?xml version="1.0" encoding="utf-8"?>
    <LinearLayout
        xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:tools="http://schemas.android.com/tools"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:orientation="vertical"
        tools:context="com.example.myapp.MainActivity">


        <EditText
            android:layout_marginTop="10dp"
            android:layout_marginBottom="10dp"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:id="@+id/et_username"
            android:hint="@string/input_username"
            />
        <EditText
            android:inputType="textPassword"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:id="@+id/et_password"
            android:hint="@string/input_password"
            />
        <RelativeLayout
            android:layout_width="match_parent"
            android:layout_height="wrap_content">
            <CheckBox
                android:layout_alignParentLeft="true"
                android:layout_centerVertical="true"
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:id="@+id/cb_rem"
                android:text="@string/rem_password"
                />
            <Button
                android:layout_centerVertical="true"
                android:layout_alignParentRight="true"
                android:paddingLeft="50dp"
                android:paddingRight="50dp"
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:id="@+id/bt_login"
                android:text="@string/login"
                />
        </RelativeLayout>
    </LinearLayout>


```

2. MainActivity.java

```java
    package com.example.myapp;

    import android.content.Context;
    import android.os.Environment;
    import android.support.v7.app.AppCompatActivity;
    import android.os.Bundle;
    import android.text.TextUtils;
    import android.text.format.Formatter;
    import android.util.Log;
    import android.view.View;
    import android.widget.Button;
    import android.widget.CheckBox;
    import android.widget.EditText;
    import android.widget.Toast;

    import java.io.File;
    import java.util.Map;

    public class MainActivity extends AppCompatActivity implements View.OnClickListener {
        private EditText et_username;
        private EditText et_password;
        private CheckBox cb_rem;
        private Button bt_login;
        private Context mContext;

        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            et_username = (EditText) findViewById(R.id.et_username);
            et_password = (EditText) findViewById(R.id.et_password);
            cb_rem = (CheckBox) findViewById(R.id.cb_rem);
            bt_login = (Button) findViewById(R.id.bt_login);

            bt_login.setOnClickListener(this);
            Map<String,String> map = UserInfoUtil.getUserInfo(mContext);
            if(map != null){
                String username = map.get("username");
                String password = map.get("password");
                et_username.setText(username);
                et_password.setText(password);
            }

            //授予sd卡读写权限
            GrantPermission gp = new GrantPermission();
            gp.requestPermissions(this);
        }

        private void login(){
            String username = et_username.getText().toString().trim();
            String password = et_password.getText().toString().trim();
            boolean isrem = cb_rem.isChecked();
            if(TextUtils.isEmpty(username) || TextUtils.isEmpty(password)){
                Toast.makeText(mContext,"用户名密码不能为空！",Toast.LENGTH_SHORT).show();
                return;
            }

            if(isrem){
                //判断sd卡状态是否正常
                if(!Environment.getExternalStorageState().equals(Environment.MEDIA_MOUNTED)){
                    Toast.makeText(mContext,"SD卡不存在或未挂载！",Toast.LENGTH_SHORT).show();
                    return;
                }

                //判断sd卡存储空间是否满足文件的存储
                File sdcard_filedir = Environment.getExternalStorageDirectory(); //得到sdcard的目录作为一个文件对象
                long usableSpace = sdcard_filedir.getUsableSpace();  //获取文件目录对象的可用空间
                String usableSpace_str = Formatter.formatFileSize(mContext,usableSpace);
                Log.i("info",usableSpace_str);

                if(usableSpace < 1024*1024*200){  //判断剩余空间是否小于200M
                    Toast.makeText(mContext,"SD卡剩余空间不足！剩余空间为：" + usableSpace_str,Toast.LENGTH_SHORT).show();
                    return;
                }

                boolean result = UserInfoUtil.saveUserInfo(mContext,username,password);
                if(result){
                    Toast.makeText(mContext,"用户名密码保存成功！",Toast.LENGTH_SHORT).show();
                }else{
                    Toast.makeText(mContext,"用户名密码保存失败！",Toast.LENGTH_SHORT).show();
                }
            }else{
                Toast.makeText(mContext,"无需保存！",Toast.LENGTH_SHORT).show();
            }
        }
        public void onClick(View v){
            switch(v.getId()){
                case R.id.bt_login:
                    login();
                    break;
                default:



            }
        }


    }


```

3. UserInfoUtil.java

```
    package com.example.myapp;

    import android.content.Context;
    import android.os.Environment;
    import android.util.Log;

    import java.io.BufferedReader;
    import java.io.File;
    import java.io.FileInputStream;
    import java.io.FileOutputStream;
    import java.io.InputStreamReader;
    import java.util.HashMap;
    import java.util.Map;

    /**
    * Created by zhouchaozhong on 2018/1/21.
    */

    public class UserInfoUtil {

        public static boolean saveUserInfo(Context context, String username, String password){

            try{
                String userinfo = username + "##" + password;
    //            String path = "/data/data/com.example.myapp/";
                //String path = context.getFilesDir().getPath();
                //String path = "/mnt/sdcard/";  //指定保存的路径
    //            String path = Environment.getExternalStorageDirectory().getPath();
    //            String state = Environment.getExternalStorageState();


                //传入私有目录的文件名称，

    //            File file = new File(path,"userinfo.txt");
    //            FileOutputStream fos = new FileOutputStream(file);

                FileOutputStream fos = context.openFileOutput("userinfo.txt",Context.MODE_PRIVATE);
                fos.write(userinfo.getBytes());
                fos.close();
                return true;
            }catch(Exception e){
                e.printStackTrace();
            }

            return false;
        }

        public static Map<String,String> getUserInfo(Context context){
            try{
                //String path = "/data/data/com.example.myapp/";
                //String path = context.getFilesDir().getPath();
                //String path = "/mnt/sdcard/";  //指定保存的路径
    //            String path = Environment.getExternalStorageDirectory().getPath();
    //            File file = new File(path,"userinfo.txt");
    //            FileInputStream fis = new FileInputStream(file);

                FileInputStream fis = context.openFileInput("userinfo.txt");
                BufferedReader br = new BufferedReader(new InputStreamReader(fis));

                String readLine = br.readLine();
                String[] split = readLine.split("##");
                Map<String,String> hashMap = new HashMap<String,String>();
                hashMap.put("username",split[0]);
                hashMap.put("password",split[1]);
                br.close();
                fis.close();
                return hashMap;
            }catch(Exception e){
                e.printStackTrace();
            }

            return null;
        }



    }


```



外置sdcard文件读写
--------------------------------------------------------

1. Activity_main.xml

```
    <?xml version="1.0" encoding="utf-8"?>
    <LinearLayout
        xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:tools="http://schemas.android.com/tools"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:orientation="vertical"
        tools:context="com.example.myapp.MainActivity">


        <EditText
            android:layout_marginTop="10dp"
            android:layout_marginBottom="10dp"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:id="@+id/et_username"
            android:hint="@string/input_username"
            />
        <EditText
            android:inputType="textPassword"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:id="@+id/et_password"
            android:hint="@string/input_password"
            />
        <RelativeLayout
            android:layout_width="match_parent"
            android:layout_height="wrap_content">
            <CheckBox
                android:layout_alignParentLeft="true"
                android:layout_centerVertical="true"
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:id="@+id/cb_rem"
                android:text="@string/rem_password"
                />
            <Button
                android:layout_centerVertical="true"
                android:layout_alignParentRight="true"
                android:paddingLeft="50dp"
                android:paddingRight="50dp"
                android:layout_width="wrap_content"
                android:layout_height="wrap_content"
                android:id="@+id/bt_login"
                android:text="@string/login"
                />
        </RelativeLayout>
    </LinearLayout>


```

2. MainActivity.java

```
    package com.example.myapp2;

    import android.content.Context;
    import android.os.Environment;
    import android.support.v7.app.AppCompatActivity;
    import android.os.Bundle;
    import android.text.TextUtils;
    import android.text.format.Formatter;
    import android.util.Log;
    import android.view.View;
    import android.widget.Button;
    import android.widget.CheckBox;
    import android.widget.EditText;
    import android.widget.Toast;

    import java.io.File;
    import java.util.Map;

    public class MainActivity extends AppCompatActivity implements View.OnClickListener {
        private EditText et_username;
        private EditText et_password;
        private CheckBox cb_rem;
        private Button bt_login;
        private Context mContext;

        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            et_username = (EditText) findViewById(R.id.et_username);
            et_password = (EditText) findViewById(R.id.et_password);
            cb_rem = (CheckBox) findViewById(R.id.cb_rem);
            bt_login = (Button) findViewById(R.id.bt_login);

            bt_login.setOnClickListener(this);
            Map<String,String> map = UserInfoUtil.getUserInfo(mContext);
            if(map != null){
                String username = map.get("username");
                String password = map.get("password");
                et_username.setText(username);
                et_password.setText(password);
            }

        }

        private void login(){
            String username = et_username.getText().toString().trim();
            String password = et_password.getText().toString().trim();
            boolean isrem = cb_rem.isChecked();
            if(TextUtils.isEmpty(username) || TextUtils.isEmpty(password)){
                Toast.makeText(mContext,"用户名密码不能为空！",Toast.LENGTH_SHORT).show();
                return;
            }

            if(isrem){
                //判断sd卡状态是否正常
                if(!Environment.getExternalStorageState().equals(Environment.MEDIA_MOUNTED)){
                    Toast.makeText(mContext,"SD卡不存在或未挂载！",Toast.LENGTH_SHORT).show();
                    return;
                }

                //判断sd卡存储空间是否满足文件的存储
                File sdcard_filedir = Environment.getExternalStorageDirectory(); //得到sdcard的目录作为一个文件对象
                long usableSpace = sdcard_filedir.getUsableSpace();  //获取文件目录对象的可用空间
                String usableSpace_str = Formatter.formatFileSize(mContext,usableSpace);
                Log.i("info",usableSpace_str);

                if(usableSpace < 1024*1024*200){  //判断剩余空间是否小于200M
                    Toast.makeText(mContext,"SD卡剩余空间不足！剩余空间为：" + usableSpace_str,Toast.LENGTH_SHORT).show();
                    return;
                }

                boolean result = UserInfoUtil.saveUserInfo(mContext,username,password);
                if(result){
                    Toast.makeText(mContext,"用户名密码保存成功！",Toast.LENGTH_SHORT).show();
                }else{
                    Toast.makeText(mContext,"用户名密码保存失败！",Toast.LENGTH_SHORT).show();
                }
            }else{
                Toast.makeText(mContext,"无需保存！",Toast.LENGTH_SHORT).show();
            }
        }
        public void onClick(View v){
            switch(v.getId()){
                case R.id.bt_login:
                    login();
                    break;
                default:



            }
        }

    }



```

3. UserInfoUtil.java

```
    package com.example.myapp2;

    import android.content.Context;
    import android.content.SharedPreferences;
    import android.content.SharedPreferences.Editor;
    import android.preference.PreferenceManager;

    import java.util.HashMap;
    import java.util.Map;

    /**
    * Created by zhouchaozhong on 2018/1/21.
    */

    public class UserInfoUtil {

        public static boolean saveUserInfo(Context context, String username, String password){
            //SharedPreferences spf = context.getSharedPreferences("userinfo",Context.MODE_PRIVATE);
            SharedPreferences spf = PreferenceManager.getDefaultSharedPreferences(context);
            Editor editor = spf.edit();
            editor.putString("username",username);
            editor.putString("password",password);
            editor.commit();
            return true;
        }

        public static Map<String,String> getUserInfo(Context context){
            //SharedPreferences spf = context.getSharedPreferences("userinfo",Context.MODE_PRIVATE);

            SharedPreferences spf = PreferenceManager.getDefaultSharedPreferences(context);
            String username = spf.getString("username","admin");
            String password = spf.getString("password","admin");
            Map<String,String> map = new HashMap<String,String>();
            map.put("username",username);
            map.put("password",password);
            return map;
        }
    }

```

对xml文件的读写操作
-----------------------------------------------------------

1. MainActivity.java
```
    package com.example.myapp3;

    import android.content.Context;
    import android.support.v7.app.AppCompatActivity;
    import android.os.Bundle;
    import android.view.View;
    import android.widget.Button;
    import android.widget.Toast;

    public class MainActivity extends AppCompatActivity implements View.OnClickListener{
        private Context mContext;
        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            Button bt_backup = (Button) findViewById(R.id.bt_backup);
            Button bt_restore = (Button) findViewById(R.id.bt_restore);

            bt_backup.setOnClickListener(this);
            bt_restore.setOnClickListener(this);

        }

        @Override
        public void onClick(View view) {
            switch (view.getId()){
                case R.id.bt_backup:
                    if(SmsUtils.backupSms(mContext)){
                        Toast.makeText(mContext, "备份成功！", Toast.LENGTH_SHORT).show();
                    }else{
                        Toast.makeText(mContext, "备份失败！", Toast.LENGTH_SHORT).show();
                    }
                    break;
                case R.id.bt_restore:
                    int result = SmsUtils.restoreSms(mContext);
                    Toast.makeText(mContext,"成功恢复" + result + "条短信",Toast.LENGTH_SHORT).show();
                    break;
                default:
                    break;
            }
        }


    }

```

2. SmsUtils.java

```
    package com.example.myapp3;

    import android.content.Context;
    import android.content.res.AssetManager;
    import android.util.Xml;

    import org.xmlpull.v1.XmlPullParser;
    import org.xmlpull.v1.XmlSerializer;

    import java.io.IOException;
    import java.io.InputStream;
    import java.util.ArrayList;
    import java.util.List;

    /**
    * Created by zhouchaozhong on 2018/1/22.
    */

    class SmsUtils {
        public static boolean backupSms(Context context) {
            /*
                //通过写入字符的方式创建xml文件
            StringBuffer sb = new StringBuffer();
            sb.append("<?xml version='1.0' encoding='utf-8' standalone='yes' ?>");
            sb.append("<Smss>");
            sb.append("</Smss>");
            try {
                FileOutputStream fos = context.openFileOutput("backupsms.xml",Context.MODE_PRIVATE);
                fos.write(sb.toString().getBytes());
                fos.close();
                return true;
            } catch (Exception e) {
                e.printStackTrace();
            }

            return false;

            */

            List<SmsBean> list = new ArrayList<SmsBean>();
            SmsBean sb1 = new SmsBean("13525689752","你好吗？","2018-01-22",1);
            SmsBean sb2 = new SmsBean("15565987562","我很好！","2018-01-23",2);
            SmsBean sb3 = new SmsBean("15632568966","吃晚饭了吗？","2018-01-25",3);
            list.add(sb1);
            list.add(sb2);
            list.add(sb3);

            try {
                //创建一个XmlSerializer对象
                XmlSerializer xs = Xml.newSerializer();
                //设置XmlSerializer的参数
                xs.setOutput(context.openFileOutput("backupsms2.xml",Context.MODE_PRIVATE),"utf-8");
                //序列化一个xml的声明头，(encoding,standalone)编码和是否独立
                xs.startDocument("utf-8",true);

                //序列化一个根节点的开始节点（namespace,name）命名空间，标签名称
                xs.startTag(null,"Smss");

                //循环遍历list集合中的信息
                for(SmsBean obj : list){
                    xs.startTag(null,"Sms");
                    xs.attribute(null,"id",obj.id+"");

                    xs.startTag(null,"num");
                    xs.text(obj.num);
                    xs.endTag(null,"num");

                    xs.startTag(null,"msg");
                    xs.text(obj.msg);
                    xs.endTag(null,"msg");

                    xs.startTag(null,"date");
                    xs.text(obj.date);
                    xs.endTag(null,"date");

                    xs.endTag(null,"Sms");


                }
                //序列化一个根节点的结束节点
                xs.endTag(null,"Smss");

                //将xml写入文件中，完成xml的序列化
                xs.endDocument();
                return true;
            } catch (IOException e) {
                e.printStackTrace();
            }

            return false;


        }

        //解析xml文件
        public static int restoreSms(Context context) {
            ArrayList<SmsBean> arrayList = null;
            SmsBean smsBean = null;
            try {
                //通过Xml获取一个XmlPullParser对象
                XmlPullParser xpp = Xml.newPullParser();

                //通过context获取资产管理者对象
                AssetManager assets = context.getAssets();

                //通过资产管理者对象获取一个文件读取流
                InputStream is = assets.open("backupsms2.xml");

                //设置XmlPullParser对象的参数
                //xpp.setInput(context.openFileInput("backupsms2.xml"), "utf-8");

                xpp.setInput(is,"utf-8");
                //获取当前xml行的事件类型
                int type = xpp.getEventType();
                //判断事件类型是否是文档结束的事件类型，如果不是，循环遍历解析每一行，解析一行后，获取下一行的事件类型
                while(type != XmlPullParser.END_DOCUMENT){

                    String currentTagName = xpp.getName();
                    //判断当前行的事件类型是开始标签还是结束标签
                    switch(type){
                        case XmlPullParser.START_TAG:
                            if(currentTagName.equals("Smss")){
                                //如果当前标签是Smss，需要创建一个集合
                                arrayList = new ArrayList<SmsBean>();
                            }else if(currentTagName.equals("Sms")){
                                smsBean = new SmsBean();
                                smsBean.id = Integer.valueOf(xpp.getAttributeValue(null,"id"));
                            }else if(currentTagName.equals("num")){
                                smsBean.num = xpp.nextText();
                            }else if(currentTagName.equals("msg")){
                                smsBean.num = xpp.nextText();
                            }else if(currentTagName.equals("date")){
                                smsBean.num = xpp.nextText();
                            }
                            break;
                        case XmlPullParser.END_TAG:
                            if(currentTagName.equals("Sms")){
                                //将解析的数据封装到list中
                                arrayList.add(smsBean);
                            }
                            break;
                    }
                    type = xpp.next();
                }

                return arrayList.size();
            }catch(Exception e){
                e.printStackTrace();
            }

            return 0;

        }
    }

```

3. SmsBean.java

```
    package com.example.myapp3;

    /**
    * Created by zhouchaozhong on 2018/1/22.
    */

    public class SmsBean {
        public String num ;
        public String msg;
        public String date;
        public int id;

        public SmsBean(){

        }

        public SmsBean(String num,String msg,String date,int id){
            this.num = num;
            this.msg = msg;
            this.date = date;
            this.id = id;
        }

    }

```


数据库的CRUD操作
----------------------------------------------------------------------

1. 增删改查
```
    package com.example.myapp4;

    import android.content.Context;
    import android.database.sqlite.SQLiteDatabase;
    import android.database.sqlite.SQLiteOpenHelper;

    /**
    * Created by zhouchaozhong on 2018/1/22.
    */

    public class MySqliteOpenHelper extends SQLiteOpenHelper {

        public MySqliteOpenHelper(Context context){
            super(context,"info.db",null,1);
        }


        //数据库第一次创建时会被调用，适合做表结构的初始化
        @Override
        public void onCreate(SQLiteDatabase sqLiteDatabase) {
            sqLiteDatabase.execSQL("create table info(_id integer primary key autoincrement,name varchar(20),phone varchar(11))");
        }

        //数据库版本号发生改变才会执行，适合做表结构的修改
        @Override
        public void onUpgrade(SQLiteDatabase sqLiteDatabase, int i, int i1) {

        }
    }






    package com.example.myapp4;

    import android.content.ContentValues;
    import android.content.Context;
    import android.database.Cursor;
    import android.database.sqlite.SQLiteDatabase;

    /**
    * Created by zhouchaozhong on 2018/1/22.
    */

    public class InfoDao {
        private MySqliteOpenHelper mySqliteOpenHelper;
        public InfoDao(Context context){
            //创建一个帮助类对象
        mySqliteOpenHelper = new MySqliteOpenHelper(context);

        }

        public boolean add(InfoBean bean){
            //调用getReadableDatabase方法类初始化数据库
            SQLiteDatabase db = mySqliteOpenHelper.getWritableDatabase();
            //db.execSQL("insert into info(name,phone) values (?,?);",new Object[]{bean.name,bean.phone});
            ContentValues cv = new ContentValues();
            cv.put("name",bean.name);
            cv.put("phone",bean.phone);
            long result = db.insert("info",null,cv);

            //关闭数据库对象
            db.close();

            if(result != -1){  //-1代表添加失败
                return true;
            }else{
                return false;
            }
        }

        public int delete(String name){
            //调用getReadableDatabase方法类初始化数据库
            SQLiteDatabase db = mySqliteOpenHelper.getReadableDatabase();

            //db.execSQL("delete from info where name=?;",new Object[]{name});

            int result = db.delete("info","name=?",new String[]{name});
            //关闭数据库对象
            db.close();

            return result;
        }

        public int update(InfoBean bean){
            //调用getReadableDatabase方法类初始化数据库
            SQLiteDatabase db = mySqliteOpenHelper.getReadableDatabase();
            //db.execSQL("update info set phone=? where name=?;",new Object[]{bean.phone,bean.name});

            ContentValues cv = new ContentValues();
            cv.put("phone",bean.phone);  //要更新的值
            int result = db.update("info",cv,"name=?",new String[]{bean.name});
            //关闭数据库对象
            db.close();

            return result;
        }

        public void query(String name){
            //调用getReadableDatabase方法类初始化数据库
            SQLiteDatabase db = mySqliteOpenHelper.getReadableDatabase();
            //Cursor cursor = db.rawQuery("select _id,name,phone from info where name=?",new String[]{name});

            //table:表名 columns:查询的列名  如果null代表查询所有列 selection:查询条件 selectionArgs:条件占位符的参数值
            //groupBy按什么字段分组 having:分组条件  orderBy按什么字段排序
            Cursor cursor = db.query("info",new String[]{"_id","name","phone"},"name=?",new String[]{name},null,null,"_id desc");

            //解析cursor中的数据
            if(cursor != null && cursor.getCount() > 0){ //判断cursor中是否存在数据
                //循环遍历结果集，获取每一行的内容
                while(cursor.moveToNext()){
                    //获取数据
                    int id = cursor.getInt(0);
                    String name_str = cursor.getString(1);
                    String phone = cursor.getString(2);

                    System.out.println("_id:"+id+"  name:"+name_str+"  phone:"+phone);
                }
                cursor.close();
            }
            //关闭数据库对象
            db.close();
        }
    }

```

2. 事务操作

```
    package com.example.myapp4;

    import android.content.Context;
    import android.database.sqlite.SQLiteDatabase;
    import android.database.sqlite.SQLiteOpenHelper;

    /**
    * Created by zhouchaozhong on 2018/1/23.
    */

    public class BankOpenHelper extends SQLiteOpenHelper {
        public BankOpenHelper(Context context){
            super(context,"bank.db",null,1);
        }

        @Override
        public void onCreate(SQLiteDatabase db) {
            db.execSQL("create table account (_id integer primary key autoincrement,name varchar(20),money varchar(20))");
            db.execSQL("insert into account ('name','money') values ('张三','2000')");
            db.execSQL("insert into account ('name','money') values ('李四','5000')");
        }

        @Override
        public void onUpgrade(SQLiteDatabase sqLiteDatabase, int i, int i1) {

        }
    }



    
    package com.example.myapp4;

    import android.content.Context;
    import android.database.sqlite.SQLiteDatabase;

    /**
    * Created by zhouchaozhong on 2018/1/23.
    */

    public class BankDao {
        private BankOpenHelper myBankOpenHelper;
        public BankDao(Context context){
            //创建一个帮助类对象
            myBankOpenHelper = new BankOpenHelper(context);
        }

        public void transtation(){
            SQLiteDatabase db = myBankOpenHelper.getReadableDatabase();
            db.beginTransaction();//开启一个数据库事务
            try{
                db.execSQL("update account set money= money-200 where name=?",new String[]{"李四"});
                int i = 100/0;//模拟一个异常
                db.execSQL("update account set money= money+200 where name=?",new String[]{"张三"});
                db.setTransactionSuccessful();  //标记事务中的sql语句全部成功执行！
            }finally {
                db.endTransaction();  //判断事务的标记是否成功，如果不成功，回滚错误之前sql语句
            }

        }
    }


```

ListView的使用
--------------------------------------------------------

* 布局

    1. activity_main.xml

    ```
        <?xml version="1.0" encoding="utf-8"?>
        <RelativeLayout xmlns:android="http://schemas.android.com/apk/res/android"
            xmlns:app="http://schemas.android.com/apk/res-auto"
            xmlns:tools="http://schemas.android.com/tools"
            android:layout_width="match_parent"
            android:layout_height="match_parent"
            tools:context="com.example.myapp5.MainActivity">

            <ListView
                android:id="@+id/lv_news"
                android:layout_width="wrap_content"
                android:layout_height="match_parent"></ListView>



        </RelativeLayout>

    ```

    2. item_news_layout.xml

    ```
        <?xml version="1.0" encoding="utf-8"?>
        <LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
            android:orientation="horizontal"
            android:paddingTop="20dp"
            android:paddingBottom="20dp"
            android:layout_width="match_parent"
            android:layout_height="match_parent">
            <ImageView
                android:id="@+id/item_img_icon"
                android:src="@drawable/icon"
                android:layout_gravity="center"
                android:layout_marginRight="10dp"
                android:layout_width="68dp"
                android:layout_height="68dp" />
            <LinearLayout
                android:orientation="vertical"
                android:gravity="center"
                android:layout_width="match_parent"
                android:layout_height="match_parent">
                <TextView
                    android:id="@+id/item_tv_title"
                    android:lines="1"
                    android:ellipsize="end"
                    android:text="title"
                    android:textSize="16sp"
                    android:textColor="#000000"
                    android:layout_marginBottom="3dp"
                    android:layout_width="match_parent"
                    android:layout_height="wrap_content" />
                <TextView
                    android:id="@+id/item_tv_des"
                    android:maxLines="2"
                    android:text="des"
                    android:textSize="13sp"
                    android:textColor="#666666"
                    android:layout_width="match_parent"
                    android:layout_height="wrap_content" />
            </LinearLayout>
        </LinearLayout>

    ```

* MainActivity.java

```
    package com.example.myapp5;

    import android.content.Context;
    import android.content.Intent;
    import android.database.DataSetObserver;
    import android.graphics.Color;
    import android.net.Uri;
    import android.support.v7.app.AppCompatActivity;
    import android.os.Bundle;
    import android.view.View;
    import android.view.ViewGroup;
    import android.widget.AdapterView;
    import android.widget.BaseAdapter;
    import android.widget.Button;
    import android.widget.ListView;
    import android.widget.TextView;
    import android.widget.Toast;

    import java.util.ArrayList;
    import java.util.Random;

    public class MainActivity extends AppCompatActivity implements AdapterView.OnItemClickListener{
        private Context mContext;
        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            //找到ListView
            ListView lv_news= (ListView) findViewById(R.id.lv_news);
            ArrayList<NewsBean> allNews = NewsUtil.getAllNews(mContext);

            //创建一个Adapter，并设置给lv_news
            NewsAdapter newsAdapter = new NewsAdapter(mContext,allNews);
            lv_news.setAdapter(newsAdapter);

            //设置listview条目的点击事件
            lv_news.setOnItemClickListener(this);

        }


        //listview条目点击时会调用该方法
        @Override
        public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
            //获取条目中bean对象上的url做跳转
            NewsBean bean = (NewsBean) adapterView.getItemAtPosition(i);
            String url = bean.news_url;
            //跳转到浏览器
            Intent intent = new Intent();
            intent.setAction(Intent.ACTION_VIEW);
            intent.setData(Uri.parse(url));
            startActivity(intent);
        }
    }


```

* NewsAdapter.java

```
    package com.example.myapp5;

    import android.content.Context;
    import android.view.View;
    import android.view.ViewGroup;
    import android.widget.BaseAdapter;
    import android.widget.ImageView;
    import android.widget.TextView;

    import java.lang.reflect.Array;
    import java.util.ArrayList;

    /**
    * Created by zhouchaozhong on 2018/1/24.
    */

    public class NewsAdapter extends BaseAdapter {
        private ArrayList<NewsBean> list;
        private Context context;

        //通过构造方法接收要显示的新闻的数据集合

        public NewsAdapter(Context context, ArrayList<NewsBean> list){
            this.list = list;
            this.context = context;
        }
        @Override
        public int getCount() {
            return list.size();
        }

        @Override
        public Object getItem(int i) {
            //return null;
            return list.get(i);
        }

        @Override
        public long getItemId(int i) {
        //return 0;
            return i;
        }

        @Override
        public View getView(int i, View view, ViewGroup viewGroup) {
            View myView = null;
            if(view != null){
                myView = view;
            }else{
                myView = View.inflate(context,R.layout.item_news_layout,null);

            }

            //获取myView上的子控件对象
            ImageView item_img_icon = (ImageView) myView.findViewById(R.id.item_img_icon);
            TextView item_tv_title = (TextView) myView.findViewById(R.id.item_tv_title);
            TextView item_tv_des = (TextView) myView.findViewById(R.id.item_tv_des);

            //获取对应条目的新闻数据
            NewsBean bean = list.get(i);

            //给子控件设置数据
            item_img_icon.setImageDrawable(bean.icon);  //设置ImageView的图片
            item_tv_title.setText(bean.title);  //设置标题
            item_tv_des.setText(bean.des);   //设置描述
            return myView;
        }
    }


```  

* NewsUtil.java

```
    package com.example.myapp5;

    import android.content.Context;
    import android.support.v4.content.ContextCompat;

    import java.util.ArrayList;

    /**
    * Created by zhouchaozhong on 2018/1/24.
    */

    class NewsUtil {
        public static ArrayList<NewsBean> getAllNews(Context context) {
            ArrayList<NewsBean> arrayList = new ArrayList<NewsBean>();
            for(int i = 0 ;i <100;i++)
            {
                NewsBean newsBean = new NewsBean();
                newsBean.title ="谢霆锋经纪人偷拍系侵权行为谢霆锋经纪人偷拍系侵权行为谢霆锋经纪人偷拍系侵权行为谢霆锋经纪人：偷拍系侵权行为";
                newsBean.des= "称谢霆锋隐私权收到侵犯，将保留追究法律责任";
                newsBean.news_url= "http://www.sina.cn";
                //newsBean.icon = context.getResources().getDrawable(R.drawable.ic_launcher);//通过context对象将一个资源id转换成一个Drawable对象。
                newsBean.icon = ContextCompat.getDrawable(context,R.drawable.ic_launcher);   //上述方法已过时，这个方法是谷歌推荐使用的
                arrayList.add(newsBean);

                NewsBean newsBean1 = new NewsBean();
                newsBean1.title ="知情人：王菲是谢霆锋心头最爱的人";
                newsBean1.des= "身边的人都知道谢霆锋最爱王菲，二人早有复合迹象";
                newsBean1.news_url= "http://www.baidu.cn";
            //newsBean1.icon = context.getResources().getDrawable(R.drawable.icon);//通过context对象将一个资源id转换成一个Drawable对象。
                newsBean1.icon = ContextCompat.getDrawable(context,R.drawable.icon); //上述方法已过时，这个方法是谷歌推荐使用的

                arrayList.add(newsBean1);



                NewsBean newsBean2 = new NewsBean();
                newsBean2.title ="热烈祝贺黑马74高薪就业";
                newsBean2.des= "74期平均薪资20000，其中有一个哥们超过10万，这些It精英都迎娶了白富美.";
                newsBean2.news_url= "http://www.itheima.com";
            // newsBean2.icon = context.getResources().getDrawable(R.drawable.icon2);//通过context对象将一个资源id转换成一个Drawable对象。
                newsBean2.icon = ContextCompat.getDrawable(context,R.drawable.icon2); //上述方法已过时，这个方法是谷歌推荐使用的

                arrayList.add(newsBean2);
            }
            return arrayList;

        }
    }


```  

* NewsBean.java

```
    package com.example.myapp5;

    import android.graphics.drawable.Drawable;

    /**
    * Created by zhouchaozhong on 2018/1/24.
    */

    class NewsBean {
        public String title;
        public String des;
        public Drawable icon;
        public String news_url;
    }

```


网络请求和消息处理
----------------------------------------------------------

    1. 网路请求数据（新闻app案例）

    ```
        MainActivity.java


        package com.example.myapp7;

        import android.content.Context;
        import android.content.Intent;
        import android.database.DataSetObserver;
        import android.graphics.Color;
        import android.net.Uri;
        import android.os.Handler;
        import android.os.Message;
        import android.support.v7.app.AppCompatActivity;
        import android.os.Bundle;
        import android.view.View;
        import android.view.ViewGroup;
        import android.widget.AdapterView;
        import android.widget.BaseAdapter;
        import android.widget.Button;
        import android.widget.ListView;
        import android.widget.TextView;
        import android.widget.Toast;

        import java.util.ArrayList;
        import java.util.Random;

        public class MainActivity extends AppCompatActivity implements AdapterView.OnItemClickListener{
            private Context mContext;
            private ListView lv_news;
            @Override
            protected void onCreate(Bundle savedInstanceState) {
                super.onCreate(savedInstanceState);
                setContentView(R.layout.activity_main);

                mContext = this;
                //找到ListView
                lv_news = (ListView) findViewById(R.id.lv_news);
                //获取新闻数据用list封装，获取网络数据需要在子线程中做
                new Thread(new Runnable() {
                    @Override
                    public void run() {
                        //请求网络数据
                        ArrayList<NewsBean> allNews = NewsUtil.getAllNewsForNetwork(mContext);
                        //通过handler将Message发送给主线程去更新UI
                        Message msg = Message.obtain();
                        msg.obj = allNews;
                        handler.sendMessage(msg);
                    }
                }).start();


                //设置listview条目的点击事件
                lv_news.setOnItemClickListener(this);

            }

            private Handler handler = new Handler(){
                @Override
                public void handleMessage(Message msg) {
                    ArrayList<NewsBean> allNews = (ArrayList<NewsBean>) msg.obj;

                    if(allNews != null && allNews .size()>0)
                    {
                        //3.创建一个adapter设置给listview
                        NewsAdapter newsAdapter = new NewsAdapter(mContext, allNews);
                        lv_news.setAdapter(newsAdapter);
                    }
                }
            };


            //listview条目点击时会调用该方法
            @Override
            public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
                //获取条目中bean对象上的url做跳转
                NewsBean bean = (NewsBean) adapterView.getItemAtPosition(i);
                String url = bean.news_url;
                //跳转到浏览器
                Intent intent = new Intent();
                intent.setAction(Intent.ACTION_VIEW);
                intent.setData(Uri.parse(url));
                startActivity(intent);
            }
        }





        NewsAdapter.java


        package com.example.myapp7;

        import android.content.Context;
        import android.view.View;
        import android.view.ViewGroup;
        import android.widget.BaseAdapter;
        import android.widget.ImageView;
        import android.widget.TextView;


        import com.loopj.android.image.SmartImageView;

        import java.lang.reflect.Array;
        import java.util.ArrayList;

        /**
        * Created by zhouchaozhong on 2018/1/24.
        */

        public class NewsAdapter extends BaseAdapter {
            private ArrayList<NewsBean> list;
            private Context context;

            //通过构造方法接收要显示的新闻的数据集合

            public NewsAdapter(Context context, ArrayList<NewsBean> list){
                this.list = list;
                this.context = context;
            }
            @Override
            public int getCount() {
                return list.size();
            }

            @Override
            public Object getItem(int i) {
                //return null;
                return list.get(i);
            }

            @Override
            public long getItemId(int i) {
            //return 0;
                return i;
            }

            @Override
            public View getView(int i, View view, ViewGroup viewGroup) {
                View myView = null;
                if(view != null){
                    myView = view;
                }else{
                    myView = View.inflate(context,R.layout.item_news_layout,null);

                }

                //获取myView上的子控件对象
                SmartImageView item_img_icon = (SmartImageView) myView.findViewById(R.id.item_img_icon);
                TextView item_tv_title = (TextView) myView.findViewById(R.id.item_tv_title);
                TextView item_tv_des = (TextView) myView.findViewById(R.id.item_tv_des);

                //获取对应条目的新闻数据
                NewsBean bean = list.get(i);

                //给子控件设置数据
                //item_img_icon.setImageDrawable(bean.icon);  //设置ImageView的图片
                item_img_icon.setImageUrl(bean.icon_url);
                item_tv_title.setText(bean.title);  //设置标题
                item_tv_des.setText(bean.des);   //设置描述
                return myView;
            }
        }



        newsdata.php


        <?php

        $newsArr = array();
        $newsArr['newss'] = array(
            array(
                "id" => 2,
                "time" => "2015-08-07",
                "des" => "7月29日，历经9个月数百万人内测完善之后，微软终于发布Win10正式版系统。但是可能对于部分用户而言，Win7仍然是绝对的经典、游戏玩家的不二之选，为何非要升级到Win10系统呢？Windows10性能和功能相比Windows7，有提升吗？下面IT之家就为大家带来Win7与Win10功能与性能的正面PK，相信还在犹豫不决的用户看完本文心里就会有了答案。",
                "title" => "升还是不升：Win7、Win10全面对比评测",
                "news_url" => "http://toutiao.com/a5229867988/",
                "icon_url" => "http://p2.pstatp.com/large/6850/6105376239",
                "comment" => 5000,
                "type" => 1
            ),
            array(
                "id" => 1,
                "time" => "2015-08-09",
                "des" => "苹果iPhone6s发布的具体时间越传越近了，但至今还没有官方的准信儿。今天还是让我们关注一下火狐漏洞吧。火狐浏览器开发商Mozilla提醒用户立即升级到最新版本，最好还要修改密码和登录信息。",
                "title" => "苹果或于9月9日发布iPhone6s 火狐爆严重漏洞",
                "news_url" => "http://m.jiemian.com/article/347958.html",
                "icon_url" => "http://img.jiemian.com/101/original/20150808/143899303035536900_a580x330.jpg",
                "comment" => 1200,
                "type" => 3
            ),
            array(
                "id" => 0,
                "time" => "2015-08-10",
                "des" => "凤凰科技讯 8月8日消息，今天有网友爆料，京东创始人兼CEO刘强东已与奶茶妹妹章泽天在朝阳区民政局领证结婚。",
                "title" => "刘强东与奶茶妹妹领证结婚 有图为证",
                "news_url" => "http://i.ifeng.com/news/sharenews.f?aid=100435430",
                "icon_url" => "http://d.ifengimg.com/mw604/y3.ifengimg.com/ifengimcp/pic/20150808/ce1b80056cfc584fafbf_size20_w450_h800.jpg",
                "comment" => 3000,
                "type" => 2
            )

        );

        $data = json_encode($newsArr);
        
        echo $data;


        ?>



        NewsUtil.java

        package com.example.myapp7;

        import android.content.Context;
        import android.support.v4.content.ContextCompat;

        import org.json.JSONArray;
        import org.json.JSONObject;

        import java.io.InputStream;
        import java.net.HttpURLConnection;
        import java.net.MalformedURLException;
        import java.net.URL;
        import java.util.ArrayList;

        /**
        * Created by zhouchaozhong on 2018/1/24.
        */

        class NewsUtil {
            public static String newsPathUrl = "http://192.168.33.10/newsdata.php";
            public static ArrayList<NewsBean> getAllNewsForNetwork(Context context) {
                ArrayList<NewsBean> arrayList = new ArrayList<NewsBean>();

                try {
                    //请求服务器获取新闻数据
                    URL url = new URL(newsPathUrl);
                    HttpURLConnection connection = (HttpURLConnection) url.openConnection();
                    connection.setRequestMethod("GET");
                    connection.setConnectTimeout(1000*10);
                    int code = connection.getResponseCode();
                    if(code == 200){
                        InputStream in = connection.getInputStream();
                        String result = StreamUtils.streamToString(in);
                        //解析获取到的新闻数据到List集合中
                        JSONObject root_json = new JSONObject(result);
                        //获取root_json中的newss作为JSONArray的对象
                        JSONArray jsonArray = root_json.getJSONArray("newss");

                        for(int i = 0;i < jsonArray.length();i++){
                            JSONObject news_json = jsonArray.getJSONObject(i);
                            NewsBean newsBean = new NewsBean();
                            newsBean.id = news_json.getInt("id");
                            newsBean.comment = news_json.getInt("comment");
                            newsBean.type = news_json.getInt("type");  //新闻类型
                            newsBean.time = news_json.getString("time");
                            newsBean.des = news_json.getString("des");
                            newsBean.title = news_json.getString("title");
                            newsBean.news_url = news_json.getString("news_url");
                            newsBean.icon_url = news_json.getString("icon_url");
                            arrayList.add(newsBean);
                        }
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                }
                //解析获取的新闻数据到List
                return arrayList;

            }
        }




        NewsBean.java

        package com.example.myapp7;

        import android.graphics.drawable.Drawable;

        /**
        * Created by zhouchaozhong on 2018/1/24.
        */

        class NewsBean {
        public int id;
        public int comment;
        public int type;
        public String time;
        public String des;
        public String title;
        public String news_url;
        public String icon_url;
        public Drawable icon;
        }


    ```

    2. GET方式和POST方式请求服务器数据

    ```
    package com.example.myapp8;

    import android.content.Context;
    import android.os.Handler;
    import android.os.Message;
    import android.support.v7.app.AppCompatActivity;
    import android.os.Bundle;
    import android.view.View;
    import android.widget.Button;
    import android.widget.EditText;
    import android.widget.Toast;

    import java.io.InputStream;
    import java.net.HttpURLConnection;
    import java.net.URL;

    public class MainActivity extends AppCompatActivity implements View.OnClickListener {
        private Context mContext;
        private EditText et_username;
        private EditText et_password;
        private Button bt_submit;

        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            et_username = (EditText) findViewById(R.id.et_username);
            et_password = (EditText) findViewById(R.id.et_password);
            bt_submit = (Button) findViewById(R.id.bt_submit);
            bt_submit.setOnClickListener(this);
        }

        private Handler handler = new Handler(){
            @Override
            public void handleMessage(Message msg) {
                boolean isLogin = (boolean) msg.obj;
                if(isLogin){
                    Toast.makeText(mContext,"登录成功！",Toast.LENGTH_SHORT).show();
                }else{
                    Toast.makeText(mContext,"登录失败！",Toast.LENGTH_SHORT).show();
                }
            }
        };

        @Override
        public void onClick(View view) {
            final String username = et_username.getText().toString().trim();
            final String password = et_password.getText().toString().trim();
            new Thread(new Runnable() {
                @Override
                public void run() {
                    //boolean isLogin = requestNetForGetLogin(username,password);
                    boolean isLogin = requestNetForPostLogin(username,password);
                    Message msg = Message.obtain();
                    msg.obj = isLogin;
                    handler.sendMessage(msg);

                }
            }).start();
        }

        private boolean requestNetForGetLogin(String username,String password){
            try{
                URL url = new URL("http://192.168.33.10/androidlogin.php?username="+username+"&password="+password);
                HttpURLConnection connection = (HttpURLConnection) url.openConnection();
                connection.setRequestMethod("GET");
                connection.setConnectTimeout(10*1000);
                int code = connection.getResponseCode();
                if(code == 200){
                    InputStream inputStream = connection.getInputStream();
                    String result = StreamUtils.streamToString(inputStream);
                    if(result.contains("success")){
                        return true;
                    }
                }
            }catch (Exception e){
                e.printStackTrace();
            }

            return false;
        }

        private boolean requestNetForPostLogin(String username,String password){
            try{
                URL url = new URL("http://192.168.33.10/androidlogin.php");
                HttpURLConnection connection = (HttpURLConnection) url.openConnection();
                connection.setRequestMethod("POST");
                connection.setConnectTimeout(10*1000);
                //设置一些请求头信息
                String body = "username="+username+"&password="+password;
                connection.setRequestProperty("Content-Length",body.length()+"");

                //设置URLConnection可以写请求的内容
                connection.setDoOutput(true);
                //获取一个OutputStream，并将内容写入该流
                connection.getOutputStream().write(body.getBytes());
                int code = connection.getResponseCode();
                if(code == 200){
                    InputStream inputStream = connection.getInputStream();
                    String result = StreamUtils.streamToString(inputStream);
                    System.out.println("----------------------------------------------------------------"+result);
                    if(result.contains("success")){
                        return true;
                    }
                }
            }catch (Exception e){
                e.printStackTrace();
            }

            return false;
        }


    }


    ```


Activity的创建和跳转
--------------------------------------------------------------

1. 显式意图与隐式意图

```
    public void onClick(View view) {
            switch (view.getId()){
                case R.id.bt_call:
                    Intent intent = new Intent();
                    intent.setAction(Intent.ACTION_CALL);
                    intent.setData(Uri.parse("tel:"+112));
                    startActivity(intent);
                    break;
                case R.id.bt_redirect:
                    //隐式意图跳转
                    Intent intent1 = new Intent();
                    intent1.setAction("com.example.test");
                    intent1.setData(Uri.parse("info:helloinfo"));
                    intent1.putExtra("name","张三");
                    intent1.putExtra("sex",1);
                    intent1.addCategory("android.intent.category.DEFAULT");
                    startActivity(intent1);
                    break;
                case R.id.bt_redirect2:
                    //显式意图跳转
                    //Intent intent2 = new Intent();
                    Intent intent2 = new Intent(this,Test2Activity.class);
                    //指定包名和类名 开启页面
                    //intent2.setClassName("com.example.myapp9","com.example.myapp9.Test2Activity");
                    startActivity(intent2);
                    break;
                default:
                    break;

            }

        }



        //接收数据

         Intent intent = getIntent();
        //Uri data = intent.getData();
        //String str = data.getScheme();
        String name = intent.getStringExtra("name");
        int sex = intent.getIntExtra("sex",0);

        String str = "name:" + name + "   sex:"+sex;

```

2. AndroidManifest.xml配置

```
   <application
        android:allowBackup="true"
        android:icon="@mipmap/ic_launcher"
        android:label="@string/app_name"
        android:roundIcon="@mipmap/ic_launcher_round"
        android:supportsRtl="true"
        android:theme="@style/AppTheme">
        <activity android:name=".MainActivity">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />

                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
        <activity android:name="com.example.myapp9.TestActivity"
            android:label="第二个页面"
            >
            <intent-filter>
                <action android:name="com.example.test" />
                <data android:scheme="info"/>
                <category android:name="android.intent.category.DEFAULT" />
            </intent-filter>
        </activity>
        <activity android:name="com.example.myapp9.Test2Activity"
            android:label="第三个页面">

        </activity>
    </application>


    <!--
    
       activity android:name="com.example.myapp9.TestActivity"  这里指要配置的activity名
       <action android:name="com.example.test" />   要配置的activity  的 action名
    
    
    
    -->

```


广播接收者
-----------------------------------------------------

* MainActivity

```
    package com.example.myapp10;

    import android.content.Context;
    import android.content.Intent;
    import android.support.v7.app.AppCompatActivity;
    import android.os.Bundle;
    import android.view.View;
    import android.widget.Button;

    public class MainActivity extends AppCompatActivity  implements View.OnClickListener{
        private Button bt_send;
        private Button bt_order_send;
        private Context mContext;

        @Override
        protected void onCreate(Bundle savedInstanceState){
            mContext = this;
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            GrantPermission gp = new GrantPermission();
            gp.requestPermissions(this);

            Button bt_send = (Button) findViewById(R.id.bt_send);
            Button bt_order_send = (Button) findViewById(R.id.bt_order_send);
            bt_send.setOnClickListener(this);
            bt_order_send.setOnClickListener(this);
        }

        @Override
        public void onClick(View view) {
            switch(view.getId()){
                case R.id.bt_send:
                    //发送一条无序广播
                    Intent intent = new Intent();
                    intent.setAction("com.example.myapp10.custombroadcast");
                    intent.putExtra("name","每天晚上七点准时直播！");
                    sendBroadcast(intent);
                    break;
                case R.id.bt_order_send:
                    //发送一条有序广播
                    Intent intent2 = new Intent();
                    intent2.setAction("com.example.myapp10.rice");
                    sendOrderedBroadcast(intent2,null,new FinalReceiver(),null,1,"发大米了，快来领大米了！每人1000斤！",null);
                    System.out.println("发大米了，快来领大米了！每人1000斤！");
                    break;
                default:
                    break;
            }
        }
    }


```


* OutGoingCallReceiver

```
    package com.example.myapp10;

    import android.content.BroadcastReceiver;
    import android.content.Context;
    import android.content.Intent;
    import android.widget.Toast;

    /**
    * Created by zhouchaozhong on 2018/2/9.
    */

    public class OutGoingCallReceiver extends BroadcastReceiver {
        @Override
        public void onReceive(Context context, Intent intent) {
            //开启MainActivity
    //        Intent intent2 = new Intent(context,MainActivity.class);
            //在广播接收者里面开启Activity，要设置任务栈环境
    //        intent2.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
    //        context.startActivity(intent2);

    //        String action = intent.getAction();
    //        System.out.println(action);
    //        String currentNumber = getResultData();
    //        if(currentNumber.startsWith("0")){
    //            setResultData("17951" + currentNumber);
    //        }

            //接收无序广播
            String content = intent.getStringExtra("name");
            System.out.println(content);
            Toast.makeText(context,content,Toast.LENGTH_LONG).show();

            //接收有序广播



        }
    }


```

* ProviceReceiver

```
    package com.example.myapp10;

    import android.content.BroadcastReceiver;
    import android.content.Context;
    import android.content.Intent;
    import android.widget.Toast;

    /**
    * Created by zhouchaozhong on 2018/2/10.
    */

    public class ProvinceReceiver extends BroadcastReceiver {
        @Override
        public void onReceive(Context context, Intent intent) {
            //获取到发送广播携带的数据
            String content = getResultData();
            Toast.makeText(context,"省："+content,Toast.LENGTH_LONG).show();
            System.out.println("省："+content);
            setResultData("快来领大米，每人500斤！");
            //终止广播
            abortBroadcast();
        }
    }


```

* CityReceiver

```
    public class CityReceiver extends BroadcastReceiver {
        @Override
        public void onReceive(Context context, Intent intent) {
            //获取到发送广播携带的数据
            String content = getResultData();
            Toast.makeText(context,"市："+content,Toast.LENGTH_LONG).show();
            System.out.println("市："+content);
        }
    }

```

* CountryReceiver

```
    public class CountryReceiver extends BroadcastReceiver {
        @Override
        public void onReceive(Context context, Intent intent) {
            //获取到发送广播携带的数据
            String content = getResultData();
            Toast.makeText(context,"乡："+content,Toast.LENGTH_LONG).show();
            System.out.println("乡："+content);
        }
    }

```

* FinalReceiver

```
    public class FinalReceiver extends BroadcastReceiver {
        @Override
        public void onReceive(Context context, Intent intent) {
            String content = getResultData();
            Toast.makeText(context,"最终接收者："+content,Toast.LENGTH_LONG).show();
        }
    }

```

* AndroidManifest

```
    <?xml version="1.0" encoding="utf-8"?>
    <manifest xmlns:android="http://schemas.android.com/apk/res/android"
        package="com.example.myapp10">

        <application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="@string/app_name"
            android:roundIcon="@mipmap/ic_launcher_round"
            android:supportsRtl="true"
            android:theme="@style/AppTheme">
            <activity android:name=".MainActivity">
                <intent-filter>
                    <action android:name="android.intent.action.MAIN" />

                    <category android:name="android.intent.category.LAUNCHER" />
                </intent-filter>
            </activity>
            <receiver android:name=".OutGoingCallReceiver">
                <intent-filter>
                    <action android:name="android.intent.action.NEW_OUTGOING_CALL"></action>
                    <action android:name="android.intent.action.MEDIA_MOUNTED"></action>
                    <action android:name="android.intent.action.PACKAGE_INSTALL"></action>
                    <action android:name="android.intent.action.PACKAGE_REMOVED"></action>
                    <action android:name="android.intent.action.BOOT_COMPLETED"></action>
                    <action android:name="com.example.myapp10.custombroadcast"></action>
                </intent-filter>
            </receiver>
            <receiver android:name=".ProvinceReceiver">
                <intent-filter android:priority="1000">
                    <action android:name="com.example.myapp10.rice"></action>
                </intent-filter>
            </receiver>
            <receiver android:name=".CityReceiver">
                <intent-filter android:priority="800">
                    <action android:name="com.example.myapp10.rice"></action>
                </intent-filter>
            </receiver>
            <receiver android:name=".CountryReceiver">
                <intent-filter android:priority="500">
                    <action android:name="com.example.myapp10.rice"></action>
                </intent-filter>
            </receiver>
        </application>
        <uses-permission android:name="android.permission.PROCESS_OUTGOING_CALLS"></uses-permission>
        <uses-permission android:name="android.permission.RECEIVE_BOOT_COMPLETED"></uses-permission>
    </manifest>

```

* 动态注册广播接收者（比较频繁的广播事件，只能动态注册）

```
        //动态注册屏幕解锁和锁屏的广播
        screenReceiver = new ScreenReceiver();

        //创建intent-filter对象
        IntentFilter filter = new IntentFilter();
        //添加要注册的action
        filter.addAction("android.intent.action.SCREEN_OFF");
        filter.addAction("android.intent.action.SCREEN_ON");
        this.registerReceiver(screenReceiver,filter);




        public class ScreenReceiver extends BroadcastReceiver {
            @Override
            public void onReceive(Context context, Intent intent) {
                //获取当前广播的事件类型
                String action = intent.getAction();
                //对当前广播的事件类型做一个判断
                if("android.intent.action.SCREEN_OFF".equals(action)){
                    System.out.println("屏幕熄屏了！");
                }else if("android.intent.action.SCREEN_ON".equals(action)){
                    System.out.println("屏幕亮屏了！");
                }
            }
        }

```

Android常用对话框
-------------------------------------------------------------

```
    public void onClick(View view) {
            switch (view.getId()){
                case R.id.bt_normal:
                    //点击按钮，弹出一个普通对话框。
                    AlertDialog.Builder builder = new AlertDialog.Builder(mContext);
                    builder.setTitle("警告");
                    builder.setMessage("世界上最遥远的距离是没有网络！");
                    builder.setPositiveButton("确定", new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialogInterface, int i) {
                            System.out.println("点击了确定按钮！");
                        }
                    });

                    builder.setNegativeButton("取消", new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialogInterface, int i) {
                            System.out.println("点击了取消按钮！");
                        }
                    });

                    builder.show();
                    break;
                case R.id.bt_single:
                    //单选对话框
                    final AlertDialog.Builder builder1 = new AlertDialog.Builder(mContext);
                    builder1.setTitle("选择您喜欢的课！");
                    final String items[] = {"Android","ios","PHP","C","C++"};
                    builder1.setSingleChoiceItems(items, -1, new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialogInterface, int i) {
                            //取出选中的条目
                            String item = items[i];
                            Toast.makeText(mContext,item,Toast.LENGTH_LONG).show();
                            //关闭当前对话框
                            dialogInterface.dismiss();
                        }
                    });
                    builder1.show();
                    break;
                case R.id.bt_multi:
                    //多选对话框
                    AlertDialog.Builder builder2 = new AlertDialog.Builder(mContext);
                    builder2.setTitle("选择您喜欢吃的水果！");
                    final String items2[] = {"苹果","香蕉","葡萄","樱桃","西瓜","桃子"};
                    final boolean checkItems[] = { true,false,true,false,false,false};
                    builder2.setMultiChoiceItems(items2, checkItems, new DialogInterface.OnMultiChoiceClickListener() {
                        @Override
                        public void onClick(DialogInterface dialogInterface, int i, boolean b) {

                        }
                    });
                    builder2.setPositiveButton("确定", new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialogInterface, int i) {
                            StringBuffer sb = new StringBuffer();
                            for(int j = 0;j < checkItems.length;j++){
                                if(checkItems[j]){
                                    String fruit = items2[j];
                                    sb.append(fruit+ " ");
                                }
                            }

                            Toast.makeText(mContext,sb.toString(),Toast.LENGTH_LONG).show();
                            dialogInterface.dismiss();
                        }
                    });
                    builder2.show();
                    break;
                case R.id.bt_progress:
                    //进度条对话框
                    final ProgressDialog dialog = new ProgressDialog(mContext);
                    dialog.setTitle("正在玩命加载中...");
                    dialog.setProgressStyle(ProgressDialog.STYLE_HORIZONTAL);
                    dialog.show();
                    //设置进度条最大值
                    dialog.setMax(100);
                    //模拟更新进度条
                    new Thread(){
                        public void run(){
                            for(int k = 0;k <= 100;k++){
                                //设置当前进度
                                SystemClock.sleep(100);
                                dialog.setProgress(k);

                            }

                            //关闭对话框
                            dialog.dismiss();
                        }
                    }.start();

                    break;
                default:
                    break;
            }
        }

```

服务
-------------------------------------------------------

* MainActivity.java

```
    public class MainActivity extends AppCompatActivity implements View.OnClickListener {
        private Context mContext;
        private Button bt_start;
        private Button bt_stop;
        private Button bt_bind;
        private Myconn conn;
        private Intent intent;
        private Intent intent1;

        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            bt_start = findViewById(R.id.bt_start);
            bt_bind = findViewById(R.id.bt_bind);
            bt_stop = findViewById(R.id.bt_stop);
            bt_start.setOnClickListener(this);
            bt_bind.setOnClickListener(this);
            bt_stop.setOnClickListener(this);
            GrantPermission gp = new GrantPermission();
            gp.requestPermissions(this);
        }

        @Override
        public void onClick(View view) {
            switch (view.getId()){
                case R.id.bt_start:
                    //开启服务
                    intent = new Intent(this,FirstService.class);
                    startService(intent);

                    break;
                case R.id.bt_bind:
                    intent1 = new Intent(this,FirstService.class);
                    conn = new Myconn();
                    bindService(intent1,conn,BIND_AUTO_CREATE);
                    break;
                case R.id.bt_stop:
                    stopService(intent);
                    break;
                default:
                    break;
            }
        }

        //用来监听服务的状态
        private class Myconn implements ServiceConnection{

            //连接成功后调用
            @Override
            public void onServiceConnected(ComponentName componentName, IBinder iBinder) {

            }

            //连接中断调用
            @Override
            public void onServiceDisconnected(ComponentName componentName) {

            }
        }


    }

```

* FirstService.java

```
    public class FirstService extends Service {
        private MediaRecorder recorder;
        @Nullable
        @Override
        public IBinder onBind(Intent intent) {
            System.out.println("onBind!");
            return null;
        }

        @Override
        public void onCreate() {
            TelephonyManager tm = (TelephonyManager) getSystemService(TELEPHONY_SERVICE);
            tm.listen(new MyPhoneStateListener(),PhoneStateListener.LISTEN_CALL_STATE);
            super.onCreate();
        }

        private class MyPhoneStateListener extends PhoneStateListener{
            @Override
            public void onCallStateChanged(int state, String incomingNumber) {
                System.out.println("onCallStateChanged方法调用");
                switch (state){
                    case TelephonyManager.CALL_STATE_IDLE:
                        if(recorder != null){
                            recorder.stop();
                            recorder.reset();
                            recorder.release();
                        }
                        break;
                    case TelephonyManager.CALL_STATE_OFFHOOK:
                        recorder.start();
                        break;
                    case TelephonyManager.CALL_STATE_RINGING:
                        recorder = new MediaRecorder();
                        recorder.setAudioSource(MediaRecorder.AudioSource.MIC);
                        recorder.setOutputFormat(MediaRecorder.OutputFormat.THREE_GPP);
                        recorder.setAudioEncoder(MediaRecorder.AudioEncoder.AMR_NB);
                        recorder.setOutputFile("/mnt/sdcard/recorder.3gp");
                        try {
                            recorder.prepare();
                        } catch (IOException e) {
                            e.printStackTrace();
                        }
                        break;
                    default:
                        break;
                }
                super.onCallStateChanged(state, incomingNumber);
            }
        }

        @Override
        public int onStartCommand(Intent intent, int flags, int startId) {
            return super.onStartCommand(intent, flags, startId);
        }

        @Override
        public void onDestroy() {
            System.out.println("onDestroy被调用了！");
            super.onDestroy();
        }
    }

```

* AndroidManifest.xml

```
    <?xml version="1.0" encoding="utf-8"?>
    <manifest xmlns:android="http://schemas.android.com/apk/res/android"
        package="com.example.myapp12">

        <application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="@string/app_name"
            android:roundIcon="@mipmap/ic_launcher_round"
            android:supportsRtl="true"
            android:theme="@style/AppTheme">
            <activity android:name=".MainActivity">
                <intent-filter>
                    <action android:name="android.intent.action.MAIN" />

                    <category android:name="android.intent.category.LAUNCHER" />
                </intent-filter>
            </activity>
            <service android:name=".FirstService"></service>
        </application>
        <uses-permission android:name="android.permission.RECORD_AUDIO"></uses-permission>
        <uses-permission android:name="android.permission.READ_PHONE_STATE"></uses-permission>
        <uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE"></uses-permission>
        <uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE"></uses-permission>
    </manifest>

```
* StartService开启服务的特点
  1. 第一次开启服务，会执行服务的onCreate方法和onStart方法
  2. 第二次开启服务，会执行服务的onStart方法
  3. 服务开启后，能在设置页面找到这个服务
  4. startService方式开启服务，服务会在后台长期运行，知道用户手工停止，或者调用stopService方法，服务才会被销毁
  

* bindService开启服务的特点
  1. 当点击按钮第一次开启服务，会执行onCreate方法和onBind方法
  2. 第二次调用bindService，服务没有响应
  3. 当activity销毁的时候，服务也销毁
  4. 通过bind方式开启服务，服务不能在设置页面找到，相当于是一个隐式服务
  5. bindService不能多次解绑，多次解绑会报错


内容提供者、内容解析者、内容观察者
----------------------------------------------------------------

* MainActivity.java

```
    public class MainActivity extends AppCompatActivity implements View.OnClickListener {
        private Context mContext;
        private Button bt_getSms;
        private Button bt_bak_contacts;

        @Override
        protected void onCreate(Bundle savedInstanceState) {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_main);

            mContext = this;
            bt_getSms = findViewById(R.id.bt_getSms);
            bt_bak_contacts = findViewById(R.id.bt_bak_contacts);
            bt_getSms.setOnClickListener(this);
            bt_bak_contacts.setOnClickListener(this);

            //动态申请权限
            ActivityCompat.requestPermissions(this,new String[]{Manifest.permission.READ_SMS,Manifest.permission.WRITE_EXTERNAL_STORAGE,Manifest.permission.READ_EXTERNAL_STORAGE,Manifest.permission.READ_CONTACTS},1000);
        }

        @Override
        public void onClick(View view) {
            switch (view.getId()){
                case R.id.bt_getSms:
                    //备份短信到xml文件

                    //获取XmlSerializer的实例
                    XmlSerializer serializer = Xml.newSerializer();
                    //设置序列化器参数
                    File file = new File(Environment.getExternalStorageDirectory().getPath(),"smsbackup.xml");
                    try{
                        FileOutputStream fos = new FileOutputStream(file);
                        serializer.setOutput(fos,"utf-8");
                        //写xml文档开头
                        serializer.startDocument("utf-8",true);
                        serializer.startTag(null,"smss");

                        //构造Uri
                        Uri uri = Uri.parse("content://sms/");
                        //通过内容解析者查询
                        Cursor cursor = getContentResolver().query(uri,new String[]{"address","date","body"},null,null,null);
                        while(cursor.moveToNext()){
                            String address = cursor.getString(0);
                            String date = cursor.getString(1);
                            String body = cursor.getString(2);
                            //写sms节点
                            serializer.startTag(null,"sms");

                            serializer.startTag(null,"address");
                            serializer.text(address);
                            serializer.endTag(null,"address");

                            serializer.startTag(null,"date");
                            serializer.text(date);
                            serializer.endTag(null,"date");

                            serializer.startTag(null,"body");
                            serializer.text(body);
                            serializer.endTag(null,"body");

                            serializer.endTag(null,"sms");
                        }

                        serializer.endTag(null,"smss");
                        serializer.endDocument();
                        fos.close();
                        fos.flush();
                    }catch (Exception e){
                        e.printStackTrace();
                    }

                    break;

                case R.id.bt_bak_contacts:
                    Uri uri2 = Uri.parse("content://com.android.contacts/raw_contacts");
                    Uri dataUri = Uri.parse("content://com.android.contacts/data");
                    //这里查询data，真正查询的是view_data视图，没有mimetype_id列
                    Cursor cursor2 = getContentResolver().query(uri2,new String[]{"contact_id"},null,null,null);
                    while(cursor2.moveToNext()){
                        String contact_id = cursor2.getString(0);
                        Cursor dataCursor = getContentResolver().query(dataUri,new String[]{ "data1","mimetype" },"raw_contact_id=?",new String[]{ contact_id },null);
                        while(dataCursor.moveToNext()){
                            String data1 = dataCursor.getString(0);
                            String mimetype= dataCursor.getString(1);
                            if("vnd.android.cursor.item/name".equals(mimetype)){
                                //姓名
                                System.out.println("姓名：" + data1);
                            }else if("vnd.android.cursor.item/phone_v2".equals(mimetype)){
                                //电话号码
                                System.out.println("电话号码：" + data1);
                            }else if("vnd.android.cursor.item/email_v2".equals(mimetype)){
                                //邮箱
                                System.out.println("邮箱：" + data1);
                            }
                        }
                    }
                    break;
                default:
                    break;
            }
        }

    }

```

* AndroidManifest.xml

```<!--lang:xml-->
    <?xml version="1.0" encoding="utf-8"?>
    <manifest xmlns:android="http://schemas.android.com/apk/res/android"
        package="com.example.myapp13">

        <application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="@string/app_name"
            android:roundIcon="@mipmap/ic_launcher_round"
            android:supportsRtl="true"
            android:theme="@style/AppTheme">
            <activity android:name=".MainActivity">
                <intent-filter>
                    <action android:name="android.intent.action.MAIN" />

                    <category android:name="android.intent.category.LAUNCHER" />
                </intent-filter>
            </activity>
        </application>
        <uses-permission android:name="android.permission.READ_SMS"></uses-permission>
        <uses-permission android:name="android.permission.READ_EXTERNAL_STORAGE"></uses-permission>
        <uses-permission android:name="android.permission.WRITE_EXTERNAL_STORAGE"></uses-permission>
        <uses-permission android:name="android.permission.READ_CONTACTS"></uses-permission>
    </manifest>

```

apk反编译
---------------------------------------------------------------

* 需要用到的工具
  apktool  dex2jar jd-gui

* apk反编译步骤

 1. apktool d -f baiduinput.apk  得到app的资源文件
 2. 将apk文件解压，得到classes.dex文件，然后将classes.dex文件复制到d2j-dex2jar.bat所在的目录下，然后执行命令

    d2j-dex2jar classes.dex   得到classes-dex2jar.jar文件
 3. 用jd-gui 打开 classes-dex2jar.jar文件，即可看到源代码
  





