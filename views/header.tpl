<!DOCTYPE html>
<html lang="en">
<head>
 <title>Big Setts Test Beego Application - Wickedly Works</title>
 <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
 <link rel="stylesheet" href="/static/css/bootstrap.min.css">
 <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
 <link href="/static/css/starter-template.css" rel="stylesheet">
</head>
<body>
  <div class="navbar navbar-inverse navbar-fixed-top" role="navigation">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="#">Beego Test App</a>
      </div>
      <div class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li><a href="{{urlfor "ManageController.Index"}}">Home</a></li>
          <li class="dropdown">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">Manage Records <b class="caret"></b></a>
            <ul class="dropdown-menu">
              <li><a href="{{urlfor "ManageController.New"}}">Add Record</a></li>
              <li><a href="{{urlfor "ManageController.Index"}}">View All</a></li>
            </ul>
          </li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>
  </div>
