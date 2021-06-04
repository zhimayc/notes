<?php

	//冒泡排序 将一个数组从小到大排序



    /*function mysort($arr){
    	for($i = 0;$i < count($arr);$i++){
    		for($j = 0;$j < count($arr)-1-$i;$j++){
    			if($arr[$j] > $arr[$j+1]){
    				$tmp = $arr[$j];
    				$arr[$j] = $arr[$j+1];
    				$arr[$j+1] = $tmp;
    			}
    		}

    	}

    	return $arr;

    }

    $arr = array(2,5,6,8,1,3,4,7);
    print_r(mysort($arr));*/




    //文件读写


   // $filename = 'test.txt';
    //$myfile = fopen($filename,'r+');
    //$tmp = fread($myfile,filesize($filename));
    //$tmp = fgets($myfile);
    //var_dump($tmp);
    /*while(!feof($myfile)){
    	echo fgetc($myfile);
    }
    fclose($myfile);
    */
    // $filename = 'test2.txt';
    // $myfile = fopen($filename,'a');
    // $txt = "Hello!\r\nHow are you?\r\nI'm fine!thank you!";
    // fwrite($myfile,$txt);
    // fclose($myfile);


    //xml文件读取
    $filename = './classes.xml';
    $dom = new DOMDocument();
    $dom->load($filename);
    $stu = $dom->getElementsByTagName('stu');
    $arr = array();
    foreach($stu as $key=>$value){
    	$name = $value->getElementsByTagName('name')->item(0)->nodeValue;
    	$age = $value->getElementsByTagName('age')->item(0)->nodeValue;
    	$sex = $value->getElementsByTagName('sex')->item(0)->nodeValue;
    	$arr[$key]['name'] = $name;
    	$arr[$key]['age'] = $age;
    	$arr[$key]['sex'] = $sex;
    }
    $info = '';
    foreach($arr as $v){
    	$info .= '姓名：'.$v["name"].'  年龄：'.$v["age"].'  性别：'.$v["sex"].'<pre></pre>';

    }

    echo $info;

    //xml文件写入
    // $filename = './test.xml';
    // $dom = new DOMDocument('1.0','utf8');
    // $dom->formatOutput = true;
    // $student = $dom->createElement('student');
    // $dom->appendChild($student);
    // $name = $dom->createElement('name');
    // $student->appendChild($name);
    // $name->nodeValue = '杨菲';

    // $age = $dom->createElement('age');
    // $student->appendChild($age);
    // $age->nodeValue = 18;

    // $student = $dom->createElement('student');
    // $dom->appendChild($student);
    // $name = $dom->createElement('name');
    // $student->appendChild($name);
    // $name->nodeValue = '杨静';

    // $age = $dom->createElement('age');
    // $student->appendChild($age);
    // $age->nodeValue = 20;

    // $dom->save($filename);


    

 



?>