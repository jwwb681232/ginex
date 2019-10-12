function mini(){
    var sidebarClass = 'page-sidebar-minified';
    var targetContainer = '#page-container';
    if ($(targetContainer).hasClass(sidebarClass)) {
        $(targetContainer).removeClass(sidebarClass);
    } else {
        $(targetContainer).addClass(sidebarClass);
    }
    $(window).trigger('resize');
}