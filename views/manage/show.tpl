<div class="container">
  <div class="row">

    <h2>Article Details</h2>

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
    <a href="{{urlfor "ManageController.Index"}}" >Back to list</a>
    <td>{{.record.Id}}</td>
    <td>{{.record.Name}}</td>
    <td>{{.record.Client}}</td>
    <td>{{.record.Url}}</td>
    <a href="{{urlfor "ManageController.Edit" ":id" .record.Id}}" >Edit</a>
    <form role="form" id="user" action="{{urlfor "ManageController.Delete" ":id" .record.Id}}" method="POST">
         <input type="hidden" name="_method" value="DELETE" />
        <input type="submit" value="Delete Article" class="btn btn-danger" tabindex="4" />
      </form>
  </div>
</div>
