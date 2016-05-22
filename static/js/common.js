$.extend($.messager, {
   ajaxTitle: function (msg) {
       if (msg == "close") {
           $("#ajaxbodymsg,#ajaxtitlemsg").remove();
       } else {
           $(".datagrid-mask,.datagrid-mask-msg").remove();
           $("<div class=\"datagrid-mask\" id=\"ajaxbodymsg\" style=\"z-index: 9000\"></div>").css({
               display: "block",
               width: "100%",
               height: $(window).height()
           }).appendTo("body");
           $("<div class=\"datagrid-mask-msg\" id=\"ajaxtitlemsg\" style=\"z-index: 9001\"></div>").html(msg).appendTo("body").css({
               display: "block",
               left: ($(document.body).outerWidth(true) - 190) / 2,
               top: ($(window).height() - 45) / 2
           });
       }
   }
});

$.ajaxSetup({
  beforeSend : function(){
    $.messager.ajaxTitle("<font style='font-size:12px;'>正在处理，请稍候...</font>");
  },
  complete :function(){
    $.messager.ajaxTitle("close");
  },
  error:function(){
    $.messager.alert("错误提示-请复制错误代码，发给技术支持", request.responseText, "error");
  }
});
