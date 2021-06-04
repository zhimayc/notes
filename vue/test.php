<?php

if(!empty($_GET)){
	$arr = $_GET;
	$json = json_encode($arr);
	echo $json;
}

if(!empty($_POST)){
	if($_POST['a'] == 3){
		$arr = array(
			'a'=>'香蕉',
			'b'=>'橘子',
			'c'=>'苹果'
		);
	}else{
		$arr = $_POST;
	}
	
	$json = json_encode($arr);
	echo $json;
}

?>