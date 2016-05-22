var AppLogListPage = (function(){

    var selectorDG = "#dg-log";

    function datagridInit(){
        $(selectorDG).datagrid({
            fit:true,
            title:"应用日志列表",
            fitColumns:true,
            singleSelect:true,
            rownumbers:true,
            toolbar:"#dg-log-toolbar",
            pagination:true,
            pageList: [20,40,50,100],
            pageSize: 20,
            columns:[[
                    {field:'time',title:'Time',width:'140'},
                    {field:'app',title:'App',width:'100'},
                    {field:'level',title:'Level',width:'100'},
                    {field:'ip',title:'IP',width:'100'},
                    {field:'msg',title:'Msg',width:100}
             ]],
             url:"/AppLog/GetListJson",
             onSelect:function(index,row){
                $("#log-detail").html(row.msg);
             }
        });
    }

    function query(){
        $(selectorDG).datagrid('load', {
            website: $("#drp-website").combobox('getValue'),
            logtype: $("#drp-logtype").combobox('getValue')
        });
        $("#log-detail").html("");
    }

    function reload(){
        $(selectorDG).datagrid('reload');
        $("#log-detail").html("");
    }

    function readfromfile(){
        $.ajax({
            url:"/AppLog/ReadFromTodayFile",
            dataType:"json",
            success:function(data){
                $.messager.show({ title:'提示', msg: data.message, timeout:1000, showType:'fade', style: {right:0}});
                if(data.success){
                    query();
                }
            }
        })
    }

    function init(){
        datagridInit();
        $("#btn-reload").click(reload);
        $("#btn-search").click(query);
        $("#btn-readfromfile").click(readfromfile);
    }

    return {
        init:init
    }
})();
