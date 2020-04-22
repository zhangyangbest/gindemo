

$('#123456').submitForm({
    url:"/login",
    dataType:"json",
    callback:function(data){
        endFileUpload();
        data=eval("("+data+")");
        alert(data.Content);
        if(data.Result > 0){
            location.href = data.Redirect;
        }
    },
    before:function(){
        startFileUpload();
        var errMsg ="";
    }
}).submit();