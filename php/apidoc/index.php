<?php

/**
 * @api {post} /article/:id 获取文章标题
 * @apiVersion 1.0
 * @apiDeprecated 该接口在2.0版本以后已被弃用，不推荐使用
 * @apiExample {post} example:
 * ajax.post{
 *      'id':1
 * }
 * 
 * @apiPermission none
 * @apiDescription 根据用户ID，获取用户的姓名，性别，年龄
 * @apiSampleRequest http://web.myapp/test.php
 * @apiName GetArticle
 * @apiGroup Article
 * @apiParam {int} id 文章唯一ID
 * @apiSuccess {string} title 文章标题
 * @apiSuccess {string} date 日期
 * @apiSuccess {string} author 作者
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 OK
 * {
 *      "title":"s8比赛火热进行中",
 *      "date":"2018-10-22",
 *      "author":"今日头条"
 * }
 * 
 * @apiError UserNotFound 没有找到该文章
 * @apiErrorExample Error-Response:
 * HTTP/1.1 404 Not Found
 * {
 *     "error":"没有找到该文章的信息"
 * }
 * 
 * 
 */
function article($id){
    echo $id;
}

test(2);
