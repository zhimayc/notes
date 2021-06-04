<?php

/**
 * @api {post} /usr/:id 获取用户信息
 * @apiDeprecated 该接口在2.0版本以后已被弃用，不推荐使用
 * @apiExample {post} example:
 * ajax.post{
 *      'id':1
 * }
 * 
 * @apiPermission none
 * @apiDescription 根据用户ID，获取用户的姓名，性别，年龄
 * @apiSampleRequest http://web.myapp/test.php
 * @apiName GetUser
 * @apiGroup User
 * @apiParam {int} id 用户唯一ID
 * @apiSuccess {string} name 姓名
 * @apiSuccess {string} sex 性别
 * @apiSuccess {int} age 年龄
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 OK
 * {
 *      "name":"杨菲",
 *      "sex":"女",
 *      "age":"18"
 * }
 * 
 * @apiError UserNotFound 没有找到该用户
 * @apiErrorExample Error-Response:
 * HTTP/1.1 404 Not Found
 * {
 *     "error":"没有找到该用户的信息"
 * }
 * 
 * 
 */
function test($id){
    echo $id;
}

test(2);




/**
 * @api /department/:id 获取用户薪水
 * @apiName GetSalary
 * @apiGroup Salary
 * @apiParam {int} id 用户唯一ID
 * @apiSuccess {string} name 姓名
 * @apiSuccess {string} salary 薪水
 * @apiSuccess {string} department 部门
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 OK
 * {
 *      "name":"杨菲",
 *      "department":"客服部",
 *      "salary":8000
 * }
 * 
 * @apiError UserNotFound 没有找到该用户
 * @apiErrorExample Error-Response:
 * HTTP/1.1 404 Not Found
 * {
 *     "error":"没有找到该用户的信息"
 * }
 * 
 * 
 */
function getUser(){

}