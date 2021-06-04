<?php
	
	//MySQLi方式操作数据库
	/*$mysqli = new MySQLi('localhost','root','root','test');
	$mysqli->select_db('test');
	$results = $mysqli->query('select * from user');
	$data = array();
	while($row = $results->fetch_assoc()){
		array_push($data,$row);
	}

	$txt = '';

	foreach($data as $key=>$value){
		foreach($value as $k=>$v){
			$txt .= $k.':'.$v.'&nbsp;&nbsp;|&nbsp;&nbsp;';
		}

		$txt .= '<pre style="background-color:silver;height:0.25em"></pre>';
	}

	echo $txt;
	$results->free();
	$mysqli->close();*/


	//PDO方式操作数据库
	$dsn = 'mysql:dbname=test;host=127.0.0.1';
	$username = 'root';
	$password = 'root';
	try {
		$dbh = new PDO($dsn,$username,$password);
	}catch(PDOException $e){
		echo 'connection failed:'.$e->getMessage();
	}

	$sql = 'select * from user';
	$results = $dbh->query($sql);
	$data = array();

	while($row = $results->fetch(PDO::FETCH_ASSOC)){
		array_push($data,$row);
	}

	$txt = '';

	foreach($data as $key=>$value){
		foreach($value as $k=>$v){
			$txt .= $k.':'.$v.'&nbsp;&nbsp;|&nbsp;&nbsp;';
		}

		$txt .= '<pre style="background-color:silver;height:0.25em"></pre>';
	}
	
	echo $txt;


?>