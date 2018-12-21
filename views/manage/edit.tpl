<div class="container">
  <div class="row">

    <h2>Article Details</h2>

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}

    <p>
      <form role="form" id="user" action="{{urlfor "ManageController.Put" ":id" .Article.Id}}" method="POST">
         <input type="hidden" name="_method" value="PUT" />
        <div class="form-group {{if .Errors.Name}}has-error has-feedback{{end}}">
          <label for="name">Article name： {{if .Errors.Name}}({{.Errors.Name}}){{end}}</label>
          <input name="name" type="text" value="{{.Article.Name}}" class="form-control" tabindex="1" />
          {{if .Errors.Name}}<span class="glyphicon glyphicon-remove form-control-feedback"></span>{{end}}
        </div>

        <div class="form-group">
          <label for="client">Client：</label>
          <input name="client" type="text" value="{{.Article.Client}}" class="form-control" tabindex="2" />
        </div>

        <div class="form-group">
          <label for="url">URL：</label>
          <input name="url" type="text" value="{{.Article.Url}}" class="form-control" tabindex="3" />
        </div>
<!--
        {#{.Form | renderform}#}
-->
        <input type="submit" value="Update Article" class="btn btn-default" tabindex="4" /> &nbsp;
        <a href="{{urlfor "ManageController.Index"}}" class="btn btn-default" title="don't update the article">Cancel</a>
      </form>
    </p>
  </div>
</div>
