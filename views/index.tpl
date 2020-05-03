<!DOCTYPE html>
<html>
   <head>
      <title>Bootstrap 模板</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <!-- 引入 Bootstrap -->
      <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">

      <!-- HTML5 Shiv 和 Respond.js 用于让 IE8 支持 HTML5元素和媒体查询 -->
      <!-- 注意： 如果通过 file://  引入 Respond.js 文件，则该文件无法起效果 -->
      <!--[if lt IE 9]>
         <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
         <script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
      <![endif]-->
   </head>
   <body>
       <h3 class="red-text text-center">用户列表</h2>

<table class="table table-bordered" style="width:80%">
  <thead>
    <tr>
      <th style='text-align: center;'>用户名</th>
      <th style='text-align: center;' >邀请码</th>
      <th style='text-align: center;' >余额</th>
      <th style='text-align: center;' >认证</th>
    </tr>
  </thead>
  <tbody>

 {{range .Users}}

 	  <tr>
          <td style='text-align: center;'>  {{.Nick_name}}</td>
          <td style='text-align: center;'>{{.Invite_code}}</td>
          <td style='text-align: center;' >{{.Balance}}</td>
          {{if eq .Verify 0}}
              <td style='text-align: center;' >未认证</td>
          {{else}}
              <td style='text-align: center;' >已认证</td>
          {{end}}

       </tr>
{{end}}
  </tbody>
</table>
      <!-- jQuery (Bootstrap 的 JavaScript 插件需要引入 jQuery) -->
      <script src="https://code.jquery.com/jquery.js"></script>
      <!-- 包括所有已编译的插件 -->
      <script src="/statics/js/bootstrap.min.js"></script>
   </body>
</html>