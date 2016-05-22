var tabs = (function(){

    function addTab(tabSelector,title,iconCls,url,id){
        if(!isExistTab(tabSelector, title)){
            $(tabSelector).tabs('add', {
                title: title, iconCls:iconCls, id: id,
                bodyCls: "tabCls",
                content:'<iframe name="iframe" id="iframe-' + id + '" src="' + url + '" frameborder="no" style="overflow: hidden;" scrolling="auto" width="100%" height="100%" allowtransparency="true"></iframe>'
            });
        }else{
            selectTab(tabSelector, title);
        }
    }

    function isExistTab(tabSelector, title){
        return $(tabSelector).tabs('exists', title);
    }

    function closeTab(){

    }

    function selectTab(tabSelector, title){
        $(tabSelector).tabs('select', title);
    }

    return {
        addTab:addTab,
        closeTab:closeTab,
        selectTab:selectTab
    }

})();

var home = (function(){

    var selectorTabs = "#global-tabs";
    var urlAppLogListPage = "/AppLog/List";

    function navInit(){
        $("#global-nav").tree({
            lines:true,
            url:"/Main/GetNavTreeJson",
            onClick:function(node){
                tabs.addTab(selectorTabs, node.text, node.iconCls, urlAppLogListPage, node.id)
            }
        });
    }

    function init(){
        navInit();
    }

    return {
        init:init
    }
})();